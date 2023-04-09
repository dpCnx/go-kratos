package repo

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewTestData(t *testing.T) (*Data, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock new err: %v", err)
	}
	// defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if nil != err {
		t.Fatalf("gorm open err: %v", err)
	}

	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis run err: %v", err)
	}

	rds := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	data := NewData(gormDB, rds)

	return data, mock
}

func TestData_Insert(t *testing.T) {

	data, mock := NewTestData(t)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `d` ").WithArgs("d").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Insert(context.Background())
	if err != nil {
		t.Fatalf("data Insert err: %v", err)
	}

	assert.Nil(t, err)
}

func TestData_SetUser(t *testing.T) {

	data, _ := NewTestData(t)

	_, err := data.SetUser(context.Background())
	assert.Nil(t, err)

	result, err := data.rds.Get(context.Background(), "demo").Result()
	assert.Nil(t, err)

	assert.Equal(t, result,"test")
}
