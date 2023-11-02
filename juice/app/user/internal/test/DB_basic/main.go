package main

import (
	"context"
	"fmt"
	"log"

	"juice/app/user/internal/data/ent"

	"juice/app/user/internal/test/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open(config.DB_driver, config.DB_source)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()
	fmt.Println("Hello World")
	if _, err := CreateUser(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.UserPassword, error) {
	fmt.Println("Hello Mysql")
	u, err := client.UserPassword.
		Create().
		SetID(123).
		SetUserID(123).
		SetSalt("123").
		SetPwd("123").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
