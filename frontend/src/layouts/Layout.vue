<template>
  <!-- ElasticView插件默认布局组件 -->
  <!-- 为插件提供统一的页面布局结构，支持页面缓存和组件动态切换 -->
  <div class="eve-plugin-app">
    <!-- Vue路由视图容器 - 使用作用域插槽获取当前路由组件 -->
    <!-- 支持插件内多个页面的路由切换和展示 -->
    <router-view v-slot="{ Component }">
      <!-- keep-alive组件 - 实现页面缓存功能 -->
      <!-- 当用户在插件内切换页面时，保持组件状态不丢失，提升用户体验 -->
      <keep-alive>
        <component :is="Component"  />
      </keep-alive>
    </router-view>
  </div>
</template>

<script lang="ts">
// ElasticView插件布局组件脚本 - 使用Vue 3组合式API
import { defineComponent } from 'vue' // 导入Vue组件定义函数

// 导出布局组件定义
// 作为插件页面的容器，提供统一的布局结构和缓存机制
export default defineComponent({
  // 组件名称，用于调试和开发工具识别
  name: 'PluginLayout'
})
</script>
