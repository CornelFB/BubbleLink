// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'

const Login = () => import('@/views/Login.vue')
const MapView = () => import('@/views/MapView.vue')
const MyProfile = () => import('@/views/MyProfile.vue')

// NEW: lazy QR views
const QrCreate = () => import('@/views/QrCreate.vue')
const QrScan = () => import('@/views/QrScan.vue')

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: '/', redirect: '/map' },
		{ path: '/login', name: 'login', component: Login, meta: { guestOnly: true } },
		{ path: '/map', name: 'map', component: MapView, meta: { requiresAuth: true } },
		{ path: '/me', name: 'profile', component: MyProfile, meta: { requiresAuth: true } },

		{ path: '/qr/create/', name: 'qr-create', component: QrCreate, meta: { requiresAuth: true } },
		{ path: '/qr/scan', name: 'qr-scan', component: QrScan, meta: { requiresAuth: true } },

		{ path: '/:pathMatch(.*)*', redirect: '/map' }
	]
})

// keep your guards as-is
router.beforeEach((to, _from, next) => {
	const authed = !!localStorage.getItem('bubble_user')
	if (to.meta.requiresAuth && !authed) return next({ name: 'login' })
	if (to.meta.guestOnly && authed) return next({ name: 'map' })
	next()
})

export default router
