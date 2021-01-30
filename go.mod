module github-muzilong/go_iris_admin

go 1.14

replace github-muzilong/go_iris_admin => ./

require (
	github.com/kataras/iris/v12 v12.2.0-alpha2.0.20210115205746-6d10b014859c
	github.com/spf13/cast v1.3.0
	github.com/spf13/viper v1.7.1
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)
