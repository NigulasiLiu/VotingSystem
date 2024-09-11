import request from '@/utils/request';

// 为投票设定候选人
const addparticipation = ({ voteid, candidateid }) => {
  return request.post(`auth/participate/${voteid}/${candidateid}`);
};

// 给server participate表增加条目，用于计算
const addparticipation0 = ({ voteid, candidateid }) => {
  return request.post(`vote/participate/${voteid}/${candidateid}`);
};

// 查看某场投票的候选人
const showparticipate = ({ voteid }) => {
  return request.get(`auth/showparticipate/${voteid}`);
};

export default {
  addparticipation, addparticipation0, showparticipate,
};
