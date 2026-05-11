import type { FastifyInstance } from 'fastify'
import { Controller } from './base.ts'

// Fix this. The routes is not bound to the class. Need to fix.
export class ProvisionCtrl extends Controller {
    public async routes(fastify: FastifyInstance): Promise<void> {
        fastify.post(`${this.base}/tenant`, async (request, reply) => {
            return { hello: 'world' }
        })
    }
}
