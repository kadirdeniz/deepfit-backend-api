package pkg

import (
	"time"
)

type Date struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	DeletedAt time.Time `bson:"deleted_at" json:"deleted_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New() *Date {
	return &Date{
		CreatedAt: time.Now(),
		DeletedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (_time *Date) IsUpdated() bool {
	return _time.UpdatedAt.IsZero()
}

func (_time *Date) IsDeleted() bool {
	return _time.DeletedAt.IsZero()
}

func (_time *Date) UpdateTime() {
	_time.UpdatedAt = time.Now()
}

func (_time *Date) DeleteTime() {
	_time.DeletedAt = time.Now()
}

func (_time *Date) CreateTime() {
	_time.CreatedAt = time.Now()
}
