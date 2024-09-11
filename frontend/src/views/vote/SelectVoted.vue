<template>
  <a-button type="primary" @click="visible = true" v-if="!isEnd">投票</a-button>
  <a-modal v-model:visible="visible" ok-text="投票" cancel-text="取消" @ok="onOk">
    <a-select v-model:value="value" style="width: 80%" :options="options"></a-select>
  </a-modal>
</template>

<script>
import { message } from 'ant-design-vue';
import { defineComponent } from 'vue';
import votedService from '@/service/votedService';
import { mapState } from 'vuex';
import dpf from '@/gen/dpf';

export default defineComponent({
  props: {
    voteid: {
      type: Number,
      required: true,
    },
    num: {
      type: Number,
      required: true,
    },
    options: {
      type: Array,
      required: true,
    },
    state: {
      type: Number,
      required: true,
    },
    n: {
      type: Number,
      required: true,
    },
    deadline: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      visible: false,
      value: null,
    };
  },
  methods: {
    doDpfGen() {
      return dpf.dpfGen(this.value, 1, this.n, this.voteid)
        .then((success) => {
          if (success) {
            message.success('投票成功');
          } else {
            message.error('投票失败');
          }
          window.location.reload(); // 刷新页面
        })
        .catch((error) => {
          console.error('dpf.dpfGen 函数出现错误：', error);
          message.error('An unexpected error occurred');
        });
    },
    onOk() {
      votedService.addvoted({ voteid: this.voteid, voteKey: this.voteKey })
        .then(() => {
          this.visible = false;
          // 显示投票成功的消息
          message.info('进入dpfGen');
          return this.doDpfGen();
        })
        .catch((err) => {
          if (err.response && err.response.data) {
            message.error(err.response.data.msg);
          }
        });
    },
  },

  computed: {
    ...mapState({
      userInfo: (state) => state.userModule.userInfo,
      voteKey: (state) => state.userModule.voteKey, // 从 vuex 获取 voteKey
      isEnd() {
        const now = new Date();
        const deadline = new Date(this.deadline);
        return now >= deadline;
      },
    }),
  },
});
</script>
