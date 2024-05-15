package repository

import (
	"errors"
	"fmt"
	"golang-csv/modal"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepoImp(db *gorm.DB) UserRepo {
	return &UserRepository{Db: db}
}

func (c *UserRepository) GetAll(value modal.Responder) (interface{}, error) {
	if err := c.Db.Find(&value).Error; err != nil {
		return nil, err
	}
	if err := c.Db.Delete(value).Error; err != nil {
		return nil, errors.New("no date in this table")
	}
	return value, nil
}

func (c *UserRepository) SetAll(value modal.Responder) error {
	fmt.Println(value)
	switch v := value.(type) {
	case []modal.FullDetails:
		result := c.Db.Create(&v)
		if result.Error != nil {
			return result.Error
		}

	case []modal.EmployeeDetails:
		result := c.Db.Create(&v)
		if result.Error != nil {
			return result.Error
		}

	case []modal.ManagerDetails:
		result := c.Db.Create(&v)
		if result.Error != nil {
			return result.Error
		}

	default:
		return errors.New("unsupported responder type")
	}
	return nil
}
