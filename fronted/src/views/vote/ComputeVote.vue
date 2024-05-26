<template>
  <a-button type="primary" @click="handleOpen" v-if="state===2">开始计票</a-button>
</template>
<script>
import { defineComponent } from 'vue';
import voteService from '@/service/voteService';
import { message } from 'ant-design-vue';

export default defineComponent({
  props: {
    voteid: {
      type: Number,
      required: true,
    },
    state: {
      type: Number,
      required: true,
    },
  },
  methods: {
    handleOpen() {
      voteService.computevote({ id: this.voteid })
        .then(() => {
          return voteService.computevote0({ id: this.voteid });
        })
        .then((res0) => {
          return voteService.computevote1({ id: this.voteid })
            .then((res1) => {
              return voteService.outs({
                id: this.voteid,
                outs0: res0.data.outs,
                outs1: res1.data.outs,
              });
            });
        })
        .then(() => {
          message.success('开始计票');
        })
        .catch((err) => {
          // 添加失败
          message.error(err.response.data.msg);
        });
    },

  },
});
</script>
