package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	api "github.com/bytebase/bytebase/backend/legacyapi"
	"github.com/bytebase/bytebase/backend/store"
)

func (s *Server) registerPolicyRoutes(g *echo.Group) {
	g.PATCH("/policy/:resourceType/:resourceID", func(c echo.Context) error {
		ctx := c.Request().Context()
		resourceType, err := api.GetPolicyResourceType(c.Param("resourceType"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		resourceID, err := getPolicyResourceID(c.Param("resourceID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		pType := api.PolicyType(c.QueryParam("type"))
		if err := api.ValidatePolicyType(pType); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		updaterID := c.Get(getPrincipalIDContextKey()).(int)

		policyUpsert := &api.PolicyUpsert{}
		if err := jsonapi.UnmarshalPayload(c.Request().Body, policyUpsert); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Malformed set policy request").SetInternal(err)
		}
		if pType == api.PolicyTypeSQLReview && policyUpsert.Payload != nil {
			policy, err := api.UnmarshalSQLReviewPolicy(*policyUpsert.Payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal policy payload").SetInternal(err)
			}
			payload := api.FlattenSQLReviewRulesWithEngine(policy)
			payloadStr, err := payload.String()
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal policy payload").SetInternal(err)
			}
			policyUpsert.Payload = &payloadStr
		}

		if err := s.hasAccessToUpsertPolicy(pType, policyUpsert); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error()).SetInternal(err)
		}
		// Validate policy.
		if err := api.ValidatePolicy(resourceType, pType, policyUpsert.Payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid policy payload %s", err.Error())).SetInternal(err)
		}
		var composedEnvironment *api.Environment
		if resourceType == api.PolicyResourceTypeEnvironment {
			composedEnvironment, err = s.store.GetEnvironmentByID(ctx, resourceID)
			if err != nil {
				return err
			}
		}

		if pType == api.PolicyTypeEnvironmentTier {
			environment, err := s.store.GetEnvironmentV2(ctx, &store.FindEnvironmentMessage{UID: &resourceID})
			if err != nil {
				return err
			}
			if environment == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "environment not found")
			}
			if policyUpsert.Payload == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "empty payload")
			}
			tierPolicy, err := api.UnmarshalEnvironmentTierPolicy(*policyUpsert.Payload)
			if err != nil {
				return err
			}
			protected := false
			if tierPolicy.EnvironmentTier == api.EnvironmentTierValueProtected {
				protected = true
			}
			if _, err := s.store.UpdateEnvironmentV2(ctx, environment.ResourceID, &store.UpdateEnvironmentMessage{Protected: &protected}, updaterID); err != nil {
				return err
			}
			composedPolicy := &api.Policy{
				ID:                resourceID,
				RowStatus:         api.Normal,
				ResourceType:      resourceType,
				ResourceID:        resourceID,
				Type:              pType,
				InheritFromParent: true,
				Payload:           *policyUpsert.Payload,
				Environment:       composedEnvironment,
			}
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			if err := jsonapi.MarshalPayload(c.Response().Writer, composedPolicy); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal create set policy response").SetInternal(err)
			}
			return nil
		}

		policy, err := s.store.GetPolicyV2(ctx, &store.FindPolicyMessage{
			ResourceType: &resourceType,
			ResourceUID:  &resourceID,
			Type:         &pType,
		})
		if err != nil {
			return err
		}
		if policy == nil {
			payload := ""
			if policyUpsert.Payload != nil {
				payload = *policyUpsert.Payload
			}
			inheritFromParent := true
			if policyUpsert.InheritFromParent != nil {
				inheritFromParent = *policyUpsert.InheritFromParent
			}
			policy, err = s.store.CreatePolicyV2(ctx, &store.PolicyMessage{
				ResourceType:      resourceType,
				ResourceUID:       resourceID,
				Type:              pType,
				Payload:           payload,
				InheritFromParent: inheritFromParent,
				// Enforce cannot be force while creating a policy.
				Enforce: true,
			}, updaterID)
			if err != nil {
				return err
			}
		} else {
			var enforce *bool
			if policyUpsert.RowStatus != nil {
				if *policyUpsert.RowStatus == string(api.Normal) {
					t := true
					enforce = &t
				} else {
					f := false
					enforce = &f
				}
			}
			policy, err = s.store.UpdatePolicyV2(ctx, &store.UpdatePolicyMessage{
				UpdaterID:         updaterID,
				ResourceType:      resourceType,
				ResourceUID:       resourceID,
				Type:              pType,
				InheritFromParent: policyUpsert.InheritFromParent,
				Payload:           policyUpsert.Payload,
				Enforce:           enforce,
			})
			if err != nil {
				return err
			}
		}

		if pType == api.PolicyTypeSQLReview {
			policy.Payload = api.MergeSQLReviewRulesWithoutEngine(policy.Payload)
		}

		composedPolicy := &api.Policy{
			ID:                policy.UID,
			RowStatus:         api.Normal,
			ResourceType:      resourceType,
			ResourceID:        resourceID,
			Type:              pType,
			InheritFromParent: policy.InheritFromParent,
			Payload:           policy.Payload,
			Environment:       composedEnvironment,
		}
		if !policy.Enforce {
			composedPolicy.RowStatus = api.Archived
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		if err := jsonapi.MarshalPayload(c.Response().Writer, composedPolicy); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal create set policy response").SetInternal(err)
		}
		return nil
	})

	g.DELETE("/policy/:resourceType/:resourceID", func(c echo.Context) error {
		ctx := c.Request().Context()
		resourceType, err := api.GetPolicyResourceType(c.Param("resourceType"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		resourceID, err := getPolicyResourceID(c.Param("resourceID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		pType := api.PolicyType(c.QueryParam("type"))
		if err := api.ValidatePolicyType(pType); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}

		policy, err := s.store.GetPolicyV2(ctx, &store.FindPolicyMessage{
			ResourceType: &resourceType,
			ResourceUID:  &resourceID,
			Type:         &pType,
		})
		if err != nil {
			return err
		}
		if policy == nil {
			return echo.NewHTTPError(http.StatusNotFound, "policy not found")
		}

		if err := s.store.DeletePolicyV2(ctx, policy); err != nil {
			return err
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return nil
	})

	g.GET("/policy/:resourceType/:resourceID", func(c echo.Context) error {
		ctx := c.Request().Context()
		resourceType, err := api.GetPolicyResourceType(c.Param("resourceType"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		resourceID, err := getPolicyResourceID(c.Param("resourceID"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		pType := api.PolicyType(c.QueryParam("type"))
		if err := api.ValidatePolicyType(pType); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		var composedEnvironment *api.Environment
		if resourceType == api.PolicyResourceTypeEnvironment {
			composedEnvironment, err = s.store.GetEnvironmentByID(ctx, resourceID)
			if err != nil {
				return err
			}
		}
		if pType == api.PolicyTypeEnvironmentTier {
			environment, err := s.store.GetEnvironmentV2(ctx, &store.FindEnvironmentMessage{UID: &resourceID})
			if err != nil {
				return err
			}
			tierPolicy := api.EnvironmentTierPolicy{
				EnvironmentTier: api.EnvironmentTierValueUnprotected,
			}
			if environment.Protected {
				tierPolicy.EnvironmentTier = api.EnvironmentTierValueProtected
			}
			payload, err := tierPolicy.String()
			if err != nil {
				return err
			}
			composedPolicy := &api.Policy{
				ID:                resourceID,
				RowStatus:         api.Normal,
				ResourceType:      resourceType,
				ResourceID:        resourceID,
				Type:              pType,
				InheritFromParent: true,
				Payload:           payload,
				Environment:       composedEnvironment,
			}
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			if err := jsonapi.MarshalPayload(c.Response().Writer, composedPolicy); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal policy response").SetInternal(err)
			}
			return nil
		}

		policy, err := s.store.GetPolicyV2(ctx, &store.FindPolicyMessage{
			ResourceType: &resourceType,
			ResourceUID:  &resourceID,
			Type:         &pType,
		})
		if err != nil {
			return err
		}
		if policy == nil {
			// Return the default policy when there is no stored policy.
			payload, err := api.GetDefaultPolicy(pType)
			if err != nil {
				return err
			}
			policy = &store.PolicyMessage{
				UID:               api.DefaultPolicyID,
				ResourceType:      resourceType,
				ResourceUID:       resourceID,
				Type:              pType,
				InheritFromParent: true,
				Payload:           payload,
				Enforce:           true,
			}
		}

		if pType == api.PolicyTypeSQLReview {
			policy.Payload = api.MergeSQLReviewRulesWithoutEngine(policy.Payload)
		}

		composedPolicy := &api.Policy{
			ID:                policy.UID,
			ResourceType:      policy.ResourceType,
			ResourceID:        policy.ResourceUID,
			Type:              policy.Type,
			InheritFromParent: policy.InheritFromParent,
			Payload:           policy.Payload,
			RowStatus:         api.Normal,
			Environment:       composedEnvironment,
		}
		if !policy.Enforce {
			composedPolicy.RowStatus = api.Archived
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		if err := jsonapi.MarshalPayload(c.Response().Writer, composedPolicy); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to marshal get policy response: %v", pType)).SetInternal(err)
		}
		return nil
	})

	g.GET("/policy", func(c echo.Context) error {
		var resourceType *api.PolicyResourceType
		var resourceID *int
		if c.QueryParam("resourceType") != "" {
			rt, err := api.GetPolicyResourceType(c.QueryParam("resourceType"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
			}
			resourceType = &rt
		}
		if c.QueryParam("resourceId") != "" {
			id, err := getPolicyResourceID(c.QueryParam("resourceID"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
			}
			resourceID = &id
		}
		pType := api.PolicyType(c.QueryParam("type"))
		if err := api.ValidatePolicyType(pType); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}

		ctx := c.Request().Context()
		policies, err := s.store.ListPoliciesV2(ctx, &store.FindPolicyMessage{
			ResourceType: resourceType,
			ResourceUID:  resourceID,
			Type:         &pType,
		})
		if err != nil {
			return err
		}
		var composedPolicies []*api.Policy
		for _, policy := range policies {
			policy := policy
			if pType == api.PolicyTypeSQLReview {
				policy.Payload = api.MergeSQLReviewRulesWithoutEngine(policy.Payload)
			}

			composedPolicy := &api.Policy{
				ID:                policy.UID,
				ResourceType:      policy.ResourceType,
				ResourceID:        policy.ResourceUID,
				Type:              policy.Type,
				InheritFromParent: policy.InheritFromParent,
				Payload:           policy.Payload,
				RowStatus:         api.Normal,
			}
			if !policy.Enforce {
				composedPolicy.RowStatus = api.Archived
			}
			if policy.ResourceType == api.PolicyResourceTypeEnvironment {
				composedEnvironment, err := s.store.GetEnvironmentByID(ctx, policy.ResourceUID)
				if err != nil {
					return err
				}
				if composedEnvironment == nil {
					return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("environment %v not found", policy.ResourceUID))
				}
				composedPolicy.Environment = composedEnvironment
			}
			composedPolicies = append(composedPolicies, composedPolicy)
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		if err := jsonapi.MarshalPayload(c.Response().Writer, composedPolicies); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to marshal list policy response: %v", pType)).SetInternal(err)
		}
		return nil
	})
}

func getPolicyResourceID(resourceID string) (int, error) {
	id, err := strconv.Atoi(resourceID)
	if err != nil {
		return 0, errors.Errorf("invalid policy resource ID %q", resourceID)
	}
	return id, nil
}

// hasAccessToUpdatePolicy checks if user can access to policy control feature.
// return nil if user has access.
func (s *Server) hasAccessToUpsertPolicy(pType api.PolicyType, policyUpsert *api.PolicyUpsert) error {
	// nil payload means user doesn't update the payload field
	if policyUpsert.Payload == nil {
		return nil
	}

	defaultPolicy, err := api.GetDefaultPolicy(pType)
	if err != nil {
		return err
	}

	if defaultPolicy == *policyUpsert.Payload {
		return nil
	}
	switch pType {
	case api.PolicyTypePipelineApproval:
		payload, err := api.UnmarshalPipelineApprovalPolicy(*policyUpsert.Payload)
		if err != nil {
			return err
		}
		if payload.Value == api.PipelineApprovalValueManualNever && !s.licenseService.IsFeatureEnabled(api.FeatureApprovalPolicy) {
			return errors.Errorf(api.FeatureApprovalPolicy.AccessErrorMessage())
		}
	case api.PolicyTypeBackupPlan:
		if !s.licenseService.IsFeatureEnabled(api.FeatureBackupPolicy) {
			return errors.Errorf(api.FeatureBackupPolicy.AccessErrorMessage())
		}
	case api.PolicyTypeSQLReview:
		return nil
	case api.PolicyTypeEnvironmentTier:
		if !s.licenseService.IsFeatureEnabled(api.FeatureEnvironmentTierPolicy) {
			return errors.Errorf(api.FeatureEnvironmentTierPolicy.AccessErrorMessage())
		}
	case api.PolicyTypeSensitiveData:
		if !s.licenseService.IsFeatureEnabled(api.FeatureSensitiveData) {
			return errors.Errorf(api.FeatureSensitiveData.AccessErrorMessage())
		}
	}
	return nil
}
