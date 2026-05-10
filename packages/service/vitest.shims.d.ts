declare module 'vitest' {
    export interface ProvidedContext {
        DATABASE_URL: string
    }
}

// mark this file as a module so augmentation works correctly
export {}
