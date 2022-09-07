package services

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"errors"
)

type KategoriServices struct {
	Repository repository.KategoriRepository
}

func (services KategoriServices) Get(id string) (*entity.Kategori, error) {
	kategori := services.Repository.FindById(id)
	if kategori == nil {
		return kategori, errors.New("category not found")
	} else {
		return kategori, nil
	}
}
