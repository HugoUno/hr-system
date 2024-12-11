package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"hr-system/internal/models"
	"time"
)

const employeeCacheKey = "employee:%d"

func SetEmployee(ctx context.Context, employee *models.Employee) error {
	key := fmt.Sprintf(employeeCacheKey, employee.ID)
	data, err := json.Marshal(employee)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, key, data, 24*time.Hour).Err()
}

func GetEmployee(ctx context.Context, id uint) (*models.Employee, error) {
	key := fmt.Sprintf(employeeCacheKey, id)
	data, err := RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var employee models.Employee
	err = json.Unmarshal(data, &employee)
	return &employee, err
}
