import request from '@/utils/request';

// 用户注册
const register = ({ telephone, password }) => {
  return request.post('auth/register', { telephone, password });
};

// 用户注册
const login = ({ telephone, password }) => {
  return request.post('auth/login', { telephone, password });
};

// 获取用户信息
const info = () => {
  return request.get('auth/info');
};

export default {
  register,
  login,
  info,
};
