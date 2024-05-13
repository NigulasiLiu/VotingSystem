<template>
  <div class="container">
    <a-card title="添加投票" style="text-align:center; width:25rem">
      <a-form :model="vote" name="basic" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
        autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="投票名" name="name">
          <a-input v-model:value="vote.name" />
        </a-form-item>
        <a-form-item name="deadline" label="截止时间">
          <a-date-picker v-model:value="vote.deadline" show-time format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss" :disabled-date="disabledDate" style="width: 100%" />
        </a-form-item>
        <a-form-item label="可投票数" name="num">
          <a-input-number v-model:value="vote.num" :min="1" :max="10" style="width: 100%" />
        </a-form-item>
        <a-form-item :wrapper-col="{span: 24 }">
          <a-button type="primary" html-type="submit" style="width: 40%">提&emsp;交</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<script setup>
import voteService from '@/service/voteService';
import dayjs from 'dayjs';
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { message } from 'ant-design-vue';

const router = useRouter();

const vote = reactive({
  name: '',
  num: 0,
  deadline: '',
});

const validateName = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入投票名'));
  }
  if (value.length > 30) {
    return Promise.reject(new Error('投票名不大于30位'));
  }
  return Promise.resolve();
};

const validateDDL = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请选择截止时间'));
  }
  if (dayjs().isAfter(dayjs(value))) {
    return Promise.reject(new Error('投票时间不早于当前时间'));
  }
  return Promise.resolve();
};

const validateNum = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入可投票数'));
  }
  if (value < 1 || value >= 10) {
    return Promise.reject(new Error('可投票数不正确'));
  }
  return Promise.resolve();
};

const rules = {
  name: {
    required: true,
    validator: validateName,
    trigger: 'blur',
  },
  num: {
    required: true,
    validator: validateNum,
    trigger: 'blur',
  },
  deadline: {
    required: true,
    validator: validateDDL,
    trigger: 'blur',
  },
};

const onFinish = (values) => {
  voteService.addvote({ name: values.name, num: values.num, deadline: values.deadline })
    .then(() => {
      message.success('添加成功');
      router.push({ path: '/votinglist' });
    }).catch((err) => {
      message.error(err.response.data.msg);
    });
};

const onFinishFailed = (errors) => {
  console.log(errors);
};

const disabledDate = (current) => {
  // Can not select days before today and today
  return current < dayjs().endOf('day');
};

</script>
<style scoped>
.container {
  text-align: -webkit-center;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 80vh;
}
</style>
