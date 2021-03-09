import Vue from 'vue'
import App from './App.vue'
import store from "./store/index"
import VueRouter from 'vue-router'
// 
import {Loading} from "element-ui"
import 'element-ui/lib/theme-chalk/index.css'
import router from "./router/routes"

Vue.use(Loading)
Vue.config.productionTip = false
Vue.use(VueRouter)

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
