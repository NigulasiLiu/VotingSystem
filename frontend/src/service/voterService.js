import request from '@/utils/request';

// 新增投票者
const addVoter = ({
  voteid, name, phone, email, key,
}) => {
  return request.post(`auth/addvoter/${voteid}`, {
    voteid,
    name,
    phone,
    email,
    key,
  });
};

// 查看某场投票的所有投票者
const showVoters = ({ voteid }) => {
  return request.get(`auth/showvoters/${voteid}`);
};

export default {
  addVoter,
  showVoters,
};
