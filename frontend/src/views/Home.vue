<template>
  <div class="flex flex-col">
    <div class="px-4 py-2 flex justify-between items-center">
      <EnvironmentTabFilter
        :include-all="true"
        :environment="selectedEnvironment?.uid ?? String(UNKNOWN_ID)"
        @update:environment="changeEnvironmentId"
      />
      <SearchBox
        :value="state.searchText"
        :placeholder="$t('issue.search-issue-name')"
        :autofocus="true"
        @update:value="changeSearchText($event)"
      />
    </div>

    <WaitingForMyApprovalIssueTable
      v-if="hasCustomApprovalFeature"
      session-key="home-waiting-approval"
      :issue-find="{
        statusList: ['OPEN'],
      }"
    >
      <template #table="{ issueList, loading }">
        <IssueTable
          :left-bordered="false"
          :right-bordered="false"
          :show-placeholder="!loading"
          :title="$t('issue.waiting-approval')"
          :issue-list="issueList.filter(keywordAndEnvironmentFilter)"
        />
      </template>
    </WaitingForMyApprovalIssueTable>

    <!-- show OPEN Assigned issues with pageSize=10 -->
    <PagedIssueTable
      session-key="home-assigned"
      :issue-find="{
        statusList: ['OPEN'],
        assigneeId: currentUser.id,
      }"
      :page-size="OPEN_ISSUE_LIST_PAGE_SIZE"
    >
      <template #table="{ issueList, loading }">
        <IssueTable
          class="-mt-px"
          :left-bordered="false"
          :right-bordered="false"
          :show-placeholder="!loading"
          :title="$t('issue.waiting-rollout')"
          :issue-list="issueList.filter(keywordAndEnvironmentFilter)"
        />
      </template>
    </PagedIssueTable>

    <!-- show OPEN Created issues with pageSize=10 -->
    <PagedIssueTable
      session-key="home-created"
      :issue-find="{
        statusList: ['OPEN'],
        creatorId: currentUser.id,
      }"
      :page-size="OPEN_ISSUE_LIST_PAGE_SIZE"
    >
      <template #table="{ issueList, loading }">
        <IssueTable
          class="-mt-px"
          :left-bordered="false"
          :right-bordered="false"
          :show-placeholder="!loading"
          :title="$t('common.created')"
          :issue-list="issueList.filter(keywordAndEnvironmentFilter)"
        />
      </template>
    </PagedIssueTable>

    <!-- show OPEN Subscribed issues with pageSize=10 -->
    <PagedIssueTable
      session-key="home-subscribed"
      :issue-find="{
        statusList: ['OPEN'],
        subscriberId: currentUser.id,
      }"
      :page-size="OPEN_ISSUE_LIST_PAGE_SIZE"
    >
      <template #table="{ issueList, loading }">
        <IssueTable
          class="-mt-px"
          :left-bordered="false"
          :right-bordered="false"
          :show-placeholder="!loading"
          :title="$t('common.subscribed')"
          :issue-list="issueList.filter(keywordAndEnvironmentFilter)"
        />
      </template>
    </PagedIssueTable>

    <!-- show the first 5 DONE or CANCELED issues -->
    <!-- But won't show "Load more", since we have a "View all closed" link below -->
    <PagedIssueTable
      session-key="home-closed"
      :issue-find="{
        statusList: ['DONE', 'CANCELED'],
      }"
      :page-size="MAX_CLOSED_ISSUE"
      :hide-load-more="true"
    >
      <template #table="{ issueList, loading }">
        <IssueTable
          class="-mt-px"
          :left-bordered="false"
          :right-bordered="false"
          :show-placeholder="!loading"
          :title="$t('project.overview.recently-closed')"
          :issue-list="issueList.filter(keywordAndEnvironmentFilter)"
        />
      </template>
    </PagedIssueTable>
  </div>
  <div class="w-full flex justify-end mt-2 px-4">
    <router-link to="/issue?status=closed" class="normal-link">
      {{ $t("project.overview.view-all-closed") }}
    </router-link>
  </div>

  <BBModal
    v-if="state.showTrialStartModal && subscriptionStore.subscription"
    :title="
      $t('subscription.trial-start-modal.title', {
        plan: $t(
          `subscription.plan.${planTypeToString(
            subscriptionStore.currentPlan
          ).toLowerCase()}.title`
        ),
      })
    "
    @close="onTrialingModalClose"
  >
    <div class="min-w-0 md:min-w-400 max-w-2xl">
      <div class="flex justify-center items-center">
        <img :src="planImage" class="w-56 px-4" />
        <div class="text-lg space-y-2">
          <p>
            <i18n-t keypath="subscription.trial-start-modal.content">
              <template #plan>
                <strong>
                  {{
                    $t(
                      `subscription.plan.${planTypeToString(
                        subscriptionStore.currentPlan
                      ).toLowerCase()}.title`
                    )
                  }}
                </strong>
              </template>
              <template #date>
                <strong>{{ subscriptionStore.expireAt }}</strong>
              </template>
            </i18n-t>
          </p>
          <p>
            <i18n-t keypath="subscription.trial-start-modal.subscription">
              <template #page>
                <router-link
                  to="/setting/subscription"
                  class="normal-link"
                  exact-active-class=""
                >
                  {{ $t("subscription.trial-start-modal.subscription-page") }}
                </router-link>
              </template>
            </i18n-t>
          </p>
        </div>
      </div>
      <div class="flex justify-end space-x-2 pb-4">
        <button
          type="button"
          class="btn-primary"
          @click.prevent="onTrialingModalClose"
        >
          {{ $t("subscription.trial-start-modal.button") }}
        </button>
      </div>
    </div>
  </BBModal>
