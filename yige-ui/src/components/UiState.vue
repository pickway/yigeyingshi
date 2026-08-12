<script setup>
import { AlertCircle, Inbox, LoaderCircle, RefreshCw } from '@lucide/vue'

defineProps({
  type: { type: String, default: 'loading' },
  title: { type: String, default: '' },
  message: { type: String, default: '' },
  compact: Boolean,
})

defineEmits(['retry'])
</script>

<template>
  <div class="ui-state" :class="{ 'ui-state--compact': compact }" role="status">
    <LoaderCircle v-if="type === 'loading'" class="ui-state__icon spin" :size="26" />
    <AlertCircle v-else-if="type === 'error'" class="ui-state__icon ui-state__icon--error" :size="26" />
    <Inbox v-else class="ui-state__icon" :size="26" />
    <div>
      <p class="ui-state__title">{{ title || (type === 'loading' ? '正在放映…' : type === 'error' ? '暂时无法加载' : '这里还没有内容') }}</p>
      <p v-if="message" class="ui-state__message">{{ message }}</p>
    </div>
    <button v-if="type === 'error'" class="ui-state__retry" type="button" @click="$emit('retry')">
      <RefreshCw :size="15" /> 再试一次
    </button>
  </div>
</template>

<style scoped>
.ui-state{min-height:220px;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:14px;padding:36px;text-align:center;border:1px dashed var(--color-border-default);border-radius:var(--radius-xl);background:rgba(255,255,255,.015)}
.ui-state--compact{min-height:140px;padding:24px}.ui-state__icon{color:var(--color-primary)}.ui-state__icon--error{color:var(--state-error)}
.ui-state__title{font-weight:650;color:var(--color-text-primary)}.ui-state__message{max-width:420px;margin-top:5px;color:var(--color-text-tertiary);font-size:.875rem;line-height:1.6}
.ui-state__retry{display:inline-flex;align-items:center;gap:7px;border:1px solid var(--color-border-default);border-radius:999px;padding:8px 14px;background:var(--color-bg-surface);color:var(--color-text-secondary);cursor:pointer}
.spin{animation:spin 1s linear infinite}@keyframes spin{to{transform:rotate(360deg)}}
</style>

