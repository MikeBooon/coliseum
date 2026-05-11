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

const addNameColumn = (ctb: CreateTableBuilder<any, any>) => {
    return ctb.addColumn('name', 'text', (col) => col.notNull())
}

export async function up(db: Kysely<any>): Promise<void> {
    await db.schema
        .createTable('tenant')
        .$call(addBaseColumns)
        .$call(addNameColumn)
        .addColumn('slug', 'text', (col) => col.notNull().unique())
        .execute()

    await db.schema.createType('user_type').asEnum(['client', 'tenant']).execute()

    await db.schema
        .createTable('role')
        .$call(addBaseColumns)
        .$call(addTenantIdColumn)
        .$call(addNameColumn)
        .addColumn('type', sql`user_type`, (col) => col.notNull())
        .addColumn('default', 'boolean', (col) => col.notNull())
        .execute()

    await db.schema
        .createTable('permission')
        .$call(addBaseColumns)
        .addColumn('key', 'text', (col) => col.notNull())
        .addColumn('role_id', 'uuid', (col) =>
            col.references('role.id').onDelete('cascade').notNull()
        )
        .execute()

    await db.schema
        .createTable('user')
        .$call(addBaseColumns)
        .$call(addTenantIdColumn)
        .$call(addNameColumn)
        .addColumn('email', 'text', (col) => col.notNull())
        .addColumn('type', sql`user_type`, (col) => col.notNull())
        .addColumn('role_id', 'uuid', (col) =>
            col.references('role.id').onDelete('restrict').notNull()
        )
        .addUniqueConstraint('email_tenant_unique', ['email', 'tenant_id'])
        .execute()
}
