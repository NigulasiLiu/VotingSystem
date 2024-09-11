import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO) ? JSON.parse(storageService.get(storageService.USER_INFO)) : null,
    voteID: storageService.get(storageService.USER_ID),
    voteInfo: {},
    voteKey: storageService.get('VOTE_KEY'), // 添加 voteKey 存储
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // 更新state
      state.userInfo = userInfo;
    },
    SET_ID(state, voteID) {
      state.voteID = voteID;
    },
    SET_VOTEKEY(state, voteKey) { // 添加 SET_VOTEKEY
      storageService.set('VOTE_KEY', voteKey);
      state.voteKey = voteKey;
    },
  },
  actions: {
    register(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ telephone, password, role: 1 }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    ballot(context, { voteKey }) {
      return new Promise((resolve, reject) => {
        // 调用 ballot API 进行密钥验证
        userService.ballot({ voteKey }).then((res) => {
          // 使用对象解构提取 voteID
          const { voteID } = res.data.data;

          // 保存 voteID 或其他相关信息
          context.commit('SET_ID', voteID);
          context.commit('SET_VOTEKEY', voteKey); // 保存 voteKey
          // 验证成功后直接跳转到指定页面
          // this.$router.push({ path: `/vote/${voteID}` });

          resolve(res); // 成功的结果返回
        }).catch((err) => {
          // 处理错误
          reject(err);
        });
      });
    },

    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.login({ telephone, password }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },

    logout({ commit }) {
      // 清理token
      commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清除用户信息
      commit('SET_USERINFO', '');
      storageService.set(storageService.USER_INFO, '');
      commit('SET_VOTEKEY', ''); // 清除 voteKey
      storageService.set('VOTE_KEY', '');
    },
  },
};

export default userModule;
