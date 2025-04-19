<script>
  import Sidebar from "./components/Sidebar.svelte";
  import "bootstrap-icons/font/bootstrap-icons.css";
  import Editor from "./components/Editor.svelte";
  import TopBar from "./components/TopBar.svelte";
  import { appState } from "./state.svelte";
  import Preview from "./components/Preview.svelte";
</script>

<div class="flex h-screen">
  <div class="min-w-48 w-2/12 max-w-64 overflow-y-scroll">
    <Sidebar />
  </div>

  <div class="flex-1 flex flex-col overflow-x-hidden">
    <TopBar />

    <hr class="bg-gray-200 border-gray-200">

    {#if appState.preview}
      <div class="overflow-y-scroll">
        <Preview />
      </div>
    {:else}
      <Editor />
    {/if}
  </div>
</div>
