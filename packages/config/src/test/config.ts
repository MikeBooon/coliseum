import type { Config } from '../index.ts'

export function getTestConfig(dbUrl: string): Config {
    return {
        databaseUrl: dbUrl,
        domain: 'https://[sub].localhost',
        restPort: 6665,
        encKey: 'test-key',
        logging: false,
    }
}
