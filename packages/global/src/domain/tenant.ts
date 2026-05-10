import type { dao } from '@coli/db'
import type { Selectable } from 'kysely'

export type Tenant = Selectable<dao.Tenant>
