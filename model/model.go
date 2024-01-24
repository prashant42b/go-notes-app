package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Text      string    `json:"text"`
	CreatedOn time.Time `json:"created_on"`
}

func (t Todo) BeforeCreate() {
	t.CreatedOn = time.Now()
}
