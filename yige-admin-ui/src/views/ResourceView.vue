<script setup>
import { computed, nextTick, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Delete, Edit, Plus, Refresh, Search } from '@element-plus/icons-vue'
import { resourceApi } from '@/api'
import { categoryNames, resourceConfigs } from '@/config/resources'

const route = useRoute()
const resource = computed(() => route.meta.resource)
const config = computed(() => resourceConfigs[resource.value])
const loading = ref(false)
const saving = ref(false)
const dialogVisible = ref(false)
const editingId = ref(null)
const formRef = ref()
const rows = ref([])
const keyword = ref('')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const form = reactive({})

const rules = computed(() => Object.fromEntries(config.value.fields
  .filter((field) => field.required)
  .map((field) => [field.key, [{ required: true, message: `请填写${field.label}`, trigger: 'blur' }]])))

function normalizeDate(value) {
  if (!value) return ''
  const date = new Date(value)
  return Number.isNaN(date.getTime()) ? value : date.toLocaleString('zh-CN', { hour12: false }).replaceAll('/', '-')
}

function resetForm(values = {}) {
  for (const key of Object.keys(form)) delete form[key]
  Object.assign(form, structuredClone(config.value.defaults), structuredClone(values))
}

async function load() {
  loading.value = true
  try {
    const params = new URLSearchParams({ page: String(page.value), pageSize: String(pageSize.value) })
    if (keyword.value.trim()) params.set('keyword', keyword.value.trim())
    const result = await resourceApi.list(resource.value, params.toString())
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

function openCreate() {
  editingId.value = null
  resetForm()
  dialogVisible.value = true
  nextTick(() => formRef.value?.clearValidate())
}

async function openEdit(row) {
  try {
    const result = await resourceApi.get(resource.value, row.id)
    editingId.value = row.id
    resetForm(result.data)
    dialogVisible.value = true
    nextTick(() => formRef.value?.clearValidate())
  } catch (error) {
    ElMessage.error(error.message)
  }
}

async function save() {
  if (!await formRef.value.validate().catch(() => false)) return
  saving.value = true
  try {
    const payload = { ...form }
    if (payload.publishedAt) payload.publishedAt = new Date(payload.publishedAt).toISOString()
    else if ('publishedAt' in payload) payload.publishedAt = null
    if (editingId.value) await resourceApi.update(resource.value, editingId.value, payload)
    else await resourceApi.create(resource.value, payload)
    ElMessage.success(`${config.value.singular}${editingId.value ? '更新' : '创建'}成功`)
    dialogVisible.value = false
    await load()
  } catch (error) {
    ElMessage.error(error.message)
  } finally {
    saving.value = false
  }
}

async function remove(row) {
  try {
    await ElMessageBox.confirm(`此操作将永久删除“${row.title || row.name}”，是否继续？`, '确认删除', { type: 'warning', confirmButtonText: '删除', confirmButtonClass: 'danger-confirm' })
    await resourceApi.remove(resource.value, row.id)
    ElMessage.success('删除成功')
    if (rows.value.length === 1 && page.value > 1) page.value--
    await load()
  } catch (error) {
    if (error !== 'cancel' && error !== 'close') ElMessage.error(error.message)
  }
}

load()
</script>

<template>
  <section>
    <div class="page-head">
      <div>
        <h1>{{ route.meta.title }}</h1>
        <p>{{ config.description }} · 共 {{ total }} 条记录</p>
      </div>
      <el-button type="primary" :icon="Plus" @click="openCreate">新建{{ config.singular }}</el-button>
    </div>

    <div class="panel resource-panel">
      <div class="toolbar">
        <el-input v-model="keyword" :placeholder="config.searchPlaceholder" clearable :prefix-icon="Search" @keyup.enter="search" @clear="search" />
        <el-button :icon="Search" @click="search">搜索</el-button>
        <el-button :icon="Refresh" circle title="刷新" @click="load" />
      </div>

      <el-table v-loading="loading" :data="rows" row-key="id" empty-text="暂无内容">
        <el-table-column type="index" label="#" width="58" :index="(index) => (page - 1) * pageSize + index + 1" />
        <el-table-column v-for="column in config.columns" :key="column.key" :label="column.label" :prop="column.key" :width="column.width" :min-width="column.minWidth" show-overflow-tooltip>
          <template #default="scope">
            <div v-if="column.primary" class="primary-cell">
              <span v-if="scope.row.icon || scope.row.posterIcon" class="row-icon" :style="{ background: scope.row.posterColor || '#f2e8d6' }">{{ scope.row.icon || scope.row.posterIcon }}</span>
              <strong>{{ scope.row[column.key] || '未命名' }}</strong>
            </div>
            <el-tag v-else-if="column.boolean" :type="scope.row[column.key] ? 'success' : 'info'" effect="plain">{{ scope.row[column.key] ? '是' : '否' }}</el-tag>
            <el-tag v-else-if="column.category" effect="plain">{{ categoryNames[scope.row[column.key]] || scope.row[column.key] || '-' }}</el-tag>
            <el-tag v-else-if="column.tag && scope.row[column.key]" type="warning" effect="plain">{{ scope.row[column.key] }}</el-tag>
            <span v-else-if="column.rating" class="rating">★ {{ Number(scope.row[column.key] || 0).toFixed(1) }}</span>
            <span v-else-if="column.date">{{ normalizeDate(scope.row[column.key]) }}</span>
            <a v-else-if="column.link && scope.row[column.key]" :href="scope.row[column.key]" target="_blank" rel="noopener noreferrer" class="table-link">{{ scope.row[column.key] }}</a>
            <span v-else>{{ scope.row[column.key] ?? '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="142" fixed="right">
          <template #default="scope">
            <el-button link type="primary" :icon="Edit" @click="openEdit(scope.row)">编辑</el-button>
            <el-button link type="danger" :icon="Delete" @click="remove(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-row">
        <span>第 {{ page }} 页</span>
        <el-pagination v-model:current-page="page" v-model:page-size="pageSize" layout="total, prev, pager, next" :total="total" @current-change="load" />
      </div>
    </div>

    <el-dialog v-model="dialogVisible" :title="`${editingId ? '编辑' : '新建'}${config.singular}`" width="min(760px, 92vw)" destroy-on-close append-to-body>
      <div class="dialog-intro">请完善内容信息，带星号的字段为必填项。</div>
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <div class="form-grid">
          <el-form-item v-for="field in config.fields" :key="field.key" :label="field.label" :prop="field.key" :class="{ wide: field.span === 2 }">
            <el-input-number v-if="field.type === 'number'" v-model="form[field.key]" :min="field.min" :max="field.max" :step="field.step || 1" controls-position="right" />
            <el-select v-else-if="field.type === 'select'" v-model="form[field.key]" placeholder="请选择" clearable>
              <el-option v-for="option in field.options" :key="option[0]" :value="option[0]" :label="option[1]" />
            </el-select>
            <el-switch v-else-if="field.type === 'switch'" v-model="form[field.key]" inline-prompt active-text="是" inactive-text="否" />
            <div v-else-if="field.type === 'color'" class="color-field">
              <el-color-picker v-model="form[field.key]" />
              <el-input v-model="form[field.key]" placeholder="#28364a" />
            </div>
            <el-date-picker v-else-if="field.type === 'datetime'" v-model="form[field.key]" type="datetime" placeholder="选择日期时间" />
            <el-input v-else v-model="form[field.key]" :type="field.type === 'textarea' ? 'textarea' : 'text'" :rows="field.rows" :placeholder="field.placeholder" maxlength="10000" :show-word-limit="field.type === 'textarea'" />
          </el-form-item>
        </div>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saving" @click="save">保存内容</el-button>
      </template>
    </el-dialog>
  </section>
</template>

<style scoped>
.resource-panel{overflow:hidden}.toolbar{display:flex;gap:10px;padding:17px 18px;border-bottom:1px solid #edf0f4}.toolbar .el-input{max-width:360px}.pagination-row{display:flex;align-items:center;justify-content:space-between;padding:17px 20px;color:var(--admin-muted);font-size:12px}.primary-cell{display:flex;align-items:center;gap:10px}.primary-cell strong{font-weight:600}.row-icon{width:32px;height:32px;display:grid;place-items:center;flex:none;border-radius:8px}.rating{color:#aa7525;font-weight:600}.table-link{color:#946723;text-decoration:none}.table-link:hover{text-decoration:underline}.dialog-intro{margin:-5px 0 20px;padding:11px 13px;border-radius:8px;background:#f7f3ea;color:#746349;font-size:12px}.form-grid{display:grid;grid-template-columns:1fr 1fr;gap:0 20px}.form-grid .wide{grid-column:1/-1}.form-grid :deep(.el-input-number),.form-grid :deep(.el-select),.form-grid :deep(.el-date-editor){width:100%}.color-field{display:flex;gap:9px;width:100%}.color-field .el-input{flex:1}@media(max-width:620px){.toolbar{flex-wrap:wrap}.toolbar .el-input{max-width:none;flex:1 0 calc(100% - 100px)}.form-grid{grid-template-columns:1fr}.form-grid .wide{grid-column:auto}.pagination-row>span{display:none}.pagination-row{justify-content:center}}
</style>
