import request from '@/utils/request';

const addvoted = ({ voteid, voteKey }) => {
  return request.post(`auth/vote/${voteid}/${voteKey}`); // 使用 voteKey 作为验证
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