</template>

<script lang="ts" setup>
import { reactive, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { activeEnvironment, isDatabaseRelatedIssueType } from "../utils";
import { UNKNOWN_ID, Issue, planTypeToString } from "../types";
import { EnvironmentTabFilter, SearchBox } from "@/components/v2";
import {
  useCurrentUser,
  useEnvironmentV1Store,
  useSubscriptionStore,
  useOnboardingStateStore,
  featureToRef,
} from "@/store";
import {
  IssueTable,
  PagedIssueTable,
  WaitingForMyApprovalIssueTable,
} from "@/components/Issue/table";

interface LocalState {
  searchText: string;
  showTrialStartModal: boolean;
}

const OPEN_ISSUE_LIST_PAGE_SIZE = 10;
const MAX_CLOSED_ISSUE = 5;

const environmentV1Store = useEnvironmentV1Store();
const subscriptionStore = useSubscriptionStore();
const onboardingStateStore = useOnboardingStateStore();
const router = useRouter();
const route = useRoute();

const state = reactive<LocalState>({
  searchText: "",
  showTrialStartModal: false,
});

const currentUser = useCurrentUser();
const hasCustomApprovalFeature = featureToRef("bb.feature.custom-approval");

const onTrialingModalClose = () => {
  state.showTrialStartModal = false;
  onboardingStateStore.consume("show-trialing-modal");
};

const planImage = computed(() => {
  return new URL(
    `../assets/plan-${planTypeToString(
      subscriptionStore.currentPlan
    ).toLowerCase()}.png`,
    import.meta.url
  ).href;
});

const selectedEnvironment = computed(() => {
  const { environment } = route.query;
  return environment
    ? environmentV1Store.getEnvironmentByUID(
        parseInt(environment as string, 10)
      )
    : undefined;
});

const keywordAndEnvironmentFilter = (issue: Issue) => {
  if (
    selectedEnvironment.value &&
    selectedEnvironment.value.uid !== String(UNKNOWN_ID)
  ) {
    if (!isDatabaseRelatedIssueType(issue.type)) {
      return false;
    }
    if (
      String(activeEnvironment(issue.pipeline).id) !==
      selectedEnvironment.value.uid
    ) {
      return false;
    }
  }
  const keyword = state.searchText.trim();
  if (keyword) {
    if (!issue.name.toLowerCase().includes(keyword)) {
      return false;
    }
  }
  return true;
};

const changeEnvironmentId = (environment: string | undefined) => {
  if (environment && environment !== String(UNKNOWN_ID)) {
    router.replace({
      name: "workspace.home",
      query: {
        ...route.query,
        environment,
      },
    });
  } else {
    router.replace({
      name: "workspace.home",
      query: {
        ...route.query,
        environment: undefined,
      },
    });
  }
};

const changeSearchText = (searchText: string) => {
  state.searchText = searchText;
};
</script>
