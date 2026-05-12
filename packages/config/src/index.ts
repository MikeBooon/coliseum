export type Config = {
    databaseUrl: string
    domain: string
    encKey: string
    restPort: number
    logging: boolean
}

export function getConfig(): Config {
    return {
        databaseUrl: getEnvOrThrow('DATABASE_URL'),
        domain: getEnvOrThrow('DOMAIN'),
        encKey: getEnvOrThrow('ENC_KEY'),
        restPort: getIntEnvOrThrow('REST_PORT'),
        logging: true,
    }
}

export class MissingEnvVarError extends Error {
    public key: string
    constructor(key: string) {
        super(`Missing env var: '${key}'`)
        this.key = key
    }
}

function getEnvOrThrow(key: string): string {
    const v = process.env[key]
    if (!v) {
        throw new MissingEnvVarError(key)
    }
    return v
}

function getIntEnvOrThrow(key: string): number {
    const v = process.env[key]
    if (!v) {
        throw new MissingEnvVarError(key)
    }
    return parseInt(v)
}

export * from './test/config.ts'
