<template>
  <div class="container">
    <a-card title="添加新的候选人" style="width: 1000px; margin: 0 auto;">
      <a-form :model="candidates" name="basic" @finish="onFinish" @finishFailed="onFinishFailed">
        <div class="table-container">
          <a-table :dataSource="candidates" :pagination="false" bordered>
            <a-table-column title="姓名" key="name">
              <template #default="{ record, index }">
                <!-- 验证姓名字段不能为空，设置trigger为change和blur -->
                <a-form-item :name="'name' + index">
                  <a-input v-model:value="record.name" placeholder="请输入姓名" />
                </a-form-item>
              </template>
            </a-table-column>
            <a-table-column title="描述" key="detail">
              <template #default="{ record, index }">
                <a-form-item :name="'detail' + index">
                  <a-textarea v-model:value="record.detail" placeholder="请输入描述" class="custom-textarea" />
                </a-form-item>
              </template>
            </a-table-column>
            <a-table-column title="图片" key="image">
              <template #default="{ record, index }">
                <a-form-item :name="'image' + index">
                  <a-upload
                    list-type="picture-card"
                    :show-upload-list="false"
                    :before-upload="(file) => beforeUpload(file, index)"
                    @change="(info) => handleUpload(info, index)"
                    class="custom-upload"
                  >
                    <div v-if="!record.imageUrl">
                      <PlusOutlined />
                      <div style="margin-top: 8px">上传图片</div>
                    </div>
                    <img v-else :src="record.imageUrl" alt="avatar" style="width: 100%" />
                  </a-upload>
                </a-form-item>
              </template>
            </a-table-column>
            <a-table-column title="操作" key="action">
              <template #default="{ index }">
                <a-button type="link" @click="removeCandidate(index)" class="delete-button">
                  <span class="delete-icon">×</span>
                </a-button>
              </template>
            </a-table-column>
          </a-table>
        </div>

        <!-- 添加新行按钮 -->
        <a-button type="dashed" @click="addCandidate" style="margin-top: 20px; width: 100%;">
          <PlusOutlined /> 添加新候选人
        </a-button>

        <!-- 提交按钮 -->
        <a-form-item :wrapper-col="{ span: 24 }" style="text-align: center; margin-top: 20px;">
          <a-button type="primary" html-type="submit" style="width: 150px;">提交</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import candidateService from '@/service/candidateService';

const candidates = reactive([
  { name: '', detail: '', imageUrl: '' }, // 初始状态包含一行
]);

// 添加新行的方法
const addCandidate = () => {
  candidates.push({ name: '', detail: '', imageUrl: '' });
};

// 删除候选人行的方法
const removeCandidate = (index) => {
  candidates.splice(index, 1); // 删除指定行
};

// 图片上传前的验证
const beforeUpload = (file) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
  if (!isJpgOrPng) {
    message.error('只能上传 JPG/PNG 文件!');
    return false;
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    message.error('图片必须小于 2MB!');
    return false;
  }
  return isJpgOrPng && isLt2M;
};

// 处理图片上传
const handleUpload = (info, index) => {
  const reader = new FileReader();
  reader.onload = (e) => {
    candidates[index].imageUrl = e.target.result;
  };
  reader.readAsDataURL(info.file.originFileObj);
};
// 表单提交处理
const onFinish = async () => {
  const validCandidates = candidates.filter((candidate) => candidate.name); // 过滤掉未填写“姓名”的行
  if (validCandidates.length === 0) {
    message.error('至少需要添加一个有效的候选人');
    return;
  }

  try {
    // 1. 发送所有 addcandidate 请求
    const addRequests = validCandidates.map((candidate) => candidateService.addcandidate({
      name: candidate.name,
      detail: candidate.detail,
    }));
    await Promise.all(addRequests); // 等待所有添加候选人请求完成

    // 2. 发送所有 updatephoto 请求（只有当图片存在时）
    const photoRequests = validCandidates
      .filter((candidate) => candidate.imageUrl) // 仅包含有图片的候选人
      .map((candidate) => candidateService.updatephoto({
        name: candidate.name,
        photo: candidate.imageUrl,
      }));

    if (photoRequests.length > 0) {
      await Promise.all(photoRequests); // 并行处理所有图片上传
    }

    message.success('所有候选人添加成功，3s内跳转到候选人列表');

    setTimeout(() => {
      window.location.href = '/manage/candidatelist';
      window.location.reload();
    }, 3000); // 3000 毫秒等于 3 秒
  } catch (err) {
    message.error(err.response.data.msg || '添加候选人或上传图片失败');
  }
};

const onFinishFailed = (errors) => {
  console.log(errors);
};

</script>

<style scoped>
.container {
  margin-top: 60px;
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 80vh;
}

.table-container {
  height: 320px; /* 固定表格的高度 */
  overflow-y: auto;  /* 表格内部滚动 */
}

/* 自定义“描述”列的高度 */
.custom-textarea {
  height: 50px !important; /* 设置自定义的高度 */
}

/* 自定义上传图片列的高度和宽度 */
.custom-upload {
  width: 80px !important; /* 设置较小的宽度 */
  height: 80px !important; /* 设置较小的高度 */
}

.ant-upload.ant-upload-select-picture-card {
  width: 80px;
  height: 80px;
}

.ant-upload-list-item {
  display: none;
}

/* 调大删除按钮样式 */
.delete-icon {
  font-size: 24px; /* 设置更大的字体 */
  color: red; /* 保持红色 */
  cursor: pointer;
}
</style>
