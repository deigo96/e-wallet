package roles

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRole(c context.Context, roleID int) (*entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (rr *roleRepository) GetRole(c context.Context, roleID int) (*entity.Role, error) {
	response := &entity.Role{}

	if err := rr.db.Where("id = ?", roleID).First(response).Error; err != nil {
		return nil, constant.ErrInternalServerError
	}

	return response, nil
}
