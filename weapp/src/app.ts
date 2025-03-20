import { createApp } from 'vue'
import { setupStore } from './stores/index'

import './app.scss'

const App = createApp({
  onShow(options) {
  },
  // 入口组件不需要实现 render 方法，即使实现了也会被 taro 所覆盖
})
setupStore(App)

export default App
