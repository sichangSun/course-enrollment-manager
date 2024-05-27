import { createRouter, createWebHistory } from 'vue-router'
import StudentHome from '../views/StudentHomeView.vue'
import LoginView from '@/views/LoginView.vue'
import CourseListALook from '../views/CourseListALook.vue'
import CourseDetail from '../views/CourseDetail.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'LoginView',
      component: LoginView
    },
    {
      //
      path: '/home',
      name: 'StudentHome',
      component: () => import('../views/StudentHomeView.vue'),
    },
    {
      //
      path: '/course',
      name: 'CourseListALook',
      component: () => import('../views/CourseListALook.vue'),
    },
    {
      //
      path: '/detail',
      name: 'CourseDetail',
      component: () => import('../views/CourseDetail.vue'),
    },
    {
      //
      path: '/change-password',
      name: 'ChangePassword',
      component: () => import('../views/ChangePassword.vue'),
    },
    {
      //
      path: '/error',
      name: 'ErrorPage',
      component: () => import('../views/ErrorPage.vue'),
    },

  ]
})

export default router
