import { type Config, getTestConfig } from '@coli/config'
import { connectDb, type DB } from '@coli/db'
import { Services } from '@coli/service'

export type TestContext = {
    config: Config
    services: Services
    db: DB
}

export function getTestContext(dbUrl: string): TestContext {
    const db = connectDb(dbUrl)
    const services = new Services({ db })
    return {
        config: getTestConfig(dbUrl),
        db: db,
        services: services,
    }
}
