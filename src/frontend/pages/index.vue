<template>
    <div class="min-h-screen flex flex-col items-center justify-center bg-gray-50 px-4">
      <div class="text-center space-y-6">
        <h1 class="text-3xl font-bold text-blue-gray-800">Welcome to Room Reservation</h1>
        
        <NuxtLink
          to="/register"
          class="inline-block rounded-lg bg-gradient-to-tr from-pink-500 to-orange-400 px-6 py-3 text-white font-bold shadow-md hover:shadow-lg transition"
        >
          สมัครสมาชิก
        </NuxtLink>
  
        <NuxtLink
          to="/login"
          class="inline-block rounded-lg bg-gradient-to-tr from-green-500 to-emerald-400 px-6 py-3 text-white font-bold shadow-md hover:shadow-lg transition"
        >
          เข้าสู่ระบบ
        </NuxtLink>
      </div>
    </div>
  </template>
  
  <script setup>
  // ยังไม่มี logic ใช้ script setup เปล่า ๆ ได้
  </script>
  