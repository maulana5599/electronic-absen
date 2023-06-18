package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/maulana5599/golang-gqlgen/config"
	"github.com/maulana5599/golang-gqlgen/graph/model"
	"github.com/maulana5599/golang-gqlgen/models"
)

// TambahGuru is the resolver for the tambahGuru field.
func (r *mutationResolver) TambahGuru(ctx context.Context, input *model.TambahGuru) (*models.Guru, error) {
	newGuru := new(models.Guru)
	newGuru.Nip = input.Nip
	newGuru.NamaGuru = input.NamaGuru
	newGuru.TahunMasuk = input.TahunMasuk

	tx := config.DB.Create(&newGuru)
	if tx.Error != nil {
		return newGuru, tx.Error
	}
	return newGuru, nil
}

// Guru is the resolver for the guru field.
func (r *queryResolver) Guru(ctx context.Context) ([]*models.Guru, error) {
	var result []*models.Guru
	config.DB.Find(&result)
	return result, nil
}
