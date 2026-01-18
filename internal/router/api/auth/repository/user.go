package repository

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = code.UserNotFound
	ErrUserAlreadyExists = code.UserAlreadyExistsOrEmailAlreadyUsed
	ErrDuplicateKey      = code.DatabaseError
)

type UserRepository interface {
	FindByID(id int64) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	ExistsByUsernameOrEmail(username, email string) (bool, error)
	DeleteByID(id int64) error
	List(offset, limit int) ([]*model.User, error)
	Count() (int64, error)
	CheckUnique(username, email string) (bool, string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CheckUnique(username, email string) (bool, string, error) {
	var existingUsers []model.User
	if err := r.db.Where("username = ? OR email = ?", username, email).
		Find(&existingUsers).Error; err != nil {
		return false, "", err
	}

	if len(existingUsers) == 0 {
		return true, "", nil
	}

	for _, user := range existingUsers {
		if user.Username == username {
			return false, "username", errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		}
		if user.Email == email {
			return false, "email", errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		}
	}

	return false, "unknown", errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
}

func (r *userRepository) Create(user *model.User) error {
	unique, conflictField, err := r.CheckUnique(user.Username, user.Email)
	if err != nil {
		return err
	}

	if !unique {
		if conflictField == "username" {
			return errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		} else if conflictField == "email" {
			return errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		}
		return errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
	}

	if err := r.db.Create(user).Error; err != nil {
		return r.handleCreateError(err)
	}

	return nil
}

func (r *userRepository) handleCreateError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "23505") {
		return errors.New(strconv.Itoa(code.DatabaseError))
	}

	return err
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(strconv.Itoa(code.UserNotFound))
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(strconv.Itoa(code.UserNotFound))
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "user_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(strconv.Itoa(code.UserNotFound))
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	if user.Username == "" && user.Email == "" {
		return r.db.Save(user).Error
	}

	var existingUser model.User
	query := r.db.Where("user_id != ?", user.UserID)

	if user.Username != "" {
		query = query.Or("username = ?", user.Username)
	}
	if user.Email != "" {
		query = query.Or("email = ?", user.Email)
	}

	if err := query.First(&existingUser).Error; err == nil {
		if existingUser.Username == user.Username {
			return errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		}
		if existingUser.Email == user.Email {
			return errors.New(strconv.Itoa(code.UserAlreadyExistsOrEmailAlreadyUsed))
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.db.Save(user).Error
}

func (r *userRepository) ExistsByUsernameOrEmail(username, email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).
		Where("username = ? OR email = ?", username, email).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) DeleteByID(id int64) error {
	result := r.db.Where("user_id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New(strconv.Itoa(code.UserNotFound))
	}
	return nil
}

func (r *userRepository) List(offset, limit int) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
