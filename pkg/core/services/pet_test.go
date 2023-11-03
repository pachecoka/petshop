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

type mockPetRepository struct {
	pets map[string]domain.Pet
}

func NewMockPetRepository() repository.PetRepositoryI {
	return &mockPetRepository{
		pets: make(map[string]domain.Pet),
	}
}

func (m mockPetRepository) Insert(pet domain.Pet) (string, error) {
	id := uuid.New().String()
	m.pets[id] = pet
	return id, nil
}

var _ = Describe("Pet", func() {
	var petService service.PetServiceI

	BeforeEach(func() {
		petRepo := NewMockPetRepository()
		petService = services.NewPetService(petRepo)
	})

	Describe("Registering Pet", func() {
		Context("with valid input", func() {
			It("should register the pet", func() {
				// when
				id, err := petService.Register(domain.Pet{
					Name:      "Kurama",
					BirthDate: "12/09/2014",
					Tutor:     "Karolina",
				})

				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(id).NotTo(BeEmpty())
			})
		})

		Context("with imissing pet name", func() {
			It("should not register the pet", func() {
				// when
				id, err := petService.Register(domain.Pet{
					BirthDate: "12/09/2014",
					Tutor:     "Karolina",
				})

				// then
				Expect(err).To(HaveOccurred())
				Expect(id).To(BeEmpty())
			})
		})
	})
})
