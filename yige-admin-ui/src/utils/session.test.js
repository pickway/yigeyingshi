import test from 'node:test'
import assert from 'node:assert/strict'
import { buildListQuery, clearToken, getToken, setToken } from './session.js'

function fakeStorage() { const values = new Map(); return { getItem:key=>values.get(key)||null,setItem:(key,value)=>values.set(key,value),removeItem:key=>values.delete(key) } }
test('session helpers persist and clear token',()=>{const store=fakeStorage();setToken('abc',store);assert.equal(getToken(store),'abc');clearToken(store);assert.equal(getToken(store),'')})
test('list query trims keyword and includes pagination',()=>{assert.equal(buildListQuery({page:2,pageSize:20,keyword:'  诺兰 '}),'page=2&pageSize=20&keyword=%E8%AF%BA%E5%85%B0')})

