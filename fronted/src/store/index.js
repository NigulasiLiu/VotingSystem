import { createStore } from 'vuex';
import userModule from './module/user';
import voteModule from './module/vote';

export default createStore({
  strict: process.env.NODE_ENV !== 'production',
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    userModule,
    voteModule,
  },
});
