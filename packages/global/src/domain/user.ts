import type { dao } from '@coli/db'
import type { Selectable } from 'kysely'

export type User = Selectable<dao.User>
