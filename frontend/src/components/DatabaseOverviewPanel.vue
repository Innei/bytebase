<template>
  <div class="space-y-6 divide-y divide-block-border">
    <div v-if="anomalySectionList.length > 0">
      <div class="text-lg leading-6 font-medium text-main mb-4 flex flex-row">
        {{ $t("common.anomalies") }}
        <span class="ml-2 textinfolabel items-center flex">
          {{
            $t(
              "database.the-list-might-be-out-of-date-and-is-refreshed-roughly-every-10-minutes"
            )
          }}
        </span>
      </div>
      <AnomalyTable :anomaly-section-list="anomalySectionList" />
    </div>
    <div
      v-else
      class="text-lg leading-6 font-medium text-main mb-4 flex flex-row"
    >
      {{ $t("database.no-anomalies-detected") }}
      <heroicons-outline:check-circle class="ml-1 w-6 h-6 text-success" />
    </div>

    <!-- Description list -->
    <dl
      class="grid grid-cols-1 gap-x-4 gap-y-4 sm:grid-cols-2 pt-4"
      data-label="bb-database-overview-description-list"
    >
      <template
        v-if="
          database.instance.engine !== 'CLICKHOUSE' &&
          database.instance.engine !== 'SNOWFLAKE' &&
          database.instance.engine !== 'MONGODB'
        "
      >
        <div class="col-span-1 col-start-1">
          <dt class="text-sm font-medium text-control-light">
            {{
              database.instance.engine == "POSTGRES"
                ? $t("db.encoding")
                : $t("db.character-set")
            }}
          </dt>
          <dd class="mt-1 text-sm text-main">
            {{ databaseSchemaMetadata.characterSet }}
          </dd>
        </div>

        <div class="col-span-1">
          <dt class="text-sm font-medium text-control-light">
            {{ $t("db.collation") }}
          </dt>
          <dd class="mt-1 text-sm text-main">
            {{ databaseSchemaMetadata.collation }}
          </dd>
        </div>
      </template>

      <div class="col-span-1 col-start-1">
        <dt class="text-sm font-medium text-control-light">
          {{ $t("database.sync-status") }}
        </dt>
        <dd class="mt-1 text-sm text-main">
          <span>{{ database.syncStatus }}</span>
        </dd>
      </div>

      <div class="col-span-1">
        <dt class="text-sm font-medium text-control-light">
          {{ $t("database.last-successful-sync") }}
        </dt>
        <dd class="mt-1 text-sm text-main">
          {{ humanizeTs(database.lastSuccessfulSyncTs) }}
        </dd>
      </div>
    </dl>

    <div class="pt-6">
      <div
        v-if="hasSchemaProperty"
        class="flex flex-row justify-start items-center mb-4"
      >
        <span class="text-lg leading-6 font-medium text-main mr-2">Schema</span>
        <BBSelect
          class="!w-auto min-w-[12rem]"
          :selected-item="state.selectedSchemaName"
          :item-list="schemaNameList"
          :placeholder="$t('database.schema.select')"
          :show-prefix-item="true"
          @select-item="(schema: string) => state.selectedSchemaName = schema"
        >
          <template #menuItem="{ item: schema }">
            {{ schema }}
          </template>
        </BBSelect>
      </div>

      <template v-if="databaseEngine !== 'REDIS'">
        <div class="text-lg leading-6 font-medium text-main mb-4">
          <span v-if="databaseEngine === 'MONGODB'">{{
            $t("db.collections")
          }}</span>
          <span v-else>{{ $t("db.tables") }}</span>
        </div>

        <TableTable
          :database="database"
          :schema-name="state.selectedSchemaName"
          :table-list="tableList"
        />

        <div class="mt-6 text-lg leading-6 font-medium text-main mb-4">
          {{ $t("db.views") }}
        </div>
        <ViewTable
          :database="database"
          :schema-name="state.selectedSchemaName"
          :view-list="viewList"
        />

        <template v-if="database.instance.engine == 'POSTGRES'">
          <div class="mt-6 text-lg leading-6 font-medium text-main mb-4">
            {{ $t("db.extensions") }}
          </div>
          <DBExtensionTable :db-extension-list="dbExtensionList" />
        </template>

        <template v-if="database.instance.engine === 'POSTGRES'">
          <div class="mt-6 text-lg leading-6 font-medium text-main mb-4">
            {{ $t("db.functions") }}
          </div>
          <FunctionTable
            :database="database"
            :schema-name="state.selectedSchemaName"
            :function-list="functionList"
          />
        </template>
      </template>
    </div>

    <!-- Hide data source list for now, as we don't allow adding new data source after creating the database. -->
    <div v-if="false" class="pt-6">
      <DataSourceTable :instance="database.instance" :database="database" />
    </div>

    <template v-if="allowViewDataSource">
      <template
        v-for="(item, index) of [
          { type: 'RW', list: readWriteDataSourceList },
          { type: 'RO', list: readonlyDataSourceList },
        ]"
        :key="index"
      >
        <div v-if="item.list.length" class="pt-6">
          <div
            v-if="hasDataSourceFeature"
            class="text-lg leading-6 font-medium text-main mb-4"
          >
            <span v-data-source-type>{{ item.type }}</span>
          </div>
          <div class="space-y-4">
            <div v-for="(ds, dsIndex) of item.list" :key="dsIndex">
              <div v-if="hasDataSourceFeature" class="relative mb-2">
                <div
                  class="absolute inset-0 flex items-center"
                  aria-hidden="true"
                >
                  <div class="w-full border-t border-gray-300"></div>
                </div>
                <div class="relative flex justify-start">
                  <router-link
                    :to="`/db/${databaseSlug}/data-source/${dataSourceSlug(
                      ds
                    )}`"
                    class="pr-3 bg-white font-medium normal-link"
                    >{{ ds.name }}</router-link
                  >
                </div>
              </div>
              <div
                v-if="allowConfigInstance"
                class="flex justify-end space-x-3"
              >
                <template v-if="isEditingDataSource(ds)">
                  <button
                    type="button"
                    class="btn-normal"
                    @click.prevent="cancelEditDataSource"
                  >
                    {{ $t("common.cancel") }}
                  </button>
                  <button
                    type="button"
                    class="btn-normal"
                    :disabled="!allowSaveDataSource"
                    @click.prevent="saveEditDataSource"
                  >
                    <!-- Heroicon name: solid/save -->
                    <heroicons-solid:save
                      class="-ml-1 mr-2 h-5 w-5 text-control-light"
                    />
                    <span>{{ $t("common.save") }}</span>
                  </button>
                </template>
                <template v-else>
                  <button
                    type="button"
                    class="btn-normal"
                    @click.prevent="editDataSource(ds)"
                  >
                    <!-- Heroicon name: solid/pencil -->
                    <heroicons-solid:pencil
                      class="-ml-1 mr-2 h-5 w-5 text-control-light"
                    />
                    <span>{{ $t("common.edit") }}</span>
                  </button>
                </template>
              </div>
              <DataSourceConnectionPanel
                :editing="isEditingDataSource(ds)"
                :data-source="
                  isEditingDataSource(ds) ? state.editingDataSource! : ds
                "
              />
            </div>
          </div>
        </div>
      </template>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { cloneDeep, head, isEqual } from "lodash-es";
import { computed, reactive, watchEffect, PropType } from "vue";
import { hasWorkspacePermission, memberListInProject } from "../utils";
import {
  Anomaly,
  Database,
  DataSource,
  DataSourcePatch,
  EngineType,
} from "../types";
import {
  featureToRef,
  useCurrentUser,
  useDataSourceStore,
  useAnomalyList,
  useDBSchemaStore,
} from "@/store";
import { BBTableSectionDataSource } from "../bbkit/types";
import AnomalyTable from "../components/AnomalyTable.vue";
import DataSourceTable from "../components/DataSourceTable.vue";
import DataSourceConnectionPanel from "../components/DataSourceConnectionPanel.vue";
import TableTable from "../components/TableTable.vue";
import ViewTable from "../components/ViewTable.vue";
import FunctionTable from "../components/FunctionTable.vue";

interface LocalState {
  selectedSchemaName: string;
  editingDataSource?: DataSource;
}

const props = defineProps({
  database: {
    required: true,
    type: Object as PropType<Database>,
  },
});

const dataSourceStore = useDataSourceStore();

const state = reactive<LocalState>({
  selectedSchemaName: "",
});

const currentUser = useCurrentUser();
const dbSchemaStore = useDBSchemaStore();

const databaseEngine = computed(() => {
  return props.database.instance.engine as EngineType;
});

const hasSchemaProperty = computed(() => {
  return (
    databaseEngine.value === "POSTGRES" ||
    databaseEngine.value === "SNOWFLAKE" ||
    databaseEngine.value === "ORACLE" ||
    databaseEngine.value === "MSSQL" ||
    databaseEngine.value === "REDSHIFT"
  );
});

const prepareDatabaseMetadata = async () => {
  await dbSchemaStore.getOrFetchDatabaseMetadataById(props.database.id);
  if (hasSchemaProperty.value && schemaList.value.length > 0) {
    state.selectedSchemaName = head(schemaList.value)?.name || "";
  }
};

watchEffect(prepareDatabaseMetadata);

const anomalyList = useAnomalyList(
  computed(() => ({ databaseId: props.database.id }))
);

const anomalySectionList = computed((): BBTableSectionDataSource<Anomaly>[] => {
  const list: BBTableSectionDataSource<Anomaly>[] = [];
  if (anomalyList.value.length > 0) {
    list.push({
      title: props.database.name,
      list: anomalyList.value,
    });
  }
  return list;
});

const hasDataSourceFeature = featureToRef("bb.feature.data-source");

const schemaList = computed(() => {
  return dbSchemaStore.getSchemaListByDatabaseId(props.database.id);
});

const schemaNameList = computed(() => {
  return schemaList.value.map((schema) => schema.name);
});

const databaseSchemaMetadata = computed(() => {
  return dbSchemaStore.getDatabaseMetadataByDatabaseId(props.database.id);
});

const tableList = computed(() => {
  if (hasSchemaProperty.value) {
    return (
      schemaList.value.find(
        (schema) => schema.name === state.selectedSchemaName
      )?.tables || []
    );
  }
  return dbSchemaStore.getTableListByDatabaseId(props.database.id);
});

const viewList = computed(() => {
  if (hasSchemaProperty.value) {
    return (
      schemaList.value.find(
        (schema) => schema.name === state.selectedSchemaName
      )?.views || []
    );
  }
  return dbSchemaStore.getViewListByDatabaseId(props.database.id);
});

const dbExtensionList = computed(() => {
  return dbSchemaStore.getExtensionListByDatabaseId(props.database.id);
});

const functionList = computed(() => {
  if (hasSchemaProperty.value) {
    return (
      schemaList.value.find(
        (schema) => schema.name === state.selectedSchemaName
      )?.functions || []
    );
  }
  return dbSchemaStore.getFunctionListByDatabaseId(props.database.id);
});

const allowConfigInstance = computed(() => {
  return hasWorkspacePermission(
    "bb.permission.workspace.manage-instance",
    currentUser.value.role
  );
});

const allowViewDataSource = computed(() => {
  if (allowConfigInstance.value) {
    return true;
  }

  return (
    memberListInProject(
      props.database.project,
      currentUser.value,
      /* empty array to "ALL" */ []
    ).length > 0
  );
});

const dataSourceList = computed(() => {
  return props.database.dataSourceList;
});

const readWriteDataSourceList = computed(() => {
  return dataSourceList.value.filter((dataSource: DataSource) => {
    return dataSource.type == "RW";
  });
});

const readonlyDataSourceList = computed(() => {
  return dataSourceList.value.filter((dataSource: DataSource) => {
    return dataSource.type == "RO";
  });
});

const isEditingDataSource = (dataSource: DataSource) => {
  return state.editingDataSource && state.editingDataSource.id == dataSource.id;
};

const allowSaveDataSource = computed(() => {
  for (const dataSource of dataSourceList.value) {
    if (dataSource.id == state.editingDataSource!.id) {
      return !isEqual(dataSource, state.editingDataSource);
    }
  }
  return false;
});

const editDataSource = (dataSource: DataSource) => {
  state.editingDataSource = cloneDeep(dataSource);
};

const cancelEditDataSource = () => {
  state.editingDataSource = undefined;
};

const saveEditDataSource = () => {
  const dataSourcePatch = {
    username: state.editingDataSource?.username,
    password: state.editingDataSource?.password,
  } as DataSourcePatch;
  dataSourceStore
    .patchDataSource({
      databaseId: state.editingDataSource?.databaseId as number,
      dataSourceId: state.editingDataSource?.id as number,
      dataSource: dataSourcePatch,
    })
    .then(() => {
      state.editingDataSource = undefined;
    });
};
</script>
