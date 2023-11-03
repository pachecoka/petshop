package service

import "petshop/pkg/core/domain"

type PetServiceI interface {
	Register(pet domain.Pet) (string, error)
	GetDetails(id string) (domain.Pet, error)
	Unregister(id string) error
}
