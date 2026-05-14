import type { DB } from '@coli/db'
import type { Static, TSchema } from 'typebox'

export type SaveTaskOptions = {
    tenantId?: string
}

export abstract class Task<Schema extends TSchema> {
    public abstract readonly tag: string
    public abstract readonly schema: Schema

    public abstract readonly maxAttempts: number

    protected readonly data: Static<Schema>

    protected tenantId?: string

    constructor(data: Static<Schema>, opts?: SaveTaskOptions) {
        this.data = data
        this.tenantId = opts?.tenantId
    }

    public async save(db: DB): Promise<void> {
        await db
            .insertInto('task')
            .values({
                tag: this.tag,
                data: this.data,
                tenantId: this.tenantId,
                maxAttempts: this.maxAttempts,
            })
            .execute()
    }

    /**
     * Gets run by worker
     */
    public abstract run(): Promise<void>
}
