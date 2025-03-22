<template>
    <nut-navbar :title="content"></nut-navbar>
    <!-- <nut-calendar-card v-model="value" @change="onChange"></nut-calendar-card> -->
    <AtCalendar
        :currentDate="recodeTime"  
        @select-date="handleDateChange"
        :marks="markDate"
     />

    <nut-form>
        <nut-form-item label="备注">
            <nut-input v-model="recodeContent"  type="text" ></nut-input>
        </nut-form-item>
        <nut-form-item label="记录时间">
            <nut-input v-model="recodeTimeString" readonly  ></nut-input>
        </nut-form-item>

        <nut-space v-if="addOrModify" style="margin-left: 40vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button type="primary" size="small" @click="addRecode">添加</nut-button>
        </nut-space>

        <nut-space v-else style="margin-left: 30vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button size="small" @click="deleteRecode" style="margin-right: 10vw;">删除</nut-button>
            <nut-button type="primary" size="small" @click="modifyRecode">修改</nut-button>
        </nut-space>

    </nut-form>

    <nut-form disabled>
        <nut-form-item label="记录天数">
            <nut-input v-model="recodeDayCount"  />
        </nut-form-item>
        <nut-form-item label="平均间隔天数">
            <nut-input v-model="recodeDaySpaced"  />
        </nut-form-item>
    </nut-form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useDidShow, useRouter } from '@tarojs/taro'
import { getAction, postAction } from 'src/http'
import Taro from '@tarojs/taro'
import { AtCalendar } from "taro-ui-vue3"
import { dateFormat } from '../utils'
import "taro-ui-vue3/dist/style/components/calendar.scss";


const recodeTime = ref(new Date())
const recodeTimeString = ref(dateFormat(recodeTime.value, 'YYYY-MM-DD'))
const recodeContent = ref('')
const content = ref('')
const recodeList = ref<any[]>([])
const markDate = ref<any[]>([])
const dateTokey = new Map()
const addOrModify = ref(true)  // true 添加  false修改
const recodeDayCount = ref(0)
const recodeDaySpaced = ref(0)
let itemsId = 0

// 页面显示时 获取reminder数据
useDidShow(() => {
    const router = useRouter()
    itemsId = Number(router.params.id) as number
    itemsByid(itemsId)
    getRecodeList()
})

const handleDateChange = (selectedDates) => {
    recodeTimeString.value = selectedDates.value.start
    recodeTime.value = new Date(recodeTimeString.value)
    
    if (dateTokey.has(recodeTimeString.value)) {
        const key = dateTokey.get(recodeTimeString.value)
        recodeContent.value = recodeList.value[key].content
        addOrModify.value = false
    } else {
        recodeContent.value = ""
        addOrModify.value = true
    }
}

const deleteRecode = () => {
    const key = dateTokey.get(recodeTimeString.value)
    deleteReq(recodeList.value[key].id)
}

const addRecode = () => {
    createReq(recodeContent.value, Math.floor(recodeTime.value.getTime()/1000))
}

const modifyRecode = () => {
    const key = dateTokey.get(recodeTimeString.value)
    modifyReq(recodeList.value[key].id, recodeContent.value, key)
}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function itemsByid(itemsId) {
  try {
    const result = await getAction('/recode/v1/items/itemsById', {
        id: itemsId
    }, {
      loadingTitle: '正在获取...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        content.value = result.data.items.content
    } else {
        Taro.showToast({
            title: '获取失败！' + result.message,
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

async function createReq(content, recodeTime) {
  try {
    const result = await postAction('/recode/v1/recode/create', {
        content: content,
        recodeTime: recodeTime,
        itemsId: itemsId,
    }, {
      loadingTitle: '正在添加...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '添加成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        getRecodeList()
    } else {
        Taro.showToast({
            title: '添加失败！' + result.message,
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

async function getRecodeList() {
  try {
    const result = await getAction('/recode/v1/recode/list', {
        lastId: 0,
        page: 0,
        pageSize: 0,
        itemsId: itemsId
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        recodeList.value = []
        markDate.value = []
        dateTokey.clear()
        const list = result.data.list
        for (let index = 0; index < list.length; index++) {
          let reminder = {
            key: index,
            id: list[index].id,
            userId: list[index].userId,
            recodeTime: dateFormat(list[index].recodeTime*1000, 'YYYY-MM-DD'),
            recodeTimeStape: list[index].recodeTime,
            content: list[index].content,
          }
          markDate.value.push({
                value: reminder.recodeTime
          })
          dateTokey.set(reminder.recodeTime, index)
          recodeList.value.push(reminder)
        }
        recodeDayCount.value = result.data.recodeDayCount
        recodeDaySpaced.value = result.data.recodeDaySpaced.toFixed(1)
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

async function modifyReq(id, content, key) {
  try {
    const result = await postAction('/recode/v1/recode/modify', {
        id: id,
        content: content,
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
        recodeList.value[key].content = content
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

async function deleteReq(id) {
  try {
    const result = await postAction('/recode/v1/recode/delete', {
        id: id,
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
        getRecodeList()
        recodeContent.value = ""
        addOrModify.value = true
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