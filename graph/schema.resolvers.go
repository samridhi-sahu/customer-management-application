package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"log"

	"github.com/samridhi-sahu/cms/database"
	"github.com/samridhi-sahu/cms/graph/model"
	pb "github.com/samridhi-sahu/cms/protobuf"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	customer, err := database.Client.CreateCustomer(ctx,
		&pb.NewCustomer{Name: input.Name, City: input.City})
	if err != nil {
		log.Println("error in creating customer ", err)
	}
	newCustomer := &model.Customer{ID: customer.Id, Name: customer.Name, City: customer.City}
	return newCustomer, nil
}

// UpdateCustomer is the resolver for the updateCustomer field.
func (r *mutationResolver) UpdateCustomer(ctx context.Context, input model.UpdateCustomer) (*model.Customer, error) {
	customer, err := database.Client.UpdateCustomer(ctx,
		&pb.Customer{Id: input.ID, Name: input.Name, City: input.City})
	if err != nil {
		log.Println("error in updating customer ", err)
	}
	updateCustomer := &model.Customer{ID: customer.Id, Name: customer.Name, City: customer.City}
	return updateCustomer, nil
}

// DeleteCustomer is the resolver for the deleteCustomer field.
func (r *mutationResolver) DeleteCustomer(ctx context.Context, input model.CustomerRequest) (*model.Customer, error) {
	_, err := database.Client.DeleteCustomer(ctx, &pb.CustomerRequest{Id: input.ID})
	if err != nil {
		log.Println("error in deleting customer")
	}
	return &model.Customer{}, nil
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, input model.CustomerRequest) (*model.Customer, error) {
	customer, err := database.Client.GetCustomer(ctx, &pb.CustomerRequest{Id: input.ID})
	if err != nil {
		log.Println("error in getting customer ", err)
	}
	getCustomer := &model.Customer{ID: customer.Id, Name: customer.Name, City: customer.City}
	return getCustomer, nil
}

// GetAllCustomers is the resolver for the getAllCustomers field.
func (r *queryResolver) GetAllCustomers(ctx context.Context) (*model.AllCustomersList, error) {
	allCustomersList := &model.AllCustomersList{}
	list, err := database.Client.GetAllCustomers(ctx, &pb.Request{})
	if err != nil {
		log.Println("error in getting all customers")
	}
	for _, value := range list.Customers {
		customer := &model.Customer{ID: value.Id, Name: value.Name, City: value.City}
		allCustomersList.Customers = append(allCustomersList.Customers, customer)
	}
	return allCustomersList, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
