<template>
    <nut-tabbar v-model="active" bottom safe-area-inset-bottom placeholder>
      <nut-tabbar-item tab-title="抽奖" name="list" @click="goList">
        <template #icon>
          <Category></Category>
        </template>
      </nut-tabbar-item>

      <nut-tabbar-item tab-title="发起" name="create" @click="goCreate">
        <template #icon>
          <Star></Star>
        </template>
      </nut-tabbar-item>
    </nut-tabbar>
  </template>
  
<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { Category, Star } from '@nutui/icons-vue-taro'

const pages = Taro.getCurrentPages()
const page = pages[pages.length-1]

const path = page.route
const parts = path.split("/"); // 使用 "/" 分割字符串
const target = parts[2]; // 获取第三个部分（索引为2）

const active = target

const goList = () => {
  Taro.redirectTo({ url: '/pages/lottery/list/index' })
}
const goCreate = () => {
  Taro.redirectTo({ url: '/pages/lottery/create/index?id=0' })
}

</script>
  