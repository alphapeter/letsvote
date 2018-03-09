import Vue from 'vue'
import AdminApp from './Admin.vue'

var VueCookie = require('vue-cookie')
Vue.use(VueCookie)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#admin',
  render: h => h(AdminApp)
})
