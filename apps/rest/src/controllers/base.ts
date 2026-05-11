import type { Services } from '@coli/service'
import type { FastifyInstance, FastifyPluginOptions } from 'fastify'

export abstract class Controller {
    protected services: Services
    protected base: string
    constructor(base: string, services: Services) {
        this.services = services
        this.base = base
    }

    public abstract routes(fastify: FastifyInstance, options?: FastifyPluginOptions): Promise<void>
}
