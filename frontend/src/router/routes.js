// ElasticView插件路由配置文件 - 定义插件的页面路由结构
// 根据ElasticView插件开发规范，路由配置需要与plugin.json中的frontend_routes保持一致

import pluginJson from '../../../plugin.json' // 导入插件配置文件，获取插件别名

// 导出路由配置数组 - ElasticView插件的页面路由定义
export const routes = [
  {
    // 插件根路径 - 必须以插件别名作为根路径，确保在ElasticView系统中的路由唯一性
    path: `/${pluginJson.plugin_alias}`, // 根据plugin.json中的plugin_alias动态生成
    component: ()=>import("@/layouts/Layout.vue"), // 懒加载默认布局组件，提供统一的页面结构
    // 子路由配置 - 插件的具体功能页面
    children: [
      {
        // HelloWorld页面路由 - ElasticSearch连接测试功能
        path: 'hello-world',                               // 路由路径，对应plugin.json中的frontend_routes配置
        component: ()=>import('@/views/helloworld/index.vue'), // 懒加载HelloWorld页面组件
        name: 'hello-world',                               // 路由名称，用于编程式导航
        meta: {
          title: '访问es',      // 页面标题，在ElasticView系统的导航菜单中显示
          icon: 'el-icon-coin'  // 页面图标，使用Element Plus图标
        }
      },
      {
        // 数据库测试页面路由 - 插件数据存储功能演示
        path: 'db-test',                                   // 路由路径，对应plugin.json中的frontend_routes配置
        component: ()=>import('@/views/db_test/DbTest.vue'), // 懒加载数据库测试页面组件
        name: 'db-test',                                   // 路由名称，用于编程式导航
        meta: {
          title: '操作数据库',   // 页面标题，展示插件的数据库操作能力
          icon: 'el-icon-coin'  // 页面图标，使用Element Plus图标
        }
      },
    ]
  },
]

// 默认导出路由配置 - 供router/index.js使用
export default routes
