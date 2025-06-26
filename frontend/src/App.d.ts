

interface Notification {
    id: string
    title: string
    message?: string
    type?: 'info' | 'error' | 'success' | 'warning'
    primaryTitle?: string
    secondaryTitle?: string
    onPrimaryClick?: () => void
    onSecondaryClick?: () => void
    time: number
}

interface Input {
    label?: string
    type?: 'text' | 'password' | 'email' | 'number' | 'select'
    value: string
    onupdate?: (value: string) => void
    validator?: (value: string) => string?
    placeholder?: string
}

interface Modal {
    title: string
    message?: string
    primaryTitle?: string
    onPrimaryClick?: () => void
    input?: Input 
}