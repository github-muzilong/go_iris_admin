package config

import "github-muzilong/go_iris_admin/bootstrap/config"

func init() {
	config.Add("pagination", config.StrMap{

		// 默认每页条数
		"perpage": 10,
	})
}
