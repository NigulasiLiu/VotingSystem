<template>
  <a-button @click="handleOpen" v-if="state===0">开放投票</a-button>
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
      voteService.openvote({ id: this.voteid })
        .then(() => {
          // 添加成功
          message.success('投票已开放');
          window.location.reload(); // 刷新页面
        })
        .catch((err) => {
          // 添加失败
          message.error(err.response.data.msg);
        });
    },
  },
});
</script>
