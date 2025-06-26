import { GetSelectedNote, SaveFileContent, SelectNote } from "../wailsjs/go/main/App";
import { appState } from "./state.svelte";

export function getSelectedNote() {
    GetSelectedNote().then((selectedDir) => {
        if (selectedDir.Path != "") {
            appState.selectedNote = selectedDir;
        } else {
            appState.selectedNote = null;
        }
    });
}

export function selectNote(filPath: string) {
    SelectNote(filPath).then(async (data) => {
        if (appState.selectedNote) {
            const res = await SaveFileContent(
                appState.selectedNote.Name,
                appState.mdContent,
            );
            if (res.Error != "") {
                console.error("Error saving file: ", res.Error);
                return;
            }
        }

        if (data.Error != "") {
            return;
        }
        console.log("Selected note: ", data.SelectedNote);
        appState.selectedNote = data.SelectedNote;
    });
}