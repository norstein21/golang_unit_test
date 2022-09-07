package services

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var kategoriRepository = &repository.KategoriRepositoryMock{Mock: mock.Mock{}}
var kategoriServices = KategoriServices{Repository: kategoriRepository}

func TestCategoryServices_GetNotFound(t *testing.T) {
	kate := entity.Kategori{
		Id:   "2",
		Name: "Habie",
	}
	// ini untuk return nill kategoriRepository.Mock.On("FindById", "1").Return(nil)
	kategoriRepository.Mock.On("FindById", "2").Return(kate)
	//kategoriRepository.Mock.On("FindById", "2").Return(kate)
	hasil, err := kategoriServices.Get("2")
	assert.NotNil(t, hasil)
	assert.Nil(t, err) //kebalik naro parameter di NIL YANG BIKIN ERROR

	assert.Equal(t, kate.Id, hasil.Id, "Harus sama")
	assert.Equal(t, kate.Name, hasil.Name, "Harus sama")
}
