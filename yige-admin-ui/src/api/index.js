import { clearToken, getToken } from '@/utils/session'

const BASE = import.meta.env.VITE_ADMIN_API_BASE || '/api/admin'

async function request(path, options = {}) {
  const token = getToken()
  const response = await fetch(`${BASE}${path}`, { ...options, headers: { 'Content-Type':'application/json', ...(token ? { Authorization:`Bearer ${token}` } : {}), ...(options.headers || {}) } })
  if (response.status === 401) { clearToken(); if (!location.pathname.endsWith('/login')) location.assign('/login') }
  if (!response.ok) { const payload=await response.json().catch(()=>null); const error=new Error(payload?.error?.message || `请求失败（${response.status}）`); error.code=payload?.error?.code; error.status=response.status; throw error }
  if (response.status === 204) return null
  return response.json()
}

export const authApi = { login: data => request('/auth/login',{method:'POST',body:JSON.stringify(data)}) }
export const dashboardApi = { get: () => request('/dashboard') }
export const resourceApi = {
  list: (resource, query='') => request(`/${resource}${query ? `?${query}` : ''}`),
  get: (resource,id) => request(`/${resource}/${id}`),
  create: (resource,data) => request(`/${resource}`,{method:'POST',body:JSON.stringify(data)}),
  update: (resource,id,data) => request(`/${resource}/${id}`,{method:'PUT',body:JSON.stringify(data)}),
  remove: (resource,id) => request(`/${resource}/${id}`,{method:'DELETE'}),
  setSubscriberStatus: (id,active) => request(`/subscribers/${id}/status`,{method:'PATCH',body:JSON.stringify({active})}),
}

