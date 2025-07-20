<script lang="ts">
    import { appState } from "../state.svelte";
    import Bold from "./icons/Bold.svelte";
    import CopyIcon from "./icons/CopyIcon.svelte";
    import DisablePreview from "./icons/DisablePreview.svelte";
    import EnablePreviewIcon from "./icons/EnablePreviewIcon.svelte";
    import Italic from "./icons/Italic.svelte";
    import Underline from "./icons/Underline.svelte";
    let noteName = $state("");
</script>

<div class="flex flex-col">
    <!-- Note title -->
    <div class="p-2 font-bold text-lg">
        <input
            type="text"
            class="rounded p-1 w-full placeholder:text-gray-500 outline-0"
            placeholder="Enter name for new note"
            value={noteName}
        />
    </div>

    <div class="flex p-2 gap-2 overflow-x-scroll justify-end">
        <button
            onclick={() => {
                appState.bold = true;
            }}
            type="button"
            class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
        >
            <Bold />
        </button>

        <button
            onclick={() => {
                appState.italic = true;
            }}
            type="button"
            class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
        >
            <Italic />
        </button>
        <button
            onclick={() => {
                appState.underline = true;
            }}
            type="button"
            class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
        >
            <Underline />
        </button>

        <button
            type="button"
            class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
        >
            <CopyIcon />
        </button>

        {#if appState.preview}
            <button
                onclick={() => (appState.preview = false)}
                type="button"
                class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
            >
                <DisablePreview />
            </button>
        {:else}
            <button
                onclick={() => (appState.preview = true)}
                type="button"
                class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
            >
                <EnablePreviewIcon />
            </button>
        {/if}

        <!-- <button
            type="button"
            class="rounded bg-darker px-2 py-1 text-xs font-semibold text-white shadow-sm hover:bg-dark focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-darker"
        >
            <DownloadIcon />
        </button> -->
    </div>
</div>
