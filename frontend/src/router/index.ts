import { createRouter, createWebHistory } from 'vue-router'
import { kitchenSinkRoutes } from './kitchen-sink'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Scrape from '../views/Scrape.vue'
import SignUp from '../views/SignUp.vue'
import Recipe from '../views/Recipe.vue'

const routes = [
    {
        path: '/',
        component: Home,
    },
    {
        path: '/scrape',
        component: Scrape,
    },
    {
        path: '/login',
        component: Login,
    },
    {
        path: '/signup',
        component: SignUp,
    },
    {
        path: '/recipe',
        component: Recipe,
    },
]

routes.push(kitchenSinkRoutes)

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
