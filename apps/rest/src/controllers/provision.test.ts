import { assert, beforeAll, describe, expect, inject, test } from 'vitest'
import { type dto } from '@coli/global'
import { buildApp } from '../app.ts'
import { getTestContext } from '@coli/testing'
import { buildTestDomain } from '../test.ts/util.ts'

const newTenantProvisionFixture: dto.ProvisionTenant = {
    email: 'test@test.com',
    name: 'tenant name',
    slug: 'tenant-slug',
}

describe('ProvisionService', () => {
    const context = getTestContext(inject('DATABASE_URL'))
    const app = buildApp(context.config, context.services)
    const fastify = app.getFastify()

    beforeAll(async () => {
        await fastify.ready()
    })

    test('can provision tenant', async () => {
        const res = await fastify.inject({
            method: 'POST',
            url: buildTestDomain('/api/v1/provision/tenant'),
            payload: newTenantProvisionFixture,
        })

        expect(res.statusCode).toBe(200)

        expect(res.body).toContain('tenant')
    })

    test('new tenant validates input', async () => {
        const response = await fastify.inject({
            method: 'POST',
            url: buildTestDomain('/api/v1/provision/tenant'),
            payload: {
                blah: 'blah',
            },
        })

        expect(response.statusCode).toBe(400)
    })
})
