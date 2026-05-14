import type { DB } from '@coli/db'
import { UserService } from './services/user.service.ts'
import { ProvisionService } from './services/provision.service.ts'
import type { Config } from '@coli/config'

export type SerivceDependencies = {
    db: DB
    config: Config
}

export class Services {
    public user: UserService
    public provision: ProvisionService
    constructor(deps: SerivceDependencies) {
        this.user = new UserService(deps)
        this.provision = new ProvisionService(deps)
    }
}
