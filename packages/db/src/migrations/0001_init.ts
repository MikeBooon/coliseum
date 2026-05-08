import { CreateTableBuilder, Kysely, sql } from 'kysely'

const addBaseColumns = (ctb: CreateTableBuilder<any, any>) => {
    return ctb
        .addColumn('id', 'uuid', (col) => col.primaryKey().defaultTo(sql`gen_random_uuid()`))
        .addColumn('created_at', 'date', (col) => col.notNull().defaultTo(sql`now()`))
}

const addTenantIdColumn = (ctb: CreateTableBuilder<any, any>) => {
    return ctb.addColumn('tenant_id', 'uuid', (col) =>
        col.references('tenant.id').onDelete('cascade').notNull()
    )
}

export async function up(db: Kysely<any>): Promise<void> {
    await db.schema.createTable('tenant').$call(addBaseColumns).$call(addTenantIdColumn).execute()

    await db.schema
        .createTable('user')
        .$call(addBaseColumns)
        .$call(addTenantIdColumn)
        .addColumn('email', 'text', (col) => col.notNull())
        .addUniqueConstraint('email_tenant_unique', ['email'])
        .execute()
}
