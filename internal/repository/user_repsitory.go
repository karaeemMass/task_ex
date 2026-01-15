package repository

import (
	"context"

	"gorm.io/gorm"
	"task_ex/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

/*func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
*/

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUser(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {

	var users []*model.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}
