package service

import "petshop/pkg/core/domain"

type TutorServiceI interface {
	Register(tutor domain.Tutor) (string, error)
	GetDetails(id string) (domain.Tutor, error)
}
