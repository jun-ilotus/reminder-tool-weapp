import { defineStore } from 'pinia'


const useTabbarStore = defineStore('tabbar', {
    state: () => ({
        tabbarSelectedName: "index",
    }),
    actions: {
      setTabbarSelectedName(name) {
        this.tabbarSelectedName = name
      },
    },
    getters: {
        getTabbarSelectedName() {
            return this.tabbarSelectedName
        },
    },
  })
  export { useTabbarStore }