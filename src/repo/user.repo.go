package repo

import (
	"context"
	"encoding/json"
	"time"

	"goboilerplate.com/src/models"
	"goboilerplate.com/src/pkg/database"
	"goboilerplate.com/src/pkg/redis"
)

const userCachePrefix = "user:email:"
const userCacheTTL = 10 * time.Minute

type IUserRepo interface {
	CreateUser(ctx context.Context, opt models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepo struct {
	db    database.IDatabase
	redis redis.IRedis
}

func NewUserRepo(db database.IDatabase, redis redis.IRedis) *UserRepo {
	return &UserRepo{db: db, redis: redis}
}

func (r *UserRepo) CreateUser(ctx context.Context, opt models.User) (*models.User, error) {
	if err := r.db.Create(ctx, "users", &opt); err != nil {
		return nil, err
	}
	// Invalidate cache after create
	if err := r.redis.Del(ctx, userCachePrefix+opt.Email); err != nil {
		return nil, err
	}
	return &opt, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	cacheKey := userCachePrefix + email

	if cached, err := r.redis.Get(ctx, cacheKey); err == nil {
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return user, nil
		}
	}

	if err := r.db.First(ctx, "users", database.Filter{"email": email}, &user); err != nil {
		return nil, err
	}

	if data, err := json.Marshal(user); err == nil {
		if err := r.redis.Set(ctx, cacheKey, string(data), userCacheTTL); err != nil {
			return nil, err
		}
	}
	return user, nil
}
