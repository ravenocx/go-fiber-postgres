package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Books struct {
	ID        string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"` // setting id as primary key
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func (b *Books) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}

// gorm gives ability to auto migrate
// in postgress we need create table manually (it doesn't auto insert like mongodb) -> using automigrate to do that

/*
Auto migrate is a process that automatically handles changes to the database schema
based on changes made to the application's data models or entities.

The auto migrate feature is beneficial for developers because it automates the process of synchronizing
the application's data models with the underlying database schema. It helps maintain the consistency and
integrity of the data while avoiding the need for manual database schema changes, which can be error-prone
and time-consuming.

t's primarily concerned with synchronizing the application's data models (entity definitions)
with the database schema (table definitions).
*/
