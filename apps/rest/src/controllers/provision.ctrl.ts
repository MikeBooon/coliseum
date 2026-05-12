import { dto } from '@coli/global'
import type { FastifyAppInstance } from '../app.ts'

export default async function routes(fastify: FastifyAppInstance) {
    fastify.post('/tenant', { schema: { body: dto.ProvisionTenant } }, async (request, reply) => {
        const tenant = await fastify.services.provision.provisionTenant(request.body)
        return { tenant }
    })
}
