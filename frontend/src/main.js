import Vue from 'vue'
import App from './App'
import {store} from './store/store'
require('./assets/css/main.css')

var VueCookie = require('vue-cookie')
Vue.use(VueCookie)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  render: h => h(App)
})
