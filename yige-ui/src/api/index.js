// API 请求封装
const BASE = '/api/v1'

async function request(path, options = {}) {
  const url = `${BASE}${path}`
  const res = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
    ...options,
  })
  if (!res.ok) {
    const payload = await res.json().catch(() => null)
    const apiError = payload?.error
    const error = new Error(apiError?.message || `请求失败（${res.status}）`)
    error.status = res.status
    error.code = apiError?.code || 'REQUEST_FAILED'
    throw error
  }
  return res.json()
}

// === Movies ===
export const movieApi = {
  list: (params = {}) => {
    const q = new URLSearchParams()
    if (params.genre && params.genre !== '全部') q.set('genre', params.genre)
    if (params.keyword) q.set('keyword', params.keyword)
    if (params.sort) q.set('sort', params.sort)
    if (params.page) q.set('page', params.page)
    if (params.pageSize) q.set('pageSize', params.pageSize)
    const qs = q.toString()
    return request(`/movies${qs ? `?${qs}` : ''}`)
  },
  featured: () => request('/movies/featured'),
  detail: (id) => request(`/movies/${id}`),
}

// === Articles ===
export const articleApi = {
  list: (params = {}) => {
    const q = new URLSearchParams()
    if (params.category) q.set('category', params.category)
    if (params.tag) q.set('tag', params.tag)
    const qs = q.toString()
    return request(`/articles${qs ? `?${qs}` : ''}`)
  },
  latest: (limit = 2) => request(`/articles/latest?limit=${limit}`),
  detail: (id) => request(`/articles/${id}`),
}

export const searchApi = {
  search: (query) => request(`/search?q=${encodeURIComponent(query)}`),
}

// === Learning ===
export const learningApi = {
  courses: (params = {}) => {
    const q = new URLSearchParams()
    if (params.category && params.category !== '全部') q.set('category', params.category)
    if (params.level) q.set('level', params.level)
    const qs = q.toString()
    return request(`/learning/courses${qs ? `?${qs}` : ''}`)
  },
  paths: () => request('/learning/paths'),
}

// === AI ===
export const aiApi = {
  tools: (params = {}) => {
    const q = new URLSearchParams()
    if (params.category) q.set('category', params.category)
    const qs = q.toString()
    return request(`/ai/tools${qs ? `?${qs}` : ''}`)
  },
  featuredTools: () => request('/ai/tools/featured'),
  subscribe: (email) => request('/ai/subscribe', {
    method: 'POST',
    body: JSON.stringify({ email }),
  }),
}
