import voteService from '@/service/voteService';

const voteModule = {
  namespaced: true,
  actions: {
    addvote(context, { name, num, deadline }) {
      return new Promise((resolve, reject) => {
        voteService.addvote({ name, num, deadline }).then((res) => {
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
  },
};

export default voteModule;
