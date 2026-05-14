import type { dao } from '@coli/db'
import { Repo } from './base.ts'
import type { Insertable } from 'kysely'
import { handleConstraintError } from './util.ts'

export type NewTenant = Insertable<dao.Tenant>

export class TenantRepo extends Repo {
    public async createTenant(data: NewTenant) {
        try {
            return await this.db
                .insertInto('tenant')
                .values(data)
                .returningAll()
                .executeTakeFirstOrThrow()
        } catch (e) {
            handleConstraintError(e, 'tenant_slug_key', 'slug')
            throw e
        }
    }
}
