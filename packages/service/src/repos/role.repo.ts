import type { dao } from '@coli/db'
import { TenantedRepo } from './base.ts'
import type { Insertable, Selectable } from 'kysely'
import type { OmitTenantId } from './util.ts'
import { err } from '@coli/global'

export type NewRole = OmitTenantId<Insertable<dao.Role>>

export class RoleRepo extends TenantedRepo {
    public async create(data: NewRole) {
        return await this.db
            .insertInto('role')
            .values({
                ...data,
                tenantId: this.tenantId,
            })
            .returningAll()
            .executeTakeFirstOrThrow()
    }

    public async removeAllPermissions(roleId: string): Promise<void> {
        await this.getRoleOrThrow(roleId)

        await this.db.deleteFrom('permission').where('roleId', '=', roleId).execute()
    }

    public async assignPermissions(roleId: string, permissions: string[]) {
        await this.getRoleOrThrow(roleId)

        await this.db
            .insertInto('permission')
            .values(
                permissions.map((p) => ({
                    roleId: roleId,
                    key: p,
                }))
            )
            .execute()
    }

    public async getRoleById(roleId: string): Promise<Selectable<dao.Role> | undefined> {
        return this.db
            .selectFrom('role')
            .selectAll()
            .where('tenantId', '=', this.tenantId)
            .where('id', '=', roleId)
            .executeTakeFirst()
    }

    /**
     * Throws NotFoundError if role does not exist
     */
    private async getRoleOrThrow(roleId: string): Promise<Selectable<dao.Role>> {
        const role = await this.getRoleById(roleId)

        if (!role) {
            throw new err.NotFoundError('role')
        }

        return role
    }
}
