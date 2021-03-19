package service

import (
	"errors"
	"hola/internal/domain"
	"hola/internal/repository"
)

type PersonajeService struct {
	repository *repository.PersonajeRepository
}

func NewPersonajeService(r *repository.PersonajeRepository) *PersonajeService {
	personajeService := &PersonajeService{}
	personajeService.repository = r
	return personajeService
}

func (s *PersonajeService) Save(data domain.Personaje) error {
	if data.Nombre == "ctm" {
		return errors.New("Tu personaje no puede tener ese nombre")
	}

	s.repository.Save(data)
	return nil

}
