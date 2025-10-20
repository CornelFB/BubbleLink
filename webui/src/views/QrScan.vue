<template>
  <div class="page">
    <header class="bar">
      <h2>Scan a QR Code</h2>
      <div class="actions">
        <button class="btn" @click="start" :disabled="running">Start</button>
        <button class="btn ghost" @click="stop" :disabled="!running">Stop</button>
        <button class="btn ghost" @click="goBack">Back</button>
      </div>
      <p v-if="error" class="err">{{ error }}</p>
    </header>

    <section class="scanner">
      <!-- html5-qrcode mounts here -->
      <div id="qr-region" ref="qrRegion" class="qr-region"></div>
    </section>

    <section v-if="result" class="result">
      <h3>Result</h3>
      <p class="mono">{{ result }}</p>
      <div class="actions">
        <button class="btn" @click="open">Open</button>
        <button class="btn ghost" @click="copy">Copy</button>
      </div>
      <p v-if="placeId" class="ok">Matched route: /placeExterior/{{ placeId }}</p>
    </section>
  </div>
</template>

<script setup>
import { onMounted, onBeforeUnmount, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Html5QrcodeScanner } from 'html5-qrcode'

const router = useRouter()

const qrRegion = ref(null)
let scanner = null
const running = ref(false)
const result = ref('')
const error = ref('')
const placeId = ref('')

// set your expected base (works in dev/prod)
const BASE = window.location.origin // e.g., http://localhost:5173

function parsePlaceUrl(text) {
  try {
    const u = new URL(text)
    // accept both same-origin and absolute http(s)
    if (!/^https?:$/.test(u.protocol)) return null
    const path = u.pathname.replace(/\/+$/, '') // trim trailing slash
    // Expect /placeExterior/:placeId
    const m = path.match(/^\/placeExterior\/([^/]+)$/)
    if (m) return { placeId: m[1], url: u.toString() }
    return null
  } catch {
    return null
  }
}

function onScanSuccess(decodedText /*, decodedResult */) {
  result.value = decodedText
  const match = parsePlaceUrl(decodedText)
  placeId.value = match ? match.placeId : ''
  // optional: auto-navigate if it matches our route and origin
  // if (match && decodedText.startsWith(BASE)) router.push(`/placeExterior/${match.placeId}`)
  // else stop only and let the user choose:
  stop()
}

function onScanError(/* err */) {
  // ignore frequent decode errors; leave console quiet
}

function start() {
  if (running.value) return
  error.value = ''
  result.value = ''
  placeId.value = ''

  // Options: https://github.com/mebjas/html5-qrcode
  const config = {
    fps: 12,
    qrbox: { width: 280, height: 280 },
    rememberLastUsedCamera: true,
    // formatsToSupport: [ Html5QrcodeSupportedFormats.QR_CODE ], // default supports QR
    // aspectRatio: 1.0,
    showTorchButtonIfSupported: true
  }

  // Build the scanner UI (camera chooser + viewfinder + status)
  scanner = new Html5QrcodeScanner(qrRegion.value.id, config, /* verbose */ false)
  scanner.render(onScanSuccess, onScanError)
  running.value = true
}

function stop() {
  if (!scanner) return
  // clear() stops camera and removes UI; returns a promise
  scanner.clear().finally(() => {
    scanner = null
    running.value = false
  })
}

function open() {
  try {
    const u = new URL(result.value)
    window.location.href = u.toString()
  } catch {
    error.value = 'Scanned content is not a valid URL.'
  }
}

async function copy() {
  try { await navigator.clipboard.writeText(result.value) } catch {}
}

function goBack() { router.back() }

onMounted(() => {
  // Optionally auto-start on mount:
  // start()
})

onBeforeUnmount(() => stop())
</script>

<style scoped>
.page { padding: 16px; color: #e9ecf1; }
.bar { display:flex; flex-direction:column; gap:12px; margin-bottom: 16px; }
.actions { display:flex; gap:8px; flex-wrap:wrap; }
.btn { background: linear-gradient(135deg, #7aaaff, #8bd0ff); color: #0c0f1a; border:0; padding:8px 12px; border-radius:10px; font-weight:700; cursor:pointer; }
.btn.ghost { background: transparent; border:1px solid rgba(255,255,255,.2); color:#cfe6ff; }
.err { color: #ffb4b4; }
.ok { color: #b4ffcd; }

.scanner { display:grid; place-items:center; }
.qr-region { width: min(520px, 95vw); }
.mono { font-family: ui-monospace, SFMono-Regular, Menlo, monospace; background: rgba(255,255,255,.06); padding: 8px 10px; border-radius: 8px; word-break: break-all; }
</style>
