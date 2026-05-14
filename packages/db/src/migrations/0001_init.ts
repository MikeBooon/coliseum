import { CreateTableBuilder, Kysely, sql } from 'kysely'

const addBaseColumns = (ctb: CreateTableBuilder<any, any>) => {
    return ctb
        .addColumn('id', 'uuid', (col) => col.primaryKey().defaultTo(sql`gen_random_uuid()`))
        .addColumn('created_at', 'timestamptz', (col) => col.notNull().defaultTo(sql`now()`))
}

const addTenantIdColumn = (ctb: CreateTableBuilder<any, any>) => {
    return ctb.addColumn('tenant_id', 'uuid', (col) =>
        col.references('tenant.id').onDelete('cascade').notNull()
    )
}

const addOptionalTenantIdColumn = (ctb: CreateTableBuilder<any, any>) => {
    return ctb.addColumn('tenant_id', 'uuid', (col) =>
        col.references('tenant.id').onDelete('cascade')
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
        .createType('task_status')
        .asEnum(['pending', 'running', 'failed', 'complete'])
        .execute()

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

    await db.schema
        .createTable('user_credential')
        .$call(addBaseColumns)
        .addColumn('password_hash', 'text')
        .addColumn('totp_secret', 'text')
        .addColumn('user_id', 'uuid', (col) =>
            col.references('user.id').onDelete('cascade').notNull().unique()
        )
        .execute()

    await db.schema
        .createTable('task')
        .$call(addBaseColumns)
        .$call(addOptionalTenantIdColumn)
        .addColumn('tag', 'text', (col) => col.notNull())
        .addColumn('data', 'jsonb', (col) => col.notNull())
        .addColumn('status', sql`task_status`, (col) => col.notNull().defaultTo(`pending`))
        .addColumn('attempts', 'integer', (col) => col.notNull().defaultTo(0))
        .addColumn('max_attempts', 'integer', (col) => col.notNull().defaultTo(1))
        .addColumn('last_run', 'timestamptz')
        .addColumn('last_error', 'text')
        .addColumn('worker_id', 'text')
        .execute()
}
