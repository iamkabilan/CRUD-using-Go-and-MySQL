package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func KeySize(redisConn *redis.Client, key string, ctx context.Context) int64 {
	size, err := redisConn.LLen(ctx, key).Result()
	if err != nil {
		fmt.Println("Error in checking the redis list size, ", err)
		return -1
	}

	return size
}

func RemoveUser(redisConn *redis.Client, key string, ctx context.Context) bool {
	_, err := redisConn.LPop(ctx, key).Result()
	if err != nil {
		fmt.Println("Error in removing the user from the redis user list, ", err)
		return false
	}

	return true
}

func InsertUser(redisConn *redis.Client, key string, user any, ctx context.Context) bool {
	_, err := redisConn.RPush(ctx, key, user).Result()
	if err != nil {
		fmt.Println("Error in inserting the user in the redis user list, ", err)
		return false
	}

	return true
}

func GetUsers(redisConn *redis.Client, key string, ctx context.Context) []string {
	users, err := redisConn.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println("Error in fetching the users, ", err)
		return []string{}
	}

	return users
}

func DeleteKey(redisConn *redis.Client, key string, ctx context.Context) bool {
	_, err := redisConn.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("Error in deleting the key, ", err)
		return false
	}

	return true
}
