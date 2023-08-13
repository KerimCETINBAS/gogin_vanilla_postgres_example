package types

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Messages uint   `json:"-"`
}

type UserCreateDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
