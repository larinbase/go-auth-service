package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/exception"
	"auth-service/internal/repository"

	"github.com/google/uuid"
	"time"
)

type UserService struct {
	userRepo                 *repository.UserRepository
	jwtService               JWTService
	refreshSessionRepository *repository.RefreshSessionRepository
}

func NewUserService(userRepository *repository.UserRepository, jwtService JWTService, refreshSessionRepository repository.RefreshSessionRepository) *UserService {
	return &UserService{
		userRepo:                 userRepository,
		jwtService:               jwtService,
		refreshSessionRepository: &refreshSessionRepository,
	}
}

func (s *UserService) Register(req *domain.CreateUserRequest) (*dto.TokenCoupleResponse, error) {
	// Валидация входных данных
	if err := domain.ValidateEmail(req.Email); err != nil {
		return nil, exception.InvalidEmail
	}
	if err := domain.ValidatePassword(req.Password); err != nil {
		return nil, exception.InvalidPassword
	}

	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, exception.UserAlreadyExists
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

	tokenCouple, err := s.jwtService.GenerateTokenCouple(*user)
	if err != nil {
		return nil, err
	}

	session := &domain.RefreshSession{}
	session.SetId(tokenCouple.RefreshToken)
	session.SetExpiresIn(time.Now().Add(time.Duration(1440) * time.Hour))

	err = s.refreshSessionRepository.Create(session)
	if err != nil {
		return nil, err
	}

	return tokenCouple, nil
}

func (s *UserService) Login(req *domain.LoginRequest) (*dto.TokenCoupleResponse, error) {

	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, exception.InvalidEmail
	}

	if !user.ValidatePassword(req.Password) {
		return nil, exception.InvalidPassword
	}

	tokenCouple, err := s.jwtService.GenerateTokenCouple(*user)
	if err != nil {
		return nil, err
	}

	session := &domain.RefreshSession{}
	session.SetId(tokenCouple.RefreshToken)
	session.SetExpiresIn(time.Now().Add(time.Duration(1440) * time.Hour))

	err = s.refreshSessionRepository.Create(session)
	if err != nil {
		return nil, err
	}

	return tokenCouple, nil
}

func (s *UserService) RefreshTokens(req *domain.TokenCoupleRequest) (*dto.TokenCoupleResponse, error) {

	email, err := s.jwtService.ExtractUsername(req.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	tokenCouple, err := s.jwtService.GenerateTokenCouple(*user)
	if err != nil {
		return nil, err
	}

	session, err := s.refreshSessionRepository.GetById(uuid.MustParse(req.RefreshToken))
	if err != nil {
		return nil, err
	}
	if session.ExpiresIn.Before(time.Now()) {
		return nil, exception.RefreshTokenIsAlreadyExpired
	}
	session.SetExpiresIn(time.Now().Add(time.Duration(1440) * time.Hour))
	session.SetId(tokenCouple.RefreshToken)

	err = s.refreshSessionRepository.Update(session)
	if err != nil {
		return nil, err
	}

	return tokenCouple, nil
}

func (s *UserService) ChangePassword(email string, req *dto.ChangePasswordRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if !user.ValidatePassword(req.OldPassword) {
		return nil, exception.InvalidPassword
	}

	hashedPassword, err := domain.HashPassword(req.NewPassword)

	user.SetPasswordHash(hashedPassword)

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	user, err = s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return dto.NewUserResponse(
		user.ID,
		user.Email,
		user.Role.Name.Value(),
		user.CreatedAt,
		user.UpdatedAt,
	), nil
}
