<template>
  <div class="container">
    <a-card title="添加候选人" style="text-align:center; width:25rem">
      <a-form :model="candidate" name="basic" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
        autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="姓名" name="name">
          <a-input v-model:value="candidate.name" />
        </a-form-item>
        <a-form-item label="描述" name="detail">
          <a-textarea v-model:value="candidate.detail" />
        </a-form-item>
        <a-form-item :wrapper-col="{span: 24 }">
          <a-button type="primary" html-type="submit" style="width: 40%">提&emsp;交</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<script setup>
import { reactive } from 'vue';
import { message } from 'ant-design-vue';
import candidateService from '@/service/candidateService';

const candidate = reactive({
  name: '',
  detail: '',
});

const validateName = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入姓名'));
  }
  if (value.length > 10) {
    return Promise.reject(new Error('姓名不大于10位'));
  }
  return Promise.resolve();
};

const validateDetail = async (_rule, value) => {
  if (value.length > 50) {
    return Promise.reject(new Error('描述不大于50个字'));
  }
  return Promise.resolve();
};

const rules = {
  name: {
    required: true,
    validator: validateName,
    trigger: 'blur',
  },
  detail: {
    required: false,
    validator: validateDetail,
    trigger: 'blur',
  },
};

const onFinish = (values) => {
  // 添加候选人
  candidateService.addcandidate({ name: values.name, detail: values.detail })
    .then(() => {
      message.success('添加成功');
    }).catch((err) => {
      message.error(err.response.data.msg);
    });
};

const onFinishFailed = (errors) => {
  console.log(errors);
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
