<template>
  <BBSelect
    :selected-item="state.selectedProject"
    :item-list="projectList"
    :disabled="disabled"
    :placeholder="$t('project.select')"
    :show-prefix-item="true"
    :error="!validate()"
    @select-item="(project) => $emit('select-project-id', project.id)"
  >
    <template #menuItem="{ item: project }">
      {{ projectName(project) }}
    </template>
  </BBSelect>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, reactive, watch } from "vue";
import {
  Project,
  UNKNOWN_ID,
  DEFAULT_PROJECT_ID,
  ProjectRoleType,
  unknown,
} from "../types";
import { useCurrentUser, useProjectStore } from "@/store";
import { isMemberOfProject } from "@/utils";

interface LocalState {
  selectedProject?: Project;
}

export enum Mode {
  Standard = 1,
  Tenant = 2,
}

export default defineComponent({
  name: "ProjectSelect",
  props: {
    selectedId: {
      default: UNKNOWN_ID,
      type: Number,
    },
    disabled: {
      default: false,
      type: Boolean,
    },
    allowedRoleList: {
      default: () => ["OWNER", "DEVELOPER"],
      type: Array as PropType<ProjectRoleType[]>,
    },
    includeDefaultProject: {
      default: false,
      type: Boolean,
    },
    mode: {
      type: Number as PropType<Mode>,
      default: Mode.Standard | Mode.Tenant,
    },
    onlyUserself: {
      type: Boolean,
      default: true,
    },
    required: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["select-project-id"],
  setup(props) {
    const state = reactive<LocalState>({
      selectedProject: undefined,
    });

    const currentUser = useCurrentUser();
    const projectStore = useProjectStore();

    const rawProjectList = computed((): Project[] => {
      let list = projectStore.projectList as Project[];

      if (props.onlyUserself) {
        list = list.filter((project) => {
          return isMemberOfProject(project, currentUser.value);
        });
      }

      list = list.filter((project) => {
        if (project.tenantMode === "DISABLED" && props.mode & Mode.Standard) {
          return true;
        }
        if (project.tenantMode === "TENANT" && props.mode & Mode.Tenant) {
          return true;
        }
        return false;
      });

      return list.filter((project: Project) => {
        // Do not show Default project in selector.
        return project.id != DEFAULT_PROJECT_ID;
      });
    });

    const selectedIdNotInList = computed((): boolean => {
      if (props.selectedId == UNKNOWN_ID) {
        return false;
      }

      return (
        rawProjectList.value.find((item) => {
          return item.id == props.selectedId;
        }) == null
      );
    });

    const projectList = computed((): Project[] => {
      const list = rawProjectList.value.filter((project) => {
        if (project.rowStatus === "NORMAL") {
          return true;
        }
        // project.rowStatus === "ARCHIVED"
        if (project.id === props.selectedId) {
          return true;
        }
        return false;
      });

      const defaultProject = projectStore.getProjectById(DEFAULT_PROJECT_ID);
      if (
        props.includeDefaultProject ||
        props.selectedId === DEFAULT_PROJECT_ID
      ) {
        // If includeDefaultProject is false but the selected project is the default
        // project, we will show it. If includeDefaultProject is true, then it's
        // already in the list, so no need to show it twice
        list.unshift(defaultProject);
      }

      if (
        props.selectedId !== DEFAULT_PROJECT_ID &&
        selectedIdNotInList.value
      ) {
        // It may happen the selected id might not be in the project list.
        // e.g. the selected project is deleted after the selection and we
        // are unable to cleanup properly. In such case, the selected project id
        // is orphaned and we just display the id
        const dummyProject = reactive(unknown("PROJECT"));
        dummyProject.id = props.selectedId;
        dummyProject.name = props.selectedId.toString();
        list.unshift(dummyProject);
      }

      return list;
    });

    const validate = () => {
      if (!props.required) {
        return true;
      }
      return !!state.selectedProject && state.selectedProject.id !== UNKNOWN_ID;
    };

    watch(
      [() => props.selectedId, projectList],
      ([selectedId, list]) => {
        state.selectedProject = list.find(
          (project) => project.id === selectedId
        );
      },
      { immediate: true }
    );

    return {
      UNKNOWN_ID,
      DEFAULT_PROJECT_ID,
      state,
      projectList,
      validate,
      selectedIdNotInList,
    };
  },
});
</script>
