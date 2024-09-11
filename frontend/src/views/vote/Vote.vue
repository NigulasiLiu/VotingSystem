<template>
  <div class="container" v-if="item">
    <a-space>
      <AddParticipate :voteid="parseInt(itemId, 10)" :state="item.state" />
      <OpenVote :voteid="parseInt(itemId, 10)" :state="item.state" />
      <SelectVoted :voteid="parseInt(itemId, 10)" :state="item.state" :n="item.count" :num="item.num" :options="options"
                   :deadline="item.deadline" />
      <ComputeVote :voteid="parseInt(itemId, 10)" :state="item.state" />
    </a-space>
    <a-divider class="small-margin-divider"></a-divider>
    <a-space direction="vertical">
      <a-card>
        <a-descriptions :column="2">
          <a-descriptions-item label="投票名称">{{ item.name }}</a-descriptions-item>
          <a-descriptions-item label="状态">{{ voteStatus }}</a-descriptions-item>
          <a-descriptions-item label="可投票数">{{ item.num }}</a-descriptions-item>
          <a-descriptions-item label="截止日期">{{ item.deadline }}</a-descriptions-item>
        </a-descriptions>
      </a-card>
      <a-divider>候选人列表</a-divider>
      <a-list v-if="participateData.list" :grid="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 3, xxl: 4 }" item-layout="vertical"
              size="large" :data-source="participateData.list">
        <template #renderItem="{ item }">
          <a-list-item :key="item.ID">
            <a-card :title="`ID&emsp;${item.ID}`" class="equal-card">
              <template #actions>
                <span>投票结果：{{ item.Out }}</span>
              </template>
              <a-list-item-meta>
                <template #title>
                  <a>{{ item.Name }}</a>
                </template>
                <template #avatar>
                  <a-avatar v-if="participateData" :src="`data:image/png;base64,${item.Photo}`" />
                </template>
              </a-list-item-meta>
              {{ item.Detail }}
            </a-card>
          </a-list-item>
        </template>
      </a-list>
    </a-space>
  </div>
</template>

<script>
import voteService from '@/service/voteService';
import participateService from '@/service/participateService';
import { message } from 'ant-design-vue';
import { mapState } from 'vuex';
import { reactive } from 'vue';
import AddParticipate from '@/views/vote/AddParticipate.vue';
import SelectVoted from '@/views/vote/SelectVoted.vue';
import OpenVote from '@/views/vote/OpenVote.vue';
import ComputeVote from './ComputeVote.vue';

export default {
  setup() {
    const participateData = reactive({
      list: reactive([]),
    });

    return {
      participateData,
    };
  },
  components: {
    AddParticipate,
    SelectVoted,
    OpenVote,
    ComputeVote,
  },
  data() {
    return {
      item: null,
    };
  },
  created() {
    this.itemId = this.$route.params.key; // 修改为通过投票密钥获取
    this.fetchVoteInfo(this.itemId);
    this.GetParticipateData(this.itemId).then((data) => {
      this.participateData.list = data;
    });
  },
  methods: {
    fetchVoteInfo(value) {
      voteService.voteinfo({ id: value })
        .then((response) => {
          this.item = response.data.data.vote;
        })
        .catch((err) => {
          message.error(err.response.data.msg);
        });
    },
    GetParticipateData(voteid) {
      return participateService.showparticipate({ voteid })
        .then((res) => {
          return res.data.data;
        })
        .catch((err) => {
          message.error(err.response.data.msg);
        });
    },
  },
  computed: {
    ...mapState({
      userInfo: (state) => state.userModule.userInfo,
      voteStatus() {
        if (!this.item) return '';
        const now = new Date();
        const deadline = new Date(this.item.deadline);
        if (this.item.state === 0 && now < deadline) return '投票未开放';
        if (this.item.state === 1 && now < deadline) return '投票中';
        if (this.item.state === 3) return '计票中';
        if (this.item.state === 4) return '已结束';
        return '投票已截止';
      },
      options() {
        return (this.participateData.list || []).map((item) => ({
          value: item.X,
          label: `${item.X} ${item.Name}`,
        }));
      },
    }),
  },
};
</script>

<style scoped>
.container {
  margin: 30px auto;
  max-width: 1350px;
  padding: 0 20px;
}

.small-margin-divider {
  margin-top: -5px;
}

.equal-card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 250px;
  max-height: 400px;
}
</style>
