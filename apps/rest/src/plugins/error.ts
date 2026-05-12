import fp from 'fastify-plugin'
import type { FastifyPluginCallback } from 'fastify'
import { err } from '@coli/global'

const REMOVED_ERROR_DATA_PROPERTIES = ['statusCode', 'stack'] as const

const handler: FastifyPluginCallback = (fastify, _, done) => {
    fastify.setErrorHandler((error, _, reply) => {
        if (error instanceof err.RequestError) {
            const data: any = { ...error, message: error.message }
            for (const p of REMOVED_ERROR_DATA_PROPERTIES) {
                delete data[p]
            }
            reply.code(error.statusCode).send(data)
        } else {
            throw error
        }
    })
    done()
}

export const errorHandler = fp(handler)
