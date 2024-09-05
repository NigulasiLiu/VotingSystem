<template>
  <div class="container">
    <a-card title="发起投票" style="text-align:center; width:1300px">
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
                  <a-radio value="pause">暂不开放投票</a-radio>
                  <a-radio value="immediate">立刻开放投票</a-radio>
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
<!--          <div class="candidate-list-container">-->
<!--            <a-space>-->
<!--              <a-button size="large" @click="$router.replace({name:'addcandidate'})">-->
<!--                <PlusOutlined /> 新增候选人-->
<!--              </a-button>-->
<!--              &lt;!&ndash; 添加候选人 &ndash;&gt;-->
<!--              <a-button type="primary" size="large" @click="visible = true">指定候选人</a-button>-->
<!--              <a-modal v-model:open="visible" ok-text="添加" cancel-text="取消" @ok="onOkStep2"-->
<!--                       class="candidate-modal"-->
<!--                       :style="{ top: '50%', transform: 'translateY(-50%)' }">-->
<!--                <a-form :model="formState" name="form_in_modal" :rules="rulesId" ref="candidateFormRef">-->
<!--                  <a-form-item name="candidateid" label="候选人ID">-->
<!--                    <a-select v-model:value="formState.candidateid" placeholder="请选择候选人" class="candidate-select">-->
<!--                      <a-select-option v-for="item in candidateData.list" :key="item.ID" :value="item.ID">-->
<!--                        {{ item.ID }} - {{ item.Name }}-->
<!--                      </a-select-option>-->
<!--                    </a-select>-->
<!--                  </a-form-item>-->
<!--                </a-form>-->
<!--              </a-modal>-->
<!--            </a-space>-->
<!--            <a-divider></a-divider>-->
<!--            <a-list-->
<!--              v-if="candidateData.list"-->
<!--              :grid="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 3, xxl: 4 }"-->
<!--              item-layout="vertical"-->
<!--              size="middle"-->
<!--              :data-source="paginatedData"-->
<!--              :pagination="pagination"-->
<!--            >-->
<!--              <template #renderItem="{ item }">-->
<!--                <a-list-item :key="item.ID">-->
<!--                  <a-card :title="`ID&emsp;${item.ID}`" :class="{'highlighted-card': item.ID === vote.candidateid}">-->
<!--                    <template #actions>-->
<!--                      <router-link :to="{ name: 'editcandidate', params: { id: item.ID }}">-->
<!--                        <EditOutlined />&emsp;<span class="edit-text">修改信息</span>-->
<!--                      </router-link>-->
<!--                    </template>-->
<!--                    <a-list-item-meta>-->
<!--                      <template #title>-->
<!--                        <a>{{ item.Name }}</a>-->
<!--                      </template>-->
<!--                      <template #avatar>-->
<!--                        <a-avatar v-if="candidateData" :src="`data:image/png;base64,${item.Photo}`" />-->
<!--                      </template>-->
<!--                    </a-list-item-meta>-->
<!--                    {{ item.Detail }}-->
<!--                  </a-card>-->
<!--                </a-list-item>-->
<!--              </template>-->
<!--            </a-list>-->
<!--          </div>-->
          <div class="candidate-list-container">
            <a-space>
              <a-button size="large" @click="$router.replace({name:'addcandidate'})">
                <PlusOutlined /> 新增候选人
              </a-button>
            </a-space>
            <a-divider></a-divider>
            <a-list
              v-if="candidateData.list"
              :grid="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 3, xxl: 4 }"
              item-layout="vertical"
              size="middle"
              :data-source="paginatedData"
              :pagination="pagination"
            >
              <template #renderItem="{ item }">
                <a-list-item :key="item.ID">
                  <a-card
                    :title="`ID&emsp;${item.ID}`"
                    :class="{'highlighted-card': item.ID === vote.candidateid}"
                    @click="toggleCandidateSelection(item.ID)"
                    style="cursor: pointer;"
                  >
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

        <!-- 第3步内容：完成 -->
        <template v-if="currentStep === 2">
          <a-table :dataSource="stepData" :columns="columns" bordered :pagination="false">
            <template #action="{ record }">
              <div v-if="record.key === '1'">
                <a-button
                  type="primary"
                  :disabled="!isStepOneValid || voteCreated"
                  @click="onFinish(vote)"
                >
                  创建投票
                </a-button>
              </div>
              <div v-else-if="record.key === '2'">
                <a-button
                  type="primary"
                  :disabled="!isStepTwoValid || !voteCreated || candidateConfirmed"
                  @click="onOk"
                >
                  确认候选人
                </a-button>
              </div>
            </template>
          </a-table>
        </template>

        <!-- 第4步内容：分享 -->
        <template v-if="currentStep === 3">
          <div class="share-content-container">
            <!-- 投票链接 -->
            <a-card title="投票链接 (分享给投票人)" style="margin-bottom: 20px;">
              <div class="voting-link-section">
                <!-- 左侧虚拟URL链接 -->
                <div class="voting-link">
                  <div class="link-item">
                    <p style="font-weight: bold;">参与投票：“{{ vote.name }}”</p>
                    <div class="link-wrapper">
                      <a href="https://vote.example.com/m/12345/abcde" target="_blank" class="styled-link">
                        https://vote.example.com/m/12345/abcde
                      </a>
                      <a-button type="link" @click="copyToClipboard('https://vote.example.com/m/12345/abcde')">复制</a-button>
                    </div>
                  </div>
                  <div class="link-item">
                    <p style="font-weight: bold;">参与投票：“{{ vote.name }}”</p>
                    <div class="link-wrapper">
                      <a href="https://vote.cn/m/12345/abcde" target="_blank" class="styled-link">
                        https://vote.cn/m/12345/abcde
                      </a>
                      <a-button type="link" @click="copyToClipboard('https://vote.cn/m/12345/abcde')">复制</a-button>
                    </div>
                  </div>
                </div>
                <!-- 右侧微信二维码 -->
                <div class="qrcode-section">
                  <img src="@/assets/qrcode.svg" alt="微信二维码" class="qrcode-img"/>
                  <p>扫码微信投票</p>
                </div>
              </div>
            </a-card>
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
        <a-button v-if="currentStep === steps.length - 1" type="primary" @click="finishAndRedirect" style="margin-left: 8px">
          完成
        </a-button>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import {
  computed, reactive, ref, watch,
} from 'vue';
import voteService from '@/service/voteService';
import dayjs from 'dayjs';
import timezone from 'dayjs/plugin/timezone';
import { message } from 'ant-design-vue';
import { candidateData } from '@/data/candidatedata';
import { EditOutlined, PlusOutlined } from '@ant-design/icons-vue';
import participateService from '@/service/participateService';

