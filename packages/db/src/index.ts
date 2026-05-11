import { CamelCasePlugin, Kysely, PostgresDialect } from 'kysely'

import * as dao from './dao.gen.ts'
import { Pool } from 'pg'

export function connectDb(databaseUrl: string): Kysely<dao.DB> {
    const dialect = new PostgresDialect({
        pool: new Pool({
            connectionString: databaseUrl,
        }),
    })
    return new Kysely<dao.DB>({
        dialect,
        plugins: [new CamelCasePlugin()],
    })
}

export { dao }
export type DB = Kysely<dao.DB>
export * from './migrate.ts'
