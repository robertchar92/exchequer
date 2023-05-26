package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"exchequer/lib/database_transaction"
	"exchequer/utils/validators"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/subosito/gotenv"
	"go.uber.org/fx"

	"exchequer/app/middleware"
	"exchequer/app/server"

	authHTTP "exchequer/service/auth/delivery/http"
	authModule "exchequer/service/auth/module"

	userHTTP "exchequer/service/user/delivery/http"
	userModule "exchequer/service/user/module"

	bookHTTP "exchequer/service/book/delivery/http"
	bookModule "exchequer/service/book/module"
)

type libs struct {
	fx.Out

	// Redis              redis.Client
	TransactionManager database_transaction.Client
}

type handlers struct {
	fx.In

	// OhlcHandler *exchequerHTTP.Handler
	AuthHandler *authHTTP.Handler
	UserHandler *userHTTP.Handler
	BookHandler *bookHTTP.Handler
}

func main() {
	log.Println("server is starting")

	loadEnv()

	// set log to file
	if os.Getenv("APP_ENV") != "development" {
		log.Println("running in ", os.Getenv("APP_ENV"), " environment")
		f, err := os.OpenFile("error-log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()

		//set output of logs to f
		log.SetOutput(f)
	}

	app := fx.New(
		fx.Provide(
			setupDatabase,
			initLibs,
		),
		authModule.Module,
		userModule.Module,
		bookModule.Module,
		fx.Invoke(
			validators.NewValidator,
			startServer,
		),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, db *gorm.DB, handlers handlers) {
	m := middleware.New(middleware.Config{
		Db: db,
	})

	h := server.BuildHandler(m,
		handlers.AuthHandler,
		handlers.UserHandler,
		handlers.BookHandler,
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(s *http.Server) {
				log.Printf("api is available at %s\n", s.Addr)
				if err := s.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}(s)
			return nil
		},
		OnStop: func(c context.Context) error {
			_ = s.Shutdown(c)
			log.Println("api gracefully stopped")
			return nil
		},
	})
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		log.Println("failed to load from .env")
	}
}

func setupDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dbDriver := os.Getenv("DATABASE_DRIVER")

	if dbDriver == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_PORT"),
		)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println(err)
			panic(err)
		}
	} else if dbDriver == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true",
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"),
		)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	sqlDB, _ := db.DB()

	// Set the maximum number of concurrently idle connections. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	sqlDB.SetMaxIdleConns(50)

	// Set the number of open connections (in-use + idle).
	sqlDB.SetMaxOpenConns(50)

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func initLibs(lc fx.Lifecycle, db *gorm.DB) libs {
	l := libs{
		// Redis: redis.NewClient(redis.Credentials{
		// 	Host:     os.Getenv("REDIS_HOST"),
		// 	Port:     os.Getenv("REDIS_PORT"),
		// 	Password: os.Getenv("REDIS_PASSWORD"),
		// }, os.Getenv("APP_ENV")),
		TransactionManager: database_transaction.New(db),
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			// _ = l.Redis.Close()

			return nil
		},
	})

	return l
}
