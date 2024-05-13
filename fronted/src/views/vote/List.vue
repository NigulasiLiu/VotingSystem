<template>
  <div class="container">
    <a-space>
      <a-button v-if="userInfo && userInfo.role===1" type="primary"
        @click="$router.replace({name:'addvote'})">新增投票</a-button>
      <a-button v-if="userInfo && userInfo.role===1" type="primary"
        @click="$router.replace({name:'candidatelist'})">管理候选人</a-button>
    </a-space>
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
  max-width: 800px;
  padding: 0 20px;
}
</style>
