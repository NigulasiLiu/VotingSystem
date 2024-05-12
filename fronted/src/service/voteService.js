import request from '@/utils/request';

// 新增投票
const addvote = ({ name, deadline, num }) => {
  return request.post('auth/addvote', { name, deadline, num });
};

export default {
  addvote,
};
