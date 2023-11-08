package repository

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MainAppEntityMock struct {
	Mock mock.Mock
}

func (mam *MainAppEntityMock) Create(model interface{}) *gorm.DB {
	args := mam.Mock.Called(model)
	return args.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Table(name string) *gorm.DB {
	argss := mam.Mock.Called(name)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Joins(query string) *gorm.DB {
	argss := mam.Mock.Called(query)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Select(query interface{}) *gorm.DB {
	argss := mam.Mock.Called(query)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Where(query interface{}, args ...interface{}) *gorm.DB {
	argss := mam.Mock.Called(query, args)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Find(dest interface{}) *gorm.DB {
	argss := mam.Mock.Called(dest)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) First(query interface{}) *gorm.DB {
	argss := mam.Mock.Called(query)
	return argss.Get(0).(*gorm.DB)
}

func (mam *MainAppEntityMock) Save(dest interface{}) *gorm.DB {
	args := mam.Mock.Called(dest)
	return args.Get(0).(*gorm.DB)
}
