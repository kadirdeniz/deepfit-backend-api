package pkg

import (
	"time"
)

var timeFormat = time.RFC822
var NanoSecond = time.Now().UnixNano()

type TimeStr struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	DeletedAt time.Time `bson:"deleted_at" json:"deleted_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New() TimeStr {
	return TimeStr{
		CreatedAt: time.Now(),
		DeletedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (_time *TimeStr) IsUpdated() bool {
	return _time.UpdatedAt.IsZero()
}

func (_time *TimeStr) IsDeleted() bool {
	return _time.DeletedAt.IsZero()
}

func (_time *TimeStr) UpdateTime() {
	_time.UpdatedAt = time.Now()
}

func (_time *TimeStr) DeleteTime() {
	_time.DeletedAt = time.Now()
}

func (_time *TimeStr) CreateTime() {
	_time.CreatedAt = time.Now()
}
