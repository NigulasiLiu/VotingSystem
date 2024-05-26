<template>
  <a-button type="primary" @click="visible = true" v-if="state===1 && !isEnd">投票</a-button>
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
      return new Promise((resolve, reject) => {
        const selectedValue = this.value;
        dpf.dpfGen(selectedValue, 1, this.n, this.voteid)
          .then((value) => {
            if (value) message.success('投票成功');
            else message.success('投票失败');
            window.location.reload(); // 刷新页面
            resolve();
          })
          .catch((error) => {
            console.error('dpf.dpfGen 函数出现错误：', error);
            reject(error);
          });
      });
    },
    onOk() {
      votedService.addvoted({ voteid: this.voteid, userid: this.userInfo.id })
        .then(() => {
          this.visible = false;
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
      isEnd() {
        const now = new Date();
        const deadline = new Date(this.deadline);
        if (now < deadline) return false;
        return true;
      },
    }),
  },
});
</script>
