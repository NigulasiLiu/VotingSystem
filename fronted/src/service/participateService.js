import request from '@/utils/request';

// 新增投票
const addparticipation = ({ voteid, candidateid }) => {
  return request.post(`auth/participate/${voteid}/${candidateid}`);
};

// 新增投票
const addparticipation0 = ({ voteid, candidateid }) => {
  return request.post(`vote/participate/${voteid}/${candidateid}`);
};

// 查看某场选举的候选人
const showparticipate = ({ voteid }) => {
  return request.get(`auth/showparticipate/${voteid}`);
};

export default {
  addparticipation, addparticipation0, showparticipate,
};
