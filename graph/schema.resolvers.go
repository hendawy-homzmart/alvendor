package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/El-Hendawy/gograph/auth"
	"github.com/El-Hendawy/gograph/graph/generated"
	"github.com/El-Hendawy/gograph/graph/model"
	"github.com/El-Hendawy/gograph/graph/repository"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateAdmin(ctx context.Context, input model.NewAdmin) (*model.Admin, error) {
	admin := &model.Admin{
		ID: strconv.Itoa(rand.Int()),
		UserInfo: &model.User{
			ID:       strconv.Itoa(rand.Int()),
			Email:    input.UserInfo.Email,
			Password: HashPassword(input.UserInfo.Password),
			Mobile:   input.UserInfo.Mobile,
			IsActive: input.UserInfo.IsActive,
			Role: &model.Role{
				ID:   input.UserInfo.Role.ID,
				Name: input.UserInfo.Role.Name,
				Ability: []*model.Abilities{
					&model.Abilities{
						ID:      strconv.Itoa(rand.Int()),
						Action:  input.UserInfo.Role.Ability.Action,
						Subject: input.UserInfo.Role.Ability.Subject,
					},
				},
			},
		},
		UserID: strconv.Itoa(rand.Int()),
		// Manages: []*model.Seller{
		// 	&model.Seller{
		// 		ID: input.Addresses.ID,
		// 	},
		// },
	}

	adminRepo.SaveAdmin(admin)
	return admin, nil
}

func (r *mutationResolver) CreateSeller(ctx context.Context, input model.NewSeller) (*model.Seller, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Seller{}, fmt.Errorf("access denied")
	}

	seller := &model.Seller{
		ID: strconv.Itoa(rand.Int()),
		UserInfo: &model.User{
			ID: strconv.Itoa(rand.Int()),
			FirstName: &model.FirstName{
				Ar: input.UserInfo.FirstName.Ar,
				En: input.UserInfo.FirstName.En,
			},
			LastName: &model.LastName{
				Ar: input.UserInfo.LastName.Ar,
				En: input.UserInfo.LastName.En,
			},
			Email:    input.UserInfo.Email,
			Password: HashPassword(input.UserInfo.Password),
			Role: &model.Role{
				ID:   input.UserInfo.Role.ID,
				Name: input.UserInfo.Role.Name,
				Ability: []*model.Abilities{
					&model.Abilities{
						ID:      strconv.Itoa(rand.Int()),
						Action:  input.UserInfo.Role.Ability.Action,
						Subject: input.UserInfo.Role.Ability.Subject,
					},
				},
			},
		},
		UserID: strconv.Itoa(rand.Int()),
		Addresses: []*model.Address{
			&model.Address{
				ID:           input.Addresses.ID,
				FloorNo:      input.Addresses.FloorNo,
				AppartmentNo: input.Addresses.AppartmentNo,
				LandMark:     input.Addresses.LandMark,
				Street:       input.Addresses.Street,
				City:         input.Addresses.City,
				State:        input.Addresses.State,
				Country:      input.Addresses.Country,
				Coordinates: &model.Coordinates{
					Lat: input.Addresses.Coordinates.Lat,
					Lng: input.Addresses.Coordinates.Lng,
				},
			},
		},

		Shops: []*model.Shops{
			&model.Shops{
				ID: strconv.Itoa(rand.Int()),
				Name: &model.BusinessName{
					Ar: input.Shops.Name.Ar,
					En: input.Shops.Name.En,
				},
				Address: &model.Address{
					ID: input.Shops.Address.ID,

					FloorNo: input.Shops.Address.FloorNo,

					AppartmentNo: input.Shops.Address.AppartmentNo,

					LandMark: input.Shops.Address.LandMark,

					Street: input.Shops.Address.Street,

					City: input.Shops.Address.City,

					State: input.Shops.Address.State,

					Country: input.Shops.Address.Country,

					Coordinates: &model.Coordinates{

						Lat: input.Shops.Address.Coordinates.Lat,

						Lng: input.Shops.Address.Coordinates.Lng,
					},
				},
			},
		},
	}

	sellerRepo.Save(seller)
	return seller, nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	customer := &model.Customer{
		ID: strconv.Itoa(rand.Int()),

		UserInfo: &model.User{
			ID: strconv.Itoa(rand.Int()),
			FirstName: &model.FirstName{
				Ar: input.UserInfo.FirstName.Ar,
				En: input.UserInfo.FirstName.En,
			},
			LastName: &model.LastName{
				Ar: input.UserInfo.LastName.Ar,
				En: input.UserInfo.LastName.En,
			},
			Email:    input.UserInfo.Email,
			Password: HashPassword(input.UserInfo.Password),
			Role: &model.Role{
				ID:   input.UserInfo.Role.ID,
				Name: input.UserInfo.Role.Name,
				Ability: []*model.Abilities{
					&model.Abilities{
						ID:      strconv.Itoa(rand.Int()),
						Action:  input.UserInfo.Role.Ability.Action,
						Subject: input.UserInfo.Role.Ability.Subject,
					},
				},
			},
		},
		UserID: strconv.Itoa(rand.Int()),
		Addresses: []*model.Address{
			&model.Address{
				ID:           input.Addresses.ID,
				FloorNo:      input.Addresses.FloorNo,
				AppartmentNo: input.Addresses.AppartmentNo,
				LandMark:     input.Addresses.LandMark,
				Street:       input.Addresses.Street,
				City:         input.Addresses.City,
				State:        input.Addresses.State,
				Country:      input.Addresses.Country,
				Coordinates: &model.Coordinates{
					Lat: input.Addresses.Coordinates.Lat,
					Lng: input.Addresses.Coordinates.Lng,
				},
			},
		},
	}

	customerRepo.SaveCustomer(customer)
	return customer, nil
}

