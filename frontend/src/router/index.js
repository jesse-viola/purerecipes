import { createRouter, createWebHistory } from 'vue-router'
import { kitchenSinkRoutes } from './kitchen-sink'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import About from '../views/About.vue'
import Scrape from '../views/Scrape.vue'

const routes = [
    {
        path: '/',
        component: Home,
    },
    {
        path: '/scrape',
        component: Scrape
    },
    {
        path: '/login',
        component: Login,
    },
    {
        path: '/about',
        component: About,
    },
]

routes.push(kitchenSinkRoutes)

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
