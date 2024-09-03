<template>
  <div class="container">
    <a-card title="添加投票" style="text-align:center; width:1300px">
      <!-- Stepper -->
      <a-steps :current="currentStep" style="margin-bottom: 30px;">
        <a-step v-for="item in steps" :key="item.title" :title="item.title" />
      </a-steps>

      <!-- Form -->
      <a-form
        :model="vote"
        name="basic"
        :rules="rules"
        :label-col="{ span: 24 }"
        :wrapper-col="{ span: 24 }"
        layout="vertical"
        autocomplete="off"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
      >
        <!-- 第1步内容 -->
        <template v-if="currentStep === 0">
          <div class="form-left-right">
            <!-- 左侧 -->
            <div class="form-left">
              <a-form-item label="投票名" name="name">
                <a-input v-model:value="vote.name" />
              </a-form-item>

              <a-form-item label="投票发起人联系方式" name="primaryContact">
                <a-input v-model:value="vote.primaryContact" placeholder="请输入投票发起人联系方式" />
              </a-form-item>

              <a-form-item name="deadline" label="截止时间">
                <a-date-picker
                  v-model:value="vote.deadline"
                  show-time
                  format="YYYY-MM-DD HH:mm:ss"
                  value-format="YYYY-MM-DD HH:mm:ss"
                  :disabled-date="disabledDate"
                  style="width: 100%"
                />
              </a-form-item>

              <a-form-item label="可投票数" name="num">
                <a-input-number v-model:value="vote.num" :min="1" :max="10" style="width: 100%" />
              </a-form-item>
            </div>

            <!-- 右侧 -->
            <div class="form-right">

              <a-form-item label="开放时间" name="openTime">
                <a-radio-group v-model:value="vote.openTimeOption">
                  <a-radio value="immediate">立刻打开投票</a-radio>
                  <a-radio value="scheduled">
                    定时释放：
                    <a-date-picker
                      v-model:value="vote.scheduledOpenTime"
                      show-time
                      format="YYYY-MM-DD HH:mm:ss"
                      value-format="YYYY-MM-DD HH:mm:ss"
                      :disabled-date="disabledDate"
                      style="width: 200px; margin-left: 10px;"
                      :disabled="vote.openTimeOption !== 'scheduled'"
                    />
                  </a-radio>
                </a-radio-group>
              </a-form-item>
              <a-form-item label="投票类型" name="voteOptions">
                <a-radio-group v-model:value="vote.voteType">
                  <a-radio value="test">测试：最多 5 名投票人</a-radio>
                  <a-radio value="live">正式：不限投票人人数，提供所有功能。</a-radio>
                </a-radio-group>
              </a-form-item>

              <!-- 结果查看 -->
              <a-form-item label="结果查看" name="resultsViewing">
                <a-radio-group v-model:value="vote.resultsViewing">
                  <a-radio value="adminAnytime" class="custom-radio">管理员：投票开始后随时查看</a-radio>
                  <a-radio value="adminAfterEnd" class="custom-radio">管理员：仅在投票结束后查看</a-radio>
                  <a-radio value="votersAfterEnd" class="custom-radio">投票人：仅在投票结束后查看</a-radio>
                  <a-radio value="votersAnytime" class="custom-radio">投票人：投票开始后随时查看</a-radio>
                  <a-radio value="adminOnly" class="custom-radio">投票人无法查看 - 仅管理员可查看</a-radio>
                </a-radio-group>
                <a-text style="margin-top: 8px; display: block;">
                  结果不会自动发送给投票人。<a href="#" class="more-info-link">了解更多</a>
                </a-text>
              </a-form-item>
            </div>
          </div>
        </template>

        <!-- 第2步内容：候选人 -->
        <template v-if="currentStep === 1">
          <div class="candidate-list-container">
            <a-space>
              <a-button size="large" @click="$router.replace({name:'addcandidate'})">
                <PlusOutlined /> 新增候选人
              </a-button>

              <!-- 添加候选人 -->
              <a-button type="primary" size="large" @click="visible = true">指定候选人</a-button>
              <a-modal v-model:open="visible" ok-text="添加" cancel-text="取消" @ok="onOk(formState)">
                <a-form :model="formState" name="form_in_modal" :rules="rulesId" ref="formRef">
                  <a-form-item name="candidateid" label="候选人ID">
                    <a-input-number v-model:value="formState.candidateid" />
                  </a-form-item>
                </a-form>
              </a-modal>
            </a-space>
            <a-divider></a-divider>
            <a-list
              v-if="candidateData.list"
              :grid="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 3, xxl: 4 }"
              item-layout="vertical"
              size="large"
              :data-source="paginatedData"
              :pagination="pagination"
            >
              <template #renderItem="{ item }">
                <a-list-item :key="item.ID">
                  <a-card :title="`ID&emsp;${item.ID}`">
                    <template #actions>
                      <router-link :to="{ name: 'editcandidate', params: { id: item.ID }}">
                        <EditOutlined />&emsp;<span class="edit-text">修改信息</span>
                      </router-link>
                    </template>
                    <a-list-item-meta>
                      <template #title>
                        <a>{{ item.Name }}</a>
                      </template>
                      <template #avatar>
                        <a-avatar v-if="candidateData" :src="`data:image/png;base64,${item.Photo}`" />
                      </template>
                    </a-list-item-meta>
                    {{ item.Detail }}
                  </a-card>
                </a-list-item>
              </template>
            </a-list>
          </div>
        </template>
      </a-form>

      <!-- Navigation Buttons -->
      <div :style="{ marginTop: '24px' }">
        <a-button v-if="currentStep > 0" @click="prev">
          上一步
        </a-button>
        <a-button v-if="currentStep < steps.length - 1" style="margin-left: 8px" type="primary" @click="next">
          下一步
        </a-button>
        <a-button v-if="currentStep === steps.length - 1" type="primary" html-type="submit">
          提交
        </a-button>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { reactive, ref, computed } from 'vue';
