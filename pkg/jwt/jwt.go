package jwt

import (
	"errors"

	"strings"
	"time"

	"github.com/Cynthia/goapi/pkg"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/gin-gonic/gin"

	jwtpkg "github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
    ErrTokenExpired           error = errors.New("令牌已过期")
    ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
    ErrTokenMalformed         error = errors.New("请求令牌格式有误")
    ErrTokenInvalid           error = errors.New("请求令牌无效")
    ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
    ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct{
	SignKey []byte
	MaxRefresh time.Duration
}

type JWTCustomClaims struct{
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	ExpireAtTime string `json:"expire_time"`
	jwtpkg.StandardClaims
}

func NewJWT()*JWT{
	return &JWT{
		SignKey: []byte(viper.GetString("app.key")),
		MaxRefresh: time.Duration(viper.GetInt64("jwt.max_refresh_time"))*time.Minute,
	}
}

func(jwt *JWT)ParserToken(c *gin.Context)(*JWTCustomClaims,error){
	tokenString,parseErr:=jwt.getTokenFromHeader(c)
	if parseErr!=nil {
		return nil,parseErr
	}
	token,err:=jwt.parseTokenString(tokenString)
	if err!=nil{
		validationErr,ok:=err.(*jwtpkg.ValidationError)
		if ok{
			if validationErr.Errors==jwtpkg.ValidationErrorMalformed{
				return nil,ErrTokenMalformed
			}else if validationErr.Errors==jwtpkg.ValidationErrorExpired{
				return nil,ErrTokenExpired
			}
		}

		return nil,ErrTokenInvalid
	}

	if claims,ok:=token.Claims.(*JWTCustomClaims);ok && token.Valid{
		return claims,err
	}
	return nil,ErrTokenInvalid
}

func (jwt *JWT)getTokenFromHeader(c *gin.Context)(string,error){
	authHeader:=c.Request.Header.Get("Authorization")
	if authHeader==""{
		return "",ErrHeaderEmpty
	}
	parts:=strings.SplitN(authHeader," ",2)
	if !(len(parts)==2&&parts[0]=="Bearer"){
		return "",ErrHeaderMalformed
	}
	return parts[1],nil
}

func(jwt *JWT)parseTokenString(tokenString string)(*jwtpkg.Token,error){
	return jwtpkg.ParseWithClaims(tokenString,&JWTCustomClaims{},func(t *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey,nil
	})
}

func (jwt *JWT)RefreshToken(c *gin.Context)(string,error){
	tokenString, parseErr := jwt.getTokenFromHeader(c)
    if parseErr != nil {
        return "", parseErr
    }
    token, err := jwt.parseTokenString(tokenString)
    if err != nil {
        validationErr, ok := err.(*jwtpkg.ValidationError)
        // 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
        if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
            return "", err
        }
    }

    // 4. 解析 JWTCustomClaims 的数据
    claims := token.Claims.(*JWTCustomClaims)

    // 5. 检查是否过了『最大允许刷新的时间』
    x := pkg.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
    if claims.IssuedAt > x {
        // 修改过期时间
        claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
        return jwt.createToken(*claims)
    }

    return "", ErrTokenExpiredMaxRefresh
}

func(jwt *JWT)createToken(claims JWTCustomClaims)(string,error){
	token:=jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256,claims)
	return token.SignedString(jwt.SignKey)
}

func(jwt *JWT)expireAtTime()int64{
	timeNow:=pkg.TimenowInTimezone()

	var expireTime int64

	if viper.GetBool("app.debug"){
		expireTime=viper.GetInt64("jwt.debug_expire_time")
	}else{
		expireTime=viper.GetInt64("jwt.expire_time")
	}

	expire:=time.Duration(expireTime)*time.Minute
	return timeNow.Add(expire).Unix()
}

func(jwt *JWT)IssueToken(userID,userName string)string{
	expireAtTime := jwt.expireAtTime()
    claims := JWTCustomClaims{
        userID,
        userName,
		cast.ToString(expireAtTime),
        jwtpkg.StandardClaims{
            NotBefore: pkg.TimenowInTimezone().Unix(), // 签名生效时间
            IssuedAt:  pkg.TimenowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
            ExpiresAt: expireAtTime,                   // 签名过期时间
            Issuer:    viper.GetString("app.name"),   // 签名颁发者
        },
    }

	token,err:=jwt.createToken(claims)
	if err!=nil{
		logger.LogIf(err)
		return ""
	}
	return token
}