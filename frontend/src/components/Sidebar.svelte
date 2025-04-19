<script lang="ts">
    import { onMount } from "svelte";

    import { GetFiles } from "../../wailsjs/go/main/App";
    import { main } from "../../wailsjs/go/models";

    let files: main.DirEntry[] = $state([]);

    function fetchFiles() {
        GetFiles().then((data) => {
            if (data.Error != "") {
                return;
            }

            files = data.DirEntries;
        });
    }

    onMount(() => {
        fetchFiles();
    });
</script>

<div class="flex-1 flex flex-col gap-4 bg-dark p-4 text-light">
    {#each files as file}
        <button
            title={file.Path}
            onclick={() => console.log(file.Path)}
            class="flex justify-start items-center gap-2 text-start mt-4"
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
            <p class="flex-1 break-all text-primary ">{file.Name}</p>
        </button>
    {/each}
</div>
