package main

import (
	"context"
	"fmt"

	"github.com/mkock/pklgo/config"
)

func main() {
	cfg, err := config.LoadFromPath(context.Background(), "config/config.pkl")
	if err != nil {
		panic(err)
	}
	fmt.Println("App:", cfg.AppConfig.AppName)
	fmt.Println("Version: ", cfg.AppConfig.AppVersion)
	fmt.Println("Database DSN", cfg.AppConfig.Database.Dsn)
	fmt.Println("Database Driver", cfg.AppConfig.Database.Driver)
	fmt.Println("Database MaxIdleConns", cfg.AppConfig.Database.MaxIdleConns)
}
