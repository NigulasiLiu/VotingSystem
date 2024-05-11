import request from '@/utils/request';

// 用户注册
const register = ({ telephone, password }) => {
  return request.post('auth/register', { telephone, password });
};

// 获取用户信息
const info = () => {
  return request.get('auth/info');
};

export default {
  register,
  info,
};
