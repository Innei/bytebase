import { getElementBounding } from "./utils";

import { App, createApp } from "vue";
import dialogVue from "./dialog.vue";

const getTargetElementBySelectors = (selectors: string[][]) => {
  let targetElement = document.body;
  for (const selector of selectors) {
    try {
      targetElement = document.body.querySelector(
        selector.join(" ")
      ) as HTMLElement;
    } catch (error) {
      // do nth
    }

    if (targetElement) {
      break;
    }
  }
  return targetElement;
};

let currentDialogInstance: App<Element>;

export const showGuideDialog = async (guideStep: any) => {
  removeGuideDialog();
  const targetElement = await waitForTargetElement(guideStep.selectors);
  if (targetElement) {
    renderHighlightWrapper(targetElement);
    currentDialogInstance = renderGuideDialog(
      targetElement,
      guideStep.title,
      guideStep.description
    );
  }
};

const renderHighlightWrapper = (targetElement: HTMLElement) => {
  const highlightWrapper = document.createElement("div");
  highlightWrapper.className = "bb-guide-highlight-wrapper";
  document.body.appendChild(highlightWrapper);
  const bounding = getElementBounding(targetElement);
  highlightWrapper.style.top = `${bounding.top}px`;
  highlightWrapper.style.left = `${bounding.left}px`;
  highlightWrapper.style.width = `${bounding.width}px`;
  highlightWrapper.style.height = `${bounding.height}px`;
};

const renderGuideDialog = (
  targetElement: HTMLElement,
  title: string,
  description: string
) => {
  const $div = document.createElement("div");
  document.body.appendChild($div);
  const $dialog = createApp(dialogVue, {
    title,
    description,
    targetElement,
  });
  $dialog.mount($div);
  return $dialog;
};

let mutationObserver: MutationObserver | null = null;
const waitForTargetElement = (selectors: string[][]): Promise<HTMLElement> => {
  return new Promise((resolve) => {
    let targetElement = getTargetElementBySelectors(selectors);
    if (targetElement) {
      return resolve(targetElement);
    }
    mutationObserver?.disconnect();

    mutationObserver = new MutationObserver(() => {
      targetElement = getTargetElementBySelectors(selectors);
      if (targetElement) {
        mutationObserver?.disconnect();
        return resolve(targetElement);
      }
    });

    mutationObserver.observe(document.body, {
      childList: true,
      subtree: true,
    });
  });
};

export const removeGuideDialog = () => {
  currentDialogInstance?.unmount();
  document.body
    .querySelectorAll(".bb-guide-highlight-wrapper")
    ?.forEach((element) => {
      element.remove();
    });
};
