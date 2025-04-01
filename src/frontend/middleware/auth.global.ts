export default defineNuxtRouteMiddleware((to, from) => {
    const auth = useAuthStore()
  
    if (import.meta.client) {
      auth.loadToken()
  
    //   const protectedRoutes = ['/booking', '/calendar', '/profile']
    //   if (protectedRoutes.includes(to.path) && !auth.isAuthenticated) {
    //     return navigateTo('/login')
    //   }
    }
  })
  