import { CLIENT_PERMISSIONS, GLOBAL_PERMISSIONS, TENANT_PERMISSIONS } from '@coli/perms'
import { RoleRepo } from '../repos/role.repo.ts'
import { TenantRepo } from '../repos/tenant.repo.ts'
import { UserRepo } from '../repos/user.repo.ts'
import { createHash } from '../util/hash.ts'
import { Service } from './base.ts'
import type { domain, dto } from '@coli/global'

const DEFAULT_TENANT_ROLE = {
    default: true,
    name: 'Admin',
    type: 'tenant',
} as const

const DEFAULT_CLIENT_ROLE = {
    default: true,
    name: 'Admin',
    type: 'client',
} as const

export class ProvisionService extends Service {
    public async provisionTenant(data: dto.ProvisionTenant): Promise<domain.Tenant> {
        return this.db.transaction().execute(async (tx) => {
            const tenantRepo = new TenantRepo(tx)
            const tenant = await tenantRepo.createTenant({
                name: data.name,
                slug: data.slug,
            })

            const roleRepo = new RoleRepo(tx, tenant.id)

            const tenantRole = await roleRepo.create(DEFAULT_TENANT_ROLE)
            const clientRole = await roleRepo.create(DEFAULT_CLIENT_ROLE)

            roleRepo.assignPermissions(tenantRole.id, [
                ...GLOBAL_PERMISSIONS,
                ...TENANT_PERMISSIONS,
            ])
            roleRepo.assignPermissions(clientRole.id, [
                ...GLOBAL_PERMISSIONS,
                ...CLIENT_PERMISSIONS,
            ])

            const userRepo = new UserRepo(tx, tenant.id)

            const name = data.email.split('@')[0]

            const user = await userRepo.createUser({
                roleId: tenantRole.id,
                email: data.email,
                name: name,
                type: 'tenant',
            })

            await userRepo.createCredential({
                userId: user.id,
                passwordHash: await createHash(data.password),
            })

            return tenant
        })
    }
}
