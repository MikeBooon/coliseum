import Type from 'typebox'
import { INPUTS } from '../constants/index.ts'

export const ProvisionTenant = Type.Object({
    name: Type.String({
        minLength: INPUTS.MIN_NAME_LENGTH,
        maxLength: INPUTS.MAX_NAME_LENGTH,
    }),
    slug: Type.String({
        minLength: INPUTS.MIN_SLUG_LENGTH,
        maxLength: INPUTS.MAX_SLUG_LENGTH,
    }),
    email: Type.String({ format: 'email', maxLength: INPUTS.MAX_EMAIL_LENGTH }),
    password: Type.String({
        minLength: INPUTS.PASSWORD_MIN_LENGTH,
        pattern: INPUTS.PASSWORD_REGEX,
    }),
})

export type ProvisionTenant = Type.Static<typeof ProvisionTenant>
