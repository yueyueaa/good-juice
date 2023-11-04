package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"juice/public/data/ent"
	"juice/public/data/ent/userpassword"
	"juice/test/config"

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
	CreateUser(ctx, client)
	find(ctx, client)
	time.Sleep(5 * time.Second)
	update(ctx, client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.UserPassword, error) {
	fmt.Println("Hello Mysql")
	u, err := client.UserPassword.
		Create().
		SetID(123).
		SetUserID(321).
		SetSalt("123123").
		SetPwd("123123123").
		SetCreateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func find(ctx context.Context, client *ent.Client) {
	u, err := client.UserPassword.Query().Where(userpassword.UserIDGT(0)).Select("pwd", "user_id").All(ctx)
	if err != nil {
		fmt.Errorf("failed creating user: %v", err)
		return
	}

	for _, elm := range u {
		fmt.Print(elm.ID)
		fmt.Print(":")
		fmt.Print(elm.Pwd)
		fmt.Print(":")
		fmt.Print(elm.Salt)
		fmt.Print(":")
		fmt.Println(elm.UserID)
	}
}

func update(ctx context.Context, client *ent.Client) {
	client.UserPassword.Update().
		SetPwd("HelloWorld").
		SetUpdateTime(time.Now()).
		Where(userpassword.UserID(321)).
		Save(ctx)
}
