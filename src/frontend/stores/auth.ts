// stores/auth.ts
import { defineStore } from 'pinia'
import API from '@/functions/api'

interface User {
  name: string
  email: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(null)
  const user = ref<User | null>(null)

  // ✅ login
  const login = async (email: string, password: string) => {
    const res = await API.apiLogin(email, password)

    token.value = res.token
    localStorage.setItem('token', res.token)

    // ดึงข้อมูล user ทันที (กรณี API.login ไม่ได้ส่ง user)
    const profile = await API.apiGetUser(res.token)
    user.value = profile
    localStorage.setItem('user', JSON.stringify(profile))
  }

  // ✅ logout
  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  // ✅ โหลดข้อมูลจาก localStorage
  const loadToken = () => {
    if (import.meta.client) {
      const savedToken = localStorage.getItem('token')
      if (savedToken) token.value = savedToken

      const savedUser = localStorage.getItem('user')
      if (savedUser) user.value = JSON.parse(savedUser)
    }
  }

  const initialize = () => {
    loadToken()
  }

  const isAuthenticated = computed(() => !!token.value)

  return {
    token,
    user,
    login,
    logout,
    loadToken,
    initialize,
    isAuthenticated
  }
})
