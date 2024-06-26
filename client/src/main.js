import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import VueQRCodeComponent from 'vue-qrcode-component'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueQrcode from 'vue-qrcode'


Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.component('qr-code', VueQrcode)


Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
