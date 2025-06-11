# ElasticView插件配置文件说明 (plugin.json)

## 概述

`plugin.json` 是ElasticView插件的核心配置文件，定义了插件的基本信息、前后端路由配置和开发参数。该文件必须位于插件项目根目录，且格式必须严格遵循JSON规范。

## 配置示例

```json
{
  "developer": "肖大侠",
  "plugin_alias": "eve-plugin-vue3-template",
  "plugin_name":"hello-world",
  "version": "0.0.1",
  "main_go_file": "main.go",
  "frontend_dev_port":7001,
  "frontend_2c": false,
  "frontend_debug": true,
  "frontend_routes": [
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
  "backend_routes": [
    {
      "path": "/api/DbSearch",
      "remark": "测试db查询",
      "needAuth": true
    }
  ]
}
```

## 字段详解

### 基本信息

#### `developer` (必需)
- **类型**: String
- **说明**: 插件开发者名称
- **示例**: `"肖大侠"`
- **用途**: 在ElasticView插件市场和管理界面中显示

#### `plugin_alias` (必需)
- **类型**: String
- **说明**: 插件唯一标识符，也称为插件ID
- **示例**: `"eve-plugin-vue3-template"`
- **要求**: 
  - 全局唯一，不能与其他插件重复
  - 只能包含字母、数字、连字符
  - 建议使用有意义的名称
- **用途**: 
  - 作为前端路由的根路径
  - 插件安装和管理的唯一标识

#### `plugin_name` (必需)
- **类型**: String
- **说明**: 插件显示名称
- **示例**: `"hello-world"`
- **用途**: 在ElasticView界面中显示的插件名称

#### `version` (必需)
- **类型**: String
- **说明**: 插件版本号
- **示例**: `"0.0.1"`
- **格式**: 建议使用语义化版本 (Semantic Versioning)
- **用途**: 版本管理和更新检测

### 后端配置

#### `main_go_file` (必需)
- **类型**: String
- **说明**: Go主程序文件路径
- **示例**: `"main.go"`
- **用途**: ElasticView构建工具需要此信息来编译插件

### 前端配置

#### `frontend_dev_port` (必需)
- **类型**: Number
- **说明**: 前端开发服务器端口
- **示例**: `7001`
- **要求**: 端口不能与其他服务冲突
- **用途**: 开发模式下前端服务器监听端口

#### `frontend_2c` (可选)
- **类型**: Boolean
- **说明**: 是否为客户端(C端)插件
- **默认值**: `false`
- **说明**: 
  - `true`: 客户端插件，面向最终用户
  - `false`: 管理端插件，面向管理员

#### `frontend_debug` (必需)
- **类型**: Boolean
- **说明**: 前端调试模式开关
- **用法**:
  - `true`: 开发模式，使用开发服务器
  - `false`: 生产模式，使用打包后的静态文件
- **注意**: 生产部署时必须设置为`false`

### 路由配置

#### `frontend_routes` (必需)
- **类型**: Array
- **说明**: 前端页面路由配置数组
- **结构**:
```json
{
  "path": "hello-world",     // 路由路径
  "name": "HelloWorld",      // 路由名称
  "meta": {
    "title": "HelloWorld",   // 页面标题
    "icon": "el-icon-coin"   // 页面图标
  }
}
```

##### 路由字段说明
- **`path`**: 
  - 路由路径，最终路径为 `/${plugin_alias}/${path}`
  - 必须与前端路由配置保持一致
- **`name`**: 
  - 路由名称，用于编程式导航
  - 建议使用PascalCase命名
- **`meta.title`**: 
  - 页面标题，显示在ElasticView导航菜单中
  - 支持国际化
- **`meta.icon`**: 
  - 页面图标，使用Element Plus图标类名
  - 可选字段，默认使用通用图标

#### `backend_routes` (可选)
- **类型**: Array
- **说明**: 后端API路由配置数组
- **结构**:
```json
{
  "path": "/api/DbSearch",   // API路径
  "remark": "测试db查询",    // 接口说明
  "needAuth": true           // 是否需要认证
}
```

##### 后端路由字段说明
- **`path`**: API接口路径
- **`remark`**: 接口功能说明
- **`needAuth`**: 是否需要用户认证

## 配置最佳实践

### 1. 命名规范
- **plugin_alias**: 使用kebab-case，包含项目前缀
- **路由名称**: 使用PascalCase或kebab-case保持一致性
- **版本号**: 遵循语义化版本规范

### 2. 开发环境配置
```json
{
  "frontend_debug": true,    // 开发时启用
  "frontend_dev_port": 7001  // 确保端口可用
}
```

### 3. 生产环境配置
```json
{
  "frontend_debug": false    // 生产时禁用调试
}
```

### 4. 路由一致性
确保以下配置保持同步：
- `plugin.json`中的`frontend_routes`
- `frontend/src/router/routes.js`中的路由定义
- 实际的Vue页面组件

## 常见问题

### 1. 路由不显示
检查以下配置：
- `frontend_routes`配置是否正确
- 前端路由文件是否与配置一致
- 页面组件是否存在

### 2. 开发服务器无法启动
检查：
- `frontend_dev_port`是否被占用
- `frontend_debug`是否设置为`true`

### 3. 插件加载失败
检查：
- JSON格式是否正确（使用JSON验证工具）
- 必需字段是否都已配置
- `plugin_alias`是否唯一

## 配置验证

在修改`plugin.json`后，建议进行以下验证：

1. **JSON格式验证**: 使用在线JSON验证工具检查语法
2. **路由一致性检查**: 确保前后端路由配置匹配
3. **端口可用性检查**: 确认开发端口未被占用
4. **插件加载测试**: 重启ElasticView服务验证配置

## 扩展配置

随着插件功能的增加，可能需要添加新的配置项：

1. **新增页面**: 在`frontend_routes`中添加路由配置
2. **新增API**: 在`backend_routes`中添加接口配置
3. **版本更新**: 修改`version`字段并遵循版本管理规范

## 注意事项

1. **备份配置**: 修改前建议备份原配置文件
2. **团队协作**: 确保团队成员使用相同的配置
3. **环境切换**: 开发和生产环境的`frontend_debug`配置不同
4. **安全考虑**: 敏感信息不要写入配置文件 