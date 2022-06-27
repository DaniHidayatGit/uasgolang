package usecase

import (
	"notification/models"
	"notification/repository"
)

type UC struct {
	queryrepo repository.Repo
}
type Usecase interface {
	GetDataUser() ([]models.User, error)
	InsertDataUser(models.User) error
	UpdateDataUser(models.User) error
	DeleteDataUser(id int) error
	NotifUser(id int) (models.NotifItemku, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
