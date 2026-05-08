import * as path from 'path'

import { Pool } from 'pg'
import { promises as fs } from 'fs'
import {
  Kysely,
  Migrator,
  PostgresDialect,
  FileMigrationProvider,
} from 'kysely'

const MIGRATIONS_FOLDER = "migrations"
const CONNECTION_STRING = process.env["DATABASE_URL"]

if (!CONNECTION_STRING) {
  throw new Error("Missing env variable: DATABASE_URL")
}

async function migrate(connectionString: string) {
  const db = new Kysely<any>({
    dialect: new PostgresDialect({
      pool: new Pool({
        connectionString
      }),
    }),
  })

  const migrator = new Migrator({
    db,
    provider: new FileMigrationProvider({
      fs,
      path,
      migrationFolder: path.join(import.meta.dirname, MIGRATIONS_FOLDER),
    }),
  })

  const { error, results } = await migrator.migrateToLatest()

  results?.forEach((it) => {
    if (it.status === 'Success') {
      console.log(`migration "${it.migrationName}" was executed successfully`)
    } else if (it.status === 'Error') {
      console.error(`failed to execute migration "${it.migrationName}"`)
    }
  })

  if (error) {
    console.error('failed to migrate')
    console.error(error)
    process.exit(1)
  }

  await db.destroy()
}

await migrate(CONNECTION_STRING!)