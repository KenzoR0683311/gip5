package db

import (
    "context"
    "time"
    "github.com/go-redis/redis/v8"
    "testProject/models"
    "encoding/json"
)

var redisClient *redis.Client

func init() {
    // Initialize Redis client
    redisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis address
        Password: "",               // No password set
        DB:       0,                // Use default DB
    })
}

// Function to save session data in Redis
func SaveSession(sessionToken string, sessionData interface{}, expiry time.Time) error {
    pBytes, err := json.Marshal(sessionData)
    ctx := context.Background()
    err = redisClient.Set(ctx, sessionToken, pBytes, expiry.Sub(time.Now())).Err()
    if err != nil {
      return err
    }
    return nil
}

// Function to retrieve session data from Redis
func GetSession(sessionToken string) (models.Session, error) {
    var sess models.Session
    ctx := context.Background()
    sessionData, err := redisClient.Get(ctx, sessionToken).Result()
    if err != nil {
        return sess, err
    }

    err = json.Unmarshal([]byte(sessionData), &sess)
    return sess, nil
}

// Function to delete session data from Redis
func DeleteSession(sessionToken string) error {
    ctx := context.Background()
    err := redisClient.Del(ctx, sessionToken).Err()
    if err != nil {
        return err
    }
    return nil
}
