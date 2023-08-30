package repository

type AppUser interface {
}

type AppSegment interface {
}

type Repository struct {
	AppUser
	AppSegment
}

func NewRepository() *Repository {
	return &Repository{}
}
