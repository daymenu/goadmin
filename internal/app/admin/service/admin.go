package service

import (
	"github.com/daymenu/goadmin/api/admin"
	"github.com/daymenu/goadmin/internal/app/admin/biz"
)

// AdminService admin service
type AdminService struct {
	adminUseCase *biz.AdminUseCase
}

// NewAdminService 新建一个 admin service
func NewAdminService(adminUseCase *biz.AdminUseCase) *AdminService {
	return &AdminService{
		adminUseCase: adminUseCase,
	}
}

// Add add a admin
func (adminService *AdminService) Add(admin *admin.Admin) {
	// dto --> adminUsecase

}
