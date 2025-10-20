<!-- src/App.vue -->
<template>
  <div class="app">
    <header class="app-header">
      <div class="app-brand">
        <img src="/favicon.ico" alt="logo" />
        <span>BubbleLink</span>
      </div>

      <nav class="app-actions">
        <button class="app-link" @click="goMap">Map</button>

        <button v-if="!user" class="app-btn" @click="goLogin">Log in</button>

        <div v-else class="app-user-chip">
          <span class="key">{{ user.name || `User #${user.userId}` }}</span>
          <span class="divider">•</span>
          <button class="app-link" @click="goProfile">My profile</button>
          <span class="divider">•</span>
          <button class="app-link" @click="onLogout">Logout</button>
        </div>
      </nav>
    </header>

    <router-view />
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

const router = useRouter()
const { user, logout } = useAuth()

function goMap()    { router.push({ name: 'map' }) }
function goLogin()  { router.push({ name: 'login' }) }
function goProfile(){ router.push({ name: 'profile' }) }
function onLogout() { logout(); router.push({ name: 'login' }) }
</script>

<style scoped>
.app {
  min-height: 100vh;
  background: radial-gradient(1200px 600px at 10% -10%, #252b46 0%, #151826 60%, #0c0f1a 100%);
  color: #e9ecf1;
}

/* Header */
.app-header {
  position: sticky; top: 0; z-index: 50;
  display: flex; align-items: center; justify-content: space-between;
  padding: 10px 16px;
  background:
    linear-gradient(180deg, rgba(14,16,28,0.38), rgba(14,16,28,0.18)),
    transparent;
  backdrop-filter: saturate(120%) blur(14px);
  -webkit-backdrop-filter: saturate(120%) blur(14px);
  border-bottom: 1px solid rgba(255,255,255,0.08);
  box-shadow: 0 8px 28px rgba(0,0,0,0.25);
}

.app-brand {
  display: flex; gap: 10px; align-items: center;
  font-weight: 700; letter-spacing: .3px;
}
.app-brand img { width: 20px; height: 20px; filter: drop-shadow(0 0 6px rgba(255,255,255,.2)); }

/* right side */
.app-actions { display: flex; align-items: center; gap: 10px; }

/* Buttons/links (namespaced to avoid collisions) */
.app-btn {
  background: linear-gradient(135deg, #7aaaff, #8bd0ff);
  color: #0c0f1a;
  border: 0;
  padding: 8px 12px;
  border-radius: 10px;
  font-weight: 700;
  box-shadow: 0 6px 18px rgba(138, 174, 255, .25);
  cursor: pointer;
}

.app-link {
  background: transparent;
  border: 0;
  color: #9cd6ff;
  cursor: pointer;
  font-weight: 700;
  padding: 6px 8px;
  border-radius: 8px;
}
.app-link:hover { text-decoration: underline; }

/* User chip */
.app-user-chip {
  display: flex; align-items: center; gap: 8px;
  padding: 6px 10px; border-radius: 999px;
  background: rgba(255,255,255,0.08);
  border: 1px solid rgba(255,255,255,0.12);
  box-shadow: inset 0 1px 0 rgba(255,255,255,0.06);
}
.app-user-chip .divider { opacity: .45; }

/* Small screens: keep header tidy */
@media (max-width: 520px) {
  .app-actions { gap: 6px; }
  .app-user-chip { display: none; } /* keeps header clean on tiny screens */
}
</style>
