package services_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"petshop/pkg/core/domain"
	"petshop/pkg/core/ports/repository"
	"petshop/pkg/core/ports/service"
	"petshop/pkg/core/services"
)

type mockTutorRepository struct {
	tutors map[string]domain.Tutor
}

func NewMockTutorRepository() repository.TutorRepositoryI {
	return &mockTutorRepository{
		tutors: make(map[string]domain.Tutor),
	}
}

func (m mockTutorRepository) Insert(tutor domain.Tutor) (string, error) {
	id := uuid.New().String()
	m.tutors[id] = tutor
	return id, nil
}

var _ = Describe("Tutor", func() {
	var tutorService service.TutorServiceI

	BeforeEach(func() {
		tutorRepo := NewMockTutorRepository()
		tutorService = services.NewTutorService(tutorRepo)
	})

	Describe("registering tutor", func() {
		Context("with valid input", func() {
			It("should register the tutor", func() {
				// when
				id, err := tutorService.Register(domain.Tutor{
					Name:    "Karolina",
					Phone:   123456,
					Address: "test address",
				})

				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(id).NotTo(BeEmpty())
			})
		})

		Context("with missing tutor name", func() {
			It("should not register the pet", func() {
				// when
				id, err := tutorService.Register(domain.Tutor{
					Phone:   123456,
					Address: "test address",
				})

				// then
				Expect(err).To(HaveOccurred())
				Expect(id).To(BeEmpty())
			})
		})
	})
})
