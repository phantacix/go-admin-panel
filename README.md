## 项目结构

- `config`      环境配置文件、系统配置表。
- `core`        通用的公共文件。
- `logs`        系统启动时`zap`日志系统生成的日志文件。
- `misc`        系统必要的一些杂项配置文件，例如：db脚本、文档、程序启动脚本。
- `module`      模块功能，目前`admin`放系统权限管理代码。
- `publish`     发布目录。
- `resources`   静态资源目录。UI资源和功能html引擎模块。


## 调试&配置

- 打开goland，`File->Settings->Go->Go Modules(vgo)`，勾选`Enable Go Modules(vgo) integration`，设置`Proxy`为`https://goproxy.cn`
- 从`misc->db_script`目录导入相关的db脚本到数据库。确认`config/env`的配置参数正确。
- main.go默认读取`config/env/local` 的环境配置，如需自定义，请在`config/env`创建自己的环境目录。并且在`Debug Configurations/Program arguments`设置`-env=xxx`。
- 开发时使用`gorm`的AutoMigrate功能更新数据库。可参照`module/admin/admin_test.go`中的`TestAdminDBMigrate()`编写脚本。


## 模板引擎

- UI框架使用的是`layui`，说明文档请访问 https://www.layui.com/doc/
- html模板引擎使用的是`golang`系统自带的`template`。请搜索`golang template`查询查关语法。

## 路由规划

- TODO