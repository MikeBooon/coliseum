import { describe, expect, test } from 'vitest'
import { connectDb } from '@coli/db'
import type { dto } from '@coli/global'
import { Services } from '../index.ts'

const newTenantProvisionFixture: dto.ProvisionTenant = {
    email: 'test@test.com',
    name: 'tenant name',
    slug: 'tenant-slug',
}

describe('ProvisionService', () => {
    console.log(process.env['DATABASE_URL_2']!)
    const db = connectDb(process.env['DATABASE_URL']!)
    const services = new Services({ db })

    test('can provision tenant', async () => {
        const tenant = await services.provision.provisionTenant(newTenantProvisionFixture)
        expect(tenant.slug).toBe('tenant-slug')
    })
})
