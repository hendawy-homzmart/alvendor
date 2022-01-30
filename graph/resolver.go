package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/El-Hendawy/gograph/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	sellers   []*model.Seller
	admins    []*model.Admin
	customers []*model.Customer
	seller    *model.Seller
	admin     *model.Admin
	customer  *model.Customer
}
