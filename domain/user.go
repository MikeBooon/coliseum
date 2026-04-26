package domain

type User struct {
	Base  `tstype:",extends"`
	Email string `json:"email"`
}
