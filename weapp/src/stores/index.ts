import { App } from 'vue'
import { createPinia } from 'pinia'
export { useTabbarStore } from './tabbar'
export { useUserStore } from './user'
export const store = createPinia()

export function setupStore(app: App) {
  app.use(store)
}