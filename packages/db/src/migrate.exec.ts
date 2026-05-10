import { migrate } from './migrate.ts'

const CONNECTION_STRING = process.env['DATABASE_URL']

if (!CONNECTION_STRING) {
    throw new Error('Missing env variable: DATABASE_URL')
}
await migrate(CONNECTION_STRING!)