import voteService from '@/service/voteService';
import participateService from '@/service/participateService';
import dayjs from 'dayjs';
import { message } from 'ant-design-vue';
import { candidateData } from '@/data/candidatedata';
import { EditOutlined, PlusOutlined } from '@ant-design/icons-vue';

const vote = reactive({
  name: '',
  num: 0,
  deadline: '',
  openTimeOption: 'immediate', // 默认选择立刻打开投票
  scheduledOpenTime: '', // 定时开放时间
  primaryContact: '',
  primaryLanguage: 'chinese',
  voteType: 'test',
  resultsViewing: 'adminAfterEnd',
});

const steps = [
  { title: '基本信息' },
  { title: '候选人' },
  { title: '链接分享' },
  { title: '完成' },
];

const currentStep = ref(0);

const visible = ref(false);
const formState = ref({ candidateid: null });
const formRef = ref(null);

const next = () => {
  if (currentStep.value < steps.length - 1) {
    currentStep.value += 1;
  }
};

const prev = () => {
  if (currentStep.value > 0) {
    currentStep.value -= 1;
  }
};

const validateID = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入候选人ID'));
  }
  if (value <= 0) {
    return Promise.reject(new Error('ID应大于0'));
  }
  return Promise.resolve();
};

const rulesId = {
  candidateid: [{ required: true, validator: validateID, trigger: 'blur' }],
};

const onOk = (values) => {
  formRef.value
    .validateFields()
    .then(() => {
      participateService
        .addparticipation({ voteid: vote.id, candidateid: values.candidateid })
        .then(() => {
          visible.value = false;
          message.success('添加成功');
          window.location.reload(); // 刷新页面
        })
        .catch((err) => {
          message.error(err.response.data.msg);
        });
    })
    .catch(() => {
      message.error('请填写正确候选人id');
    });
};

// 分页数据
const pagination = ref({
  current: 1,
  pageSize: 8, // 每页显示的候选人数量
  total: candidateData.list.length,
  showSizeChanger: true,
});

const paginatedData = computed(() => {
  const start = (pagination.value.current - 1) * pagination.value.pageSize;
  const end = pagination.value.current * pagination.value.pageSize;
  return candidateData.list.slice(start, end);
});

pagination.value.onChange = (page, pageSize) => {
  pagination.value.current = page;
  pagination.value.pageSize = pageSize;
};

const validateName = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入投票名'));
  }
  if (value.length > 30) {
    return Promise.reject(new Error('投票名不大于 30 位'));
  }
  return Promise.resolve();
};

const validateDDL = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请选择截止时间'));
  }
  if (dayjs().isAfter(dayjs(value))) {
    return Promise.reject(new Error('投票时间不能早于当前时间'));
  }
  return Promise.resolve();
};

const validateNum = async (_rule, value) => {
  if (value === '') {
    return Promise.reject(new Error('请输入可投票数'));
  }
  if (value < 1 || value >= 10) {
    return Promise.reject(new Error('可投票数必须在 1 到 10 之间'));
  }
  return Promise.resolve();
};

const rules = {
  name: {
    required: true,
    validator: validateName,
    trigger: 'blur',
  },
  num: {
    required: true,
    validator: validateNum,
    trigger: 'blur',
  },
  deadline: {
    required: true,
    validator: validateDDL,
    trigger: 'blur',
  },
};

const onFinish = (values) => {
  voteService
    .addvote({ name: values.name, num: values.num, deadline: values.deadline })
    .then(() => {
      message.success('添加成功');
      window.location.reload(); // 刷新页面
      window.location.href = '/votinglist'; // 设置目标页面的 URL
    })
    .catch((err) => {
      message.error(err.response.data.msg);
    });
};

const onFinishFailed = (errors) => {
  console.log(errors);
};

const disabledDate = (current) => {
  return current < dayjs().endOf('day');
};
</script>

<style scoped>
.container {
  text-align: -webkit-center;
  margin: 35px auto;
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
}

.form-left-right {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.form-left,
.form-right {
  width: 48%; /* 左右两侧各占48%，中间留4%的间隔 */
}

.custom-radio {
  margin-bottom: 8px; /* 增加每个单选按钮的间距 */
}

.more-info-link {
  color: #1890ff; /* 链接颜色 */
  text-decoration: none; /* 去除下划线 */
}

.more-info-link:hover {
  text-decoration: underline; /* 鼠标悬停时添加下划线 */
}

.candidate-list-container {
  margin-top: 20px;
  width: 100%;
}
</style>
