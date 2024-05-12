<template>
    <div class="container" v-if="item">
        <a-button type="link" @click="goBack">
            <ArrowLeftOutlined />返回投票列表页
        </a-button>

        <a-divider />
        <a-card>
            <a-descriptions :column="2">
                <a-descriptions-item label="投票名称">{{ item.name }}</a-descriptions-item>
                <a-descriptions-item label="可投票数">{{ item.num }}</a-descriptions-item>
                <a-descriptions-item label="截止日期">{{ item.deadline }}</a-descriptions-item>
                <a-descriptions-item label="投票结果">{{ voteStatus }}</a-descriptions-item>
            </a-descriptions>
        </a-card>
    </div>
</template>

<script>
import voteService from '@/service/voteService';
import { message } from 'ant-design-vue';
import { ArrowLeftOutlined } from '@ant-design/icons-vue';

export default {
  components: {
    ArrowLeftOutlined,
  },
  data() {
    return {
      item: null,
    };
  },
  created() {
    this.itemId = this.$route.params.id;
    this.fetchVoteInfo(this.itemId);
  },
  methods: {
    fetchVoteInfo(value) {
      voteService.voteinfo({ id: value })
        .then((response) => {
          this.item = response.data.data.vote;
        }).catch((err) => {
          message.error(err.response.data.msg);
        });
    },
    goBack() {
      this.$router.go(-1); // 返回上一页
    },
  },
  computed: {
    voteStatus() {
      if (!this.item) return '';
      const now = new Date();
      const deadline = new Date(this.item.deadline);
      if (now < deadline) {
        return '投票中';
      }
      return this.item.outs;
    },
  },
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
