package middleware

import (
	"net/http"
	"strings"

	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/repository/profile"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Authorization(config *config.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(tokenHeader) < 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenHeader[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
			c.Set("id", claims["id"])
			c.Set("role", claims["role"])
			c.Next()
		}
	}
}

func TransactionAuthorization(config *config.Configuration, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileRepository := profile.NewProfileRepository(db)
		ctxUser := utils.GetContext(c)

		newError := customError.NewError(customError.ErrUnverifiedPhone.Error())

		profile, err := profileRepository.GetProfile(c, ctxUser.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    newError.Code,
				"message": newError.Message,
			})
			return
		}

		if !profile.IsVerifiedPhone {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    newError.Code,
				"message": newError.Message,
			})
			return
		}

		c.Next()
	}
}
