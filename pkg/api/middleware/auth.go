package middleware

import (
	"errors"
	"fmt"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	response "github/MohdAhzan/internPortalAPP/pkg/utils/Response"
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func jwtMiddleware(secret string,expectedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
    tokenString:=c.GetHeader("Authorization")

    if tokenString==""{
      errMsg:=response.ClientResponse(http.StatusUnauthorized,"use authorized token",nil,errors.New("missing authorization token"))
      c.AbortWithStatusJSON(http.StatusUnauthorized,errMsg)
      return
    }

    tokenString= strings.TrimPrefix(tokenString,"Bearer ")

    token,err:=jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
      return []byte(secret),nil
    })
		if err != nil {
      errMsg:=response.ClientResponse(http.StatusUnauthorized,"enter valid token",nil,errors.New("invalid token"))
			c.AbortWithStatusJSON(http.StatusUnauthorized,errMsg)
			return
		}
    
    
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errMsg := response.ClientResponse(http.StatusUnauthorized, "Invalid token claims", nil, errors.New("invalid claims"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg)
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != expectedRole {
			errMsg := response.ClientResponse(http.StatusForbidden, "Unauthorized role", nil, errors.New("unauthorized access"))
			c.AbortWithStatusJSON(http.StatusForbidden, errMsg)
			return
		}
		id, ok := claims["id"].(float64)
		if !ok || id == 0 {
			errMsg := response.ClientResponse(http.StatusForbidden, "Invalid user ID", nil, errors.New("invalid user id"))
			c.AbortWithStatusJSON(http.StatusForbidden, errMsg)
			return
		}

		c.Set("role", role)
		c.Set("id", int(id))

		fmt.Println("Authorized:", role, "ID:", int(id))

		c.Next()
	}
}

func AdminAuthMiddleware(cfg config.Config)gin.HandlerFunc{
  return jwtMiddleware(cfg.AdminSecret,string(domain.Admin))
} 

func ReceptionistAuthMiddleware(cfg config.Config) gin.HandlerFunc {
	return jwtMiddleware(cfg.ReceptionistSecret,string(domain.Receptionist))
}

func DoctorAuthMiddleware(cfg config.Config) gin.HandlerFunc {
	return jwtMiddleware(cfg.DoctorSecret,string(domain.Doctor))
}

