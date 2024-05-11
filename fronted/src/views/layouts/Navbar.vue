<template>
  <a-layout-header>
    <div class="logo">电子投票系统</div>
    <a-menu v-model:selectedKeys="current" theme="dark" mode="horizontal" :style="{ lineHeight: '64px', float:'right'}">
      <a-menu-item key='home' @click="$router.replace({name:'home'})">
        <HomeOutlined />
        <span>首页</span>
      </a-menu-item>
      <a-menu-item key='about' @click="$router.replace({name:'about'})">
        <InfoCircleOutlined />
        <span>关于</span>
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>
<script setup>
import { InfoCircleOutlined, HomeOutlined } from '@ant-design/icons-vue';

</script>
<style scoped>
.logo {
  float: left;
  font-size: 18px;
  font-weight: bold;
  text-align: center;
  color: white;
}
</style>
