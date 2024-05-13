<template>
  <div class="container" v-if="item">
    <a-space>
      <a-button type="link" @click="goBack">
        <ArrowLeftOutlined />返回投票列表页
      </a-button>
      <a-button v-if="userInfo && userInfo.role===1" type="primary" @click="visible = true">添加候选人</a-button>
      <a-modal v-model:open="visible" title="Create a new collection" ok-text="Create" cancel-text="Cancel"
        @ok="onOk(formState)">
        <a-form :model="formState" layout="vertical" name="form_in_modal">
          <a-form-item name="candidateid" label="候选人ID" :rules="[{ required: true, message: '请输入添加候选人ID' }]">
            <a-input v-model:value="formState.candidateid" />
          </a-form-item>
        </a-form>
      </a-modal>
    </a-space>
    <a-divider v-if="userInfo && userInfo.role===1" />
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
import participateService from '@/service/participateService';
import { message } from 'ant-design-vue';
import { ArrowLeftOutlined } from '@ant-design/icons-vue';
import { mapState } from 'vuex';

export default {
  components: {
    ArrowLeftOutlined,
  },
  data() {
    return {
      formRef: null,
      visible: false,
      formState: { candidateid: 0 },
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
    onOk(values) {
      this.visible = false;
      participateService.addparticipation({ voteid: this.itemId, candidateid: values.candidateid })
        .then(() => {
          message.success('添加成功');
        }).catch((err) => {
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
        if (now < deadline) {
          return '投票中';
        }
        return this.item.outs;
      },
    }),
  },
};
</script>

<style scoped>
.container {
  margin: 0 auto;
  max-width: 800px;
  padding: 0 20px;
}
</style>
