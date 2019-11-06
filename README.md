## 项目结构

- `config`      环境配置文件、系统配置表。
- `core`        通用的公共文件。
- `logs`        系统启动时`zap`日志系统生成的日志文件。
- `misc`        系统必要的一些杂项配置文件，例如：db脚本、文档、程序启动脚本。
- `module`      模块功能，目前`admin`放系统权限管理代码。
- `publish`     发布目录。
- `resources`   静态资源目录。UI资源和功能html引擎模块。


## 调试&配置

- 请确认本机已安装go 1.13+版本。
- 打开goland，`File->Settings->Go->Go Modules(vgo)`，勾选`Enable Go Modules(vgo) integration`，设置`Proxy`为`https://goproxy.cn`
- main.go默认读取`config/env/local/admin.json`配置文件。请确认`端口`可用，`mysql`结点的数据库、帐号和密码正确。
- 如需自定义配置，请在`config/env`创建自己的环境目录。并且在`Debug Configurations/Program arguments`设置`-env=xxx`。
- 创建数据库表结构。执行`module/admin/admin_test.go`中的`TestAdminDBMigrate()`函数即可创建表结构，或手工导入`misc/db_script/go_admin_panel.sql`文件。


## 模板引擎

- UI框架使用的是`layui`，说明文档请访问 https://www.layui.com/doc/
- html模板引擎使用的是`golang`系统自带的`template`。请搜索`golang template`查询查关语法。

## 路由规划

- TODO