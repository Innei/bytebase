<template>
  <div class="flex flex-col">
    <div class="py-2 flex justify-between items-center">
      <div class="flex justify-start items-center gap-x-2">
        <BBAttention
          :style="'INFO'"
          :title="$t('setting.project.description')"
        />
      </div>

      <div class="flex justify-end items-center gap-x-2">
        <SearchBox
          :value="state.searchText"
          :placeholder="$t('project.dashboard.search-bar-placeholder')"
          :autofocus="true"
          @update:value="changeSearchText($event)"
        />
      </div>
    </div>

    <div class="py-2">
      <NCheckbox v-model:checked="state.includesArchived">
        <span class="textinfolabel">
          {{ $t("setting.project.show-archived") }}
        </span>
      </NCheckbox>
    </div>

    <ProjectV1Table :project-list="filteredProjectList" />

    <BBModal
      v-if="state.showCreateModal"
      class="relative overflow-hidden"
      :title="$t('quick-action.create-project')"
      @close="state.showCreateModal = false"
    >
      <ProjectCreate @dismiss="state.showCreateModal = false" />
    </BBModal>
  </div>
</template>

<script setup lang="ts">
import { reactive, computed } from "vue";
import { NCheckbox } from "naive-ui";

import { DEFAULT_PROJECT_V1_NAME } from "../types";
import { useProjectV1List } from "@/store";
import ProjectCreate from "../components/ProjectCreate.vue";
import { ProjectV1Table, SearchBox } from "@/components/v2";
import { State } from "@/types/proto/v1/common";
import { filterProjectV1ListByKeyword } from "@/utils";

interface LocalState {
  searchText: string;
  showCreateModal: boolean;
  includesArchived: boolean;
}

const { projectList } = useProjectV1List(true /* showDeleted */);

const state = reactive<LocalState>({
  searchText: "",
  showCreateModal: false,
  includesArchived: false,
});

const changeSearchText = (searchText: string) => {
  state.searchText = searchText;
};

const filteredProjectList = computed(() => {
  let list = [...projectList.value];
  list = filterProjectV1ListByKeyword(list, state.searchText);

  // Put "Unassigned" to the first;
  const unassignedIndex = list.findIndex(
    (project) => project.name === DEFAULT_PROJECT_V1_NAME
  );
  if (unassignedIndex >= 0) {
    const unassignedProject = list[unassignedIndex];
    list.splice(unassignedIndex, 1);
    list.unshift(unassignedProject);
  }
  if (!state.includesArchived) {
    list = list.filter((project) => project.state !== State.DELETED);
  }
  return list;
});
</script>
