package model

import "time"

type User struct {
	Id          uint          `json:"id" gorm:"primaryKey"`
	Email       string        `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password    string        `json:"password" gorm:"not null" validate:"required,min=8"`
	Name        string        `json:"name" gorm:"not null" validate:"required"`
	EmailRecord []EmailRecord `json:"emailRecord,omitempty" gorm:"foreignKey:UserID"`
}

type Status string

const (
	Pending   Status = "pending"
	Delivered Status = "delivered"
	Failed    Status = "failed"
)

type EmailSendingType string

const (
	WELCOM EmailSendingType = "welcome"
	UPDATE EmailSendingType = "update"
	NOTIFY EmailSendingType = "notify"
	PROMOTION EmailSendingType = "promotion"
	OTHER EmailSendingType = "other"
	LAUNCH EmailSendingType = "launch"
)

type EmailRecord struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null" validate:"required"`
	User      User      `gorm:"foreignKey:UserID"`
	SendTo    string    `json:"send_to" gorm:"not null" validate:"required,email"`
	EmailType EmailSendingType `json:"email_type" gorm:"not null;default:'other'" validate:"oneof=welcome update notify promotion other launch"`
	Status    Status    `json:"status" gorm:"not null;default:'pending'" validate:"oneof=pending delivered failed"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
