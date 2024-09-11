import request from '@/utils/request';

// 用户注册
const register = ({ telephone, password }) => {
  return request.post('auth/register', { telephone, password });
};

// 用户登录
const login = ({ telephone, password }) => {
  return request.post('auth/login', { telephone, password });
};

// 获取用户信息
const info = () => {
  return request.get('auth/info');
};

// 请求密钥验证
const ballot = ({ voteKey }) => {
  return request.post('auth/ballot', { key: voteKey });
};

// 获取投票详细信息
const getVoteInfo = (voteKey) => {
  return request.get(`vote/${voteKey}`);
};

export default {
  register,
  login,
  info,
  ballot,
  getVoteInfo,
};
