import voteService from '@/service/voteService';
import { message } from 'ant-design-vue';
import { reactive } from 'vue';

const voteData = reactive({
  columns: [
    {
      title: 'ID',
      dataIndex: 'ID',
      key: 'ID',
    },
    {
      title: '投票名',
      dataIndex: 'Name',
      key: 'Name',
    },
    {
      title: '截止时间',
      dataIndex: 'Deadline',
      key: 'Deadline',
    },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
    },
  ],
  list: reactive([]),
  selectedRowkeys: [],
});

function GetVoteData() {
  return voteService.showvote().then((res) => {
    return res.data.data;
  }).catch((err) => {
    message.error(err.response.data.msg);
  });
}

GetVoteData().then((data) => {
  voteData.list = data;
});

function onSelectChange(selectedkeys) {
  voteData.selectedRowkeys = selectedkeys;
}

export { voteData, GetVoteData, onSelectChange };
