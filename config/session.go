package config

import "github-muzilong/go_iris_admin/bootstrap/config"

func init() {
	config.Add("session", config.StrMap{

		// 目前只支持 Cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),

		// 会话的 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "goblog-session"),

		// 后台用户的session key
		"session_admin": config.Env("SESSION_ADMIN", "nmdshfsfasf"),
	})
}
