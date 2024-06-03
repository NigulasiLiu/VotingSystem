import request from '@/utils/request';

// 新增投票
const addcandidate = ({ name, detail }) => {
  return request.post('auth/addcandidate', { name, detail });
};

// 获取候选人列表
const showcandidate = () => {
  return request.get('auth/showcandidate');
};

// 上传图片
const updatephoto = ({ name, photo }) => {
  return request.post('auth/photo', { name, photo });
};

// 修改描述
const updatedetail = ({ id, detail }) => {
  return request.post(`auth/candidate/detail/${id}`, { detail });
};

export default {
  addcandidate,
  showcandidate,
  updatephoto,
  updatedetail,
};
