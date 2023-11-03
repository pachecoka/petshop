package repository

import "petshop/pkg/core/domain"

type PetRepositoryI interface {
	Insert(pet domain.Pet) (string, error)
}