dayjs.extend(timezone);
// 转换为本地时区
const vote = reactive({
  name: '',
  num: 0,
  deadline: '',
  openTimeOption: 'pause', // 默认选择立刻打开投票
  scheduledOpenTime: '', // 定时开放时间
  primaryContact: '',
  voteType: 'test',
  resultsViewing: 'adminAnytime',
  voteid: 0,
  candidateid: 0,
});
// 转换为本地时区
const localTime = dayjs.utc(vote.deadline).tz('Asia/Shanghai').format('YYYY-MM-DD HH:mm:ss');
console.log(localTime);
const steps = [
  { title: '基本信息' },
  { title: '候选人' },
  { title: '完成' },
  { title: '分享' },
];

const currentStep = ref(0);
const voteCreated = ref(false);
const candidateConfirmed = ref(false);
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
  voteService.addvote({ name: values.name, num: values.num, deadline: values.deadline })
    .then((response) => {
      console.log('Response Data:', response.data); // 调试信息

      if (response.data && response.data.data.id) {
        const { id } = response.data.data;
        message.success(`投票 ${id} 创建成功`);
        vote.voteid = id; // 设置 voteid
        voteCreated.value = true;
      } else {
        message.error('创建投票失败，未返回ID');
      }
    }).catch((err) => {
      message.error(err.response.data.msg || '创建投票失败');
    });
};
const onFinishFailed = (errors) => {
  console.log(errors);
};
const disabledDate = (current) => {
  return current < dayjs().endOf('day');
};
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    message.success('链接已复制到剪贴板');
  }).catch(() => {
    message.error('复制失败，请手动复制');
  });
};

// const visible = ref(false);
// const formState = reactive({ candidateid: null });
// const rulesId = {
//   candidateid: [{ required: true, message: '请选择候选人', trigger: 'change' }],
// };
const toggleCandidateSelection = (id) => {
  if (vote.candidateid === id) {
    // 如果当前候选人已经被选择，则取消选择
    vote.candidateid = null;
  } else {
    // 选择候选人
    vote.candidateid = id;
  }
};
const pagination = ref({
  current: 1,
  pageSize: 6,
  total: candidateData.list.length,
});

