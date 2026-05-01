package domain

type UserType string

const (
	TenantUserType UserType = "tenant"
	ClientUserType UserType = "client"
)
