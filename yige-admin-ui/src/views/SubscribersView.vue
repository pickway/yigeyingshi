<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Delete, Download, Refresh, Search } from '@element-plus/icons-vue'
import { resourceApi } from '@/api'

const loading = ref(false)
const rows = ref([])
const keyword = ref('')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

function formatDate(value) {
  return new Date(value).toLocaleString('zh-CN', { hour12: false }).replaceAll('/', '-')
}

async function load() {
  loading.value = true
  try {
    const params = new URLSearchParams({ page: String(page.value), pageSize: String(pageSize.value) })
    if (keyword.value.trim()) params.set('keyword', keyword.value.trim())
    const result = await resourceApi.list('subscribers', params.toString())
    rows.value = result.data
    total.value = result.total
  } catch (error) {
    ElMessage.error(error.message)
  } finally {
    loading.value = false
  }
}

function search() {
  page.value = 1
  load()
}

async function changeStatus(row) {
  const next = !row.active
  try {
    await resourceApi.setSubscriberStatus(row.id, next)
    row.active = next
    ElMessage.success(next ? '已恢复订阅' : '已暂停订阅')
  } catch (error) {
    ElMessage.error(error.message)
  }
}

async function remove(row) {
  try {
    await ElMessageBox.confirm(`确定删除订阅邮箱 ${row.email}？`, '删除订阅者', { type: 'warning', confirmButtonText: '删除' })
    await resourceApi.remove('subscribers', row.id)
    ElMessage.success('订阅者已删除')
    await load()
  } catch (error) {
    if (error !== 'cancel' && error !== 'close') ElMessage.error(error.message)
  }
}

function escapeCsv(value) {
  return `"${String(value ?? '').replaceAll('"', '""')}"`
}

async function exportCsv() {
  loading.value = true
  try {
    const allRows = []
    let exportPage = 1
    let exportTotal = 0
    do {
      const result = await resourceApi.list('subscribers', `page=${exportPage}&pageSize=100`)
      allRows.push(...result.data)
      exportTotal = result.total
      exportPage++
    } while (allRows.length < exportTotal)
    const csv = ['编号,邮箱,状态,订阅时间', ...allRows.map((row) => [row.id, escapeCsv(row.email), row.active ? '启用' : '停用', escapeCsv(formatDate(row.createdAt))].join(','))].join('\r\n')
    const url = URL.createObjectURL(new Blob([`\uFEFF${csv}`], { type: 'text/csv;charset=utf-8' }))
    const anchor = document.createElement('a')
    anchor.href = url
    anchor.download = `cineverse-subscribers-${new Date().toISOString().slice(0, 10)}.csv`
    anchor.click()
    URL.revokeObjectURL(url)
    ElMessage.success(`已导出 ${allRows.length} 位订阅用户`)
  } catch (error) {
    ElMessage.error(error.message)
  } finally {
    loading.value = false
  }
}

load()
</script>

<template>
  <section>
    <div class="page-head">
      <div>
        <h1>订阅用户</h1>
        <p>查看邮件订阅状态，处理停用与删除 · 共 {{ total }} 位订阅者</p>
      </div>
      <el-button :icon="Download" @click="exportCsv">导出名单</el-button>
    </div>
    <div class="summary-strip">
      <div><span class="status-dot active"></span><strong>正常订阅</strong><small>可接收网站内容更新</small></div>
      <div><span class="status-dot"></span><strong>暂停订阅</strong><small>保留记录但不再发送</small></div>
    </div>
    <div class="panel subscriber-panel">
      <div class="toolbar">
        <el-input v-model="keyword" placeholder="搜索邮箱地址" clearable :prefix-icon="Search" @keyup.enter="search" @clear="search" />
        <el-button :icon="Search" @click="search">搜索</el-button>
        <el-button :icon="Refresh" circle title="刷新" @click="load" />
      </div>
      <el-table v-loading="loading" :data="rows" row-key="id" empty-text="暂无订阅用户">
        <el-table-column prop="id" label="编号" width="90" />
        <el-table-column prop="email" label="邮箱地址" min-width="260">
          <template #default="scope"><div class="email-cell"><span>{{ scope.row.email.slice(0, 1).toUpperCase() }}</span><strong>{{ scope.row.email }}</strong></div></template>
        </el-table-column>
        <el-table-column prop="createdAt" label="订阅时间" min-width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column label="状态" width="140">
          <template #default="scope"><el-switch :model-value="scope.row.active" active-text="启用" inactive-text="停用" @change="changeStatus(scope.row)" /></template>
        </el-table-column>
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="scope"><el-button link type="danger" :icon="Delete" @click="remove(scope.row)">删除</el-button></template>
        </el-table-column>
      </el-table>
      <div class="pagination-row"><span>订阅数据仅用于本站内容通知</span><el-pagination v-model:current-page="page" layout="total, prev, pager, next" :page-size="pageSize" :total="total" @current-change="load" /></div>
    </div>
  </section>
</template>

<style scoped>
.summary-strip{display:flex;gap:14px;margin-bottom:18px}.summary-strip>div{display:grid;grid-template-columns:12px auto;column-gap:8px;align-items:center;padding:11px 15px;border:1px solid #e7eaf0;border-radius:10px;background:#fff}.summary-strip strong{font-size:12px}.summary-strip small{grid-column:2;color:var(--admin-muted);font-size:10px}.status-dot{width:8px;height:8px;grid-row:1/3;border-radius:50%;background:#98a2b3}.status-dot.active{background:#3a9b69}.subscriber-panel{overflow:hidden}.toolbar{display:flex;gap:10px;padding:17px 18px;border-bottom:1px solid #edf0f4}.toolbar .el-input{max-width:360px}.email-cell{display:flex;align-items:center;gap:11px}.email-cell>span{width:31px;height:31px;display:grid;place-items:center;border-radius:50%;background:#f2e8d6;color:#9a6a22;font-size:12px}.email-cell strong{font-weight:500}.pagination-row{display:flex;align-items:center;justify-content:space-between;padding:17px 20px;color:var(--admin-muted);font-size:12px}@media(max-width:620px){.summary-strip{display:none}.toolbar{flex-wrap:wrap}.toolbar .el-input{max-width:none;flex:1 0 calc(100% - 100px)}.pagination-row>span{display:none}.pagination-row{justify-content:center}}
</style>
