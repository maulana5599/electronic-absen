package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/maulana5599/golang-gqlgen/config"
	"github.com/maulana5599/golang-gqlgen/models"
)

// Absensi is the resolver for the absensi field.
func (r *queryResolver) Absensi(ctx context.Context) ([]*models.Absensi, error) {
	var result []*models.Absensi
	config.DB.Find(&result)
	return result, nil
}
