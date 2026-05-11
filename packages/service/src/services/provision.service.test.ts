import { describe, expect, inject, test } from 'vitest'
import { connectDb } from '@coli/db'
import { err, type dto } from '@coli/global'
import { Services } from '../index.ts'
import { UserRepo } from '../repos/user.repo.ts'

const newTenantProvisionFixture: dto.ProvisionTenant = {
    email: 'test@test.com',
    name: 'tenant name',
    slug: 'tenant-slug',
}

describe('ProvisionService', () => {
    const db = connectDb(inject('DATABASE_URL'))
    const services = new Services({ db })

    test('can provision tenant', async () => {
        const tenant = await services.provision.provisionTenant(newTenantProvisionFixture)

        expect(tenant.slug).toBe('tenant-slug')
        await expect(services.provision.provisionTenant(newTenantProvisionFixture)).rejects.toThrow(
            err.UniqueConstraintError
        )

        const userRepo = new UserRepo(db, tenant.id)
        const createdUser = await userRepo.getByEmail('test@test.com')

        expect(createdUser).toBeTruthy()
    })
})
