<template>
  <nut-tabs v-model="tabsValue" @change="tabChange">
    <nut-tab-pane title="待提醒" pane-key="0">

      <div v-for="reminder in reminderList">
        <nut-cell center :title="reminder.content" :sub-title="reminder.reminderTime+' ('+memberKeyToString[reminder.member]+')'">
          <template #icon>
            <IconFont name="clock"></IconFont>
          </template>
          <template #desc>
            <IconFont @click="click(reminder)" name="horizontal-n"></IconFont>
          </template>
        </nut-cell>
      </div>

    </nut-tab-pane>
    <nut-tab-pane title="已提醒" pane-key="1">
      <div v-for="reminder in reminderList">
        <nut-cell center :title="reminder.content" :sub-title="reminder.reminderTime+' ('+memberKeyToString[reminder.member]+')'">
          <template #icon>
            <IconFont name="clock"></IconFont>
          </template>
          <template #desc>
            <IconFont @click="click(reminder)" name="horizontal-n"></IconFont>
          </template>
        </nut-cell>
      </div>
    </nut-tab-pane>

    <nut-tab-pane title="已完成" pane-key="2">
      <div v-for="reminder in reminderList">
        <nut-cell center :title="reminder.content" :sub-title="reminder.reminderTime+' ('+memberKeyToString[reminder.member]+')'">
          <template #icon>
            <IconFont name="clock"></IconFont>
          </template>
          <template #desc>
            <IconFont @click="click(reminder)" name="horizontal-n"></IconFont>
          </template>
        </nut-cell>
      </div>
    </nut-tab-pane>
  </nut-tabs>

  <div style="position: fixed; bottom: 50vh; right: 5vw;">
    <nut-button @click="addReminder" type="success" size="small" style="display: flex; align-items: center; justify-content: center;">
      <IconFont name="uploader" style="vertical-align: middle;"></IconFont>
    </nut-button>
  </div>

  <Tabbar></Tabbar>
  <nut-action-sheet v-model:visible="show" :menu-items="menuItems" :title="actionSheetTitle" @choose="choose" cancel-txt="取消"/>
</template>

<script lang="ts" setup>
import Taro from "@tarojs/taro";
import Tabbar from "../../components/Tabbar.vue";
import { IconFont } from '@nutui/icons-vue-taro'
import { ref } from 'vue'
import { getAction, postAction } from "src/http";
import { dateFormat } from "../utils";
import { useDidShow } from '@tarojs/taro'

const tabsValue = ref('0')

const show = ref(false)
const actionSheetTitle = ref('')
const actionSheetId = ref(0)
const reminderList = ref<any[]>([])
let memberKeyToString: string[] = ['自己', '他/她', '一起']

const menuItems = ref<any[]>([])
const menuItems0 = [
  {
    key: 0,
    name: '修改'
  },
  {
    key: 1,
    name: '完成'
  },
  {
    key: 2,
    name: '删除',
    color: 'red',
  }
]
const menuItems2 = [
  {
    key: 0,
    name: '修改'
  },
  {
    key: 2,
    name: '删除',
    color: 'red',
  }
]
const clickReminder = ref<any>()
const click = (reminder) => {
  show.value = true
  actionSheetTitle.value = reminder.content
  actionSheetId.value = reminder.id  // 当前选中actionSheet的reminderId
  clickReminder.value = reminder
}

const choose = (item) => {
  if (item.key === 0) {  //修改
    Taro.navigateTo({ 
      url: '/pages/reminder/index?id='+ actionSheetId.value
    })
  } else if (item.key === 1) {  //完成
    console.log(clickReminder.value)
    modifyReq(clickReminder.value.id, 2)
  } else if (item.key === 2) {  //删除
    deleteReq(clickReminder.value.id)
  }
}

const addReminder = () => {
  Taro.navigateTo({ url: '/pages/reminder/index?id=0' })
}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getReminderList() {
  try {
    const result = await getAction('/reminder/v1/reminder/list', {
        lastId: 0,
        page: 0,
        pageSize: 0,
        status: tabsValue.value,
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        reminderList.value = []
        const list = result.data.list
        for (let index = 0; index < list.length; index++) {
          let reminder = {
            key: index,
            id: list[index].id,
            userId: list[index].userId,
            reminderTime: dateFormat(list[index].reminderTime*1000, 'YYYY-MM-DD HH:mm'),
            reminderTimeStape: list[index].reminderTime,
            content: list[index].content,
            status: list[index].status,
            member: list[index].member,
          }
          reminderList.value.push(reminder)
        }
    } else {
        Taro.showToast({
            title: '获取失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    
  }
}

useDidShow(() => {
  getReminderList()
  if (tabsValue.value === "2") {
    menuItems.value = menuItems2
  } else {
    menuItems.value = menuItems0
  }
})

const tabChange = () => {
  getReminderList()
  if (tabsValue.value === "2") {
    menuItems.value = menuItems2
  } else {
    menuItems.value = menuItems0
  }
}

async function modifyReq(reminderId, status) {
  try {
    const result = await postAction('/reminder/v1/reminder/done', {
        id: reminderId,
        status: status,
    }, {
      loadingTitle: '正在修改...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '修改成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        getReminderList()
    } else {
        Taro.showToast({
            title: '修改失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}

async function deleteReq(reminderId) {
  try {
    const result = await postAction('/reminder/v1/reminder/delete', {
        id: reminderId,
    }, {
      loadingTitle: '正在删除...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '删除成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        getReminderList()
    } else {
        Taro.showToast({
            title: '删除失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}
</script>