const paginatedData = ref(
  candidateData.list.slice(
    (pagination.value.current - 1) * pagination.value.pageSize,
    pagination.value.current * pagination.value.pageSize,
  ),
);

// const candidateFormRef = ref(null);
// const onOkStep2 = () => {
//   candidateFormRef.value
//     .validate()
//     .then(() => {
//       vote.candidateid = formState.candidateid; // 记录candidateid到vote对象中
//       visible.value = false;
//       message.success('候选人已指定');
//     })
//     .catch(() => {
//       message.error('请正确选择候选人');
//     });
// };

const onOk = () => {
  participateService
    .addparticipation({ voteid: vote.voteid, candidateid: vote.candidateid })
    .then(() => {
      message.success(`候选人${vote.candidateid}已确认`);
      candidateConfirmed.value = true;
    })
    .catch((err) => {
      message.error(err.response.data.msg);
    });
};
// const onOk = (values) => {
//   formRef.value
//     .validateFields()
//     .then(() => {
//       participateService
//         .addparticipation({ voteid: vote.voteid, candidateid: values.candidateid })
//         .then(() => {
//           visible.value = false;
//           message.success('候选人已确认');
//           // 如果需要，可以在这里添加页面刷新或者跳转逻辑
//         })
//         .catch((err) => {
//           message.error(err.response.data.msg);
//         });
//     })
//     .catch(() => {
//       message.error('请正确选择候选人');
//     });
// };
// 确保分页数据正确更新
watch(
  () => pagination.value.current,
  (newVal) => {
    paginatedData.value = candidateData.list.slice(
      (newVal - 1) * pagination.value.pageSize,
      newVal * pagination.value.pageSize,
    );
  },
);

const isStepOneValid = computed(() => {
  const isNameValid = vote.name.trim() !== '';
  const isDeadlineValid = dayjs().isBefore(dayjs(vote.deadline));
  const isNumValid = vote.num >= 1 && vote.num <= 10;

  return isNameValid && isDeadlineValid && isNumValid;
});

const isStepTwoValid = computed(() => {
  return vote.candidateid > 0; // 假设至少需要选择一个候选人
});

const finishAndRedirect = () => {
  if (voteCreated.value && candidateConfirmed.value) {
    window.location.href = '/votinglist';
  }
};
const stepData = [
  {
    key: '1',
    step: '基本信息',
    action: isStepOneValid.value ? '创建投票' : '',
  },
  {
    key: '2',
    step: '候选人',
    action: isStepTwoValid.value && voteCreated.value ? '确认候选人' : '',
  },
];

const columns = [
  {
    title: '步骤',
    dataIndex: 'step',
    key: 'step',
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action',
    slots: { customRender: 'action' }, // 使用Vue3的插槽定义
  },
];

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
.candidate-modal {
  display: flex;
  align-items: center;
  justify-content: center;
}

.equal-card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 250px; /* 根据需要调整最小高度 */
  max-height: 400px; /* 根据需要调整最大高度 */
  padding-left: 16px; /* 增加左内边距 */
}

.equal-card .ant-card-body {
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 使得内容靠左对齐 */
}

.candidate-select {
  width: 90%; /* 根据需要调整宽度 */
}

.highlighted-card {
  background: linear-gradient(to right, rgba(144, 238, 144, 0.5), rgba(250, 250, 250, 0.5)); /* 从浅绿到半透明的渐变 */
  border: 2px solid lightgreen; /* 绿色边框 */
}

.ant-table {
  background-color: #f6f9fc;
  margin: auto;
  justify-content: center;
  align-content: center;
  width: 75%;
}

.ant-table-thead > tr > th {
  background-color: #1f2937;
  color: white;
  text-align: center;
}

.ant-table-tbody > tr > td {
  text-align: center;
  padding: 16px;
}

.creation-section {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  margin-top: 20px;
}

.share-content-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.voting-link-section {
  display: flex;
  justify-content: space-around; /* 调整为居中对齐 */
  align-items: center;
}

.voting-link {
  text-align: left;
  flex: 1; /* 链接部分占用一定的宽度 */
}

.qrcode-section {
  text-align: center;
  flex: 1; /* 二维码部分占用一定的宽度 */
}

.link-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 16px;
}

.qrcode-img {
  width: 120px; /* 根据实际二维码大小调整 */
  height: 120px;
}
.styled-link {
  font-size: 16px;
  color: #1890ff;
  text-decoration: underline;
}

.styled-link:hover {
  color: #40a9ff;
}

.link-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

</style>
