<script lang="ts">
  import { appState } from "../../src/state.svelte";
  import Check from "./icons/Check.svelte";
  import Warning from "./icons/Warning.svelte";
  let { notification }: { notification: Notification } = $props();

  let border = "";

  if (notification.type == "warning") {
    border = "border-2 border-yellow-500";
  } else if (notification.type == "success") {
    border = "border-2 border-green-500";
  }

  function closeNotification() {
    // Remove the notification from the list
    appState.notifications = appState.notifications.filter(
      (n, index) => n.id != notification.id,
    );
  }

  setTimeout(() => {
    closeNotification();
  }, notification.time * 1000);
</script>

<div
  class="{border} pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black/5"
>
  <div class="p-4">
    <div class="flex items-start">
      <div class="ml-3 w-0 flex-1 pt-0.5">
        <p class="text-sm font-medium text-gray-900">
          {notification.title}
        </p>

        {#if notification.message}
          <p class="mt-1 text-sm text-gray-500">{notification.message}</p>
        {/if}

        {#if (notification.onPrimaryClick && notification.onPrimaryClick) || (notification.onSecondaryClick && notification.onSecondaryClick)}
          <div class="mt-3 flex space-x-7">
            {#if notification.primaryTitle && notification.onPrimaryClick}
              <button
                onclick={notification.onPrimaryClick}
                type="button"
                class="rounded-md bg-white text-sm font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2"
                >{notification.primaryTitle}</button
              >
            {/if}

            {#if notification.secondaryTitle && notification.onSecondaryClick}
              <button
                onclick={notification.onSecondaryClick}
                type="button"
                class="rounded-md bg-white text-sm font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2"
                >{notification.secondaryTitle}</button
              >
            {/if}
          </div>
        {/if}
      </div>

      <!-- Close button -->
      <div class="ml-4 flex shrink-0">
        <button
          onclick={closeNotification}
          type="button"
          class="inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2"
        >
          <span class="sr-only">Close</span>
          <svg
            class="size-5"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
            data-slot="icon"
          >
            <path
              d="M6.28 5.22a.75.75 0 0 0-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 1 0 1.06 1.06L10 11.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L11.06 10l3.72-3.72a.75.75 0 0 0-1.06-1.06L10 8.94 6.28 5.22Z"
            />
          </svg>
        </button>
      </div>
    </div>
  </div>
</div>
