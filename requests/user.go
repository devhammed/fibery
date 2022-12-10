package requests

type CreateUserRequest struct {
	Name string `validate:"required,min=2"`
}
