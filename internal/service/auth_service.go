package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"task_ex/internal/auth"
	"task_ex/internal/repository"
	pb "task_ex/service/pb"
	"time"

	"github.com/golang-jwt/jwt/v5"
	goredis "github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	jwtManager  *auth.JWT
	userRepo    *repository.UserRepository
	redisClient *goredis.Client
}

func NewAuthService(userRepo *repository.UserRepository, jwtManager *auth.JWT, redisClient *goredis.Client) *AuthService {
	if redisClient != nil {
		fmt.Println("Using provided Redis client for AuthService")
	}
	return &AuthService{userRepo: userRepo, jwtManager: jwtManager, redisClient: redisClient}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {

	if email == "" || password == "" {
		return "", "", errors.New("invalid credentials")
	}

	if s.userRepo == nil {
		return "", "", errors.New("user repository is not configured")
	}

	user, err := s.userRepo.Login(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Accept both hashed and plain passwords to keep compatibility with existing rows.
	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil || user.Password == password
	if !passwordMatch {
		return "", "", errors.New("invalid credentials")
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)
	accessToken, err := auth.GenerateToken(userID, time.Minute*15)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := auth.GenerateToken(userID, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	// print log
	fmt.Printf("User %s logged in successfully\n", email)

	// store user data and token in redis
	if s.redisClient != nil {
		// Store refresh token with user ID for later validation
		err = s.redisClient.Set(ctx, fmt.Sprintf("refresh_token:%s", refreshToken), userID, time.Hour*24*7).Err()
		if err != nil {
			fmt.Printf("Warning: Failed to store refresh token in Redis: %v\n", err)
		}

		// Store user's current session info
		err = s.redisClient.Set(ctx, fmt.Sprintf("user_session:%s", userID), email, time.Hour*24*7).Err()
		if err != nil {
			fmt.Printf("Warning: Failed to store user session in Redis: %v\n", err)
		}
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.AuthResponse, error) {

	token, err := auth.ValidateToken(req.RefreshToken)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	rawUserID, ok := claims["user_id"]
	if !ok {
		return nil, errors.New("user_id missing in refresh token")

	}

	userID := fmt.Sprintf("%v", rawUserID)

	newAccess, err := auth.GenerateToken(userID, time.Minute*15)
	if err != nil {
		return nil, err
	}

	newRefresh, err := auth.GenerateToken(userID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	// print log
	fmt.Printf("Refreshed token for user_id: %s\n", userID)

	return &pb.AuthResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}
