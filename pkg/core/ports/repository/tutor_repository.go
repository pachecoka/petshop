package repository

import "petshop/pkg/core/domain"

type TutorRepositoryI interface {
	Insert(tutor domain.Tutor) (string, error)
}
