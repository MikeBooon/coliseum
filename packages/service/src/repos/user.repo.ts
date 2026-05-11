import type { dao } from '@coli/db'
import { TenantedRepo } from './base.ts'
import type { Insertable, Selectable } from 'kysely'
import { handleConstraintError, type OmitTenantId } from './util.ts'

export type NewUser = OmitTenantId<Insertable<dao.User>>

export class UserRepo extends TenantedRepo {
    public async getByEmail(email: string): Promise<Selectable<dao.User> | undefined> {
        return this.db
            .selectFrom('user')
            .selectAll()
            .where('tenantId', '=', this.tenantId)
            .where('email', '=', email.toLowerCase())
            .executeTakeFirst()
    }

    public async create(data: NewUser): Promise<Selectable<dao.User>> {
        try {
            return await this.db
                .insertInto('user')
                .values({
                    ...data,
                    email: data.email.toLowerCase(),
                    tenantId: this.tenantId,
                })
                .returningAll()
                .executeTakeFirstOrThrow()
        } catch (e) {
            handleConstraintError(e, 'email_tenant_unique', 'email')
            throw e
        }
    }
}
