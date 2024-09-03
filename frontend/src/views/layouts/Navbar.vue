<template>
  <a-layout-header class="header">
    <a-button class="logo" @click="$router.replace({ name: 'home' })">
      <img src="@/assets/election.png" alt="Logo" class="logo-image" />
      在线投票
    </a-button>
    <a-menu class="menu" theme="light" mode="horizontal">
      <a-menu-item key='login' @click="$router.replace({ name: 'login' })" v-if="!userInfo">
        <user-outlined />
        <span class="login-text">登录</span>
      </a-menu-item>
      <a-menu-item key='logout' @click="handleLogout" v-if="userInfo">
        <user-outlined />
        <span>退出</span>
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import { UserOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';

export default {
  components: {
    UserOutlined,
  },
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: {
    ...mapActions('userModule', ['logout']),
    async handleLogout() {
      try {
        await this.logout(); // 确保登出操作完成
        message.success('已退出');
        this.$router.replace({ name: 'home' }); // 使用路由导航进行页面跳转
      } catch (error) {
        message.error('退出失败，请重试');
      }
    },
  },
};
</script>

<style scoped>
.header {
  background-color: #ffffff; /* 设置白色背景 */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* 设置下边缘阴影 */
  display: flex;
  align-items: center;
  padding: 0 16px; /* 添加内边距 */
  position: fixed; /* 固定位置 */
  top: 0; /* 定位到顶部 */
  width: 100%; /* 宽度 100% */
  z-index: 1000; /* 设置较高的 z-index 以确保在其他内容之上 */
}

.logo {
  font-size: 18px;
  font-weight: bold;
  color: #000000; /* 黑色字体 */
  border: none;
  background: none;
  display: flex; /* 使用 flexbox 布局 */
  align-items: center; /* 垂直居中对齐 */
  margin-right: auto; /* 将按钮推到左边 */
  box-shadow: none; /* 去除阴影 */
  transform: translateX(60px);
}

.logo-image {
  width: 40px; /* 设置图片的宽度 */
  height: auto; /* 高度自动调整 */
  margin-right: 8px; /* 图片与文字之间的间距 */
}

.logo:hover {
  color: rgba(24, 144, 255, 0.5); /* 悬停时颜色为半透明蓝色 */
}

.menu {
  margin-right: 60px;
  line-height: 64px;
  float: right;
}

.login-text {
  color: #ffffff; /* 白色字体 */
  background-color: #1890ff; /* 蓝色背景 */
  padding: 4px 8px;
  border-radius: 4px; /* 添加圆角 */
}
</style>
