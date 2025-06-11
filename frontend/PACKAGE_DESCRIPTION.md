# ElasticView插件前端项目依赖说明

## 基本信息
- **项目名称**: ev-plugin (ElasticView插件)
- **版本**: 0.0.0 (开发版本)
- **类型**: ES模块 (module)
- **私有项目**: true (不发布到npm)

## 构建脚本 (scripts)
```json
{
  "dev": "vite",                              // 启动开发服务器，支持热重载调试
  "build": "vite build && node build_ev.js", // 构建生产版本并执行后处理脚本
  "preview": "vite preview"                   // 预览构建产物
}
```

## 生产依赖 (dependencies)

### ElasticView插件系统
- **@elasticview/plugin-sdk**: 1.0.14
  - ElasticView官方插件SDK
  - 提供插件生命周期管理、API调用、路由集成等核心功能
  - 是插件与ElasticView基座通信的桥梁

### UI组件库
- **element-plus**: ^2.7.8
  - Vue 3版本的Element Plus UI组件库
  - 提供丰富的企业级UI组件
- **@element-plus/icons-vue**: ^2.3.1
  - Element Plus官方图标库
  - Vue 3版本的图标组件
- **element-ui**: ^2.15.14
  - Vue 2版本的Element UI（兼容性考虑）

### Vue生态系统
- **vue**: ^3.2.45
  - Vue 3框架核心
  - 提供响应式数据绑定和组件系统
- **vue-router**: ^4.1.6
  - Vue 3官方路由库
  - 用于插件内部页面路由管理
- **vue-i18n**: 10.0.0-beta.5
  - Vue国际化库
  - 支持多语言切换
- **vue3-eventbus**: ^2.0.0
  - Vue 3事件总线
  - 组件间通信工具
- **vue-facing-decorator**: ^3.0.4
  - Vue装饰器库
  - 支持类组件语法

### 工具库
- **@vueuse/core**: ^10.11.1
  - Vue组合式API工具集
  - 提供常用的响应式工具函数
- **@vueuse/components**: ^10.11.1
  - VueUse组件库
  - 基于组合式API的实用组件
- **lodash**: ^4.17.21
  - JavaScript工具库
  - 提供数据处理和操作方法
- **jsonlint-mod**: ^1.7.6
  - JSON验证库
  - 用于JSON格式校验和错误检查

### 代码编辑器
- **monaco-editor**: ^0.50.0
  - VS Code同款代码编辑器
  - 提供语法高亮、自动补全等功能
  - 用于插件中的代码编辑功能

## 开发依赖 (devDependencies)

### 构建工具
- **vite**: ^4.1.0
  - 现代化前端构建工具
  - 支持快速热重载和ES模块
- **@vitejs/plugin-vue**: ^4.0.0
  - Vite的Vue 3官方插件
  - 提供.vue文件编译支持

### TypeScript支持
- **typescript**: ^5.5.4
  - TypeScript编译器
  - 提供类型检查和ES6+转换
- **vue-tsc**: ^2.0.29
  - Vue文件的TypeScript类型检查工具

### 样式处理
- **sass**: ^1.77.8
  - Sass CSS预处理器
  - 支持嵌套语法和变量

### Vite插件
- **vite-plugin-qiankun**: ^1.0.15
  - qiankun微前端集成插件
  - 用于集成到ElasticView插件系统
- **vite-plugin-dynamic-base**: ^1.1.0
  - 动态基础路径插件
  - 支持运行时路径调整
- **vite-plugin-monaco-editor**: ^1.1.0
  - Monaco编辑器Vite插件
  - 优化Monaco编辑器的打包配置
- **vite-plugin-string-replace**: ^1.1.3
  - 字符串替换插件
  - 用于构建时的文本替换操作

## 架构说明

这个配置支持：
1. **微前端架构**: 通过qiankun插件集成到ElasticView系统
2. **现代化开发**: 使用Vite + Vue 3 + TypeScript技术栈
3. **企业级UI**: Element Plus组件库提供专业的用户界面
4. **代码编辑**: Monaco编辑器支持代码编辑功能
5. **国际化**: vue-i18n支持多语言
6. **类型安全**: TypeScript提供完整的类型检查 