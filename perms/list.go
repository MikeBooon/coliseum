package perms

type ClientPermissions string
type TenantPermissions string
type GlobalPermissions string

const (
	ReadClientUser  GlobalPermissions = "r.client.user"
	WriteClientUser GlobalPermissions = "w.client.user"
)

func AllGlobalPermissions() []GlobalPermissions {
	return []GlobalPermissions{
		ReadClientUser,
		WriteClientUser,
	}
}

const (
	ReadTenantOptions  TenantPermissions = "r.tenant.options"
	WriteTenantOptions TenantPermissions = "w.tenant.options"
)

func AllTenantPermissions() []TenantPermissions {
	return []TenantPermissions{
		ReadTenantOptions,
		WriteTenantOptions,
	}
}
