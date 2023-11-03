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

func (m mockPetRepository) Find(id string) (domain.Pet, error) {
	pet, exists := m.pets[id]
	if !exists {
		return domain.Pet{}, errors.New("no pet found")
	}

	return pet, nil
}

var _ = Describe("Pet", func() {
	var petService service.PetServiceI

	BeforeEach(func() {
		petRepo := NewMockPetRepository()
		petService = services.NewPetService(petRepo)
	})

	Describe("registering pet", func() {
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

		Context("with missing pet name", func() {
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

	Describe("getting pet's details", func() {
		Context("pet is registered", func() {
			var id string
			BeforeEach(func() {
				var err error
				id, err = petService.Register(domain.Pet{
					Name:      "Kurama",
					BirthDate: "12/09/2014",
					Tutor:     "Karolina",
				})

				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(id).NotTo(BeEmpty())
			})

			It("should return the pet's details", func() {
				// when
				details, err := petService.GetDetails(id)
				Expect(err).NotTo(HaveOccurred())

				// then
				Expect(details.Name).To(Equal("Kurama"))
				Expect(details.Tutor).To(Equal("Karolina"))
			})
		})

		Context("pet is not registered", func() {
			It("should return an error", func() {
				// when
				_, err := petService.GetDetails("1233")

				// then
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
