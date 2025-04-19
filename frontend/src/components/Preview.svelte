<script lang="ts">
    import { marked } from "marked";
    import { onMount } from "svelte";
    import { appState } from "../state.svelte";

    onMount(() => {
        marked.setOptions({
            gfm: true,
            breaks: true,
        });
    });
</script>
<div class="prose p-2">{@html marked(appState.mdContent)}</div>
