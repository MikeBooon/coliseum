import type { TAny } from 'typebox'
import type { Task } from './base.ts'

/**
 * Use for lookup by worker
 */
export const TASK_TAG_MAP: Readonly<Record<string, Task<TAny>>> = {}
