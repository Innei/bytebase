import { defineStore } from "pinia";
import axios from "axios";
import { computed, onBeforeMount, unref, watch } from "vue";
import {
  DataSource,
  empty,
  EMPTY_ID,
  Environment,
  EnvironmentId,
  Instance,
  InstanceCreate,
  InstanceId,
  InstanceMigration,
  InstancePatch,
  InstanceState,
  InstanceUserId,
  INSTANCE_OPERATION_TIMEOUT,
  MaybeRef,
  MigrationHistory,
  MigrationHistoryId,
  ResourceIdentifier,
  ResourceObject,
  RowStatus,
  SQLResultSet,
  unknown,
  UNKNOWN_ID,
} from "@/types";
import { InstanceUser } from "@/types/InstanceUser";
import { useLegacyEnvironmentStore } from "./environment";
import { useDataSourceStore } from "./dataSource";
import { useSQLStore } from "./sql";

function convert(
  instance: ResourceObject,
  includedList: ResourceObject[]
): Instance {
  const environmentId = (
    instance.relationships!.environment.data as ResourceIdentifier
  ).id;
  let environment: Environment = unknown("ENVIRONMENT") as Environment;
  environment.id = parseInt(environmentId);

  const dataSourceIdList = instance.relationships!.dataSourceList
    .data as ResourceIdentifier[];
  const dataSourceList: DataSource[] = [];
  for (const item of dataSourceIdList) {
    const dataSource = unknown("DATA_SOURCE") as DataSource;
    dataSource.id = parseInt(item.id);
    dataSourceList.push(dataSource);
  }

  const instancePartial = {
    ...(instance.attributes as Omit<
      Instance,
      "id" | "environment" | "dataSourceList"
    >),
    id: parseInt(instance.id),
    environment,
    dataSourceList: [],
  };

  const legacyEnvironmentStore = useLegacyEnvironmentStore();
  const dataSourceStore = useDataSourceStore();
  for (const item of includedList || []) {
    if (
      item.type == "environment" &&
      (instance.relationships!.environment.data as ResourceIdentifier).id ==
        item.id
    ) {
      environment = legacyEnvironmentStore.convert(item, includedList);
    }

    if (
      item.type == "dataSource" &&
      item.attributes.instanceId == instancePartial.id
    ) {
      const i = dataSourceList.findIndex(
        (dataSource: DataSource) => parseInt(item.id) == dataSource.id
      );
      if (i != -1) {
        dataSourceList[i] = dataSourceStore.convert(item);
      }
    }
  }

  return {
    ...(instancePartial as Omit<Instance, "environment" | "dataSourceList">),
    environment,
    dataSourceList,
  };
}

function convertInstanceUser(instanceUser: ResourceObject): InstanceUser {
  return {
    ...(instanceUser.attributes as Omit<InstanceUser, "id">),
    id: instanceUser.id,
  };
}

function convertMigrationHistory(history: ResourceObject): MigrationHistory {
  const payload = history.attributes.payload
    ? JSON.parse((history.attributes.payload as string) || "{}")
    : {};
  return {
    ...(history.attributes as Omit<
      MigrationHistory,
      "id" | "issueId" | "payload"
    >),
    id: history.id,
    // This issueId is special since it's stored in the migration history table
    // and may refer to the issueId from the external system in the future.
    issueId: parseInt(history.attributes.issueId as string),
    payload,
  };
}

