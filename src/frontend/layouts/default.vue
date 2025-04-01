<template>
    <div>
      <Navbar />
      <NuxtPage />
    </div>
  </template>
  
  <script setup>
  import Navbar from '@/components/layout/Navbar.vue' // ✅ path ที่ถูก
  
  import { useAuthStore } from '@/stores/auth'
  const auth = useAuthStore()
  if (import.meta.client) {
    auth.loadToken()
  }
  </script>
  