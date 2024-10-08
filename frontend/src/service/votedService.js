import request from '@/utils/request';

// 新增投票
const addvoted = ({ voteid, userid }) => {
  return request.post(`auth/vote/${voteid}/${userid}`);
};

const vote0 = ({ id, k }) => {
  return request.post(`vote/server0/${id}`, k);
};

const vote1 = ({ id, k, nid }) => {
  return request.post(`vote/server1/${id}/${nid}`, k);
};

export default {
  addvoted,
  vote0,
  vote1,
};
