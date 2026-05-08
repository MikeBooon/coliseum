import { getConfig } from '@coli/config'
import { buildApp } from './app.ts'

const config = getConfig()

const app = buildApp(config)

app.start()

async function shutdown() {
    try {
        await app.stop()
        process.exit(0)
    } catch (e) {
        console.error(e)
        process.exit(1)
    }
}

process.on('SIGTERM', shutdown)
process.on('SIGINT', shutdown)
