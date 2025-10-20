import { ref } from 'vue'

function readUser() {
    try {
        const raw = localStorage.getItem('bubble_user')
        return raw ? JSON.parse(raw) : null
    } catch {
        return null
    }
}

// one shared ref across all components importing this module
const userRef = ref(readUser())

function setUser(u) {
    userRef.value = u
    if (u) localStorage.setItem('bubble_user', JSON.stringify(u))
    else localStorage.removeItem('bubble_user')
}

function logout() {
    setUser(null)
}

export function useAuth() {
    return { user: userRef, setUser, logout }
}
