package perms

type ClientPermissions string
type TenantPermissions string
type GlobalPermissions string

type Permissions interface {
	ClientPermissions | TenantPermissions | GlobalPermissions
}

const (
	CreateClientUser TenantPermissions = "create.client.user"
	DeleteClientUser TenantPermissions = "delete.client.user"
)
