package repository

import (
	"encoding/json"
	"fmt"
	"github.com/ashishkujoy/paper-trading-backend/internal/model"
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
	"time"
)

type UserRepository struct {
	redisClient *redis.Client
}

func NewUserRepository(redisClient *redis.Client) UserRepository {
	return UserRepository{redisClient: redisClient}
}

func (u *UserRepository) GetUserAccount(username string) (model.Account, error) {
	account := model.Account{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	bytes, err := u.redisClient.Get(ctx, fmt.Sprintf("USER_%s", username)).Bytes()
	if err != nil {
		return account, err
	}

	err = json.Unmarshal(bytes, &account)
	return account, err
}

func (u *UserRepository) Save(account model.Account) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	bytes, err := json.Marshal(account)
	if err != nil {
		return err
	}

	return u.redisClient.Set(ctx, fmt.Sprintf("USER_%s", account.Username), bytes, redis.KeepTTL).Err()
}
