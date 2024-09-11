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

              <a-form-item label="投票活动描述" name="description">
                <a-input v-model:value="vote.description" placeholder="请输入投票活动描述" />
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
                  <a-radio value="adminAfterEnd" class="custom-radio">管理员：仅在投票结束后查看</a-radio>
                  <a-radio value="adminAnytime" class="custom-radio">管理员：投票开始后随时查看</a-radio>
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
            <a-divider></a-divider>
            <a-list
              v-if="candidateData.list.length > 0"
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
                    :class="{'highlighted-card': vote.candidateid.includes(item.ID)}"
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
        <!-- 第3步：输入投票者数据 -->
        <template v-if="currentStep === 2">
          <a-form :model="step3Data" name="basic">
            <div class="table-container">
              <a-table :dataSource="step3Data" :pagination="false" bordered>
                <a-table-column title="姓名" key="name">
                  <template #default="{ record, index }">
                    <a-form-item :name="'name' + index">
                      <a-input v-model:value="record.name" placeholder="请输入姓名" />
                    </a-form-item>
                  </template>
                </a-table-column>
                <a-table-column title="电话" key="phone">
                  <template #default="{ record, index }">
                    <a-form-item :name="'phone' + index">
                      <a-input v-model:value="record.phone" placeholder="请输入电话"/>
                    </a-form-item>
                  </template>
                </a-table-column>
                <a-table-column title="邮箱" key="email">
                  <template #default="{ record, index }">
                    <a-form-item :name="'email' + index">
                      <a-input v-model:value="record.email" placeholder="请输入邮箱"/>
                    </a-form-item>
                  </template>
                </a-table-column>
                <a-table-column title="操作" key="action">
                  <template #default="{ index }">
                    <a-button type="link" @click="removeVoter(index)" class="delete-button">
                      <span class="delete-icon">×</span>
                    </a-button>
                  </template>
                </a-table-column>
              </a-table>
            </div>
            <a-button type="dashed" @click="addVoter" style="margin-top: 20px; margin-right: 20px;width: 200px;">
              <PlusOutlined /> 添加新候选人
            </a-button>
            <!-- 存储第三步的数据 -->
            <a-button type="primary" @click="storeStep3Data" style="margin-top: 20px;">
              保存
            </a-button>
          </a-form>

        </template>

        <!-- 第4步内容：完成 -->
        <template v-if="currentStep === 3">
          <a-table :dataSource="stepData" :columns="columns" bordered :pagination="false" style="width: 1000px; margin: 0 auto;"> <!-- 居中表格 -->
            <!-- 步骤列 -->
            <template #step="{ record }">
              <div style="font-weight: bold; font-size: 21px;">
                <div v-if="record.key === '1'">
                  基本信息
                  <div style="font-style: italic; color: gray; font-size: 14px;">
                    包括投票名称、投票截至时间以及每个投票者的可投票数。
                  </div>
                </div>
                <div v-else-if="record.key === '2'">
                  候选人
                  <div style="font-style: italic; color: gray; font-size: 14px;">
                    参与投票的被投票者(候选人)ID。
                  </div>
                </div>
                <div v-else-if="record.key === '3'">
                  投票者
                  <div style="font-style: italic; color: gray; font-size: 14px;">
                    参与投票的投票者联系方式。
                  </div>
                </div>
              </div>
            </template>

            <!-- 状态列 -->
            <template #status="{ record }">
              <div style="font-size: 18px; display: flex; justify-content: space-between;">
                <div v-if="record.key === '1'" style="flex-grow: 1;">
                  <a-list bordered>
                    <!-- 投票名 -->
                    <a-list-item>
                      <a-list-item-meta
                        title="投票名"
                        :description="record.name ? '名称：' + record.name : '未设置投票名'"
                      />
                      <CheckCircleFilled v-if="record.name" style="color: green; font-size: 24px; margin-right: 30px;" />
                      <CloseCircleFilled v-else style="color: red; font-size: 24px; margin-right: 30px;" />
                    </a-list-item>

                    <!-- 截至时间 -->
                    <a-list-item>
                      <a-list-item-meta
                        title="截至时间"
                        :description="record.deadline ? '时间：' + record.deadline : '未设置截止时间'"
                      />
                      <CheckCircleFilled v-if="record.deadline" style="color: green; font-size: 24px; margin-right: 30px;" />
                      <CloseCircleFilled v-else style="color: red; font-size: 24px; margin-right: 30px;" />
                    </a-list-item>

                    <!-- 可投票数 -->
                    <a-list-item>
                      <a-list-item-meta
                        title="可投票数"
                        :description="record.num ? '数量：' + record.num : '未设置可投票数'"
                      />
                      <CheckCircleFilled v-if="record.num >= 1 && record.num <= 10" style="color: green; font-size: 24px; margin-right: 30px;" />
                      <CloseCircleFilled v-else style="color: red; font-size: 24px; margin-right: 30px;" />
                    </a-list-item>
                  </a-list>
                </div>

                <!-- 第二步：显示候选人信息 -->
                <div v-else-if="record.key === '2'" style="flex-grow: 1;">
                  <a-list bordered>
                    <a-list-item>
                      <a-list-item-meta
                        title="候选人"
                        description="参与投票的候选人"
                      />
                      <div v-if="record.candidateid && record.candidateid.length > 0">
                        <div v-for="id in record.candidateid" :key="id" style="display: flex; justify-content: space-between;">
                          <a-tag color="blue">候选人ID：{{ id }}</a-tag>
                          <CheckCircleFilled style="color: green; font-size: 24px; margin-right: 30px;" />
                        </div>
                      </div>
                      <div v-else>
                        <a-tag color="red">没有候选人</a-tag>
                        <CloseCircleFilled style="color: red; font-size: 24px; margin-right: 30px;" />
                      </div>
                    </a-list-item>
                  </a-list>
                </div>

                <!-- 显示投票者信息 -->
                <div v-else-if="record.key === '3'" style="flex-grow: 1;">
                  <div v-if="record.voters && record.voters.length > 0">
                    <a-list
                      :dataSource="record.voters"
                      bordered
                      style="width: 100%;"
                    >
                      <template #renderItem="{ item }">
                        <a-list-item>
                          <a-list-item-meta
                            :title="item.name"
                          />
                          <div>
                            <a-tag
                              :color="isPhoneValid(item.phone) ? 'green' : 'red'"
                              style="margin-right: 8px;"
                            >
                              电话：{{ item.phone }}
                            </a-tag>
                            <a-tag
                              :color="isEmailValid(item.email) ? 'green' : 'red'"
                            >
                              邮箱：{{ item.email }}
                            </a-tag>
                          </div>
                        </a-list-item>
                      </template>
                    </a-list>

                  </div>
                  <div v-else style="margin-top: 8px; display: flex; justify-content: space-between;">
                    <span>未添加投票者或未保存</span>
                    <CloseCircleFilled style="color: red; font-size: 24px; margin-right: 30px;" />
                  </div>
                </div>

              </div>
            </template>

            <!-- 操作列 -->
            <template #action="{ record }">
              <div style="text-align: center;">
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
                    @click="onCandidateOk"
                  >
                    确认候选人
                  </a-button>
                </div>
                <div v-else-if="record.key === '3'">
                  <a-button
                    type="primary"
                    :disabled="!isStepThreeValid || !candidateConfirmed || voterConfirmed"
                    @click="onVoterOk"
                  >
                    确认投票者
                  </a-button>
                </div>
              </div>
            </template>

          </a-table>
        </template>
        <!-- 第5步内容：分享 -->
        <template v-if="currentStep === 4">
          <div class="share-content-container">
            <!-- 投票链接 -->
            <a-card title="投票链接 (分享给投票人)" style="margin-bottom: 20px;">
              <div class="voting-link-section">

                <!-- 左侧虚拟URL链接行 -->
                <div class="voting-link" style="display: block;"> <!-- 确保独占一行 -->
                  <div class="link-item">
                    <p style="font-weight: bold;">参与投票：“{{ vote.name }}”</p>
                    <div class="link-wrapper">
                      <a href="https://vote.example.com/m/12345/abcde" target="_blank" class="styled-link">
                        https://vote.example.com/m/12345/abcde
                      </a>
                      <a-button type="link" @click="copyToClipboard('https://vote.example.com/m/12345/abcde')">复制</a-button>
                    </div>
                  </div>
                </div>

                <!-- 右侧微信二维码 -->
                <div class="qrcode-section">
                  <p style="font-weight: bold;">扫码微信投票</p>
                  <img src="@/assets/qrcode.svg" alt="微信二维码" class="qrcode-img"/>
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
        <!-- 下一步按钮 -->
        <a-button
          v-if="currentStep < steps.length - 1"
          style="margin-left: 8px"
          type="primary"
          :disabled="
      (currentStep === 0 && !isStepOneValid) ||
      (currentStep === 1 && !isStepTwoValid) ||
      (currentStep === 2 && !isStepThreeValid)
    "
          @click="next"
        >
          下一步
        </a-button>

        <!-- 完成按钮 -->
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
import { message } from 'ant-design-vue';
import { candidateData } from '@/data/candidatedata';
import {
  EditOutlined,
  CheckCircleFilled,
  CloseCircleFilled, PlusOutlined,
} from '@ant-design/icons-vue';
import participateService from '@/service/participateService';
import voterService from '@/service/voterService';

