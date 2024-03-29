package models

type Status string

const (
	Email Status = "email"
	Phone Status = "phone"
)

// main
type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name" validate:"required, min=5, max=50"`
	Password string `json:"password" validate:"required, min=5 max=15"`
	Phone    string `json:"phone" validate:"min=5,max=15"`
	Email    string `json:"email" validate:"min=5,max=15"`
}

// request
type UserRegister struct {
	Id                uint64 `json:"id"`
	Name              string `json:"name" validate:"required, min=5, max=50"`
	Password          string `json:"password" validate:"required, min=5 max=15"`
	CredentialsValues string `json:"credentialValues" validate:"required"`
	CredentialsType   Status `json:"credentialType" validate:"required"`
}

type UserLogin struct {
	CredentialsType   Status `json:"credentialType" validate:"required"`
	CredentialsValues string `json:"credentialValues" validate:"required"`
	Password          string `json:"password" validate:"required,min=5, max=15"`
}

type UserResLog struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
