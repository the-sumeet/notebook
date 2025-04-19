interface AppState {
    preview: boolean;
    mdContent: string;
}

export const appState: AppState = $state({
    preview: false,
    mdContent: ``
})