package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"max=100, min=1"`
}
