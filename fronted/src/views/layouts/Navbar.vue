<template>
  <a-layout-header>
    <a-button style="link" class="logo" @click="$router.replace({name:'home'})">电子投票系统</a-button>
    <a-menu theme="dark" mode="horizontal" :style="{ lineHeight: '64px', float:'right'}">
      <a-menu-item key='login' @click="$router.replace({name:'login'})" v-if="userInfo===null">
        <user-outlined />
        <span>登录</span>
      </a-menu-item>
      <a-menu-item key='login' @click="logout();back()" v-if="userInfo">
        <user-outlined />
        <span>退出</span>
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>
<script>
import { ref } from 'vue';
import { mapState, mapActions } from 'vuex';
import { useRouter } from 'vue-router';
import { UserOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';

export default {
  components: {
    UserOutlined,
  },
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: mapActions('userModule', ['logout']),
  setup() {
    const current = ref(['home']);
    const router = useRouter();
    const back = () => {
      message.success('已退出');
      router.push('/');
      // eslint-disable-next-line no-restricted-globals
      location.reload();
    };
    return {
      current,
      back,
    };
  },
};
</script>
<style scoped>
.logo {
  font-size: 18px;
  font-weight: bold;
  background-color: #141414;
  color: white;
  border-color: #141414;
}

.logo:hover {
  border-color: #141414;
  color: lightgray;
}
</style>
