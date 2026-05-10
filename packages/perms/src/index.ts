export const GLOBAL_PERMISSIONS = ['read.client.user', 'write.client.user'] as const
export type GlobalPermission = (typeof GLOBAL_PERMISSIONS)[number]

export const CLIENT_PERMISSIONS = [] as const
export type ClientPermission = (typeof CLIENT_PERMISSIONS)[number]

export const TENANT_PERMISSIONS = ['read.tenant.options', 'write.tenant.options'] as const
export type TenantPermission = (typeof TENANT_PERMISSIONS)[number]

export type AllPermissions = GlobalPermission | TenantPermission | ClientPermission
