<template>
  <a-layout>
    <navbar />
    <a-layout-content :style="{padding: '50px',background: '#fff'}">
      <router-view />
    </a-layout-content>
  </a-layout>
</template>

<script>
import Navbar from './views/layouts/Navbar.vue';

export default {
  components: { Navbar },
  data() {
  },
};
</script>

<style lang="scss"></style>
