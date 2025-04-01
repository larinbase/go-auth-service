package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/repository"

	"errors"
)

type UserService struct {
	userRepo   *repository.UserRepository
	jwtService JWTService
}

func NewUserService(userRepository *repository.UserRepository, jwtService JWTService) *UserService {
	return &UserService{
		userRepo:   userRepository,
		jwtService: jwtService,
	}
}

func (s *UserService) Register(req *domain.CreateUserRequest) (*dto.JWTResponse, error) {
	// Валидация входных данных
	if err := domain.ValidateEmail(req.Email); err != nil {
		return nil, err
	}
	if err := domain.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := domain.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	generatedToken, err := s.jwtService.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return dto.NewJWTResponse(generatedToken), nil
}

func (s *UserService) Login(req *domain.LoginRequest) (*dto.JWTResponse, error) {

	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.ValidatePassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	generatedToken, err := s.jwtService.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return dto.NewJWTResponse(generatedToken), nil
}
