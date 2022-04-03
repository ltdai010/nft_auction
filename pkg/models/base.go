package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string          `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatorID uuid.UUID       `json:"creator_id"`
	UpdaterID uuid.UUID       `json:"updater_id"`
	CreatedAt time.Time       `json:"created_at" sql:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `json:"updated_at" sql:"default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" sql:"index"`
}

type ErrorHttp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErr(code int, message string) *ErrorHttp {
	return &ErrorHttp{
		Code:    code,
		Message: message,
	}
}

func (e *ErrorHttp) Error() string {
	return e.Message
}
