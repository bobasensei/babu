package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type Page struct {
	Id        string       `json:"id" gorm:"primaryKey"`
	Document  pgtype.JSONB `gorm:"type:jsonb" json:"document"`
	CreatedAt time.Time    // Automatically managed by GORM for creation time
	UpdatedAt time.Time    // Aut
}