package repository

import (
	"belajar-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type KategoriRepositoryMock struct {
	Mock mock.Mock
}

func (reposit *KategoriRepositoryMock) FindById(id string) *entity.Kategori {
	argumen := reposit.Mock.Called(id)
	if argumen.Get(0) == nil {
		return nil
	}

	hasil := argumen.Get(0).(entity.Kategori)
	return &hasil
}
