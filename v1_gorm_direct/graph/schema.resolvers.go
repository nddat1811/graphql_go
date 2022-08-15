package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"v1_gorm_direct/config"
	"v1_gorm_direct/controller"
	"v1_gorm_direct/graph/generated"
	"v1_gorm_direct/graph/model"
	model2 "v1_gorm_direct/model"
	"v1_gorm_direct/service"
)

// CreateNewUser is the resolver for the createNewUser field.
func (r *mutationResolver) CreateNewUser(ctx context.Context, input model.NewUser) (*model.Respone, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	DB := config.DB
	userRepository := model2.NewUserModel(DB)
	ctr := controller.NewUserController(userRepository)
	re := ctr.GetAllUsers()
	return re, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input model.UserLogin) (*model.LoginResponse, error) {
	DB := config.DB
	userRepository := model2.NewUserModel(DB)
	ctr := controller.NewAuthController(userRepository, service.NewAuthService())

	re := ctr.Login(input)
	return re, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
