<template>
  <div  class="app-container">
    <div  class="search-container">
      <el-form>
        <el-form-item>
          <el-tag>
            数据库操作 demo
          </el-tag>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="list"
        style="width: 100%;"
        border
    >

      <el-table-column
          prop="key"
          label="键"
      >
        <template #default="{ row, $index }">
          <el-input
              v-model="row.key"
              placeholder="键"
          />
        </template>
      </el-table-column>

      <el-table-column
          prop="value"
          label="值"
      >
        <template #default="{ row, $index }">
          <el-input
              v-model="row.value"
              placeholder="值"
          />
        </template>
      </el-table-column>

      <el-table-column
          label="操作"
          width="300"
      >
        <template #header>

          <el-button  @click="searchTest(1)">
            {{ $t('刷新') }}
          </el-button>
          <el-button
              type="primary"
              @click="addHeader"
          >
            新增
          </el-button>
          <el-button type="danger" @click="clean">
            {{ $t('清空') }}
          </el-button>

        </template>
        <template #default="{ row }">
          <el-button
              @click="save(row)"
              type="info"
          >保存</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination
          :current-page="input.page"
          :page-sizes="[10, 20, 30, 50,100,150,200]"
          :page-size="input.limit"
          layout="total, sizes, prev, pager, next"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="searchTest"
      />
    </div>
  </div>
</template>

<script>
import {ElMessage} from "element-plus";
import {ref, onMounted,onActivated} from 'vue';
import { CleanAction, ListAction,DbInsert} from '@/api/db_test';

export default {
  name: 'DbTest',
  setup(props, { emit }) {
    const input = ref({
      page: 1,
      limit: 10,
    });
    const list = ref([]);
    const total = ref(0);
    const searchTest = async (page = 1) => {

      input.value.page = page

      const { data, code, msg } = await ListAction(input.value);
      if (code !== 0) {
        ElMessage({ type: 'error', message: msg });
      } else {
        if(data.list == null) data.list = []
        list.value = data.list;
        total.value = data.count;
      }
    };

    const handleSizeChange = (v) => {
      input.value.limit = v;
      searchTest(1);
    };

    const clean = async () => {
      const { code, msg } = await CleanAction();
      if (code !== 0) {
        ElMessage({ type: 'error', message: msg });
      } else {
        ElMessage({ type: 'success', message: msg });
        searchTest();
      }
    };

    const addHeader = ()=>{
      list.value.unshift({key:'',value:''})
    }

    const close = () => {
      emit('close', false);
    };

    onMounted(() => {
      searchTest();
    });
    onActivated(()=>{

    })

    const save = async (row)=>{

      const { code, msg } = await DbInsert({
        key:row.key,
        value:row.value,
      });
      if (code !== 0) {
        ElMessage({ type: 'error', message: msg });
      } else {
        ElMessage({ type: 'success', message: msg });
      }
    }

    return {
      save,
      input,
      list,
      total,
      handleSizeChange,
      clean,
      close,
      searchTest,
      addHeader,
    };
  },
};
</script>

<style scoped>

</style>
