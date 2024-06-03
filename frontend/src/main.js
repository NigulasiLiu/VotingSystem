import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import axios from 'axios';
import VueAxios from 'vue-axios';
import App from './App.vue';
import router from './router';
import store from './store';

createApp(App).use(store).use(router).use(Antd)
  .use(VueAxios, axios)
  .mount('#app');
