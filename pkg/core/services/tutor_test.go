package services_test

import (
	"errors"
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

func (m mockTutorRepository) Find(id string) (domain.Tutor, error) {
	tutor, exists := m.tutors[id]
	if !exists {
		return domain.Tutor{}, errors.New("tutor not found")
	}
	return tutor, nil
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

				Expect(err).NotTo(HaveOccurred())
				Expect(id).NotTo(BeEmpty())
			})
		})

		Context("with missing tutor name", func() {
			It("should not register the tutor", func() {
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

	Describe("getting tutor's details", func() {
		Context("tutor is registered", func() {
			var id string
			BeforeEach(func() {
				var err error
				id, err = tutorService.Register(domain.Tutor{
					Name:    "Karolina",
					Phone:   123456,
					Address: "test address",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(id).NotTo(BeEmpty())
			})

			It("should return the tutor's details", func() {
				// when
				details, err := tutorService.GetDetails(id)
				Expect(err).NotTo(HaveOccurred())

				// then
				Expect(details.Name).To(Equal("Karolina"))
			})
		})

		Context("tutor is not registered", func() {
			It("should return an error", func() {
				// when
				_, err := tutorService.GetDetails("1233")

				// then
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
