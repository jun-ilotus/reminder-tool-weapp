<template>
  <nut-tabs v-model="value">
    <nut-tab-pane title="待提醒" pane-key="1">

      <nut-cell title="事项" sub-title="今天12:00">
        <template #desc>
          <IconFont @click="click" name="horizontal-n"></IconFont>
        </template>
      </nut-cell>

      <nut-cell title="事项" sub-title="今天12:00">
        <template #desc>
          <IconFont @click="click" name="horizontal-n"></IconFont>
        </template>
      </nut-cell>

    </nut-tab-pane>
    <nut-tab-pane title="已提醒" pane-key="2"> Content 2 </nut-tab-pane>
    <nut-tab-pane title="已完成" pane-key="3"> Content 3 </nut-tab-pane>
  </nut-tabs>
  <Tabbar></Tabbar>
  <nut-action-sheet v-model:visible="show" :menu-items="menuItems" title="事项" @choose="choose" cancel-txt="取消"/>
</template>

<script lang="ts" setup>
import Tabbar from "../../components/Tabbar.vue";
import { IconFont } from '@nutui/icons-vue-taro'
import { ref } from 'vue'

const value = ref('1')

const show = ref(false)
const val = ref('')
const menuItems = [
  {
    name: '修改'
  },
  {
    name: '完成'
  },
  {
    name: '删除',
    color: 'red',
  }
]
const click = () => {
  show.value = true
}
const choose = (item) => {
  val.value = item.name
}
</script>
