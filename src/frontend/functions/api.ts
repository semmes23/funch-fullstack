// frontend/functions/api.ts

const request = async (path: string, options?: RequestInit) => {
    const config = useRuntimeConfig()
    const baseURL = config.public.apiUrl
  
    const response = await fetch(`${baseURL}${path}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...(options?.headers || {})
      }
    })
  
    if (!response.ok) {
      const errorBody = await response.json().catch(() => ({}))
      throw new Error(errorBody.message || 'API request failed')
    }
  
    return response.json()
  }
  
  const API = {
    apiLogin(email: string, password: string) {
      return request('/login', {
        method: 'POST',
        body: JSON.stringify({ email, password })
      })
    },
  
    apiRegister(name: string, email: string, password: string) {
      return request('/register', {
        method: 'POST',
        body: JSON.stringify({ name, email, password })
      })
    },
  
    apiGetUser(token: string) {
      return request('/user', {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
    }
  }
  
  export default API
  