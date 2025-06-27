package main

import (
	"context"
	"db-playground/dao"
	"db-playground/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	db, err := setupDatabase()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
		return
	}
	// Create UserDAO instance
	userDAO := dao.NewUserDaoPGDB(db)

	var users []*model.User
	var userId []string
	for i := 0; i < 1000000; i++ {
		users = append(users, &model.User{
			Id:    uuid.NewString(),
			Name:  fmt.Sprintf("User %d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Age:   20 + i%50,
		})
		userId = append(userId, users[i].Id)
	}
	batchSize := 100000

	//// Measure time for batch insert
	start := time.Now()
	for i := 0; i < len(users); i += batchSize {
		if err := userDAO.BulkCreate(context.Background(), users[i:min(len(users), i+batchSize)]); err != nil {
			//log.Printf("Failed to batch insert users: %v", err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time for batch insert: %s\n", elapsed)

	time.Sleep(10 * time.Second)

	start = time.Now()
	for i := 0; i < len(userId); i += batchSize {
		if err := userDAO.Delete(context.Background(), userId[i:min(len(userId), i+batchSize)]); err != nil {
			log.Printf("Failed to delete users: %v", err)
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Time for batch delete: %s\n", elapsed)

}

func setupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=prajjwal@07 dbname=test_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
