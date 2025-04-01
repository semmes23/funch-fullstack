// nuxt.config.ts
export default defineNuxtConfig({
  css: ['@/assets/css/tailwind.css'],
  modules: ['@pinia/nuxt', '@nuxtjs/tailwindcss'],

  runtimeConfig: {
    public: {
      apiUrl: process.env.NUXT_PUBLIC_API_URL || 'http://localhost:8080'
    }
  },

  compatibilityDate: '2025-03-25',
})