package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"context"
	"fmt"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakService struct {
	ServerURL     string
	Realm         string
	ClientID      string
	AdminUsername string
	AdminPassword string
	UserRepo      *repository.UserRepository
}

func NewKeycloakService(
	serverURL, realm, clientID, adminUsername, adminPassword string,
	userRepo *repository.UserRepository,
) *KeycloakService {
	return &KeycloakService{
		ServerURL:     serverURL,
		Realm:         realm,
		ClientID:      clientID,
		AdminUsername: adminUsername,
		AdminPassword: adminPassword,
		UserRepo:      userRepo,
	}
}

func (s *KeycloakService) getAdminClient() *gocloak.GoCloak {
	return gocloak.NewClient(s.ServerURL)
}

func (s *KeycloakService) getAdminToken(ctx context.Context) (*gocloak.JWT, error) {
	client := s.getAdminClient()
	token, err := client.LoginAdmin(ctx, s.AdminUsername, s.AdminPassword, "master")
	if err != nil {
		return nil, fmt.Errorf("failed to get admin token: %w", err)
	}
	return token, nil
}

func (s *KeycloakService) RegisterUser(ctx context.Context, req *domain.CreateUserRequest) error {
	log.Printf("Registering user %s", req.Email)

	client := s.getAdminClient()
	token, err := s.getAdminToken(ctx)
	if err != nil {
		return err
	}

	cred := gocloak.CredentialRepresentation{
		Type:      gocloak.StringP("password"),
		Value:     gocloak.StringP(req.Password),
		Temporary: gocloak.BoolP(false),
	}

	user := gocloak.User{
		Username:    gocloak.StringP(req.Email),
		Email:       gocloak.StringP(req.Email),
		Enabled:     gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{cred},
	}

	_, err = client.CreateUser(ctx, token.AccessToken, s.Realm, user)
	if err != nil {
		return fmt.Errorf("failed to create user in keycloak: %w", err)
	}

	hashedPassword, err := domain.HashPassword(req.Password)
	if err != nil {
		return err
	}

	userSave := &domain.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	err = s.UserRepo.Create(userSave)
	if err != nil {
		return err
	}

	return nil
}
