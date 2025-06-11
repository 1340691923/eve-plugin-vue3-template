<template>
  <!-- ElasticView插件根组件模板 -->
  <!-- 使用Element Plus配置提供器包装整个应用，确保组件样式和尺寸的一致性 -->
  <div>
    <!-- Element Plus全局配置组件 -->
    <!-- size属性根据ElasticView系统的布局设置动态调整组件尺寸 -->
    <el-config-provider :size="size">
      <!-- Vue路由视图容器 - 插件页面内容的渲染位置 -->
      <!-- 不同的路由对应不同的插件功能页面 -->
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>
// ElasticView插件根组件脚本 - 使用Vue 3 Composition API

import {sdk} from '@elasticview/plugin-sdk' // 导入ElasticView插件SDK
import {computed} from "vue";               // 导入Vue响应式计算属性

// 计算属性 - 动态获取Element Plus组件的尺寸设置
// 根据ElasticView系统的全局布局配置自动调整组件大小
// 确保插件UI与ElasticView主系统的视觉风格保持一致
const size = computed(()=>{
  return sdk.getLayoutSize() // 调用SDK方法获取当前布局尺寸设置（如'small', 'default', 'large'）
})

</script>
