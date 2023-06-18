package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/maulana5599/golang-gqlgen/config"
	"github.com/maulana5599/golang-gqlgen/graph/generated"
	"github.com/maulana5599/golang-gqlgen/graph/model"
	"github.com/maulana5599/golang-gqlgen/handler"
	"github.com/maulana5599/golang-gqlgen/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", 111),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateJurusan is the resolver for the createJurusan field.
func (r *mutationResolver) CreateJurusan(ctx context.Context, input model.NewJurusan) (*models.Jurusan, error) {
	jurusan := &models.Jurusan{
		Jurusan: input.Jurusan,
	}
	fmt.Println(jurusan)

	config.DB.Create(jurusan)
	return jurusan, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Link is the resolver for the link field.
func (r *queryResolver) Link(ctx context.Context) ([]*model.Link, error) {
	return r.link, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, baganID float64, juriID float64) ([]*model.User, error) {
	fmt.Println(baganID, juriID)
	result := handler.GetAllUser()

	return result, nil
}

// Customuser is the resolver for the customuser field.
func (r *queryResolver) Customuser(ctx context.Context) ([]*models.Learn, error) {
	result := handler.GetCustomUser()

	return result, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*models.Books, error) {
	var result []*models.Books
	config.DB.Preload("BookDetail").Find(&result)
	return result, nil
}

// Bookdetail is the resolver for the bookdetail field.
func (r *queryResolver) Bookdetail(ctx context.Context) ([]*models.BookDetail, error) {
	var result []*models.BookDetail
	config.DB.Find(&result)
	return result, nil
}

// Jurusan is the resolver for the jurusan field.
func (r *queryResolver) Jurusan(ctx context.Context) ([]*models.Jurusan, error) {
	var result []*models.Jurusan
	config.DB.Preload("JurusanDetail").Find(&result)
	return result, nil
}

// Jurusandetail is the resolver for the jurusandetail field.
func (r *queryResolver) Jurusandetail(ctx context.Context) ([]*models.JurusanDetail, error) {
	var result []*models.JurusanDetail
	config.DB.Find(&result)
	return result, nil
}

// Siswa is the resolver for the siswa field.
func (r *queryResolver) Siswa(ctx context.Context) ([]*models.Siswa, error) {
	var result []*models.Siswa
	config.DB.Find(&result)
	return result, nil
}

// Siswadetail is the resolver for the siswadetail field.
func (r *queryResolver) Siswadetail(ctx context.Context, nis float64) (*models.SiswaDetail, error) {
	var result *models.SiswaDetail
	config.DB.Preload("Abensi").Where("nis = ?", nis).First(&result)
	return result, nil
}

// AbsensiDetail is the resolver for the absensi_detail field.
func (r *siswaDetailResolver) AbsensiDetail(ctx context.Context, obj *models.SiswaDetail) ([]*models.Absensi, error) {
	nis := obj.NIS
	var result []*models.Absensi
	config.DB.Where("nis = ?", nis).Find(&result)
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SiswaDetail returns generated.SiswaDetailResolver implementation.
func (r *Resolver) SiswaDetail() generated.SiswaDetailResolver { return &siswaDetailResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type siswaDetailResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	return r.link, nil
}
