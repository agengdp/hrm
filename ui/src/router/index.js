import {
    createRouter,
    createWebHistory
} from 'vue-router'

import Home from '../pages/Home.vue'
import Login from '../pages/Login.vue'

import profileRoutes from "../modules/profile/router/index.js";
import unitRoutes from "../modules/organization/router/index.js";
import employeeRoutes from "../modules/employee/router/index.js";

const baseRoutes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: {
            title: 'Home HRM',
            layout: 'default'
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            title: 'Login',
            layout: 'blank'
        }
    },
]

const routes = baseRoutes.concat(profileRoutes, unitRoutes, employeeRoutes)

export default createRouter({
    history: createWebHistory(),
    routes
})
