import { main } from "../wailsjs/go/models";

interface AppState {
    preview: boolean;
    mdContent: string;
    selectedNote?: main.DirEntry | null;
    notifications: Notification[];
    modal?: Modal | null;
}

export const appState: AppState = $state({
    preview: false,
    mdContent: ``,
    selectedNote: null,
    notifications: [],
    modal: null
})