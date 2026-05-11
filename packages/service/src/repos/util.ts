import { err } from '@coli/global'

function isObject(val: any): val is Record<string, any> {
    return val !== null && typeof val === 'object' && !Array.isArray(val)
}

export function handleConstraintError(e: any, constraint: string, property: string) {
    if (isObject(e) && e['constraint'] && e['constraint'] === constraint) {
        throw new err.UniqueConstraintError(property)
    }
}

export type OmitTenantId<T> = Omit<T, 'tenantId'>