const vote = reactive({
  name: '',
  num: 0,
  deadline: '',
  openTimeOption: 'pause', // 默认选择立刻打开投票
  scheduledOpenTime: '', // 定时开放时间
  description: '',
  voteType: 'test',
  resultsViewing: 'adminAfterEnd',
  voteid: 0,
  candidateid: [], // 用数组存储候选人ID

  voters: [], // 新增投票者字段
});

const steps = [
  { title: '基本信息' },
  { title: '候选人' },
  { title: '投票者' },
  { title: '完成' },
  { title: '分享' },
];

const currentStep = ref(0);
const voteCreated = ref(false);
const candidateConfirmed = ref(false);
const voterConfirmed = ref(false);
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
  const currentTime = dayjs();
  const selectedTime = dayjs(value);

  // 判断是否早于当前时间 + 15分钟
  if (currentTime.add(15, 'minute').isAfter(selectedTime)) {
    return Promise.reject(new Error('投票时间至少要比当前时间晚15分钟'));
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
  // 获取当前时间，并加上15分钟
  const minSelectableTime = dayjs().add(1, 'minute');

  // 判断选择的时间是否在当前时间的15分钟之前
  return current && current < minSelectableTime;
};

const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    message.success('链接已复制到剪贴板');
  }).catch(() => {
    message.error('复制失败，请手动复制');
  });
};

