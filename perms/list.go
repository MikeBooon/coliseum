package perms

type ClientPermissions string
type AdminPermissions string
type GlobalPermissions string

type Permissions interface {
	ClientPermissions | AdminPermissions | GlobalPermissions
}

const (
	CreateClientUser AdminPermissions = "create.client.user"
	DeleteClientUser AdminPermissions = "delete.client.user"
)
