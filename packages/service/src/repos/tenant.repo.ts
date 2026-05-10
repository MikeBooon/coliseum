import type { dao } from '@coli/db'
import { Repo } from './base.ts'
import type { Insertable } from 'kysely'

export type NewTenant = Insertable<dao.Tenant>

export class TenantRepo extends Repo {
    public create(data: NewTenant) {
        return this.db.insertInto('tenant').values(data).returningAll().executeTakeFirstOrThrow()
    }
}
