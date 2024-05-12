<template>
  <div class="container">
    <a-button v-if="userInfo && userInfo.role===1" type="primary" @click="$router.replace({name:'addvote'})">管理</a-button>
    <a-divider v-if="userInfo && userInfo.role===1" />
    <a-table :columns="voteData.columns" :data-source="voteData.list">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <span>
            <router-link :to="{ name: 'detail', params: { id: record.ID }}">详情</router-link>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>
<script>
import { voteData } from '@/data/votedata';
import { mapState } from 'vuex';

export default {
  data() {
    return {
      voteData,
    };
  },
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>
<style scoped>
.container {
  margin: 0 auto;
  /* 让容器水平居中 */
  max-width: 800px;
  /* 设置最大宽度 */
  padding: 0 20px;
  /* 左右边距 */
}
</style>
