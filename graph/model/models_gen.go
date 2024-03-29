// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AllCustomersList struct {
	Customers []*Customer `json:"customers"`
}

type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type CustomerRequest struct {
	ID string `json:"id"`
}

type Mutation struct {
}

type NewCustomer struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Query struct {
}

type UpdateCustomer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}
