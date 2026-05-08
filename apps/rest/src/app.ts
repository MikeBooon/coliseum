import Fastify from 'fastify'
import type { Config } from '@coli/config'

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

export function buildApp(config: Config): App {
    const fastify = Fastify({
        logger: true,
    })

    fastify.get('/api', async (_, reply) => {
        reply.type('application/json').code(200)
        return { hello: 'world' }
    })

    return new App(config, fastify)
}
