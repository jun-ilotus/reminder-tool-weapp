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
  
const pages = Taro.getCurrentPages()
const page = pages[pages.length-1]

const path = page.route
const parts = path.split("/"); // 使用 "/" 分割字符串
const target = parts[1]; // 获取第三个部分（索引为2）

const active = target

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
  