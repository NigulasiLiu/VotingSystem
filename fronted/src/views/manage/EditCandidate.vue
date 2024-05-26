<template>
  <div class="container">
    <a-card :title="`修改候选人信息&emsp;ID: ${itemId}`" style="text-align:center; width:25rem">
      <a-upload v-model:file-list="fileList" list-type="fileList" :max-count="1"
        :action='`http://localhost:1016/api/auth/photo/${itemId}`' name="photo">
        <a-button>
          <upload-outlined></upload-outlined>
          Upload (Max: 1)
        </a-button>
      </a-upload>
      <a-divider></a-divider>
      <a-form :model="candidate" name="basic" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
        autocomplete="off" @finish="onFinish">
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

<script>
import { message } from 'ant-design-vue';
import candidateService from '@/service/candidateService';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      fileList: [],
      candidate: {
        detail: '',
      },
      rules: {
        detail: [
          { required: false, message: '请输入描述', trigger: 'blur' },
          { max: 50, message: '描述不大于50个字', trigger: 'blur' },
        ],
      },
      router: useRouter(),
    };
  },
  created() {
    this.itemId = this.$route.params.id;
  },
  methods: {
    onFinish(values) {
      // 更新信息
      candidateService.updatedetail({ id: this.itemId, detail: values.detail })
        .then(() => {
          message.success('添加成功');
          window.location.reload(); // 刷新页面
          window.location.href = '/manage/candidatelist'; // 设置目标页面的 URL
        }).catch((err) => {
          message.error(err.response.data.msg);
        });
    },
  },
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
