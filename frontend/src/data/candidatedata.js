import candidateService from '@/service/candidateService';
import { message } from 'ant-design-vue';
import { reactive } from 'vue';

// 创建一个响应式对象来存储候选人数据
const candidateData = reactive({
  list: reactive([]),
});

// 定义一个函数，用于从候选人服务中获取数据
function GetCandidateData() {
  return candidateService.showcandidate().then((res) => {
    // 返回候选人数据
    return res.data.data;
  }).catch((err) => {
    // 如果请求失败，显示错误消息
    message.error(err.response.data.msg);
  });
}

// 调用 GetCandidateData 函数，并将获取到的数据赋值给 candidateData.list
GetCandidateData().then((data) => {
  candidateData.list = data;
});

// 导出 candidateData 和 GetCandidateData
export { candidateData, GetCandidateData };
