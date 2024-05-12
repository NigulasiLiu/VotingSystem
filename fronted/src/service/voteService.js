import request from '@/utils/request';

// 新增投票
const addvote = ({ name, deadline, num }) => {
  return request.post('auth/addvote', { name, deadline, num });
};

// 获取投票列表
const showvote = () => {
  return request.get('auth/showvote');
};

// 获取投票信息
const voteinfo = ({ id }) => {
  return request.get(`auth/voteinfo/${id}`);
};

export default {
  addvote,
  showvote,
  voteinfo,
};
