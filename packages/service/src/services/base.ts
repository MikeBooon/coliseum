import type { DB } from '@coli/db'
import type { SerivceDependencies } from '../index.ts'

export class Service {
    protected db: DB
    constructor(deps: SerivceDependencies) {
        this.db = deps.db
    }
}
