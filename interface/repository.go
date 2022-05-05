package iface

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
}

type WhereOption interface {
	Model
	Where(db *gorm.DB) *gorm.DB
	Preload(db *gorm.DB) *gorm.DB
	IsEmptyWhereOpt() bool
	Page(db *gorm.DB) *gorm.DB
	Sort(db *gorm.DB) *gorm.DB
	WithoutCount() bool
}

type UpdateColumns interface {
	Columns() interface{}
}

type IRepository interface {
	GetDB() *gorm.DB
	List(ctx context.Context, tx *gorm.DB, data interface{}, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (uint64, error)
	GetOne(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, data interface{}, scopes ...func(*gorm.DB) *gorm.DB) error
	Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error
	Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}

type IMongoRepository interface {
	InsertOne(ctx context.Context, collection string, data interface{}) error
	List(ctx context.Context, collection string, data interface{}, filter bson.M, page int64, pageSize int64, opts ...*options.FindOptions) (int64, error)
	ListAll(ctx context.Context, collection string, data interface{}, filter bson.M, opts ...*options.FindOptions) error
}
