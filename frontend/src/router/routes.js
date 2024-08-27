import pluginJson from '../../../plugin.json'

export const routes = [
  {
    path: `/${pluginJson.plugin_alias}`,//根目录必须为插件名
    component: ()=>import("../layouts/Layout.vue"),
    children: [
      {
        path: 'hello-world',
        component: ()=>import('../views/helloworld/index.vue'),
        name: 'hello-world',
        meta: {
          title: '访问es',
          icon: 'el-icon-coin'
        }
      },
      {
        path: 'db-test',
        component: ()=>import('../views/db_test/DbTest.vue'),
        name: 'db-test',
        meta: {
          title: '操作数据库',
          icon: 'el-icon-coin'
        }
      },
    ]
  },
]

export default routes
