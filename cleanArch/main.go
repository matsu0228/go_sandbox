package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" //mssql接続のためのドライバ
	"github.com/joho/godotenv"
	cpRepos "github.com/matsu0228/go_sandbox/cleanArch/campaign/repositories"
	"github.com/matsu0228/go_sandbox/cleanArch/campaign/usecase"
	commonRepo "github.com/matsu0228/go_sandbox/cleanArch/common/repositories"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	envLoad()
	db, err := commonRepo.NewDB(
		os.Getenv("server"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("databasename"),
	)
	if err != nil {
		log.Fatal("[Error] creating connection pool:", err)
	}
	// defer db.Close()

	cr := cpRepos.NewMssqlRepositories(db)
	cu := usecase.NewCampaignUsecase(cr)
	err = cu.Set(2915, "updated title test")
	if err != nil {
		log.Fatal("[Error] cient.Set:", err)
	}

	fmt.Printf("finish")
}
