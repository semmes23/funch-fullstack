<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100 px-4">
    <div class="w-full max-w-md bg-white p-8 rounded-xl shadow-md space-y-6">
      <h2 class="text-2xl font-bold text-center text-gray-800">เข้าสู่ระบบ</h2>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-sm text-gray-600 mb-1">อีเมล</label>
          <input v-model="email" type="email" placeholder="your@email.com"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" required />
        </div>

        <div>
          <label class="block text-sm text-gray-600 mb-1">รหัสผ่าน</label>
          <input v-model="password" type="password" placeholder="********"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" required />
        </div>

        <button type="submit" class="w-full bg-green-500 text-white py-2 rounded-lg hover:bg-green-600 transition">
          เข้าสู่ระบบ
        </button>

        <p v-if="errorMessage" class="text-red-500 text-sm text-center">{{ errorMessage }}</p>
      </form>

      <p class="text-sm text-center text-gray-600">
        ยังไม่มีบัญชี? <NuxtLink to="/register" class="text-blue-500 underline">สมัครสมาชิก</NuxtLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'guest'
})

import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()
const auth = useAuthStore()

const handleLogin = async () => {
  try {
    await auth.login(email.value, password.value)
    router.push('/')
  } catch (err: any) {
    errorMessage.value = err.message || 'Login failed'
  }
}
</script>