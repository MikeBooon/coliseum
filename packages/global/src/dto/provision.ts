import Type from 'typebox'

export const ProvisionTenant = Type.Object({
    name: Type.String(),
    slug: Type.String(),
    email: Type.String(),
})

export type ProvisionTenant = Type.Static<typeof ProvisionTenant>
