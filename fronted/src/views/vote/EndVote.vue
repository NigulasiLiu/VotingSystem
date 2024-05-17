<template>
  <a-button @click="handleEnd" v-if="state==1 && !isEnd">截止投票</a-button>
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
      type: Boolean,
      required: true,
    },
    deadline: {
      type: String,
      required: true,
    },
  },
  methods: {
    handleEnd() {
      voteService.endvote({ id: this.voteid })
        .then(() => {
          // 添加成功
          message.success('已截止');
        })
        .catch((err) => {
          // 添加失败
          message.error(err.response.data.msg);
        });
    },
  },
  computed: {
    isEnd() {
      const now = new Date();
      const deadline = new Date(this.deadline);
      if (now < deadline) return false;
      return true;
    },
  },
});
</script>
