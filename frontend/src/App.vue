<template>
  <div>
    <el-config-provider :size="size">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>

import {sdk} from '@elasticview/plugin-sdk'
import {computed, onMounted} from "vue";

const size = computed(()=>{
  return sdk.getLayoutSize()
})

onMounted(()=>{
  sdk.subToChannel("live",(res)=>{
    console.log(res)
  })
  setTimeout(()=>{
    sdk.callToChannel("live",{
      test:520
    })
  })
})


</script>
