<script lang="ts">
    import { onMount } from "svelte";

    import {
        GetFiles,
        GoBack,
        SaveFileContent,
        SelectDir,
        SelectNote,
    } from "../../wailsjs/go/main/App";
    import { main } from "../../wailsjs/go/models";
    import { appState } from "../../src/state.svelte";
    import { selectNote } from "../../src/utils";

    let files: main.DirEntry[] = $state([]);

    function goBack() {
        GoBack().then((result) => {
            if (result.Error != "") {
                appState.notifications.push({
                    id: Date.now().toString(),
                    type: "warning",
                    title: result.Error,
                    time: 1
                });
                return;
            }

            fetchFiles();
        });
    }

    function selectDir(dirName: string) {
        console.log("Selecting directory: ", dirName);
        SelectDir(dirName).then(async (data) => {
            if (appState.selectedNote && !appState.selectedNote.IsDir) {
                console.log("Saving current note before changing directory", appState.selectedNote);
                const res = await SaveFileContent(
                    appState.selectedNote.Name,
                    appState.mdContent,
                );
                if (res.Error != "") {
                    console.error("Error saving file: ", res.Error);
                    return;
                }
            }
            fetchFiles();
        });
    }


    function fetchFiles() {
        GetFiles().then((data) => {
            if (data.Error != "") {
                return;
            }
            console.log("Fetched files: ", data.DirEntries);
            files = data.DirEntries;
        });
    }

    onMount(() => {
        fetchFiles();
    });
</script>

<div class="h-full flex flex-col">
    <div class="flex-1 flex flex-col gap-4 bg-darker p-4 text-light">
        <!-- Back button -->
        <button
            onclick={goBack}
            class="flex justify-start items-center gap-2 text-start mt-4 hover:text-white hover:cursor-pointer"
            ><p class="flex-1 break-all text-primary">..</p></button
        >

        {#each files as file}
            <button
                title={file.Path}
                onclick={() => {
                    if (file.IsDir) {
                        selectDir(file.Name);
                    } else {
                        selectNote(file.Name);
                    }
                }}
                class="flex justify-start items-center gap-2 text-start mt-4 hover:text-white hover:cursor-pointer"
            >
                {#if file.IsDir}
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        class="text-primary bi bi-folder"
                        viewBox="0 0 16 16"
                    >
                        <path
                            d="M.54 3.87.5 3a2 2 0 0 1 2-2h3.672a2 2 0 0 1 1.414.586l.828.828A2 2 0 0 0 9.828 3h3.982a2 2 0 0 1 1.992 2.181l-.637 7A2 2 0 0 1 13.174 14H2.826a2 2 0 0 1-1.991-1.819l-.637-7a2 2 0 0 1 .342-1.31zM2.19 4a1 1 0 0 0-.996 1.09l.637 7a1 1 0 0 0 .995.91h10.348a1 1 0 0 0 .995-.91l.637-7A1 1 0 0 0 13.81 4zm4.69-1.707A1 1 0 0 0 6.172 2H2.5a1 1 0 0 0-1 .981l.006.139q.323-.119.684-.12h5.396z"
                        />
                    </svg>
                {/if}
                <p class="flex-1 break-all text-primary hover:text-lighter">{file.Name}</p>
            </button>
        {/each}
    </div>
</div>
