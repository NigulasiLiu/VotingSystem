import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/login/Login.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: Login,
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/register/Register.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
