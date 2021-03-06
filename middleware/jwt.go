package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/request"

	"github.com/dgrijalva/jwt-go"
)

const TokenName = "token"
const Issuer = "Money-Hub"
const SecretKey = "MoneyDodo"

/*
*	Role
* 	0 - admin
* 	1 - student
* 	2 - enterprise
*
 */

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Id   string `json:"id"`
	Role int    `json:"role"`
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(SecretKey []byte, issuer string, id string, role int) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 24 * 365 * 100).Unix()),
			Issuer:    issuer,
		},
		id,
		role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.MapClaims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return claims, err
	}
	claims = token.Claims.(jwt.MapClaims)
	return
}

/**
*	提取 request内的token信息
 */
func GetTokenInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var mapClaims jwt.MapClaims
		myToken := ""
		// 如果token存在于Authorization中
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		checkErr(err)
		if token != nil {
			var ok bool
			mapClaims, ok = token.Claims.(jwt.MapClaims)
			if ok {
				myToken = strings.Split(r.Header["Authorization"][0], " ")[1]
			}
		} else {
			// 如果token存在于header中
			for k, v := range r.Header {
				if strings.ToLower(k) == "token" {
					myToken = v[0]
					break
				}
			}
			if myToken != "" {
				mapClaims, err = ParseToken(myToken, []byte(SecretKey))
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		// fmt.Println(myToken)
		if myToken == "" {
			next.ServeHTTP(w, r)
		} else {
			ctx := context.WithValue(r.Context(), "id", mapClaims["id"].(string))
			ctx = context.WithValue(ctx, "role", int(mapClaims["role"].(float64)))
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
	})
}
