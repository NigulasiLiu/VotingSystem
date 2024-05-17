import request from '@/utils/request';

// 新增投票
const addvote = ({ name, deadline, num }) => {
  return request.post('auth/addvote', { name, deadline, num });
};

// 新增投票（服务器）
const addvote0 = ({ id }) => {
  return request.post(`vote/${id}`);
};

// 获取投票列表
const showvote = () => {
  return request.get('auth/showvote');
};

// 获取投票信息
const voteinfo = ({ id }) => {
  return request.get(`auth/voteinfo/${id}`);
};

// 开放投票
const openvote = ({ id }) => {
  return request.post(`auth/openvote/${id}`);
};

// 截止投票
const endvote = ({ id }) => {
  return request.post(`auth/endvote/${id}`);
};

// 开始计票
const computevote = ({ id }) => {
  return request.post(`auth/computevote/${id}`);
};

const computevote0 = ({ id }) => {
  return request.post(`compute/server0/${id}`);
};
const computevote1 = ({ id }) => {
  return request.post(`compute/server1/${id}`);
};

const outs = ({ id, outs0, outs1 }) => {
  return request.post(`auth/outs/${id}`, { outs0, outs1 });
};

export default {
  addvote,
  addvote0,
  showvote,
  voteinfo,
  openvote,
  endvote,
  computevote,
  computevote0,
  computevote1,
  outs,
};
