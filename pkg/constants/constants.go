package constants

const (
	ServerAddr = "192.168.1.104"

	UserServiceName     = "user"
	VideoServiceName    = "video"
	InteractServiceName = "interact"
	IdentityKey         = "id"
	SecretKey           = "secret key"
	UsernameQueryKey    = "username"
	PasswordQueryKey    = "password"
	MySQLDefaultDsn     = "gorm:gorm@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	VideoSaveDirectory  = "/Users/rdstihz/nginx/static/video/"
	PlayURLPrefix       = "http://" + ServerAddr + ":9002/static/video/"
	TokenQueryKey       = "token"
	VodeoIdQueryKey     = "video_id"
	ActionTypeQueryKey  = "action_type"
	UserIdQueryKey      = "user_id"
	CommentText         = "comment_text"
	CommentId           = "comment_id"
)
