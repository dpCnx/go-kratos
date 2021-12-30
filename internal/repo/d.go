package repo

import (
	"context"

	"go-kratos/internal/model"
)

func (d *Data) Insert(ctx context.Context) error {
	return d.db.WithContext(ctx).Create(&model.D{
		Name: "d",
	}).Error
}

func (d *Data) SetUser(ctx context.Context) (string, error) {
	return d.rds.Set(ctx, "demo", "test", -1).Result()
}
