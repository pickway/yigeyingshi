export function buildMovieQuery({ keyword = '', genre = '全部', sort = 'rating_desc' } = {}) {
  const query = {}
  const normalizedKeyword = String(keyword).trim()
  if (normalizedKeyword) query.keyword = normalizedKeyword
  if (genre && genre !== '全部') query.genre = genre
  if (sort) query.sort = sort
  return query
}

export function createPageRange(totalPages, currentPage, maxVisible = 5) {
  const total = Math.max(1, Number(totalPages) || 1)
  const current = Math.min(total, Math.max(1, Number(currentPage) || 1))
  const size = Math.min(total, Math.max(1, Number(maxVisible) || 5))
  let start = Math.max(1, current - Math.floor(size / 2))
  const end = Math.min(total, start + size - 1)
  start = Math.max(1, end - size + 1)
  return Array.from({ length: end - start + 1 }, (_, index) => start + index)
}

export function normalizeExternalUrl(value) {
  if (!value) return '#'
  try {
    const url = new URL(value)
    return ['http:', 'https:'].includes(url.protocol) ? url.toString() : '#'
  } catch {
    return '#'
  }
}

export function splitList(value) {
  return String(value || '')
    .split(',')
    .map((item) => item.trim())
    .filter(Boolean)
}

export function formatChineseDate(value) {
  if (!value) return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return ''
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }).format(date)
}

