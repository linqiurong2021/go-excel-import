import Vue from 'vue'
import App from './App.vue'
import store from "./store/index"
// 
import {Loading} from "element-ui"
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(Loading)
Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App),
}).$mount('#app')
