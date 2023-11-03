package services

import (
	"errors"
	"petshop/pkg/core/domain"
	"petshop/pkg/core/ports/repository"
	"petshop/pkg/core/ports/service"
)

type petService struct {
	petRepo repository.PetRepositoryI
}

func NewPetService(petRepository repository.PetRepositoryI) service.PetServiceI {
	return &petService{petRepository}
}

func (p petService) Register(pet domain.Pet) (string, error) {
	if pet.Name == "" {
		return "", errors.New("failed to register pet, name is mandatory")
	}
	return p.petRepo.Insert(pet)
}

func (p petService) GetDetails(id string) (domain.Pet, error) {
	return p.petRepo.Find(id)
}

func (p petService) Unregister(id string) error {
	return p.petRepo.Delete(id)
}
