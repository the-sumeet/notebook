<script lang="ts">
  import { appState } from "../../src/state.svelte";
  import Input from "./Input.svelte";

  let { modal }: { modal: Modal } = $props();
</script>

<div
  class="relative z-10"
  aria-labelledby="modal-title"
  role="dialog"
  aria-modal="true"
>
  <!--
      Background backdrop, show/hide based on modal state.
  
      Entering: "ease-out duration-300"
        From: "opacity-0"
        To: "opacity-100"
      Leaving: "ease-in duration-200"
        From: "opacity-100"
        To: "opacity-0"
    -->
  <div
    class="fixed inset-0 bg-gray-500/75 transition-opacity"
    aria-hidden="true"
  ></div>

  <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
    <div
      class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
    >
      <!--
          Modal panel, show/hide based on modal state.
  
          Entering: "ease-out duration-300"
            From: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            To: "opacity-100 translate-y-0 sm:scale-100"
          Leaving: "ease-in duration-200"
            From: "opacity-100 translate-y-0 sm:scale-100"
            To: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
        -->
      <div
        class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"
      >
        <div>
          <div class="mt-3 text-center sm:mt-5">
            <h3 class="text-base font-semibold text-gray-900" id="modal-title">
              {modal.title}
            </h3>
            {#if modal.message}
              <div class="mt-2">
                <p class="text-sm text-gray-500">{modal.message}</p>
              </div>
            {/if}
          </div>

          {#if modal.input}
            <Input input={modal.input} />
          {/if}
        </div>
        <div class="flex mt-5 sm:mt-6 sm:gap-3">
          {#if modal.primaryTitle && modal.onPrimaryClick}
            <button
              onclick={modal.onPrimaryClick}
              type="button"
              class="inline-flex w-full justify-center rounded-md bg-dark px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-darker focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 sm:col-start-2"
              >{modal.primaryTitle}</button
            >
          {/if}
          <button
            onclick={() => {
              appState.modal = null;
            }}
            type="button"
            class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:col-start-1 sm:mt-0"
            >Cancel</button
          >
        </div>
      </div>
    </div>
  </div>
</div>
