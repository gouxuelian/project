package main

import (
	"context"
)

// Resolver is the root resolver
type Resolver struct{}

// GetPet resolves the getPet query
func (r *Resolver) GetEmployee(ctx context.Context, args struct{ Id int32 }) (*Employee, error) {
	employee, err := db.getEmployee(ctx, args.Id)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// petInput has everything needed to do adds and updates on a pet
type employeeInput struct {
	Id                int32
	Empno             string
	Name              string
	Gender            bool
	BirthDate         string
	IdNumber          string
	Nationality       string
	Birthplace        string
	PhoneNumber       string
	WorkDate          string
	Note              string
	Address           string
	Education         string
	Email             string
	ProfessionalTitle string
}

// UpdatePet takes care of updating any field on the pet
func (r *Resolver) UpdateEmployee(ctx context.Context, args struct{ Employee employeeInput }) (*Employee, error) {
	return db.updateEmployee(ctx, &args.Employee)
}
