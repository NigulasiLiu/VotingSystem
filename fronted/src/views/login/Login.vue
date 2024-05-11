<template>
  <div class="container">
    <a-card title="登录" style="text-align: center;width:25rem">
      <a-form :model="userlogin" name="basic" :rules="loginrules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
        autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="手机号" name="tel">
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
</template>
<script>
import { defineComponent, reactive } from 'vue';

const userlogin = reactive({
  tel: '',
  pass: '',
});

const validateTel = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入手机号'));
  }
  if (value.length !== 11) {
    return Promise.reject(new Error('手机号应为11位'));
  }
  if (value !== '' && !/^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$/.test(value)) {
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

export default defineComponent({

  setup() {
    const handleFinish = (values) => {
      console.log(values);
    };
    const handleFinishFailed = (errors) => {
      console.log(errors);
    };

    return {
      userlogin,
      loginrules,
      handleFinish,
      handleFinishFailed,
    };
  },
});
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
