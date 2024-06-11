package command

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"m1-article-service/application/command/requests"
	"m1-article-service/domain/repository/address/pgx"
	userPgx "m1-article-service/domain/repository/user/pgx"
	"m1-article-service/domain/service/address_user"
	"m1-article-service/infrastructure/godotenv"
	"m1-article-service/infrastructure/log/zerolog"
	pgxInfra "m1-article-service/infrastructure/pgx"
	"os"
	"strconv"
	"strings"
	"time"
)

func Boot(filePath string) {
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()
	conn, err := pgxInfra.SetupPool(env.DATABASE_HOST)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	addressRepo := pgx.NewAddressRepository(env, conn)
	userRepo := userPgx.NewUserRepository(env, conn)
	addrService := address_user.NewService(logger, addressRepo, userRepo)

	for {
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading JSON file: %v", err)
		}

		var users []*requests.User
		if err := json.Unmarshal(jsonData, &users); err != nil {
			log.Fatalf("Error unmarshaling JSON data: %v", err)
		}

		count, err := getCountFromStdin()
		if err != nil {
			log.Fatalf("Error getting count from stdin: %v", err)
		}

		startTime := time.Now()

		// Insert randomly from the users slice
		for i := 0; i < count; i++ {
			user := users[i%len(users)] // Select a user randomly from the users slice
			if err := insert(addrService, user); err != nil {
				logger.Error(err)
			}
		}

		elapsedTime := time.Since(startTime)

		fmt.Printf("Time taken: %s\n", elapsedTime)
	}
}

func insert(addrService *address_user.Service, user *requests.User) error {
	if err := addrService.Create(user.Addresses.ToEntities(),
		user.ToEntity()); err != nil {
		if errors.Is(err, address_user.ErrServiceUnavailable) {
			return insert(addrService, user)
		}
		return err
	}
	return nil
}

func getCountFromStdin() (int, error) {
	fmt.Print("Enter the count number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	countStr := strings.TrimSpace(input)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}
	if count < 0 {
		return 0, errors.New("count number must be positive")
	}
	return count, nil
}
