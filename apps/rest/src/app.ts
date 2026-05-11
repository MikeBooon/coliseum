import Fastify from 'fastify'
import type { Config } from '@coli/config'
import { ProvisionCtrl } from './controllers/provision.ctrl.ts'
import { connectDb, type DB } from '@coli/db'
import type { Services } from '@coli/service'

export class App {
    private config: Config
    private fastify: Fastify.FastifyInstance

    constructor(config: Config, fastify: Fastify.FastifyInstance) {
        this.config = config
        this.fastify = fastify
    }

    public getFastify() {
        return this.fastify
    }

    public start() {
        this.fastify.listen({ port: this.config.restPort, host: '0.0.0.0' }, (err, _) => {
            if (err) throw err
        })
    }

    public async stop() {
        this.fastify.log.info('Shutting down...')
        await this.fastify.close()
    }
}

export function buildApp(config: Config, services: Services): App {
    const fastify = Fastify({
        logger: true,
    })

    const provisionCtrl = new ProvisionCtrl('/api/v1/provision', services)
    fastify.register(provisionCtrl.routes)

    fastify.get('/api', async (_, reply) => {
        reply.type('application/json').code(200)
        return { hello: 'world' }
    })

    return new App(config, fastify)
}
