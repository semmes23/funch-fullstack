<template>
    <nav class="bg-white shadow-md py-4 px-6 flex justify-between items-center">
        <NuxtLink to="/" class="text-xl font-bold text-blue-600">
            🏠 RoomRes
        </NuxtLink>

        <div class="flex items-center gap-4">
            <!-- ถ้า login -->
            <template v-if="auth.isAuthenticated">
                <span class="text-gray-700">👤 {{ auth.user?.name || 'ผู้ใช้งาน' }}</span>
                <button @click="handleLogout"
                    class="bg-red-500 text-white px-4 py-1 rounded hover:bg-red-600 transition">
                    ออกจากระบบ
                </button>
            </template>

            <!-- ถ้าไม่ได้ login -->
            <template v-else>
                <NuxtLink to="/login" class="text-blue-500 hover:underline">เข้าสู่ระบบ</NuxtLink>
                <NuxtLink to="/register" class="text-green-500 hover:underline">สมัครสมาชิก</NuxtLink>
            </template>
        </div>
    </nav>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()

const handleLogout = () => {
    auth.logout()
    navigateTo('/')
}

</script>