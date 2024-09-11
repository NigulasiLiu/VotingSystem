<template>
  <div class="container">
    <a-space>
      <a-button
        v-if="userInfo && userInfo.role === 1"
        type="primary"
        size="large"
        @click="$router.replace({ name: 'addvote' })">
        <PlusOutlined /> 新增投票
      </a-button>
      <a-button
        v-if="userInfo && userInfo.role === 1"
        size="large"
        @click="$router.replace({ name: 'candidatelist' })">
        <ToolOutlined /> 管理候选人
      </a-button>
    </a-space>
    <a-divider v-if="userInfo && userInfo.role === 1"/>
    <a-table :columns="mergedColumns" :data-source="filteredVoteData">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'details'">
          <div class="details-column">
            <div class="vote-name">{{ record.Name }}</div>
            <div class="state-id-row">
              <a-tag :color="record.State === 1 ? 'green' : 'red'" class="custom-tag">
                {{ record.State === 1 ? '开放' : '未开放' }}
              </a-tag>
              <span class="vote-id">ID: {{ record.ID }}</span>
            </div>
          </div>
        </template>
        <template v-if="column.key === 'deadline'">
          <div class="deadline-column">
            <CalendarOutlined class="deadline-icon" />
            <span class="deadline-text">{{ record.Deadline }}</span>
          </div>
        </template>
        <template v-if="column.key === 'action'">
          <span class="operation-column">
            <router-link :to="{ name: 'detail', params: { id: record.ID } }">详情</router-link>
            <a-button type="link" @click="handleDelete(record.ID)" class="delete-button">删除</a-button>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script>
