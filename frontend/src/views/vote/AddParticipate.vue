<template>
  <div class="container">
    <a-button type="primary" @click="visible = true" v-if="state === 0" class="add-candidate-button">指定候选人</a-button>
    <a-modal
      v-model:open="visible"
      title="选择候选人"
      ok-text="添加"
      cancel-text="取消"
      @ok="onOk(formState)"
      class="candidate-modal"
      :style="{ top: '50%', transform: 'translateY(-50%)' }"
    >
      <a-form :model="formState" name="form_in_modal" :rules="rules" ref="formRef">
        <a-form-item
          name="candidateid"
          label="候选人ID"
          class="form-item"
          :label-col="{ span: 12 }"
          :wrapper-col="{ span: 12 }"
        >
          <a-select
            v-model:value="formState.candidateid"
            placeholder="请选择候选人"
            class="candidate-select"
            :style="{ transform: 'translateX(10%)' }"
          >
            <a-select-option
              v-for="item in candidateData.list"
              :key="item.ID"
              :value="item.ID"
            >
              {{ item.ID }} - {{ item.Name }}
            </a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import participateService from '@/service/participateService';
import { message } from 'ant-design-vue';
import { defineComponent } from 'vue';
import { candidateData } from '@/data/candidatedata';

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
      if (!value) {
        return Promise.reject(new Error('请选择候选人'));
      }
      return Promise.resolve();
    };

    const rules = {
      candidateid: [{ required: true, validator: validateID, trigger: 'change' }],
    };

    return {
      rules,
      candidateData,
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
          // 数据校验通过后执行添加候选人操作
          participateService.addparticipation({ voteid: this.voteid, candidateid: values.candidateid })
            .then(() => {
              this.visible = false;
              message.success('添加成功');
              window.location.reload(); // 刷新页面
            })
            .catch((err) => {
              message.error(err.response.data.msg);
            });
        })
        .catch(() => {
          message.error('请正确填写表单');
        });
    },
  },
});
</script>

<style scoped>
.container {
  margin: 0px auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

.candidate-modal {
  display: flex;
  align-items: center;
  justify-content: center;
}

.form-item {
  display: flex;
  align-items: center;
}

.candidate-select {
  width: 300%; /* 根据需要调整宽度 */
}
</style>
