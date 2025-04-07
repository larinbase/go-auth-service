package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"github.com/google/uuid"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

/**
 * @author Dmitriy Larin
 **/
type JWTService interface {
	GenerateTokenCouple(user domain.User) (*dto.TokenCoupleResponse, error)
	ExtractUsername(token string) (string, error)
}

type jwtService struct {
	jwtSigningKey []byte
	expiration    uint32
}

func NewJWTService(jwtSigningKey string, expiration uint32) JWTService {
	return &jwtService{
		jwtSigningKey: []byte(jwtSigningKey),
		expiration:    expiration,
	}
}

/**
 * Генерация токена.
 *
 * @param userDetails данные пользователя
 * @return токен
 */
func (s *jwtService) GenerateTokenCouple(user domain.User) (*dto.TokenCoupleResponse, error) {
	roleName := ""
	if user.Role != nil {
		roleName = user.Role.Name.Value()
	}
	claims := jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Duration(s.expiration) * time.Millisecond).Unix(),
		"role":     roleName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.jwtSigningKey)
	if err != nil {
		return dto.NewTokenCoupleResponse("", ""), err
	}

	refreshToken := uuid.New().String()

	return dto.NewTokenCoupleResponse(signedToken, refreshToken), nil
}

/**
 * Проверка токена на валидность
 *
 * @param token
 * @return userDetails
 */
func (s *jwtService) isTokenValid(token string, user domain.User) (bool, error) {
	username, err := s.ExtractUsername(token)
	if err != nil {
		return false, err
	}
	isExpired, err := s.isTokenExpired(token)
	if err != nil {
		return false, err
	}
	return username == user.Email && !isExpired, nil
}

/**
* Проверка токена на просроченность
*
* @param token токен
* @return true, если токен просрочен
 */
func (s *jwtService) isTokenExpired(token string) (bool, error) {
	exp, err := s.extractExpiration(token)
	if err != nil {
		return false, err
	}
	return exp.Before(time.Now()), nil
}

/**
* Извлечение даты истечения токена
*
* @param token токен
* @return дата истечения
 */
func (s *jwtService) extractExpiration(token string) (*jwt.NumericDate, error) {
	claims, err := s.extractAllClaims(token)
	if err != nil {
		return nil, err
	}
	return claims.GetExpirationTime()
}

/**
* Извлечение имени пользователя из токена
*
* @param token токен
* @return имя пользователя
 */
func (s *jwtService) ExtractUsername(token string) (string, error) {
	claims, err := s.extractAllClaims(token)
	if err != nil {
		return "", err
	}

	return claims["username"].(string), nil
}

/**
* Извлечение всех данных из токена
*
* @param token токен
* @return данные
 */
func (s *jwtService) extractAllClaims(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return s.jwtSigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
