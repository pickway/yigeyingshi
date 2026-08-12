const TOKEN_KEY = 'cineverse_admin_token'

export function getToken(storage = sessionStorage) { return storage.getItem(TOKEN_KEY) || '' }
export function setToken(token, storage = sessionStorage) { if (token) storage.setItem(TOKEN_KEY, token); else storage.removeItem(TOKEN_KEY) }
export function clearToken(storage = sessionStorage) { storage.removeItem(TOKEN_KEY) }
export function buildListQuery(params = {}) { const q = new URLSearchParams(); if (params.page) q.set('page', params.page); if (params.pageSize) q.set('pageSize', params.pageSize); if (String(params.keyword || '').trim()) q.set('keyword', String(params.keyword).trim()); return q.toString() }

