import { getConfig } from '@coli/config'
import { buildApp } from './app.ts'
import { Services } from '@coli/service'
import { connectDb } from '@coli/db'

const config = getConfig()

const db = connectDb(config.databaseUrl)
const services = new Services({
    db: db,
    config: config,
})

const app = buildApp(config, services)

app.start()

async function shutdown() {
    try {
        await app.stop()
        await db.destroy()
        process.exit(0)
    } catch (e) {
        console.error(e)
        process.exit(1)
    }
}

process.on('SIGTERM', shutdown)
process.on('SIGINT', shutdown)
