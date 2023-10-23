package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
)

var JwtStr = []byte("这是jwt认证密钥")

const (
	expiration = 60 * time.Minute
)

type Claims struct {
	UserName string
	Phone    string
	Role     string
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if !isTokenExist(token) {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if restoreToken(token) != nil {
			ctx.Abort()
			response.Failed(ctx, response.ErrRedis)
			return
		}
		claims, err := parseToken(token)
		if err != nil {
			newlog.Logger.Errorf("failed to parse token, err:%+v\n", err)
		}
		ctx.Request.Header.Set("userName", claims.UserName)
		ctx.Request.Header.Set("role", claims.Role)
		ctx.Request.Header.Set("userID", claims.Phone)
		newlog.Logger.Infof("user:%s, auth successfully", claims.UserName)
		ctx.Next()

	}
}

func GenerateToken(phone, userName string) (string, error) {
	claim := &Claims{
		UserName: userName,
		Phone:    phone,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(JwtStr)
}

func parseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtStr, nil
	})
	if _, ok := tokenClaim.Claims.(*Claims); ok {
		if tokenClaim.Claims.Valid() == nil {
			return tokenClaim.Claims.(*Claims), nil
		}
	}
	return nil, err
}

func StoreToken(token string) error {
	//return newlog.RedisClient.Set(token, nil, expiration).Err()
	return nil
}

func restoreToken(token string) error {
	return StoreToken(token)
}

func isTokenExist(token string) bool {
	return true
	//return newlog.RedisClient.Get(token).Err() == nil
}
