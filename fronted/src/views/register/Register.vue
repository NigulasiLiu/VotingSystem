<template>
  <div class="container">
    <a-card title="注册" style="text-align: center;width:25rem">
      <a-form :model="user" name="basic" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
        @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="手机号" name="Telephone">
          <a-input v-model:value="user.Telephone" />
        </a-form-item>

        <a-form-item label="密码" name="Password">
          <a-input-password v-model:value="user.Password" />
        </a-form-item>

        <a-form-item label="确认密码" name="checkPass">
          <a-input-password v-model:value="user.checkPass" />
        </a-form-item>

        <a-form-item :wrapper-col="{span: 24 }">
          <a-button type="primary" html-type="submit" style="width: 40%">注&emsp;册</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<script setup>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';

const router = useRouter();
const store = useStore();

const user = reactive({
  Telephone: '',
  Password: '',
  checkPass: '',
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

const validatePass2 = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请再次输入密码'));
  }
  if (value.length < 6 || value.length > 30) {
    return Promise.reject(new Error('密码应为6到30位'));
  }
  if (value !== user.Password) {
    return Promise.reject(new Error('两次输入不一致'));
  }
  return Promise.resolve();
};

const rules = {
  Telephone: {
    required: true,
    validator: validateTel,
    trigger: 'blur',
  },
  Password: {
    required: true,
    validator: validatePass,
    trigger: 'blur',
  },
  checkPass: {
    required: true,
    validator: validatePass2,
    trigger: 'blur',
  },

};

const onFinish = (values) => {
  store.dispatch('userModule/register', { telephone: values.Telephone, password: values.Password }).then(() => {
    message.success('注册成功');
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
