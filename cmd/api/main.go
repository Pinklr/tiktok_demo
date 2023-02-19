package main

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/api/handler"
	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"log"
	"time"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	// jwt身份验证中间件
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			log.Println(c.QueryArgs())
			username := c.Query("username")
			password := c.Query("password")
			if len(username) == 0 || len(password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			userID, err := rpc.CheckUser(ctx, &user.CheckUserRequest{
				UserName: username,
				Password: password,
			})
			if err != nil {
				return nil, err
			}
			return userID, nil
		},

		// PayloadFunc 向jwt token的payload中写入用户id
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			log.Printf("%T %v\n", data, data)
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"status_code": errno.Success.ErrCode,
				"status_msg":  errno.Success.ErrMsg,
				"user_id":     0,
				"token":       token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": errno.AuthorizationFailedErr.ErrMsg,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, param: token, form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		IdentityKey:   constants.IdentityKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	v1 := r.Group("douyin/")
	user := v1.Group("user/")
	user.POST("login/", authMiddleware.LoginHandler)
	user.POST("register/", handler.RegisterHandler)
	user.Use(authMiddleware.MiddlewareFunc())
	user.GET("/", handler.GetUserHandler)
	user.POST("refresh/", authMiddleware.RefreshHandler)
	r.Spin()
}
