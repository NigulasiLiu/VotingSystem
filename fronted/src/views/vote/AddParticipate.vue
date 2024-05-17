<template>
  <a-button type="primary" @click="visible = true" v-if="state===0">添加候选人</a-button>
  <a-modal v-model:open="visible" ok-text="添加" cancel-text="取消" @ok="onOk(formState)">
    <a-form :model="formState" name="form_in_modal" :rules="rules" ref="formRef">
      <a-form-item name="candidateid" label="候选人ID">
        <a-input-number v-model:value="formState.candidateid" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script>
import participateService from '@/service/participateService';
import { message } from 'ant-design-vue';
import { defineComponent } from 'vue';

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
  setup() {
    const validateID = async (_rule, value) => {
      if (value === '') {
        return Promise.reject(new Error('请输入候选人ID'));
      }
      if (value <= 0) {
        return Promise.reject(new Error('ID应大于0'));
      }
      return Promise.resolve();
    };
    const rules = {
      candidateid: [
        { required: true, validator: validateID, trigger: 'blur' },
      ],
    };
    return {
      rules,
    };
  },
  data() {
    return {
      formRef: null,
      visible: false,
      formState: { candidateid: null },
    };
  },
  methods: {
    onOk(values) {
      this.$refs.formRef
        .validateFields()
        .then(() => {
        // 数据校验通过后执行添加投票操作
          participateService.addparticipation({ voteid: this.voteid, candidateid: values.candidateid })
            .then(() => {
              this.visible = false;
              message.success('添加成功');
              return participateService.addparticipation0({ voteid: this.voteid, candidateid: values.candidateid });
            }).catch((err) => {
              message.error(err.response.data.msg);
            });
        })
        .catch(() => {
        // 数据校验失败
          message.error('请正确填写表单');
        });
    },
  },
});
</script>
