<template>
  <div
    class="bb-guide-dialog"
    :style="`top: ${bounding.top + bounding.height}px;left: ${bounding.left}px`"
  >
    <p class="bb-guide-title-text">{{ title }}</p>
    <p class="bb-guide-description-text">{{ description }}</p>

    <div class="bb-guide-btns-container">
      <button class="button">Back</button>
      <button class="button" @click="() => targetElement.click()">Next</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from "vue";
import { getElementBounding } from "./utils";
import * as storage from "./storage";
import merge from "lodash-es/merge";
const props = defineProps<{
  targetElement: HTMLElement;
  title: string;
  description: string;
}>();

const bounding = getElementBounding(props.targetElement);
onMounted(() => {
  const handler = () => {
    const { guide } = storage.get(["guide"]);
    storage.set({
      guide: merge(guide, {
        stepIndex: (guide?.stepIndex ?? 0) + 1,
      }),
    });
    storage.emitStorageChangedEvent();
  };
  props.targetElement.addEventListener("click", handler);

  onUnmounted(() => {
    props.targetElement.removeEventListener("click", handler);
  });
});
</script>

<style lang="postcss">
.bb-guide-dialog {
  @apply absolute bg-white w-72 mt-1 text-black p-2 px-3 rounded-lg text-left;
  box-shadow: 0 3px 30px rgb(33 33 33 / 30%);
  z-index: 10000000;
}

.bb-guide-title-text {
  @apply text-lg text-black leading-6 py-1 mb-1;
}

.bb-guide-description-text {
  @apply text-base text-gray-700;
}

.bb-guide-btns-container {
  @apply mt-3 w-full flex flex-row justify-between items-center;
}

.bb-guide-btns-container > button {
  @apply px-2 py-1 border text-sm rounded hover:opacity-80 hover:border-gray-400;
}
</style>
