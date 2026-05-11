import type { DB } from '@coli/db'

export class Repo {
    protected db: DB
    constructor(db: DB) {
        this.db = db
    }
}

export class TenantedRepo extends Repo {
    protected tenantId: string

    constructor(db: DB, tenantId: string) {
        super(db)
        this.tenantId = tenantId
    }
}
