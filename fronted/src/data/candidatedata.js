import candidateService from '@/service/candidateService';
import { message } from 'ant-design-vue';
import { reactive } from 'vue';

const candidateData = reactive({
  list: reactive([]),
});

function GetCandidateData() {
  return candidateService.showcandidate().then((res) => {
    return res.data.data;
  }).catch((err) => {
    message.error(err.response.data.msg);
  });
}

GetCandidateData().then((data) => {
  candidateData.list = data;
});

export { candidateData, GetCandidateData };
