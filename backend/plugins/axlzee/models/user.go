package models

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/common"
)

type AxlzeeUser struct {
	common.NoPKModel `json:"-"`

	ConnectionId uint64 `gorm:"primaryKey"`
	UserId       string `json:"user_id" gorm:"primaryKey"` // A new unique field

	Gender       string `json:"gender"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	Phone        string `json:"phone"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (AxlzeeUser) TableName() string {
	return "_tool_axlzee_users"
}
