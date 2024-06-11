package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/application/http/handlers"
	"m1-article-service/domain/repository/address/pgx"
	userPgx "m1-article-service/domain/repository/user/pgx"
	"m1-article-service/domain/service/address_user"
	"m1-article-service/infrastructure/godotenv"
	"m1-article-service/infrastructure/log/zerolog"
)

func Boot() {
	r := gin.Default()
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()
	conn, err := pgxpool.New(context.Background(), env.DATABASE_HOST)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	addressRepo := pgx.NewAddressRepository(env, conn)
	userRepo := userPgx.NewUserRepository(env, conn)
	addrService := address_user.NewService(logger, addressRepo, userRepo)
	addressUser := handlers.NewAddressUser(addrService)

	//I must define router struct but for lack of time I call handler(controller) directly
	r.POST("/api/users", addressUser.CreateUser)
	r.GET("/api/users/:id", addressUser.DetailUser)

	r.Run()
}