func (r *mutationResolver) LoginAdmin(ctx context.Context, input model.NewLogin) (*model.UserData, error) {
	panic(fmt.Errorf("not implemented"))
	// seller := &model.Login{
	// 	Email: input.Email,
	// 	Password: input.Password,
	// }
	// sellerRepo.AuthenticateSeller(seller)
}

func (r *mutationResolver) LoginSeller(ctx context.Context, input model.NewLogin) (*model.UserData, error) {
	//panic(fmt.Errorf("not implemented"))
	seller := &model.Login{
		Email:    input.Email,
		Password: input.Password,
	}
	token, err := sellerRepo.AuthenticateSeller(seller)

	return token, err
}

func (r *mutationResolver) LoginCustomer(ctx context.Context, input model.NewLogin) (*model.UserData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	//panic(fmt.Errorf("not implemented"))
	product := &model.Product{

		ProductName: &model.Name{
			Ar: input.ProductName.Ar,
			En: input.ProductName.En,
		},
		Unit:   input.Unit,
		Sku:    input.Sku,
		MinQty: input.MinQty,
		MaxQty: input.MaxQty,
		Status: input.Status,
		
		ProductImages: &model.ProductImages{
			Thumbnail: input.ProductImages.Thumbnail,
			Gallery: &model.ProductGallery{
				[]*model.Image{
					&model.Image{
						ID:   input.ProductImages.Gallery.Image.ID,
						Path: input.ProductImages.Gallery.Image.Path,
						Name: input.ProductImages.Gallery.Image.Name,
					},
					
				},
			},
		},
		// Variant: []*model.Variant{
		// 	&model.Variant{
		// 		Name: input.Variant,
		// 		Value: inp,
		// 	},
		// },

		Alias:input.Alias,
		Warranty: input.Warranty,
		Discount: input.Discount,
		Seo: &model.SeoData{
			Title: input.Seo.Title,
			Desc: input.Seo.Desc,
			Image: &model.Image{
				ID: input.Seo.Image.ID,
				Path: input.Seo.Image.Path,
				Name: input.Seo.Image.Name,

			},
		},
		Brand: &model.ProductBrand{
			ID: input.Brand.ID,
			Name: input.Brand.Name,
			Image: &model.Image{
				ID: input.Brand.Image.ID,
				Path: input.Brand.Image.Path,
				Name: input.Brand.Image.Name,

			},
		},
		Categories: &model.ProductCategories{
			ID: input.Categories.ID,
		},

	}

	productRepo.SaveProduct(product)

	return product, nil
}

func (r *queryResolver) Seller(ctx context.Context) ([]*model.Seller, error) {
	//	panic(fmt.Errorf("not implemented"))
	return sellerRepo.FindAll(), nil
}

func (r *queryResolver) Admin(ctx context.Context) ([]*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var sellerRepo repository.SellerRepository = repository.New()
var customerRepo repository.CustomerRepository = repository.NewCustomer()
var adminRepo repository.AdminRepository = repository.NewAdmin()
var productRepo repository.ProductRepository = repository.NewProduct()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
