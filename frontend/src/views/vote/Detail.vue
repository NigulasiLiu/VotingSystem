<template>
  <div class="container" v-if="item">
    <a-space>
      <a-button type="link" @click="goBack">
        <ArrowLeftOutlined />返回投票列表页
      </a-button>
      <AddParticipate :voteid="parseInt(itemId,10)" :state="item.state" v-if="userInfo && userInfo.role===1" />
      <OpenVote :voteid="parseInt(itemId,10)" :state="item.state" v-if="userInfo && userInfo.role===1" />
      <EndVote :voteid="parseInt(itemId,10)" :state="item.state" :deadline="item.deadline"
        v-if="userInfo && userInfo.role===1" />
<!--      <SelectVoted :voteid="parseInt(itemId,10)" :state="item.state" :n="item.count" :num="item.num" :options="options"-->
<!--        :deadline="item.deadline" />-->
      <ComputeVote :voteid="parseInt(itemId,10)" :state="item.state" v-if="userInfo && userInfo.role===1" />
    </a-space>
    <a-divider :voteid="parseInt(itemId,10)" :state="item.state" v-if="userInfo && userInfo.role===1"  class="small-margin-divider"></a-divider>
    <a-space direction="vertical" class="card-space">
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
                size="large" :data-source="sortedList">
          <template #renderItem="{ item, index }">
            <a-list-item :key="item.ID">
              <a-card :title="`ID&emsp;${item.ID}`" class="equal-card">
                <template #actions>
                  <span class="getResultClass(index)">
                    <span v-if="index === 0 && item.Out" class="gold-medal">🥇</span>
                    <span v-else-if="index === 1 && item.Out" class="silver-medal">🥈</span>
                    <span v-else-if="index === 2 && item.Out" class="bronze-medal">🥉</span>
                    投票结果：{{ item.Out ? item.Out : '进行中' }}
                  </span>
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
import { ArrowLeftOutlined } from '@ant-design/icons-vue';
import { mapState } from 'vuex';
import { reactive } from 'vue';
import AddParticipate from '@/views/vote/AddParticipate.vue';
import OpenVote from '@/views/vote/OpenVote.vue';
import EndVote from './EndVote.vue';
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
    ArrowLeftOutlined,
    OpenVote,
    EndVote,
    ComputeVote,
  },
  data() {
    return {
      item: null,
    };
  },
  created() {
    this.itemId = this.$route.params.id;
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
        }).catch((err) => {
          message.error(err.response.data.msg);
        });
    },
    GetParticipateData(voteid) {
      return participateService.showparticipate({ voteid }).then((res) => {
        return res.data.data;
      }).catch((err) => {
        message.error(err.response.data.msg);
      });
    },
    goBack() {
      this.$router.go(-1); // 返回上一页
    },
    // 根据排名返回不同的样式类名
    getResultClass(index) {
      if (index === 0) {
        return 'gold-medal-text';
      } if (index === 1) {
        return 'silver-medal-text';
      } if (index === 2) {
        return 'bronze-medal-text';
      }
      return '';
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
        return (this.participateData.list || []).map((item) => ({ value: item.X, label: `${item.X} ${item.Name}` }));
      },
    }),
    // 对列表进行排序，Out值从大到小排列
    sortedList() {
      return [...this.participateData.list].sort((a, b) => b.Out - a.Out);
    },
  },
};
</script>

<style scoped>
.container {
  margin: 30px auto ; /* 调整上边距 */
  max-width: 1350px;
  padding: 0 20px;
}

.small-margin-divider {
  margin-top: -5px;
}

.card-space {
  margin-top: 15px;
}
.equal-card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 250px; /* 根据需要调整最小高度 */
  max-height: 400px; /* 根据需要调整最大高度 */
}
.gold-medal-text {
  font-weight: bold;
  //color: #ffd700; /* 金色 */
}

.silver-medal-text {
  font-weight: bold;
  //color: #c0c0c0; /* 银色 */
}

.bronze-medal-text {
  font-weight: bold;
  //color: #cd7f32; /* 铜色 */
}
</style>
