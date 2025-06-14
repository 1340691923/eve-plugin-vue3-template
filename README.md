# eve-plugin-vue3-template

## Here is a template project for developing ElasticView plugins

 [中文](./README-cn.md) | English 

```
Project start-up preparation:

Base start:
Make sure the ElasticView base program has been started

Environment configuration:
golang version >= 1.20
node version >= 20.14.0

install gowatch:
go install github.com/silenceper/gowatch@latest
go install github.com/1340691923/eve-plugin-sdk-go/cmd/ev_plugin_builder@v0.0.15
go install github.com/1340691923/eve-plugin-sdk-go/cmd/ev_plugin_zip@v0.0.15

install pnpm:
npm install -g pnpm

```

### Project infrastructure

```
-backend                      Backend project directory
    -api                      Controller layer
    -dto                      Web request structure
    -migrate                  SQLite data table version control module
    -versions                 Stores each version of the data table structure upgrade rollback
    -model                    Data access layer
    -my_error                 Template custom exception package
    -response                 Template custom response package
    -router                   Backend routing definition module
    -vo                       Web response structure
-frontend Frontend            project directory
    -dist                     Frontend final package file
    -src
    -api                      Interface access layer
    -lang                     Language package
    -layouts                  Default layout
    -plugin_sdk               Plugin sdk
    -router                   Frontend routing
    -views                    Page file
```

### plugin.json：

```json
{
  "developer": "xiaowenlong", //Developer Name
  "plugin_alias": "eve-plugin-vue3-template", //Plugin id is also called plugin alias
  "plugin_name":"hello-world", //plugin display name
  "frontend_debug": false, // Whether to enable front-end page debugging
  "version": "0.0.1",     //current version
  "main_go_file": "main.go", //  main.go file position
  "frontend_dev_port":7001, //Front-end project startup port
  "frontend_routes": [  // The front-end routing is consistent with routes
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
  "backend_routes": [   //Backend Routing
    {
      "path": "/api/HelloWorld",  //Interface access path
      "remark": "HelloWorld测试接口", 
      "needAuth": false       //Is authentication required?
    }
  ]
}





```

#### Backend plugin development

```
Project root directory execution:
1. go mod tidy
2. gowatch

The following indicates successful startup:

正常链接ev基座

```

#### Front-end plugin development
```

Key point: Modify frontend_debug in plugin.json to true

1. cd frontend && pnpm i

2. npm run dev

```

#### Front-end plugin packaging
```

1. cd frontend && npm run build

```

#### Back-end plugin packaging

```
Key point: Modify frontend_debug in plugin.json to false

Project root directory execution:
./ev_plugin_builder

A compressed package will appear in the dist directory. After decompression, it will be the plugin binary for different operating systems.

```
