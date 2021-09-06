import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Toaster from '@meforma/vue-toaster'

var app = createApp(App)
app.use(router)
app.use(Toaster, {
  position: 'top-right'
})

app.use(VueAxios, axios)

app.mount('#page-top')
