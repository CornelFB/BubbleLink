<template>
  <div class="qr-page">
    <header class="qr-bar">
      <h2>QR Code Generator</h2>

      <div class="qr-row">
        <label>Target URL</label>
        <div class="qr-url-row">
          <input
            v-model.trim="urlStr"
            placeholder="https://your-host/placeExterior/123"
            @keydown.enter.prevent="generate"
          />
          <button class="qr-btn ghost" @click="copyUrl" :disabled="!urlStr">Copy</button>
          <button class="qr-btn ghost" @click="useOrigin">Use {{ origin }}</button>
        </div>
        <p v-if="error" class="qr-err">{{ error }}</p>
      </div>

      <div class="qr-row qr-grid-3">
        <div>
          <label>Size (px)</label>
          <input class="qr-input-num" type="number" min="128" max="1024" step="32" v-model.number="size" />
        </div>
        <div>
          <label>Margin</label>
          <input class="qr-input-num" type="number" min="0" max="8" step="1" v-model.number="margin" />
        </div>
        <div>
          <label>Error Correction</label>
          <select v-model="ecc">
            <option value="L">L (low)</option>
            <option value="M">M (mid)</option>
            <option value="Q">Q (quartile)</option>
            <option value="H">H (high)</option>
          </select>
        </div>
      </div>

      <div class="qr-actions">
        <button class="qr-btn" @click="generate">Generate</button>
        <button class="qr-btn ghost" @click="downloadPng" :disabled="!generated">Download PNG</button>
        <button class="qr-btn ghost" @click="printPage" :disabled="!generated">Print</button>
        <button class="qr-btn ghost" @click="goBack">Back</button>
      </div>
    </header>

    <section class="qr-canvas-wrap">
      <canvas ref="canvasEl" :width="size" :height="size"></canvas>
      <div class="qr-caption" v-if="generated">{{ urlStr }}</div>
      <div class="qr-muted" v-else>Enter a URL above and click Generate.</div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import QRCode from 'qrcode'
import { useRouter } from 'vue-router'

const router = useRouter()

const origin = window.location.origin
const urlStr = ref('')
const size = ref(512)
const margin = ref(1)
const ecc = ref('M')
const canvasEl = ref(null)
const generated = ref(false)
const error = ref('')

async function generate() {
  error.value = ''
  generated.value = false

  try {
    const u = new URL(urlStr.value)
    if (!/^https?:$/.test(u.protocol)) throw new Error('URL must be http(s)')
    await QRCode.toCanvas(canvasEl.value, u.toString(), {
      width: size.value,
      margin: margin.value,
      errorCorrectionLevel: ecc.value
    })
    generated.value = true
  } catch (e) {
    error.value = 'Please enter a valid http(s) URL.'
    console.warn(e)
  }
}

function downloadPng() {
  const a = document.createElement('a')
  a.download = 'qr.png'
  a.href = canvasEl.value.toDataURL('image/png')
  a.click()
}

function printPage() { window.print() }
function goBack() { router.back() }

async function copyUrl() {
  try { await navigator.clipboard.writeText(urlStr.value) } catch {}
}

function useOrigin() {
  urlStr.value = origin + '/placeExterior/'
}
</script>

<style scoped>
/* ===== Layout ===== */
.qr-page {
  max-width: 960px;
  margin: 0 auto;
  padding: 20px 16px 40px;
  color: #e9ecf1;
}

.qr-bar {
  display: grid;
  gap: 14px;
  margin-bottom: 16px;
}

.qr-row { display: grid; gap: 8px; }
.qr-grid-3 { grid-template-columns: repeat(3, minmax(0, 1fr)); gap: 12px; }

/* url input & helpers */
.qr-url-row {
  display: grid;
  grid-template-columns: 1fr auto auto;
  gap: 8px;
}

/* ===== Controls ===== */
input, select, .qr-btn {
  border-radius: 10px;
  border: 1px solid rgba(255,255,255,.15);
}

input, select {
  width: 100%;
  padding: 10px 12px;
  background: rgba(255,255,255,.06);
  color: #fff;
}
.qr-input-num {
  font-variant-numeric: tabular-nums;
}

.qr-actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.qr-btn {
  background: linear-gradient(135deg, #7aaaff, #8bd0ff);
  color: #0c0f1a;
  border: 0;
  padding: 8px 12px;
  font-weight: 700;
  cursor: pointer;
}
.qr-btn.ghost {
  background: transparent;
  border: 1px solid rgba(255,255,255,.2);
  color: #cfe6ff;
}

/* ===== Canvas card ===== */
.qr-canvas-wrap {
  display: grid;
  place-items: center;
  gap: 10px;
  padding: 24px;
  background: rgba(255,255,255,.03);
  border: 1px solid rgba(255,255,255,.08);
  border-radius: 12px;
  min-height: 420px;
}

/* keep canvas responsive-ish without blurring */
.qr-canvas-wrap canvas {
  width: min(520px, 90vw);
  height: auto;
  max-width: 100%;
}

.qr-caption {
  font-size: 12px;
  opacity: .8;
  word-break: break-all;
  text-align: center;
  max-width: 540px;
}

.qr-muted { opacity: .7; font-size: 13px; }
.qr-err   { color: #ffb4b4; font-size: 13px; }

/* ===== Print ===== */
@media print {
  .qr-bar { display: none; }
  .qr-canvas-wrap {
    border: 0;
    background: none;
    padding: 0;
    min-height: 0;
  }
}
</style>