const toggleCandidateSelection = (id) => {
  const index = vote.candidateid.indexOf(id);
  if (index > -1) {
    // 如果候选人已经被选中，则取消选择
    vote.candidateid.splice(index, 1);
  } else {
    // 如果未被选中，则添加到数组中
    vote.candidateid.push(id);
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

const onCandidateOk = () => {
  // 获取候选人 ID 数组
  const candidateIdsArray = vote.candidateid;

  // 如果候选人数组为空，给出提示
  if (candidateIdsArray.length === 0) {
    message.error('请至少选择一个候选人');
    return;
  }
  // 使用 Promise.all 发送多个请求，每个请求包含一个候选人 ID
  const requests = candidateIdsArray.map((candidateId) => {
    return Promise.all([
      participateService.addparticipation({ voteid: vote.voteid, candidateid: candidateId }),
      participateService.addparticipation0({ voteid: vote.voteid, candidateid: candidateId }),
    ]);
  });

  // 处理所有请求的结果
  Promise.all(requests)
    .then(() => {
      message.success(`候选人 ${candidateIdsArray.join(', ')} 已确认`);
      candidateConfirmed.value = true;
    })
    .catch((err) => {
      message.error(err.response.data.msg || '确认候选人时出错');
    });
};

const generateKey = () => {
  // 生成一个16位的随机字符串作为密钥
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let key = '';
  // eslint-disable-next-line no-plusplus
  for (let i = 0; i < 16; i++) {
    key += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return key;
};

const onVoterOk = () => {
  // 获取投票者数组
  const voterArray = vote.voters;
  // 如果投票者数组为空，给出提示
  if (voterArray.length === 0) {
    message.error('请至少添加一个投票者');
    return;
  }
  // 使用 Promise.all 发送多个请求，每个请求包含一个投票者的信息
  const requests = voterArray.map((voter) => {
    const secretKey = generateKey();
    return voterService.addVoter({
      voteid: vote.voteid,
      name: voter.name,
      phone: voter.phone,
      email: voter.email,
      key: secretKey,
    });
  });
  // 处理所有请求的结果
  Promise.all(requests)
    .then(() => {
      message.success('所有投票者信息已确认并保存');
      voterConfirmed.value = true;
    })
    .catch((err) => {
      console.error(err.response.data);
      message.error(err.response.data.msg || '确认投票者时出错');
    });
};

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
  return vote.candidateid.length > 0; // 假设至少需要选择一个候选人
});
const step3Data = reactive([
  {
    name: '', phone: '', email: '', phoneValid: true, emailValid: true,
  }, // 每一行代表一个投票者
]);

// 添加投票者
const addVoter = () => {
  step3Data.push({
    phone: '', email: '', phoneValid: true, emailValid: true,
  });
};

// 删除投票者
const removeVoter = (index) => {
  step3Data.splice(index, 1);
};

// 检查所有投票者信息是否有效
const isStepThreeValid = computed(() => {
  // 如果没有任何数据，直接返回 false 禁用下一步按钮
  if (step3Data.length === 0) return false;

  // 确保所有数据行的验证都通过
  return step3Data.every((data) => data.phoneValid && data.emailValid);
});
const isPhoneValid = (phone) => {
  // 简单的电话验证，只允许10-15位数字
  const phoneRegex = /^[0-9]{10,15}$/;
  return phoneRegex.test(phone);
};

const isEmailValid = (email) => {
  // 简单的邮箱验证
  const emailRegex = /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/;
  return emailRegex.test(email);
};

// const finishAndRedirect = () => {
//   if (voteCreated.value && candidateConfirmed.value) {
//     window.location.href = '/votinglist';
//   }
// };
const finishAndRedirect = () => {
  // 检查投票是否已创建，并确认候选人
  if (voteCreated.value && candidateConfirmed.value) {
    // 调用开放投票的服务
    voteService.openvote({ id: vote.voteid })
      .then(() => {
        message.success('投票已开放');
        // 跳转到投票列表页面
        window.location.href = '/votinglist';
      })
      .catch((err) => {
        // 处理开放投票的错误
        message.error(err.response.data.msg || '开放投票时出错');
      });
  }
};

const stepData = computed(() => [
  {
    key: '1',
    step: '基本信息',
    name: vote.name, // 从 vote 动态获取投票名
    deadline: vote.deadline, // 从 vote 动态获取截止时间
    num: vote.num, // 从 vote 动态获取可投票数
    action: isStepOneValid.value ? '创建投票' : '',
  },
  {
    key: '2',
    step: '候选人',
    candidateid: vote.candidateid, // 从 vote 动态获取候选人ID
    action: isStepTwoValid.value && voteCreated.value ? '确认候选人' : '',
  },
  {
    key: '3',
    step: '投票者',
    voters: vote.voters, // 从 vote 动态获取投票者信息
    action: isStepThreeValid.value && voterConfirmed.value ? '确认投票者' : '',
  },
]);

const storeStep3Data = () => {
  // 将 step3Data 整合到 vote.voters 中
  vote.voters = [...step3Data];
  //
  // // 更新 stepData 中的投票者数据
  // const step3Index = stepData.value.findIndex((item) => item.key === '3');
  // if (step3Index !== -1) {
  //   stepData[step3Index].voters = vote.voters; // 更新 stepData 的 voters
  // }

  message.success('投票者信息已保存');
};
// 列定义
// const Votercolumns = [
//   {
//     title: '联系电话',
//     dataIndex: 'phone',
//     key: 'phone',
//     slots: { customRender: 'phone' },
//     width: 350,
//   },
//   {
//     title: '邮箱',
//     dataIndex: 'email',
//     key: 'email',
//     slots: { customRender: 'email' },
//   },
//   {
//     title: '操作',
//     dataIndex: 'action',
//     key: 'action',
//     slots: { customRender: 'action' },
//     width: 150,
//   },
// ];
const columns = [
  // {
  //   title: '步骤',
  //   dataIndex: 'step',
  //   key: 'step',
  //   slots: { customRender: 'step' },
  //   width: 290,
  // },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    slots: { customRender: 'status' },
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action',
    slots: { customRender: 'action' },
    width: 150,
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
  width: 1000px;
  margin: auto;
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

.delete-icon {
  font-size: 24px; /* 设置更大的字体 */
  color: red; /* 保持红色 */
  cursor: pointer;
}
</style>
