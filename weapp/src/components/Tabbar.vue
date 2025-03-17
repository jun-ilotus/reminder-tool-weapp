<template>
    <nut-tabbar v-model="active" @tab-switch="tabSwitch" bottom safe-area-inset-bottom placeholder>
      <nut-tabbar-item tab-title="首页" name="index" @click="goIndex">
        <template #icon>
          <Home></Home>
        </template>
      </nut-tabbar-item>
      <nut-tabbar-item tab-title="我的" name="my" @click="goMy">
        <template #icon>
          <My></My>
        </template>
      </nut-tabbar-item>
    </nut-tabbar>
  </template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import Taro from '@tarojs/taro'
import { Home, Category, Find, Cart, My } from '@nutui/icons-vue-taro'
import { useTabbarStore } from "../stores/tabbar";
  
const store = useTabbarStore()

const active = store.getTabbarSelectedName
  
const tabSwitch = (item: Record<string, unknown>, index: number) => {
    store.setTabbarSelectedName(item.name)
}

const goIndex = () => {
    if (store.getTabbarSelectedName !== "index") {
        Taro.redirectTo({ url: '/pages/index/index' })
    }
}
const goMy = () => {
    if (store.getTabbarSelectedName !== "my") {
        Taro.redirectTo({ url: '/pages/my/index' })
    }
}
</script>
  