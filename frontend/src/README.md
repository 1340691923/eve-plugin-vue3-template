# ElasticView插件前端源码结构说明

## 目录结构

```
frontend/src/
├── api/                    # API接口定义目录
│   ├── helloworld.ts      # HelloWorld功能API接口
│   ├── db_test.ts         # 数据库测试API接口
│   └── 接口管理           # API接口管理规范说明
├── layouts/               # 布局组件目录
│   └── Layout.vue         # 默认布局组件
├── router/                # 路由配置目录
│   ├── index.js          # 路由实例创建
│   └── routes.js         # 路由规则定义
├── views/                 # 页面视图组件目录
│   ├── helloworld/       # HelloWorld功能页面
│   │   └── index.vue     # ES连接测试页面
│   └── db_test/          # 数据库测试页面
│       └── DbTest.vue    # 数据库CRUD操作页面
├── lang/                  # 国际化语言包目录
│   ├── en.ts             # 英文语言包
│   └── zh-cn.ts          # 中文简体语言包
├── App.vue               # Vue应用根组件
└── main.ts               # 应用入口文件
```

## 文件功能说明

### 核心文件

#### `main.ts` - 应用入口
- **功能**: ElasticView插件前端应用的初始化入口
- **关键作用**:
  - 集成ElasticView插件SDK
  - 配置Vue 3应用实例
  - 注册Element Plus UI组件库
  - 设置国际化支持
  - 启动插件应用

#### `App.vue` - 根组件
- **功能**: Vue应用的根组件
- **关键作用**:
  - 提供Element Plus全局配置
  - 集成ElasticView布局尺寸设置
  - 包含路由视图容器

### 路由系统

#### `router/index.js` - 路由实例
- **功能**: 创建Vue Router实例
- **关键作用**:
  - 使用Hash路由模式适配微前端环境
  - 初始化路由系统

#### `router/routes.js` - 路由配置
- **功能**: 定义插件的页面路由规则
- **关键作用**:
  - 配置插件根路径（基于plugin_alias）
  - 定义子页面路由
  - 设置页面元信息（标题、图标）

### 布局系统

#### `layouts/Layout.vue` - 默认布局
- **功能**: 插件页面的统一布局容器
- **关键作用**:
  - 提供页面缓存功能（keep-alive）
  - 支持组件动态切换
  - 统一的页面结构

### API层

#### `api/` 目录
- **功能**: 管理与后端通信的API接口
- **组织原则**:
  - 按功能模块划分文件
  - 使用ElasticView插件SDK的request方法
  - 提供完整的TypeScript类型定义

#### API文件列表
- `helloworld.ts`: ES连接测试相关接口
- `db_test.ts`: 数据库操作相关接口
- `接口管理`: API开发规范和使用说明

### 视图层

#### `views/` 目录
- **功能**: 存放插件的功能页面组件
- **组织原则**:
  - 按功能模块创建子目录
  - 每个页面组件包含完整的业务逻辑
  - 使用Vue 3 Composition API

#### 页面列表
- `helloworld/index.vue`: ElasticSearch连接测试页面
- `db_test/DbTest.vue`: 数据库CRUD操作演示页面

### 国际化

#### `lang/` 目录
- **功能**: 提供多语言支持
- **支持语言**:
  - `en.ts`: 英文语言包
  - `zh-cn.ts`: 中文简体语言包

#### 使用方式
```vue
<!-- 在模板中 -->
{{ $t('click me') }}

<!-- 在脚本中 -->
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const text = t('click me')
```

## 开发规范

### 1. 文件命名
- **组件文件**: 使用PascalCase（如：`HelloWorld.vue`）
- **页面文件**: 使用kebab-case目录 + 描述性文件名
- **API文件**: 使用snake_case（如：`db_test.ts`）
- **工具文件**: 使用camelCase

### 2. 代码组织
- **单一职责**: 每个文件专注于单一功能
- **模块化**: 合理拆分组件和工具函数
- **可维护性**: 提供完整的注释和文档

### 3. Vue 3 最佳实践
- **优先使用 Composition API**: 更好的逻辑复用和类型推导
- **响应式数据**: 合理使用ref和reactive
- **生命周期**: 正确使用onMounted、onActivated等钩子

### 4. ElasticView集成
- **插件SDK**: 充分利用`@elasticview/plugin-sdk`提供的功能
- **路由规范**: 遵循ElasticView插件路由约定
- **样式适配**: 适配ElasticView主题和布局

## 技术栈

### 核心框架
- **Vue 3**: 渐进式JavaScript框架，使用Composition API
- **Vue Router 4**: Vue 3官方路由库
- **Vue I18n**: 国际化解决方案

### UI组件库
- **Element Plus**: Vue 3版本的企业级UI组件库
- **Element Plus Icons**: 官方图标库

### 工具库
- **@elasticview/plugin-sdk**: ElasticView官方插件SDK
- **@vueuse/core**: Vue Composition API工具集
- **TypeScript**: 类型安全的JavaScript超集

### 构建工具
- **Vite**: 现代化前端构建工具
- **vite-plugin-qiankun**: 微前端集成插件

## 开发流程

### 1. 新增页面
1. 在`views/`目录创建页面组件
2. 在`router/routes.js`中添加路由配置
3. 更新`plugin.json`中的`frontend_routes`
4. 添加对应的国际化文本

### 2. 新增API接口
1. 在`api/`目录创建或更新接口文件
2. 定义接口函数和类型
3. 在页面组件中导入使用
4. 处理错误和加载状态

### 3. 调试和测试
1. 启动开发服务器：`npm run dev`
2. 在浏览器中访问ElasticView并查看插件
3. 使用开发者工具调试
4. 检查网络请求和控制台日志

## 注意事项

1. **路由配置**: 必须与`plugin.json`中的配置保持一致
2. **微前端兼容**: 使用Hash路由模式，避免路径冲突
3. **资源引用**: 使用相对路径或别名引用资源
4. **样式隔离**: 使用scoped样式避免样式污染
5. **性能优化**: 合理使用懒加载和缓存机制 