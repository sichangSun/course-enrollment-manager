import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'

import App from './App.vue'
import router from './router'
// import vuetify from '../plugins/vuetify/vuetify.js'

// Vuetify
import 'vuetify/styles'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'


const vuetify = createVuetify({
  components,
  directives,
})



const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(vuetify).mount('#app')
