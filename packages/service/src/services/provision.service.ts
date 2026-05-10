import { TenantRepo } from '../repos/tenant.repo.ts'
import { Service } from './base.ts'
import type { domain, dto } from '@coli/global'

export class ProvisionService extends Service {
    public async provisionTenant(data: dto.ProvisionTenant): Promise<domain.Tenant> {
        return this.db.transaction().execute(async (tx) => {
            const tenantRepo = new TenantRepo(tx)
            const tenant = tenantRepo.create({
                name: data.name,
                slug: data.slug,
            })

            return tenant
        })
    }
}
