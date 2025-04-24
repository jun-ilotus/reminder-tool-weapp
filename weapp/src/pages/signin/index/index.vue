<template>
    <!-- <nut-calendar-card v-model="value" @change="onChange"></nut-calendar-card> -->
    <AtCalendar
        :currentDate="recodeTime"  
        :marks="markDate"
     />
    <nut-form disabled>
        <nut-form-item label="已连续签到">
            <nut-input v-model="conDays"  />
        </nut-form-item>
    </nut-form>
    <nut-form>
        <nut-button v-if="isToday===0" block type="success" @click="addRecode()">签到</nut-button>
        <nut-button v-else block plain disabled type="success">已签到</nut-button>
    </nut-form>
    <nut-form>
      <nut-form-item body-align="right" label="签到提醒">
        <nut-switch v-model="signRemindSwitch" @change="change"></nut-switch>
      </nut-form-item>
    </nut-form>
    <nut-divider style="margin-top: 5vh;"> 积分任务 </nut-divider>
    <nut-form disabled v-for="task in taskList">
        <nut-cell :title="task.title" :sub-title="task.points" :desc="task.isFinished === 1 ? '已完成':''"></nut-cell>
    </nut-form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useReady } from '@tarojs/taro'
import { getAction, postAction } from 'src/http'
import Taro from '@tarojs/taro'
import { AtCalendar } from "taro-ui-vue3"
import { dateFormat } from '../../utils'
import "taro-ui-vue3/dist/style/components/calendar.scss";


const recodeTime = ref(new Date())
const recodeList = ref<any[]>([])
const markDate = ref<any[]>([])
const conDays = ref(0)
const taskList = ref<any[]>([])
const isToday = ref(0)
const signRemindSwitch = ref(false)
var serverChange = false

// 页面显示时 获取reminder数据
useReady(() => {
    getRecodeList()
    getTaskList()
    getRemindStatus()
})

const addRecode = () => {
    createReq()
}

const change = (value) => {
  if (!serverChange) {
    if (value) {
      Taro.requestSubscribeMessage({
        tmplIds:["T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI"],
        entityIds:["T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI"],
        success: function (res) { 
          console.log(res)
          if (res.T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI === "accept") {
            changeSignRemind()
          }
        }
      })
    } else {
      changeSignRemind()
    }
  }
  serverChange = false
}


interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function createReq() {
  try {
    const result = await postAction('/signin/v1/recode/addToday', {}, {
      loadingTitle: '正在签到...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '签到成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        getRecodeList()
        getTaskList()
    } else {
        Taro.showToast({
            title: '签到失败，' + result.message,
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
    const result = await getAction('/signin/v1/recode/list', {
        lastId: 0,
        page: 0,
        pageSize: 0,
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        recodeList.value = []
        markDate.value = []
        const list = result.data.list
        for (let index = 0; index < list.length; index++) {
          let reminder = {
            key: index,
            id: list[index].id,
            userId: list[index].userId,
            recodeTime: dateFormat(list[index].signDate*1000, 'YYYY-MM-DD'),
            recodeTimeStape: list[index].signDate,
          }
          markDate.value.push({
                value: reminder.recodeTime
          })
          recodeList.value.push(reminder)
        }
        conDays.value = result.data.days
        isToday.value = result.data.isToday
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

async function getTaskList() {
  try {
    const result = await getAction('/signin/v1/task/list', {
        lastId: 0,
        page: 0,
        pageSize: 0,
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        taskList.value = []
        const list = result.data.list
        for (let index = 0; index < list.length; index++) {
          let task = {
            key: index,
            id: list[index].id,
            title: list[index].title,
            type: list[index].type,
            points: `积分+${list[index].points}`,
            isFinished: list[index].isFinished,
          }
          taskList.value.push(task)
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

async function getRemindStatus() {
  try {
    const result = await postAction('/signin/v1/remind/getStatus', {}, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        const status = result.data.status
        if (status === 1) {
          serverChange = true
          signRemindSwitch.value = true
        } else {
          signRemindSwitch.value = false
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

async function changeSignRemind() {
  try {
    const result = await postAction('/signin/v1/remind/change', {
      status: signRemindSwitch.value ? 1 : 0
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        const status = result.data.status
        if (status === 1 && signRemindSwitch.value) {
          Taro.showToast({
              title: '开启成功',
              icon: 'success', // 'error' 'success' 'loading' 'none'
              duration: 1500
          })
        } else if (status === 0 && !signRemindSwitch.value) {
          Taro.showToast({
              title: '关闭成功',
              icon: 'success', // 'error' 'success' 'loading' 'none'
              duration: 1500
          })
        } else {
          Taro.showToast({
            title: '设置失败',
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
          })
          serverChange = true
          signRemindSwitch.value = status === 1 ? true : false
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
</script>