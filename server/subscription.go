package server

import (
	"net/http"
	"time"

	"github.com/bytebase/bytebase/api"
	"github.com/bytebase/bytebase/common"
	enterpriseAPI "github.com/bytebase/bytebase/enterprise/api"
	"github.com/google/jsonapi"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (s *Server) registerSubscriptionRoutes(g *echo.Group) {
	g.GET("/subscription", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		if err := jsonapi.MarshalPayload(c.Response().Writer, s.subscription); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal subscription response").SetInternal(err)
		}
		return nil
	})

	g.PATCH("/subscription", func(c echo.Context) error {
		patch := &enterpriseAPI.SubscriptionPatch{}
		if err := jsonapi.UnmarshalPayload(c.Request().Body, patch); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Malformatted create subscription request").SetInternal(err)
		}

		if err := s.LicenseService.StoreLicense(patch.License); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create license").SetInternal(err)
		}

		s.subscription = s.loadSubscription()

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		if err := jsonapi.MarshalPayload(c.Response().Writer, s.subscription); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal subscription response").SetInternal(err)
		}
		return nil
	})
}

// loadLicense will load current subscription by license.
// Return subscription with free plan if no license found.
func (server *Server) loadSubscription() *enterpriseAPI.Subscription {
	subscription := &enterpriseAPI.Subscription{
		Plan: api.FREE,
		// -1 means not expire, just for free plan
		ExpiresTs:     -1,
		InstanceCount: 5,
	}
	license, _ := server.loadLicense()
	if license != nil {
		subscription = &enterpriseAPI.Subscription{
			Plan:          license.Plan,
			ExpiresTs:     license.ExpiresTs,
			InstanceCount: license.InstanceCount,
		}
	}

	return subscription
}

// loadLicense will get and parse valid license from file.
func (server *Server) loadLicense() (*enterpriseAPI.License, error) {
	license, err := server.LicenseService.LoadLicense()
	if err != nil {
		if common.ErrorCode(err) == common.NotFound {
			server.l.Debug("Failed to find license", zap.String("error", err.Error()))
		} else {
			server.l.Warn("Failed to load valid license", zap.String("error", err.Error()))
		}
		return nil, err
	}

	server.l.Info(
		"Load valid license",
		zap.String("plan", license.Plan.String()),
		zap.Time("expiresAt", time.Unix(license.ExpiresTs, 0)),
		zap.Int("instanceCount", license.InstanceCount),
	)

	return license, nil
}

func (s *Server) feature(feature api.FeatureType) bool {
	return api.FeatureMatrix[feature][s.subscription.Plan]
}
