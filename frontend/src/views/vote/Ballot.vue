<template>
  <div class="container">
    <h1 class="title">登录以投票！</h1>
    <a-form :model="voteForm" @finish="onSubmit" @finishFailed="onSubmitFailed" :rules="voterules" class="form" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }">
      <a-form-item label="投票密钥" name="accessKey" rules="[ { required: true, message: '请输入访问密钥' } ]" class="form-item">
        <a-input v-model:value="voteForm.accessKey" placeholder="请输入16位访问密钥" class="input-box" />
      </a-form-item>
      <p class="access-instruction">投票密钥是发起者提供的16字符代码</p>
      <a-form-item>
        <a-button type="primary" html-type="submit" class="submit-button">立即投票</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { message } from 'ant-design-vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router'; // 引入路由

const store = useStore();
const router = useRouter(); // 创建路由实例

const voteForm = reactive({
  accessKey: '',
});

const validateKey = async (_rule, value) => {
  if (!value) {
    return Promise.reject(new Error('请输入投票密钥'));
  }
  if (value.length === 16) {
    return Promise.resolve(); // 验证通过
  }
  return Promise.reject(new Error('访问密钥必须为16位'));
};

const onSubmit = () => {
  // console.log('Submitted voteKey:', voteForm.accessKey);
  store.dispatch('userModule/ballot', { voteKey: voteForm.accessKey })
    .then((res) => {
      const { voteID } = res.data.data;
      console.log('Submitted voteKey:', res.data.voteID);
      message.success('密钥验证成功，正在跳转...');
      // 路由跳转到 'vote/key' 页面
      router.push({ path: `/vote/${voteID}` });
    })
    .catch((err) => {
      console.error(err);
      message.error(err.response.data.msg || '密钥验证失败，请重试');
    });
};

const onSubmitFailed = (errors) => {
  console.log(errors);
};

const voterules = {
  accessKey: {
    required: true,
    validator: validateKey,
    trigger: 'blur',
  },
};
</script>

<style scoped>
/* 样式保持不变 */
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 50vh;
  padding: 20px;
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.title {
  font-size: 2.5rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 20px;
  margin-left: 18px;
  text-align: center;
  width: 100%;
}

.form {
  width: 100%;
  max-width: 300px;
}

.form-item {
  margin-bottom: 20px;
}

.input-box {
  width: 100%;
  height: 45px;
  font-size: 1rem;
}

.access-instruction {
  font-size: 0.9rem;
  color: #666;
  margin-bottom: 20px;
  text-align: center;
}

.submit-button {
  width: 50%;
  background-color: orange;
  border-color: orange;
  font-size: 1.2rem;
  height: 45px;
  margin-left: 6ch;
}

.submit-button:hover {
  background-color: darkorange;
  border-color: darkorange;
}
</style>
