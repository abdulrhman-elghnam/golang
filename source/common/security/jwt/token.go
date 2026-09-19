package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/abdulrhman-elghnam/golang/source/configuration"
	"github.com/abdulrhman-elghnam/golang/source/database"
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "ACCESS"
	RefreshToken TokenType = "REFRESH"
)

type SystemRole string

const (
	UserRole  SystemRole = "USER"
	AdminRole SystemRole = "ADMIN"
)

type TokenClaims struct {
	UserID uint       `json:"sub"`
	Role   SystemRole `json:"role"`
	jwt.RegisteredClaims
}

type LoginCredential struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GenerateToken(userID uint) (string, error) {
	if configuration.ACCESS_USER_TOKEN_SECRET == "" {
		return "", errors.New("missing access token secret")
	}

	claims := TokenClaims{
		UserID: userID,
		Role:   UserRole,
	}

	return CreateToken(
		claims,
		configuration.ACCESS_USER_TOKEN_SECRET,
		configuration.ACCESS_USER_TOKEN_EXPIRY,
	)
}

func CreateToken(
	claims TokenClaims,
	secret string,
	expiresIn time.Duration,
) (string, error) {
	now := time.Now()
	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(expiresIn))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(
	tokenString string,
	secret string,
) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims == nil {
		return nil, errors.New("missing token claims")
	}

	if claims.UserID == 0 {
		return nil, errors.New("missing token payload")
	}

	return claims, nil
}

func GetTokenExpiration(tokenString string) (time.Time, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &TokenClaims{})
	if err != nil {
		return time.Time{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || claims.ExpiresAt == nil {
		return time.Time{}, errors.New("token expiration is missing")
	}

	return claims.ExpiresAt.Time, nil
}

func GetTokenSignature(role SystemRole) (string, string, error) {
	switch role {
	case AdminRole:
		return configuration.ACCESS_ADMIN_TOKEN_SECRET,
			configuration.REFRESH_ADMIN_TOKEN_SECRET,
			nil
	case UserRole:
		return configuration.ACCESS_USER_TOKEN_SECRET,
			configuration.REFRESH_USER_TOKEN_SECRET,
			nil
	default:
		return "", "", errors.New("invalid system role")
	}
}

func GetSignatureAccessAndRefresh(role SystemRole, tokenType TokenType) (string, error) {
	accessSecret, refreshSecret, err := GetTokenSignature(role)
	if err != nil {
		return "", err
	}

	switch tokenType {
	case AccessToken:
		return accessSecret, nil
	case RefreshToken:
		return refreshSecret, nil
	default:
		return "", errors.New("invalid token type")
	}
}

func GetToken(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("missing authorization token")
	}

	parts := strings.Split(authorization, " ")
	if len(parts) != 2 {
		return "", errors.New("invalid authorization format")
	}
	if parts[0] != "Bearer" {
		return "", errors.New("invalid authorization format")
	}
	if parts[1] == "" {
		return "", errors.New("missing authorization token")
	}

	return parts[1], nil
}

func CreateLoginCredential(id uint, role SystemRole) (*LoginCredential, error) {
	accessSecret, refreshSecret, err := GetTokenSignature(role)
	if err != nil {
		return nil, err
	}

	accessClaims := TokenClaims{UserID: id, Role: role}
	accessToken, err := CreateToken(accessClaims, accessSecret, configuration.ACCESS_USER_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	refreshClaims := TokenClaims{UserID: id, Role: role}
	refreshToken, err := CreateToken(refreshClaims, refreshSecret, configuration.REFRESH_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	return &LoginCredential{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func DecodeToken(authorization string, tokenType TokenType) (*model.User, *TokenClaims, error) {
	tokenString, err := GetToken(authorization)
	if err != nil {
		return nil, nil, err
	}

	unverifiedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, &TokenClaims{})
	if err != nil {
		return nil, nil, errors.New("invalid token")
	}

	unverifiedClaims, ok := unverifiedToken.Claims.(*TokenClaims)
	if !ok {
		return nil, nil, errors.New("invalid token payload")
	}

	role := unverifiedClaims.Role
	if role != UserRole && role != AdminRole {
		return nil, nil, errors.New("invalid token role")
	}

	secret, err := GetSignatureAccessAndRefresh(role, tokenType)
	if err != nil {
		return nil, nil, err
	}

	claims, err := VerifyToken(tokenString, secret)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, nil, errors.New("token expired")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, nil, errors.New("invalid token signature")
		}
		return nil, nil, errors.New("invalid token")
	}

	if claims.UserID == 0 {
		return nil, nil, errors.New("missing token payload")
	}

	user, err := database.FindUserByID(claims.UserID)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, errors.New("invalid user")
	}
	if SystemRole(user.Role) != claims.Role {
		return nil, nil, errors.New("token role does not match user role")
	}

	return user, claims, nil
}
