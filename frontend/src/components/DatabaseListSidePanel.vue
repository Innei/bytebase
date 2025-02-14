<template>
  <BBOutline
    id="database"
    :title="$t('common.databases')"
    :item-list="mixedDatabaseList"
    :allow-collapse="false"
  />
</template>

<script lang="ts">
import { computed, defineComponent, watchEffect } from "vue";
import { cloneDeep, uniqBy } from "lodash-es";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import { Action, defineAction, useRegisterActions } from "@bytebase/vue-kbar";
import type { BBOutlineItem } from "@/bbkit/types";
import { Database, EnvironmentId, UNKNOWN_ID } from "@/types";
import { databaseSlug, environmentV1Name, projectSlug } from "@/utils";
import {
  useEnvironmentV1List,
  useCurrentUser,
  useDatabaseStore,
} from "@/store";

export default defineComponent({
  name: "DatabaseListSidePanel",
  setup() {
    const { t } = useI18n();
    const databaseStore = useDatabaseStore();
    const router = useRouter();

    const currentUser = useCurrentUser();

    const rawEnvironmentList = useEnvironmentV1List();

    // Reserve the environment list, put "Prod" to the top.
    const environmentList = computed(() =>
      cloneDeep(rawEnvironmentList.value).reverse()
    );

    const prepareList = () => {
      // It will also be called when user logout
      if (currentUser.value.id !== UNKNOWN_ID) {
        databaseStore.fetchDatabaseList();
      }
    };

    watchEffect(prepareList);

    // Use this to make the list reactive when project is transferred.
    const databaseList = computed((): Database[] => {
      return databaseStore
        .getDatabaseListByPrincipalId(currentUser.value.id)
        .filter((db) => db.syncStatus === "OK");
    });

    const databaseListByEnvironment = computed(() => {
      const envToDbMap: Map<EnvironmentId, BBOutlineItem[]> = new Map();
      for (const environment of environmentList.value) {
        envToDbMap.set(environment.uid, []);
      }
      const list = [...databaseList.value].filter(
        (db) => db.project.tenantMode !== "TENANT"
      );
      list.sort((a: any, b: any) => {
        return a.name.localeCompare(b.name);
      });
      for (const database of list) {
        const dbList = envToDbMap.get(
          String(database.instance.environment.id)
        )!;
        // dbList may be undefined if the environment is archived
        if (dbList) {
          dbList.push({
            id: `bb.database.${database.id}`,
            name: database.name,
            link: `/db/${databaseSlug(database)}`,
          });
        }
      }

      return environmentList.value
        .filter((environment) => {
          const items = envToDbMap.get(environment.uid) ?? [];
          return items.length > 0;
        })
        .map((environment): BBOutlineItem => {
          return {
            id: `bb.env.${environment.uid}`,
            name: environmentV1Name(environment),
            childList: envToDbMap.get(environment.uid),
            childCollapse: true,
          };
        });
    });

    const tenantDatabaseListByProject = computed((): BBOutlineItem[] => {
      const list = databaseList.value.filter(
        (db) => db.project.tenantMode === "TENANT"
      );
      // In case that each `db.project` is not reference equal
      // we run a uniq() on the list by project.id
      const projectList = uniqBy(
        list.map((db) => db.project),
        (project) => project.id
      );
      // Sort the list as what <ProjectListSidePanel /> does
      projectList.sort((a, b) => a.name.localeCompare(b.name));
      // Then group databaseList by project
      const databaseListGroupByProject = projectList.map((project) => {
        const databaseList = list.filter((db) => db.project.id === project.id);
        return {
          project,
          databaseList,
        };
      });
      // Map groups to `BBOutlineItem[]`
      const itemList = databaseListGroupByProject.map(
        ({ project, databaseList }) => {
          return {
            id: `bb.project.${project.id}.databases`,
            name: project.name,
            childList: databaseList.map<BBOutlineItem>((db) => ({
              id: `bb.project.${project.id}.database.${db.name}`,
              name: db.name,
              link: `/project/${projectSlug(project)}#databases`,
            })),
            childCollapse: true,
          } as BBOutlineItem;
        }
      );
      return itemList;
    });

    const mixedDatabaseList = computed(() => {
      return [
        ...databaseListByEnvironment.value,
        ...tenantDatabaseListByProject.value,
      ];
    });

    const kbarActions = computed((): Action[] => {
      const actions = mixedDatabaseList.value.flatMap((group: BBOutlineItem) =>
        group.childList!.map((item) =>
          defineAction({
            // `item.id` is namespaced already
            // so here `id` looks like
            // "bb.database.7001" for non-tenant databases
            // "bb.project.3007.database.db3" for tenant databases
            id: item.id,
            section: t("common.databases"),
            name: item.name,
            // `group.name` is also a keyword to provide better search
            // e.g. "blog" under "staged" now can be searched by "bl st"
            // also "blog" under "HR system" (a project) can be searched by "bl hr"
            keywords: `database db ${group.name}`,
            data: {
              tags: [group.name],
            },
            perform: () => {
              router.push(item.link!);
            },
          })
        )
      );
      return actions;
    });
    useRegisterActions(kbarActions);

    return {
      mixedDatabaseList,
    };
  },
});
</script>
