import test from 'node:test'
import assert from 'node:assert/strict'

import { buildMovieQuery, createPageRange, normalizeExternalUrl } from './content.js'

test('buildMovieQuery trims keyword and omits default genre', () => {
  assert.deepEqual(buildMovieQuery({ keyword: '  诺兰  ', genre: '全部', sort: 'year_desc' }), {
    keyword: '诺兰',
    sort: 'year_desc',
  })
})

test('createPageRange keeps current page visible inside a compact range', () => {
  assert.deepEqual(createPageRange(8, 5), [3, 4, 5, 6, 7])
  assert.deepEqual(createPageRange(3, 1), [1, 2, 3])
})

test('normalizeExternalUrl only allows http and https links', () => {
  assert.equal(normalizeExternalUrl('https://example.com'), 'https://example.com/')
  assert.equal(normalizeExternalUrl('javascript:alert(1)'), '#')
  assert.equal(normalizeExternalUrl(''), '#')
})

