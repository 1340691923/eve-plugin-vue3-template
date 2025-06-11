<template>
  <!-- ElasticView插件HelloWorld页面 - ElasticSearch连接测试功能演示 -->
  <div>
    <!-- 应用容器 - 提供统一的页面样式和布局 -->
    <div class="app-container">
      <!-- 搜索操作区域 - 包含ES连接测试按钮 -->
      <div class="search-container">
        <!-- Element Plus表单组件 - 内联布局显示操作按钮 -->
        <el-form :inline="true">
          <!-- 表单项 - ES连接测试按钮容器 -->
          <el-form-item label="ping es">
            <!-- 测试按钮 - 点击执行ES连接测试 -->
            <!-- disabled: 在请求进行中禁用按钮，防止重复点击 -->
            <!-- v-loading: 显示加载状态，提升用户体验 -->
            <!-- @click: 点击事件处理，调用init函数执行测试 -->
            <el-button 
              :disabled="loading" 
              v-loading="loading" 
              @click="init"
            >
              {{ $t('click me') }}
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
// ElasticView插件HelloWorld页面脚本 - 使用Vue 3 Composition API

// Vue 3 组合式API导入
import {onMounted, onActivated, ref} from "vue";  // 导入生命周期钩子和响应式API
// API接口导入
import {HelloAction} from "@/api/helloworld";     // 导入HelloWorld API接口
// UI组件库导入
import {ElMessage} from "element-plus";           // 导入Element Plus消息提示组件
// ElasticView插件SDK导入
import {sdk} from '@elasticview/plugin-sdk'       // 导入插件SDK，获取ES连接信息

// 响应式数据定义
const loading = ref(false) // 加载状态标识，控制按钮的禁用和loading效果

/**
 * init函数 - 执行ElasticSearch连接测试
 * 
 * 功能说明：
 * - 获取当前选中的ES连接ID
 * - 调用后端HelloWorld接口执行Ping测试
 * - 根据测试结果显示成功或失败消息
 * - 处理加载状态，提升用户体验
 */
const init = async () => {
  loading.value = true  // 设置加载状态，禁用按钮并显示loading效果
  
  // 调用HelloWorld API接口，传递ES连接ID
  let res = await HelloAction({
    es_connect: sdk.GetSelectEsConnID() // 获取用户在ElasticView中选中的ES连接ID
  })
  
  loading.value = false // 请求完成，取消加载状态

  // 检查API响应结果
  if (res.code !== 0) {
    // 请求失败，显示错误消息
    ElMessage({
      type: 'error',
      message: res.msg     // 显示后端返回的错误信息
    })
    return
  }
  
  // 请求成功，显示ES连接测试结果
  ElMessage({
    type: 'success',
    // 将ES响应的JSON数据格式化显示，提供详细的连接信息
    message: JSON.stringify(res.data.text, null, '\t')
  })
}

// Vue生命周期钩子
onMounted(() => {
  // 组件挂载完成后自动执行一次连接测试
  // 让用户进入页面即可看到当前ES连接状态
  init()
})

onActivated(() => {
  // 组件激活时的钩子（当使用keep-alive缓存时会触发）
  // 这里可以添加页面重新激活时需要执行的逻辑
})

</script>

<style scoped>
/* 组件样式 - 使用scoped确保样式只作用于当前组件 */
/* 可以在这里添加页面特定的样式定义 */
</style>
