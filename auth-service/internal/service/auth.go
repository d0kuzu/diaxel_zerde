package service

import (
	"auth-service/internal/config"
	"auth-service/internal/crypto"
	"auth-service/internal/jwt"
	"auth-service/internal/repository"
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	users   *repository.UserRepo
	refresh *repository.RefreshRepo
	cfg     *config.Config
}

func NewAuthService(
	users *repository.UserRepo,
	refresh *repository.RefreshRepo,
	cfg *config.Config,
) *AuthService {
	return &AuthService{users, refresh, cfg}
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, ok := s.users.FindByEmail(email)
	if !ok || !crypto.Compare(user.Password, password) {
		return "", "", ErrInvalidCredentials
	}

	access, _ := jwt.GenerateAccessToken(
		user.ID,
		user.Role,
		s.cfg.AccessTokenTTL,
		s.cfg.AccessSecret,
	)

	refreshToken, _ := jwt.GenerateRefreshToken(
		user.ID,
		s.cfg.RefreshTokenTTL,
		s.cfg.RefreshSecret,
	)

	s.refresh.Save(refreshToken, user.ID)

	return access, refreshToken, nil
}

func (s *AuthService) Refresh(token string) (string, string, error) {
	userID, ok := s.refresh.Get(token)
	if !ok {
		return "", "", ErrInvalidCredentials
	}

	sub, err := jwt.ParseRefreshToken(token, s.cfg.RefreshSecret)
	if err != nil || sub != userID {
		return "", "", ErrInvalidCredentials
	}

	s.refresh.Delete(token)
	user, ok := s.users.FindByID(userID)
	if !ok {
		return "", "", ErrInvalidCredentials
	}

	newAccess, _ := jwt.GenerateAccessToken(
		userID,
		user.Role,
		s.cfg.AccessTokenTTL,
		s.cfg.AccessSecret,
	)

	newRefresh, _ := jwt.GenerateRefreshToken(
		userID,
		s.cfg.RefreshTokenTTL,
		s.cfg.RefreshSecret,
	)

	s.refresh.Save(newRefresh, userID)

	return newAccess, newRefresh, nil
}

func (s *AuthService) Logout(token string) {
	s.refresh.Delete(token)
}

func (s *AuthService) Register(email, password string) (string, string, error) {
	_, ok := s.users.FindByEmail(email)
	if ok {
		return "", "", errors.New("user already exists")
	}

	hash, err := crypto.Hash(password)
	if err != nil {
		return "", "", err
	}

	userID := uuid.NewString()
	role := "user"

	ok = s.users.Create(repository.User{
		ID:       userID,
		Email:    email,
		Password: hash,
		Role:     role,
	})
	if !ok {
		return "", "", err
	}

	access, err := jwt.GenerateAccessToken(
		userID,
		role,
		s.cfg.AccessTokenTTL,
		s.cfg.AccessSecret,
	)
	if err != nil {
		return "", "", err
	}

	refresh, err := jwt.GenerateRefreshToken(
		userID,
		s.cfg.RefreshTokenTTL,
		s.cfg.RefreshSecret,
	)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
