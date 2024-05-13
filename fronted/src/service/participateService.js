import request from '@/utils/request';

// 新增投票
const addparticipation = ({ voteid, candidateid }) => {
  return request.post(`auth/participate/${voteid}/${candidateid}`);
};

export default {
  addparticipation,
};
