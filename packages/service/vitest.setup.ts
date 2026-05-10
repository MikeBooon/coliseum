import type { TestProject } from 'vitest/node'
import { PostgreSqlContainer, StartedPostgreSqlContainer } from '@testcontainers/postgresql'
import { migrate } from '@coli/db'

let pgContainer: StartedPostgreSqlContainer

export default async function setup(project: TestProject) {
    pgContainer = await new PostgreSqlContainer('postgres:18').start()
    const pgUri = pgContainer.getConnectionUri()
    await migrate(pgUri)
    await pgContainer.snapshot()

    project.provide('DATABASE_URL', pgUri)

    project.onTestsRerun(async () => {
        await pgContainer.restoreSnapshot()
    })
}

export async function teardown() {
    await pgContainer.stop()
}
