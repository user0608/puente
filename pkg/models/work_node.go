package models

import "encoding/json"

type WorkNode struct {
	ID          string          `json:"id" gorm:"primaryKey"`
	Description string          `json:"description" gorm:"not null"`
	Prefix      string          `json:"prefix" gorm:"not null;unique;uniqueIndex"`
	Props       json.RawMessage `json:"props" gorm:"type:json;default:'{}'"`
}
