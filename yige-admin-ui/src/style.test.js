import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import test from 'node:test'

const globalStyles = readFileSync(new URL('./style.css', import.meta.url), 'utf8')
const loginStyles = readFileSync(new URL('./views/LoginView.vue', import.meta.url), 'utf8')

test('global and Element Plus text share a Windows-safe Chinese font stack', () => {
  assert.match(globalStyles, /--admin-sans:"Microsoft YaHei UI","Microsoft YaHei"/)
  assert.match(globalStyles, /--el-font-family:var\(--admin-sans\)/)
  assert.match(globalStyles, /html,body,#app\{font-family:var\(--admin-sans\)\}/)
})

test('Chinese headings do not fall back through Latin serif fonts', () => {
  assert.match(loginStyles, /\.login-brand h1\{[^}]*font-family:var\(--admin-sans\)/)
  assert.doesNotMatch(loginStyles, /\.login-brand h1\{[^}]*Georgia/)
})
