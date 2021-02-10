package repository

import "github.com/daymenu/goadmin/internal/ent"

// AdminRepo define a admin repository
type AdminRepo struct {
	client *ent.Client
}

// NewAdminRepo new a amdin repository
func NewAdminRepo(client *ent.Client) (*AdminRepo, error) {
	return nil, nil
}

// Add add a admin
func (adminRepo *AdminRepo) Add(adminEntiy *ent.Admin) {

}
