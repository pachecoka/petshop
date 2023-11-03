package service

import "petshop/pkg/core/domain"

type PetServiceI interface {
	Register(pet domain.Pet) (string, error)
}
