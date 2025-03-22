package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

// CustomClaims 自定义 Claims 结构体
type CustomClaims struct {
	JwtUserId int64 `json:"jwtUserId"` // 假设你在生成 Token 时使用了 "user_id" 作为用户 ID 的键
	jwt.StandardClaims
}

// ParseJWT 解析 JWT Token 并获取用户信息
func ParseJWT(tokenStr string, secretKey string) int64 {
	// 解析 JWT Token
	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		// 替换为你的 JWT 签名密钥
		return []byte(secretKey), nil
	})

	if err != nil {
		logx.Errorf("ParseJWT err : %+v", err)
	}

	if !token.Valid {
		logx.Errorf("ParseJWT err : %+v", err)
	}

	return claims.JwtUserId
}
