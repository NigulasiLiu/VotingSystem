<template>
    <a-row class="mt-5">
        <a-col md="8" offset-md="2" lg="6" offset-lg="3">
            <a-card title="注册" style="text-align: center;">
                <a-form :model="user" name="basic" :rules="rules" :label-col="{ span: 7 }" :wrapper-col="{ span: 17 }"
                    autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
                    <a-form-item label="手机号" name="tel">
                        <a-input v-model:value="user.tel" />
                    </a-form-item>

                    <a-form-item label="密码" name="pass">
                        <a-input-password v-model:value="user.pass" />
                    </a-form-item>

                    <a-form-item label="确认密码" name="checkPass">
                        <a-input-password v-model:value="user.checkPass" />
                    </a-form-item>

                    <a-form-item :wrapper-col="{span: 24 }">
                        <a-button type="primary" html-type="submit" style="width: 40%">注&emsp;册</a-button>
                    </a-form-item>
                </a-form>
            </a-card>
        </a-col>
    </a-row>
</template>
<script>
import { defineComponent, reactive } from 'vue';

const user = reactive({
  tel: '',
  pass: '',
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
  if (value !== user.pass) {
    return Promise.reject(new Error('两次输入不一致'));
  }
  return Promise.resolve();
};

const rules = {
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
  checkPass: {
    required: true,
    validator: validatePass2,
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
      user,
      rules,
      handleFinish,
      handleFinishFailed,
    };
  },
});
</script>