import { voteData } from '@/data/votedata';
import { mapState } from 'vuex';
import { PlusOutlined, ToolOutlined, CalendarOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';

export default {
  data() {
    return {
      voteData,
      deletedIds: [], // 存储删除的ID
      mergedColumns: [
        {
          title: '投票详情',
          key: 'details',
          scopedSlots: { customRender: 'details' },
        },
        {
          title: '截止时间',
          key: 'deadline',
          scopedSlots: { customRender: 'deadline' },
        },
        {
          title: '操作',
          key: 'action',
          scopedSlots: { customRender: 'action' },
        },
      ],
    };
  },
  computed: {
    ...mapState({
      userInfo: (state) => state.userModule.userInfo,
    }),
    // 过滤voteData.list，排除掉deletedIds中的ID
    filteredVoteData() {
      return this.voteData.list.filter((item) => !this.deletedIds.includes(item.ID));
    },
  },
  methods: {
    handleDelete(id) {
      // 将删除的ID存入deletedIds数组
      this.deletedIds.push(id);
      // 这里可以加入API调用来删除后台的数据
      message.success('投票已删除');
    },
  },
  components: {
    PlusOutlined,
    ToolOutlined,
    CalendarOutlined,
  },
};
</script>

<style scoped>
.container {
  margin: 50px auto; /* 调整上边距 */
  max-width: 1350px;
  padding: 0 20px;
}

/* 自定义 Details 栏样式 */
.details-column {
  text-align: left;
  font-size: 16px;
}

.vote-name {
  font-size: 18px;
  font-weight: bold;
}

.state-id-row {
  display: flex;
  align-items: center;
  margin-top: 8px;
}

.state-id-row .custom-tag {
  margin-right: 16px;
  width: 60px; /* 固定标签的宽度，确保对齐 */
  text-align: center; /* 使标签内容居中 */
}

.vote-id {
  font-size: 14px;
  color: #666;
}

/* 自定义 a-tag 样式 */
.custom-tag {
  font-size: 14px; /* 调整标签的字体大小 */
  padding: 2px 8px; /* 调整标签的内边距 */
  height: auto; /* 让标签的高度自动适应内容 */
  width: 60px; /* 固定标签宽度 */
}

/* 美化截止时间 */
.deadline-column {
  display: flex;
  font-size: 16px;
}

.deadline-icon {
  color: #1890ff;
  margin-right: 8px;
}

.deadline-text {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

/* 美化操作列 */
.operation-column {
  display: flex;
  align-items: center;
  gap: 10px; /* 添加间距 */
  font-size: 16px;
}

/* 美化删除按钮 */
.delete-button {
  color: #ff4d4f;
}
</style>

<!--<template>-->
<!--  <div class="container">-->
<!--    <a-space>-->
<!--      <a-button-->
<!--        v-if="userInfo && userInfo.role === 1"-->
<!--        type="primary"-->
<!--        size="large"-->
<!--        @click="$router.replace({ name: 'addvote' })">-->
<!--        <PlusOutlined /> 新增投票-->
<!--      </a-button>-->
<!--      <a-button-->
<!--        v-if="userInfo && userInfo.role === 1"-->
<!--        size="large"-->
<!--        @click="$router.replace({ name: 'candidatelist' })">-->
<!--        <ToolOutlined /> 管理候选人-->
<!--      </a-button>-->
<!--    </a-space>-->
<!--    <a-divider v-if="userInfo && userInfo.role === 1"/>-->
<!--    <a-table :columns="mergedColumns" :data-source="voteData.list">-->
<!--      <template #bodyCell="{ column, record }">-->
<!--        <template v-if="column.key === 'details'">-->
<!--          <div class="details-column">-->
<!--            <div class="vote-name">{{ record.Name }}</div>-->
<!--            <div class="state-id-row">-->
<!--              <a-tag :color="record.State === 1 ? 'green' : 'red'" class="custom-tag">-->
<!--                {{ record.State === 1 ? '开放' : '未开放' }}-->
<!--              </a-tag>-->
<!--              <span class="vote-id">ID: {{ record.ID }}</span>-->
<!--            </div>-->
<!--          </div>-->
<!--        </template>-->
<!--        <template v-if="column.key === 'deadline'">-->
<!--          <div class="deadline-column">-->
<!--            <CalendarOutlined class="deadline-icon" />-->
<!--            <span class="deadline-text">{{ record.Deadline }}</span>-->
<!--          </div>-->
<!--        </template>-->
<!--        <template v-if="column.key === 'action'">-->
<!--          <span class="operation-column">-->
<!--            <router-link :to="{ name: 'detail', params: { id: record.ID } }">详情</router-link>-->
<!--            <a-button type="link" @click="handleDelete(record.ID)" class="delete-button">删除</a-button>-->
<!--          </span>-->
<!--        </template>-->
<!--      </template>-->
<!--    </a-table>-->
<!--  </div>-->
<!--</template>-->

<!--<script>-->
<!--import { voteData } from '@/data/votedata';-->
<!--import { mapState } from 'vuex';-->
<!--import { PlusOutlined, ToolOutlined, CalendarOutlined } from '@ant-design/icons-vue';-->

<!--export default {-->
<!--  data() {-->
<!--    return {-->
<!--      voteData,-->
<!--      mergedColumns: [-->
<!--        {-->
<!--          title: '投票详情',-->
<!--          key: 'details',-->
<!--          scopedSlots: { customRender: 'details' },-->
<!--        },-->
<!--        {-->
<!--          title: '截止时间',-->
<!--          key: 'deadline',-->
<!--          scopedSlots: { customRender: 'deadline' },-->
<!--        },-->
<!--        {-->
<!--          title: '操作',-->
<!--          key: 'action',-->
<!--          scopedSlots: { customRender: 'action' },-->
<!--        },-->
<!--      ],-->
<!--    };-->
<!--  },-->
<!--  computed: mapState({-->
<!--    userInfo: (state) => state.userModule.userInfo,-->
<!--  }),-->
<!--  methods: {-->
<!--    handleDelete(id) {-->
<!--      // 通过过滤掉要删除的ID来更新列表-->
<!--      this.voteData.list = this.voteData.list.filter((item) => item.ID !== id);-->
<!--      // 这里可以加入API调用来删除后台的数据-->
<!--    },-->
<!--  },-->
<!--  components: {-->
<!--    PlusOutlined,-->
<!--    ToolOutlined,-->
<!--    CalendarOutlined,-->
<!--  },-->
<!--};-->
<!--</script>-->

<!--<style scoped>-->
<!--.container {-->
<!--  margin: 50px auto; /* 调整上边距 */-->
<!--  max-width: 1350px;-->
<!--  padding: 0 20px;-->
<!--}-->

<!--/* 自定义 Details 栏样式 */-->
<!--.details-column {-->
<!--  text-align: left;-->
<!--  font-size: 16px;-->
<!--}-->

<!--.vote-name {-->
<!--  font-size: 18px;-->
<!--  font-weight: bold;-->
<!--}-->

<!--.state-id-row {-->
<!--  display: flex;-->
<!--  align-items: center;-->
<!--  margin-top: 8px;-->
<!--}-->

<!--.state-id-row .custom-tag {-->
<!--  margin-right: 16px;-->
<!--  width: 60px; /* 固定标签的宽度，确保对齐 */-->
<!--  text-align: center; /* 使标签内容居中 */-->
<!--}-->

<!--.vote-id {-->
<!--  font-size: 14px;-->
<!--  color: #666;-->
<!--}-->

<!--/* 自定义 a-tag 样式 */-->
<!--.custom-tag {-->
<!--  font-size: 14px; /* 调整标签的字体大小 */-->
<!--  padding: 2px 8px; /* 调整标签的内边距 */-->
<!--  height: auto; /* 让标签的高度自动适应内容 */-->
<!--  width: 60px; /* 固定标签宽度 */-->
<!--}-->

<!--/* 美化截止时间 */-->
<!--.deadline-column {-->
<!--  display: flex;-->
<!--  font-size: 16px;-->
<!--}-->

<!--.deadline-icon {-->
<!--  color: #1890ff;-->
<!--  margin-right: 8px;-->
<!--}-->

<!--.deadline-text {-->
<!--  font-size: 16px;-->
<!--  font-weight: bold;-->
<!--  color: #333;-->
<!--}-->

<!--/* 美化操作列 */-->
<!--.operation-column {-->
<!--  display: flex;-->
<!--  align-items: center;-->
<!--  gap: 10px; /* 添加间距 */-->
<!--  font-size: 16px;-->
<!--}-->

<!--/* 美化删除按钮 */-->
<!--.delete-button {-->
<!--  color: #ff4d4f;-->
<!--}-->
<!--</style>-->
