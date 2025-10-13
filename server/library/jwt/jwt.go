package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type myCustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

// IssueUserIdToken 签发用户ID为载核的token
func IssueUserIdToken(userId int64, signingKey string, expireSecond int64) (string, error) {
	timeNow := time.Now()
	// Create the Claims
	claims := myCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(time.Duration(expireSecond) * time.Second)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                             // 生效时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(signingKey))
	return t, err
}

// ParseUserId 解析用户id
func ParseUserId(t, signingKey string) (int64, error) {
	if t == "" {
		return 0, nil
	}
	token, err := jwt.ParseWithClaims(t, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {

		return 0, err
	}

	if claims, ok := token.Claims.(*myCustomClaims); ok && token.Valid {
		return claims.UserId, nil
	}
	return 0, err
}
