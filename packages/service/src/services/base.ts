import type { DB } from '@coli/db'
import type { SerivceDependencies } from '../index.ts'
import type { Config } from '@coli/config'

export class Service {
    protected db: DB
    protected config: Config
    constructor(deps: SerivceDependencies) {
        this.db = deps.db
        this.config = deps.config
    }
}
