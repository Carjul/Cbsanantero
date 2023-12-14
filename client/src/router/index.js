// router.js
import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LogincbView.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue'),
  },
  {
    path: '/principal',
    name: 'princ',
    component: () => import('../views/PrincView.vue'),
  },
  {
    path: '/artes',
    name: 'art',
    component: () => import('../views/ArteView.vue'),
  },
  {
    path: '/bard',
    name: 'bar',
    component: () => import('../views/BardView.vue'),
  },
  {
    path: '/hotels',
    name: 'hote',
    component: () => import('../views/HotelsView.vue'),
  },
  {
    path: '/hospe',
    name: 'hopeda',
    component: () => import('../views/HospeView.vue'),
  },
  {
    path: '/recre',
    name: 'recrat',
    component: () => import('../views/RecreaView.vue'),
  },
  {
    path: '/restau',
    name: 'restaur',
    component: () => import('../views/RestourView.vue'),
  },
  {
    path: '/tourn',
    name: 'tures',
    component: () => import('../views/TourView.vue'),
  },
  {
    path: '/tranport',
    name: 'transp',
    component: () => import('../views/TransporView.vue'),
  },
  {
    path: '/admin',
    name: 'ad',
    component: () => import('../views/AdminView.vue'),
  },
 
  {
    path: '/login',
    name: 'log',
    component: LoginView,
  },
  {
    path: '/usuario',
    name: 'us',
    component: () => import('../views/ListarUsuario.vue'),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
