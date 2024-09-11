<template>
  <div class="container">
    <div class="image-container">
      <img src="@/assets/login-image.png" alt="Login Image" class="login-image" />
    </div>
    <div class="form-container">
      <a-card title="发起者登录" style="text-align: center;width:25rem" body-style="{ paddingTop: '12px' }">
        <div style="margin-top: 0.1rem; margin-bottom: 0.5rem; text-align: center;">
          <span>需要投票，请访问</span>
          <a-button type="link" @click="openVotingPage" style="padding-left: 0;">投票页面</a-button>
        </div>
        <a-form :model="userlogin" name="basic" :rules="loginrules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
                autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
          <a-form-item label="手机号" name="tel" style="margin-bottom: 8px;">
            <a-input v-model:value="userlogin.tel" />
          </a-form-item>

          <a-form-item label="密码" name="pass">
            <a-input-password v-model:value="userlogin.pass" />
          </a-form-item>

          <a-form-item :wrapper-col="{span: 24 }">
            <a-button type="primary" html-type="submit" style="width: 50%">登&emsp;录</a-button>
          </a-form-item>
        </a-form>
        <a-button type="link" @click="$router.replace({name:'register'})">还没注册？</a-button>
      </a-card>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';

const router = useRouter();
const store = useStore();

const userlogin = reactive({
  tel: '',
  pass: '',
});

const openVotingPage = () => {
  const votingUrl = router.resolve({ name: 'ballot' }).href;
  window.open(votingUrl, '_blank');
};

const validateTel = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入手机号'));
  }
  if (value.length !== 11) {
    return Promise.reject(new Error('手机号应为11位'));
  }
  if (value !== '' && !/^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9]|19[0|1|2|3|4|5|6|7|8|9])\d{8}$/.test(value)) {
    return Promise.reject(new Error('手机号格式不正确'));
  }
  return Promise.resolve();
};

const validatePass = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入密码'));
  }
  if (value.length < 6 || value.length > 30) {
    return Promise.reject(new Error('密码应为6到30位'));
  }
  return Promise.resolve();
};

const loginrules = {
  tel: {
    required: true,
    validator: validateTel,
    trigger: 'blur',
  },
  pass: {
    required: true,
    validator: validatePass,
    trigger: 'blur',
  },
};

const onFinish = (values) => {
  store.dispatch('userModule/login', { telephone: values.tel, password: values.pass }).then(() => {
    message.success('登录成功');
    router.push({ path: '/' });
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
  display: flex;
  align-items: center;
  justify-content: center;
}

.image-container {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-image {
  max-width: 80%;
  height: auto;
}

.form-container {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
