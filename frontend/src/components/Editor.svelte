<script lang="ts">
    import ace from "ace-builds";
    import "ace-builds/src-noconflict/theme-github"; // Example theme
    import "ace-builds/src-noconflict/mode-markdown"; // Markdown syntax support
    import { onMount, onDestroy } from "svelte";
    import { appState } from "../state.svelte";


    let theme = "github";
    let mode = "markdown";
    let fontSize = 14;

    let editor: ace.Editor;
    let editorElement: HTMLElement;

    // Initialize editor
    onMount(async () => {
        editor = ace.edit(editorElement, {
            mode: `ace/mode/${mode}`,
            theme: `ace/theme/${theme}`,
            fontSize: `${fontSize}px`,
            value: appState.mdContent,
        });

        // Sync with parent component
        editor.session.on("change", () => {
            appState.mdContent = editor.getValue();
        });

        // Handle window resize
        const resizeObserver = new ResizeObserver(() => editor.resize());
        resizeObserver.observe(editorElement);

        onDestroy(() => {
            resizeObserver.disconnect();
            editor.destroy();
        });
    });

    $effect(() => {
        if (editor && appState.mdContent !== editor.getValue()) {
            editor.setValue(appState.mdContent, -1);
        }
    });
</script>

<div bind:this={editorElement} class="flex-1 ace-editor"></div>