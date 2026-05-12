// db-plugin.js

import fp from 'fastify-plugin'
import { Services } from '@coli/service'
import type { FastifyPluginCallback } from 'fastify'

const serviceProviderCallback: FastifyPluginCallback<{ services: Services }> = (
    fastify,
    options,
    done
) => {
    fastify.decorate('services', options.services)
    done()
}

export const serviceProvider = fp(serviceProviderCallback)

declare module 'fastify' {
    interface FastifyInstance {
        services: Services
    }
}
