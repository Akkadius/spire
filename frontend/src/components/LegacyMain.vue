<template>
  <div id="legacy-root" />
</template>

<script setup>
import { onMounted, nextTick } from 'vue'

onMounted(async () => {
  const mount = () => {
    if (window.Vue2App) {
      window.Vue2App.mountLegacyApp('#legacy-root', window.location.pathname)
    }
  }

  if (!window.__vue2_loaded__) {
    const script = document.createElement('script')
    script.src = '/legacy/vue2-app.umd.js'
    script.onload = () => {
      window.__vue2_loaded__ = true
      nextTick(mount)
    }
    document.body.appendChild(script)
  } else {
    nextTick(mount)
  }
})
</script>
