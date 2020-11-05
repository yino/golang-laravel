# [Golang-Laravel](https://github.com/sk532359025/golang-laravel)


模仿PHP-Laravel，开箱即用

#### package
* gin
* gorm
---
#### 支持功能
* 路由
* 中间件
* 路由中间件
* 控制器
* 模型
* 验证器（中文）
---
#### 目录结构
与laravel几乎一模一样

    app								   应用目录
    |---- exception				       返回http状态及异常接收
	|----|---- baseException.go 		  panic异常接收文件
    |---- http							web核心目录
	|----|---- controller 		 	   控制器存放目录
	|----|---- middleware 		 	   路由中间件存放目录
    |---- models						  模型存放目录
    |---- service						 业务编写目录
    |---- validate						验证器目录
    |---- helpers.go 					 常用函数文件
    config
    |----|----app.go  					应用配置 如：ip、port等
    |----|----databases.go  			  数据库配置
    routes								路由文件夹
	|----|---- route.go				   路由控制目录
	|----|---- .....				      更多路由
    sql								   测试sql文件存放目录
    vendor								扩展包目录
	|----|---- .....				      包存放


* 还在完善中
* 目前发现不合理地方 在返回success状态码 目前是用panic抛出异常返回，有为常理。