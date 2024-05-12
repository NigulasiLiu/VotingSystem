import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login/Login.vue'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/register/Register.vue'),
  },
  {
    path: '/manage/addvote',
    name: 'addvote',
    component: () => import('../views/manage/AddVote.vue'),
  },
  {
    path: '/votinglist',
    name: 'votinglist',
    component: () => import('../views/vote/List.vue'),
  },
  {
    path: '/detail/:id',
    name: 'detail',
    component: () => import('../views/vote/Detail.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
