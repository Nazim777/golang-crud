package models

import "time"

type Todo struct {
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
