package userApi

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"open-api-backend/database"
)

type UserApi struct {
}

var db *gorm.DB

var client *redis.Client
var ctx context.Context

func init() {
	db = database.DB
	client = database.Client
	ctx = database.Ctx
}
