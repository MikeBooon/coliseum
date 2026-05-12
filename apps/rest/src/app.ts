import Fastify, { fastify } from 'fastify'
import type { Config } from '@coli/config'
import provisionCtrl from './controllers/provision.ctrl.ts'
import type { Services } from '@coli/service'
import { serviceProvider } from './plugins/services.ts'
import type { TypeBoxTypeProvider } from '@fastify/type-provider-typebox'

export type FastifyAppInstance = ReturnType<typeof getFastify>

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
    const fastify = getFastify(config.logging)

    fastify.register(serviceProvider, { services: services })

    const V1_PREFIX = '/api/v1'

    fastify.register(provisionCtrl, { prefix: `${V1_PREFIX}/provision` })

    fastify.get('/api', async (_, reply) => {
        reply.type('application/json').code(200)
        return { hello: 'world' }
    })

    return new App(config, fastify)
}

function getFastify(withLogger: boolean) {
    return Fastify({
        logger: withLogger,
    }).withTypeProvider<TypeBoxTypeProvider>()
}
