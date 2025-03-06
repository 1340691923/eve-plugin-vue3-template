<template>
  <div>
    <div class="app-container">

      <div class="search-container">
        <el-form :inline="true">
          <el-form-item label="ping es">
            <el-button :disabled="loading" v-loading="loading" @click="init">{{$t('click me')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>

  </div>
</template>

<script setup>

import {onMounted, ref} from "vue";
import {HelloAction} from "@/api/helloworld";
import {ElMessage} from "element-plus";
import {sdk} from '@elasticview/plugin-sdk'

const loading = ref(false)

const init = async ()=>{
  loading.value = true
  let res = await HelloAction({
    es_connect:sdk.GetSelectEsConnID()
  })
  loading.value = false

  if (res.code !== 0) {
    ElMessage({
      type: 'error',
      message: res.msg
    })
    return
  }
  ElMessage({
    type: 'success',
    message: JSON.stringify(res.data.text, null, '\t')
  })


}

onMounted(()=>{
  init()
})

</script>

<style scoped>

</style>
