# eve-plugin-vue3-template

## 这里用于开发ElasticView插件的模板工程

 [英文](./README.md) | 简体中文 

```
项目启动准备工作：

基座启动：
确保已经启动ElasticView基座程序

环境配置：
golang版本 >= 1.20
node版本 >= 20.14.0

安装gowatch:
go install github.com/silenceper/gowatch@latest

go install github.com/1340691923/ElasticView/cmd/ev_plugin_builder@v0.0.12
go install github.com/1340691923/ElasticView/cmd/ev_plugin_zip@v0.0.12

安装pnpm:
npm install -g pnpm

```

### 项目基础架构

```
-backend        后端项目目录
    -api        控制器层
    -dto        web请求结构
    -migrate    sqlite数据表版本控制模块
        -versions 存放各版本数据表结构升级回退
    -model      数据访问层
    -my_error   模板自定义异常包
    -response   模板自定义响应包
    -router     后端路由定义模块
    -vo         web响应结构
-frontend      前端项目目录
    -dist      前端最终打包文件
    -src
        -api   接口访问层
        -lang   语言包
        -layouts    默认布局
        -plugin_sdk 插件sdk
        -router     前端路由
        -views      页面文件
```

### plugin.json配置：

```json
{
  "developer": "xiaowenlong", //开发者名称
  "plugin_alias": "eve-plugin-vue3-template", //插件id 也叫插件别名
  "plugin_name":"hello-world", //插件显示名称
  "frontend_debug": false, // 是否开启前端页面调试
  "version": "0.0.1",     //当前版本号
  "main_go_file": "main.go", // 后端 main.go文件位置
  "frontend_dev_port":7001, //前端项目启动端口
  "frontend_routes": [  //前端路由 与 routes中保持一致
    {
      "path": "hello-world", 
      "name": "HelloWorld",
      "meta": {
        "title": "HelloWorld",
        "icon": "el-icon-coin"
      }
    },
    {
      "path": "db-test",
      "name": "db-test",
      "meta": {
        "title": "操作数据库",
        "icon": "el-icon-coin"
      }
    }
  ],
  "backend_routes": [   //后端路由
    {
      "path": "/api/HelloWorld",  //接口访问路径
      "remark": "HelloWorld测试接口", //备注
      "needAuth": false       //是否需要鉴权
    }
  ]
}





```

#### 后端插件开发

```
项目根目录执行:
1. go mod tidy
2. gowatch

显示如下为启动成功：

正常链接ev基座

```




#### 前端插件开发
```

重点：修改plugin.json 中 frontend_debug 为 true

1. cd frontend && pnpm i 

2. npm run dev

```

#### 前端插件打包
```

1. cd frontend && npm run build

```

#### 后端插件打包

```
重点：修改plugin.json 中 frontend_debug 为 false

项目根目录执行:
ev_plugin_builder

在dist目录会出现一个压缩包，解压后是各个不同操作系统的插件二进制，
按自己实际情况移动至ElasticView的plugin文件夹中即可

```
