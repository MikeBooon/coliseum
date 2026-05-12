import type { Services } from '@coli/service'
import type { FastifyInstance, FastifyPluginOptions } from 'fastify'

export abstract class Controller {
    protected services: Services
    constructor(services: Services) {
        this.services = services
    }

    public abstract routes(fastify: FastifyInstance, options?: FastifyPluginOptions): Promise<void>
}
