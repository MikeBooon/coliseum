import { type Config, getTestConfig } from '@coli/config'
import { connectDb, type DB } from '@coli/db'
import { Services } from '@coli/service'

export type TestContext = {
    config: Config
    services: Services
    db: DB
}

export function getTestContext(dbUrl: string): TestContext {
    const config = getTestConfig(dbUrl)
    const db = connectDb(config.databaseUrl)
    const services = new Services({ db, config })
    return {
        config: config,
        db: db,
        services: services,
    }
}
