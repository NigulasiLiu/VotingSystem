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
    path: '/manage/addcandidate',
    name: 'addcandidate',
    component: () => import('../views/manage/AddCandidate.vue'),
  },
  {
    path: '/votinglist',
    name: 'votinglist',
    component: () => import('../views/vote/List.vue'),
  },
  {
    path: '/manage/candidatelist',
    name: 'candidatelist',
    component: () => import('../views/manage/CandidateList.vue'),
  },
  {
    path: '/detail/:id',
    name: 'detail',
    component: () => import('../views/vote/Detail.vue'),
  },
  {
    path: '/ballot',
    name: 'ballot',
    component: () => import('../views/vote/Ballot.vue'),
  },
  {
    path: '/candidate/:id',
    name: 'editcandidate',
    component: () => import('../views/manage/EditCandidate.vue'),
  },
  {
    path: '/vote/:key',
    name: 'vote',
    component: () => import('../views/vote/Vote.vue'), // 引入你的投票页面组件
    // props: true, // 允许路由参数作为 props 传递到组件
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
