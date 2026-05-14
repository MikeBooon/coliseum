import { describe, expect, inject, test } from 'vitest'
import { err, type dto } from '@coli/global'
import { UserRepo } from '../repos/user.repo.ts'
import { getTestContext } from '@coli/testing'

const newTenantProvisionFixture: dto.ProvisionTenant = {
    email: 'test@test.com',
    name: 'tenant name',
    slug: 'tenant-slug',
    password: 'testPass123123@',
}

describe('ProvisionService', () => {
    const { db, services } = getTestContext(inject('DATABASE_URL'))

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
