package repository

import "belajar-unit-test/entity"

type KategoriRepository interface {
	FindById(id string) *entity.Kategori
}
