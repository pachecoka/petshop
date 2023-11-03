package services

import (
	"errors"
	"petshop/pkg/core/domain"
	"petshop/pkg/core/ports/repository"
	"petshop/pkg/core/ports/service"
)

type tutorService struct {
	tutorRepo repository.TutorRepositoryI
}

func NewTutorService(tutorRepo repository.TutorRepositoryI) service.TutorServiceI {
	return &tutorService{tutorRepo: tutorRepo}
}

func (t tutorService) Register(tutor domain.Tutor) (string, error) {
	if tutor.Name == "" {
		return "", errors.New("failed to register tutor, name is mandatory")
	}
	return t.tutorRepo.Insert(tutor)
}

func (t tutorService) GetDetails(id string) (domain.Tutor, error) {
	return t.tutorRepo.Find(id)
}
