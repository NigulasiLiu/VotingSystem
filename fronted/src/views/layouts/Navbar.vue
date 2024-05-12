<template>
  <a-layout-header>
    <a-button style="link" class="logo" @click="$router.replace({name:'home'})">电子投票系统</a-button>
    <a-menu v-model:selectedKeys="current" theme="dark" mode="horizontal" :style="{ lineHeight: '64px', float:'right'}">
      <a-menu-item key='home' @click="$router.replace({name:'home'})">
        <span>首页</span>
      </a-menu-item>
      <a-menu-item key='about' @click="$router.replace({name:'about'})">
        <span>关于</span>
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>
<script>
import { mapState } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
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
}
</style>