export const useInstanceStore = defineStore("instance", {
  state: (): InstanceState => ({
    instanceById: new Map(),
    instanceUserListById: new Map(),
    migrationHistoryById: new Map(),
    migrationHistoryListByIdAndDatabaseName: new Map(),
  }),
  actions: {
    convert(
      instance: ResourceObject,
      includedList: ResourceObject[]
    ): Instance {
      return convert(instance, includedList);
    },
    getInstanceList(rowStatusList?: RowStatus[]): Instance[] {
      const list = [];
      for (const [_, instance] of this.instanceById) {
        if (
          (!rowStatusList && instance.rowStatus == "NORMAL") ||
          (rowStatusList && rowStatusList.includes(instance.rowStatus))
        ) {
          list.push(instance);
        }
      }
      return list.sort((a: Instance, b: Instance) => {
        return b.id - a.id;
      });
    },
    getInstanceListByEnvironmentId(
      environmentId: EnvironmentId,
      rowStatusList?: RowStatus[]
    ): Instance[] {
      const list = this.getInstanceList(rowStatusList);
      return list.filter((item: Instance) => {
        return item.environment.id == environmentId;
      });
    },
    getInstanceById(instanceId: InstanceId): Instance {
      if (instanceId == EMPTY_ID) {
        return empty("INSTANCE") as Instance;
      }

      return (
        this.instanceById.get(instanceId) || (unknown("INSTANCE") as Instance)
      );
    },
    getInstanceUserListById(instanceId: InstanceId): InstanceUser[] {
      return this.instanceUserListById.get(instanceId) || [];
    },
    formatEngine(instance: Instance): string {
      switch (instance.engine) {
        case "POSTGRES":
          return "PostgreSQL";
        // Use MySQL as default engine.
        default:
          return "MySQL";
      }
    },
    getMigrationHistoryById(
      migrationHistoryId: MigrationHistoryId
    ): MigrationHistory | undefined {
      return this.migrationHistoryById.get(migrationHistoryId);
    },
    getMigrationHistoryListByInstanceIdAndDatabaseName(
      instanceId: InstanceId,
      databaseName: string
    ): MigrationHistory[] {
      return (
        this.migrationHistoryListByIdAndDatabaseName.get(
          [instanceId, databaseName].join(".")
        ) || []
      );
    },
    setInstanceList(instanceList: Instance[]) {
      instanceList.forEach((instance) => {
        this.instanceById.set(instance.id, instance);
      });
    },
    setInstanceById({
      instanceId,
      instance,
    }: {
      instanceId: InstanceId;
      instance: Instance;
    }) {
      this.instanceById.set(instanceId, instance);
    },
    setInstanceUserListById({
      instanceId,
      instanceUserList,
    }: {
      instanceId: InstanceId;
      instanceUserList: InstanceUser[];
    }) {
      this.instanceUserListById.set(instanceId, instanceUserList);
    },
    setMigrationHistoryById({
      migrationHistoryId,
      migrationHistory,
    }: {
      migrationHistoryId: MigrationHistoryId;
      migrationHistory: MigrationHistory;
    }) {
      this.migrationHistoryById.set(migrationHistoryId, migrationHistory);
    },
    setMigrationHistoryListByInstanceIdAndDatabaseName({
      instanceId,
      databaseName,
      historyList,
    }: {
      instanceId: InstanceId;
      databaseName: string;
      historyList: MigrationHistory[];
    }) {
      this.migrationHistoryListByIdAndDatabaseName.set(
        [instanceId, databaseName].join("."),
        historyList
      );
    },
    async fetchInstanceList(rowStatusList?: RowStatus[]) {
      const path =
        "/api/instance" +
        (rowStatusList ? "?rowstatus=" + rowStatusList.join(",") : "");
      const data = (await axios.get(path)).data;
      const instanceList: Instance[] = data.data.map(
        (instance: ResourceObject) => {
          return convert(instance, data.included);
        }
      );

      this.setInstanceList(instanceList);

      return instanceList;
    },
    async fetchInstanceById(instanceId: InstanceId) {
      const data = (await axios.get(`/api/instance/${instanceId}`)).data;
      const instance = convert(data.data, data.included);

      this.setInstanceById({
        instanceId,
        instance,
      });
      return instance;
    },
    async getOrFetchInstanceById(instanceId: InstanceId) {
      const storedInstance = this.getInstanceById(instanceId);
      if (storedInstance.id !== UNKNOWN_ID) {
        return storedInstance;
      }
      return this.fetchInstanceById(instanceId);
    },
    async createInstance(newInstance: InstanceCreate) {
      const data = (
        await axios.post(
          `/api/instance`,
          {
            data: {
              type: "InstanceCreate",
              attributes: newInstance,
            },
          },
          {
            timeout: INSTANCE_OPERATION_TIMEOUT,
          }
        )
      ).data;
      const createdInstance = convert(data.data, data.included);

      this.setInstanceById({
        instanceId: createdInstance.id,
        instance: createdInstance,
      });

      return createdInstance;
    },
    async patchInstance({
      instanceId,
      instancePatch,
    }: {
      instanceId: InstanceId;
      instancePatch: InstancePatch;
    }) {
      const data = (
        await axios.patch(
          `/api/instance/${instanceId}`,
          {
            data: {
              type: "instancePatch",
              attributes: instancePatch,
            },
          },
          {
            timeout: INSTANCE_OPERATION_TIMEOUT,
          }
        )
      ).data;
      const updatedInstance = convert(data.data, data.included);

      this.setInstanceById({
        instanceId: updatedInstance.id,
        instance: updatedInstance,
      });

      return updatedInstance;
    },
    async deleteInstanceById(instanceId: InstanceId) {
      await axios.delete(`/api/instance/${instanceId}`);
      this.instanceById.delete(instanceId);
    },
    async fetchInstanceUser(instanceId: InstanceId, userId: InstanceUserId) {
      const data = (
        await axios.get(`/api/instance/${instanceId}/user/${userId}`)
      ).data;
      return convertInstanceUser(data.data);
    },
    async fetchInstanceUserListById(instanceId: InstanceId) {
      const data = (await axios.get(`/api/instance/${instanceId}/user`)).data;
      const instanceUserList = data.data.map((instanceUser: ResourceObject) => {
        return convertInstanceUser(instanceUser);
      });

      this.setInstanceUserListById({
        instanceId,
        instanceUserList,
      });
      return instanceUserList;
    },
    async checkMigrationSetup(
      instanceId: InstanceId
    ): Promise<InstanceMigration> {
      const data = (
        await axios.get(`/api/instance/${instanceId}/migration/status`, {
          timeout: INSTANCE_OPERATION_TIMEOUT,
        })
      ).data.data;

      return {
        status: data.attributes.status,
        error: data.attributes.error,
      };
    },
    async createMigrationSetup(instanceId: InstanceId): Promise<SQLResultSet> {
      const res = (
        await axios.post(`/api/instance/${instanceId}/migration`, undefined, {
          timeout: INSTANCE_OPERATION_TIMEOUT,
        })
      ).data;

      return useSQLStore().convert(res.data) as SQLResultSet;
    },
    async fetchMigrationHistoryById({
      instanceId,
      migrationHistoryId,
      sdl,
    }: {
      instanceId: InstanceId;
      migrationHistoryId: MigrationHistoryId;
      sdl?: boolean;
    }) {
      let url = `/api/instance/${instanceId}/migration/history/${migrationHistoryId}`;
      if (sdl) {
        url += "?sdl=true";
      }
      const data = (
        await axios.get(url, {
          timeout: INSTANCE_OPERATION_TIMEOUT,
        })
      ).data;
      const migrationHistory = convertMigrationHistory(data.data);

      this.setMigrationHistoryById({
        migrationHistoryId,
        migrationHistory,
      });
      return migrationHistory;
    },
    async fetchMigrationHistoryByVersion({
      instanceId,
      databaseName,
      version,
    }: {
      instanceId: InstanceId;
      databaseName: string;
      version: string;
    }) {
      const data = (
        await axios.get(
          `/api/instance/${instanceId}/migration/history?database=${databaseName}&version=${version}`,
          {
            timeout: INSTANCE_OPERATION_TIMEOUT,
          }
        )
      ).data.data;
      const historyList: MigrationHistory[] = data.map(
        (history: ResourceObject) => {
          return convertMigrationHistory(history);
        }
      );

      if (historyList.length > 0) {
        this.setMigrationHistoryById({
          migrationHistoryId: historyList[0].id,
          migrationHistory: historyList[0],
        });
        return historyList[0];
      }
      throw new Error(
        `Migration version ${version} not found in database ${databaseName}.`
      );
    },
    async fetchMigrationHistory({
      instanceId,
      databaseName,
      limit,
    }: {
      instanceId: InstanceId;
      databaseName: string;
      limit?: number;
    }): Promise<MigrationHistory[]> {
      let url = `/api/instance/${instanceId}/migration/history?database=${databaseName}`;
      if (limit) {
        url += `&limit=${limit}`;
      }
      const data = (
        await axios.get(url, {
          timeout: INSTANCE_OPERATION_TIMEOUT,
        })
      ).data.data;
      const historyList: MigrationHistory[] = data.map(
        (history: ResourceObject) => {
          return convertMigrationHistory(history);
        }
      );

      this.setMigrationHistoryListByInstanceIdAndDatabaseName({
        instanceId,
        databaseName,
        historyList,
      });

      return historyList;
    },
    async getSamplePostgresInstance() {
      const data = (
        await axios.get<{
          host: string;
          port: number;
          username: string;
        }>("/api/instance/sample-pg")
      ).data;

      return data;
    },
  },
});

export const useInstanceList = (rowStatusList?: RowStatus[]) => {
  const store = useInstanceStore();
  // SQL Editor will visit instanceList very early.
  // Using `watchEffect` here might get a data race here, which leads a vue's
  // internal error.
  // So we fetch data when "before mount" - trying to be early but not too early.
  onBeforeMount(() => store.fetchInstanceList(rowStatusList));
  return computed(() => store.getInstanceList(rowStatusList));
};

export const useInstanceById = (instanceId: MaybeRef<InstanceId>) => {
  const store = useInstanceStore();
  watch(
    () => unref(instanceId),
    (id) => {
      if (id !== UNKNOWN_ID) {
        if (store.getInstanceById(id).id === UNKNOWN_ID) {
          store.fetchInstanceById(id);
        }
      }
    },
    { immediate: true }
  );

  return computed(() => store.getInstanceById(unref(instanceId)));
};
