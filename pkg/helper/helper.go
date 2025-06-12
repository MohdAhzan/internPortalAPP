package helper

import (
	"errors"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	response "github/MohdAhzan/internPortalAPP/pkg/utils/Response"
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Helper struct {
	cfg config.Config
}

type AuthCustomClaims struct {
	Id    uuid.UUID   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
  jwt.RegisteredClaims
}

func NewHelper(config config.Config) *Helper {
	return &Helper{
		cfg: config,
	}
}

func (h *Helper) PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *Helper)GenerateAdminToken(admin response.UserDetailsResponse,cfg config.Config) (accessTokenString string, refreshTokenString string, err error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  string(domain.Receptionist),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err = accessToken.SignedString([]byte(cfg.AdminSecret))
	if err != nil {
		return "","", err
	}
  
	refreshTokenClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   admin.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
  refreshTokenString, err = refreshToken.SignedString([]byte(cfg.AdminRefreshSecret)) 
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

func (h *Helper)GenerateReceptionistToken(user response.UserDetailsResponse,cfg config.Config) (accessTokenString string, refreshTokenString string, err error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    user.ID,
		Email: user.Email,
		Role:  string(domain.Receptionist),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err = accessToken.SignedString([]byte(cfg.ReceptionistSecret))
	if err != nil {
		return "","", err
	}

	refreshTokenClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   user.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
  refreshTokenString, err = refreshToken.SignedString([]byte(cfg.ReceptionistRefreshSecret)) 
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

func (h *Helper)GenerateDoctorToken(user response.UserDetailsResponse,cfg config.Config) (accessTokenString string, refreshTokenString string, err error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    user.ID,
		Email: user.Email,
		Role:  string(domain.Doctor),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
  
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err = accessToken.SignedString([]byte(cfg.DoctorSecret))
	if err != nil {
		return "","", err
	}

	refreshTokenClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), 
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   user.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
  refreshTokenString, err = refreshToken.SignedString([]byte(cfg.DoctorRefreshSecret)) 
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil


}


