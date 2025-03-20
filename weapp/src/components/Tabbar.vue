<template>
    <nut-tabbar v-model="active" @tab-switch="tabSwitch" bottom safe-area-inset-bottom placeholder>
      <nut-tabbar-item tab-title="首页" name="index" @click="goIndex">
        <template #icon>
          <Home></Home>
        </template>
      </nut-tabbar-item>

      <nut-tabbar-item tab-title="日期记录" name="recode" @click="goRecode">
        <template #icon>
          <Date></Date>
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
import Taro from '@tarojs/taro'
import { Home, My, Date } from '@nutui/icons-vue-taro'
import { useTabbarStore } from "../stores/tabbar";
  
const store = useTabbarStore()

const active = store.getTabbarSelectedName
  
const tabSwitch = (item: Record<string, unknown>) => {
    store.setTabbarSelectedName(item.name)
}

const goIndex = () => {
  Taro.redirectTo({ url: '/pages/index/index' })
}
const goMy = () => {
  Taro.redirectTo({ url: '/pages/my/index' })
}

const goRecode = () => {
  Taro.redirectTo({ url: '/pages/recode/index' })
}
</script>
  