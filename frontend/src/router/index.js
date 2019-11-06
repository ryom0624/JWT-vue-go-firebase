import Vue from 'vue'
import VueRouter from 'vue-router'
import firebase from "firebase";
import Home from '../views/Home.vue'
import API from '../views/API.vue'
import SignUp from "../views/SignUp";
import SignIn from "../views/SignIn";

Vue.use(VueRouter)

const routes = [
  {
    path: '*',
    redirect: 'signin'
  },
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }, {
    path: '/api',
    name: 'api',
    component: API,
    meta: {requiresAuth: true}
  }, {
    path: '/signup',
    name: 'signup',
    component: SignUp
  }, {
    path: '/signin',
    name: 'signin',
    component: SignIn
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) =>{
  let currentUser = firebase.auth().currentUser
  let requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (requiresAuth && !currentUser) next('signin')
  else if (!requiresAuth && currentUser) next()
  else next()
})

export default router
