package repository

type SubscribeActions interface {
}
type Repository struct {
	SubscribeActions
}

func NewRepository() *Repository {
	return &Repository{}
}
