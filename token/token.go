package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	// ErrMissingHeader 表示 Header中 Authorization 为空
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
// claims 包含保留claims(如，iss、sub)和自定义claims(如，下面的结构体的ID和Username)
// 自定义PayLoad字段结构体
type Context struct {
	ID       uint
	Username string
}

// secretFunc validates the secret format.
// secret格式化，secretFunc()函数返回的是jwt三方库的Keyfunc函数类型 func(*Token) (interface{}, error)
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil // secret被转换为byte单字节数组
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
// 使用指定的secret解析token字符串
// 如果解析验证token字符串是有效的，返回自定义上下文结构体Context
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
		//
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
// 验证Request里面的token
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
// 生成指定用户的token
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
