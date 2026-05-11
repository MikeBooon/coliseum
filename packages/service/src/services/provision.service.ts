import { RoleRepo } from '../repos/role.repo.ts'
import { TenantRepo } from '../repos/tenant.repo.ts'
import { UserRepo } from '../repos/user.repo.ts'
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
            const tenant = await tenantRepo.create({
                name: data.name,
                slug: data.slug,
            })

            const roleRepo = new RoleRepo(tx, tenant.id)

            const tenantRole = await roleRepo.create(DEFAULT_TENANT_ROLE)
            await roleRepo.create(DEFAULT_CLIENT_ROLE)

            const userRepo = new UserRepo(tx, tenant.id)

            const name = data.email.split('@')[0]

            await userRepo.create({
                roleId: tenantRole.id,
                email: data.email,
                name: name,
                type: 'tenant',
            })

            return tenant
        })
    }
}
