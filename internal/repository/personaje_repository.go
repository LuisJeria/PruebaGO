package repository

import (
	"context"
	"hola/internal/domain"

	"cloud.google.com/go/firestore"
)

type PersonajeRepository struct {
	collection *firestore.CollectionRef
}

func NewPersonajeRepository(c *firestore.Client) *PersonajeRepository {
	personajeRepository := &PersonajeRepository{}

	personajeRepository.collection = c.Collection("personaje")

	return personajeRepository

}

func (r *PersonajeRepository) Save(data domain.Personaje) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return r.collection.Add(context.Background(), data)
}
