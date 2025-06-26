<script lang="ts">
  import Sidebar from "./components/Sidebar.svelte";
  import Editor from "./components/Editor.svelte";
  import TopBar from "./components/TopBar.svelte";
  import { appState } from "./state.svelte";
  import Preview from "./components/Preview.svelte";
  import { onMount } from "svelte";
  import { getSelectedNote } from "./utils";
  import Notification from "./components/Notification.svelte";
  import Modal from "./components/Modal.svelte";
  import app from "./main";
  import { DeleteNoteFile, NewFolder, NewNote } from "../wailsjs/go/main/App";

  let newNoteName = "";
  let newFolderName = "";
  let refreshFiles:boolean = $state(false);

  function showDeleteNoteModal() {
    if (!appState.selectedNote) {
      return;
    }

    appState.modal = {
      title: "Delete Note",
      message: `Are you sure you want to delete the note "${appState.selectedNote.Name}"? This action cannot be undone.`,
      primaryTitle: "Delete",
      onPrimaryClick: () => {
        DeleteNoteFile(appState.selectedNote!.Path)
          .then((result) => {
            if (result.Error) {
              appState.notifications.push({
                id: Date.now().toString(),
                type: "error",
                title: result.Error,
                time: 3,
              });
            }

            appState.selectedNote = null; // Clear the selected note after deletion
            appState.modal = null; // Close the modal after deletion
            refreshFiles = !refreshFiles; // Trigger a refresh of the sidebar

          })
          .catch((error) => {
            console.error("Error deleting note:", error);
            appState.modal = null; // Close the modal on error
          });
      },
    };
  }

  function showNewNoteModal() {
    appState.modal = {
      title: "New Note",
      input: {
        value: "",
        label: "Note Title",
        type: "text",
        onupdate: (value: string) => {newNoteName = value;},
        validator: (value: string) => {
          if (value.length < 3) {
            return "Title must be at least 3 characters long.";
          }
          return "";
        },
      },
      primaryTitle: "Create",
      onPrimaryClick: () => {
        NewNote(newNoteName)
          .then((result) => {
            console.log("New note created:", result);
            appState.selectedNote = result.SelectedNote;
            appState.modal = null; // Close the modal after creating the note
            refreshFiles = !refreshFiles; // Trigger a refresh of the sidebar
          })
          .catch((error) => {
            console.error("Error creating new note:", error);
            appState.modal = null; // Close the modal on error
          });
      },
    };
  }

  function showNewFolderModal() {
    appState.modal = {
      title: "New Folder",
      input: {
        label: "Folder Title",
        type: "text",
        onupdate: (value: string) => {
          newFolderName = value;
        },
      },
      primaryTitle: "Create",
      onPrimaryClick: () => {
        NewFolder(newFolderName)
          .then((result) => {
            refreshFiles = !refreshFiles; // Trigger a refresh of the sidebar
          })
          .catch((error) => {
            console.error("Error creating new folder:", error);
          });
        appState.modal = null;
      },
    };
  }

  onMount(() => {
    getSelectedNote();

    window.runtime.EventsOn("deleteNote", () => {
      showDeleteNoteModal();
    });

    window.runtime.EventsOn("newNote", () => {
      showNewNoteModal();
    });

    window.runtime.EventsOn("newFolder", () => {
      showNewFolderModal();
    });
  });
</script>

<div class="flex h-screen">
  <div class="min-w-48 w-2/12 max-w-64 overflow-y-scroll">
    {#key refreshFiles}
    <Sidebar />
    {/key}
  </div>

  <div class="flex-1 flex flex-col">
    <TopBar />

    <hr class="bg-gray-200 border-gray-200" />

    {#if appState.preview}
      <div class="overflow-y-scroll">
        <Preview />
      </div>
    {:else}
      <Editor />
    {/if}
  </div>

  <div
    aria-live="assertive"
    class="pointer-events-none fixed inset-0 flex items-end px-4 py-6 sm:items-start sm:p-6"
  >
    <div class="flex w-full flex-col items-center space-y-4 sm:items-end">
      {#each appState.notifications as notification, i}
        <Notification {notification} />
      {/each}
    </div>
  </div>
</div>

{#if appState.modal}
  <Modal modal={appState.modal} />
{/if}
