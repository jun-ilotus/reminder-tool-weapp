<template>
    <nut-form>
        <nut-textarea v-model="content" limit-show :max-length="20" placeholder="提醒内容" />
        <nut-form-item label="提醒时间">
            <nut-input v-model="reminderTimeString" readonly   @click="show = true" />
        </nut-form-item>
        <nut-cell title="提醒谁" :desc="memberKeyToString[member]" @click="click"></nut-cell>
        <nut-space style="margin-left: 30vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button size="small" @click="reset" style="margin-right: 10vw;">重置</nut-button>
            <nut-button type="primary" size="small" @click="submit">提交</nut-button>
        </nut-space>
    </nut-form>

    <nut-popup v-model:visible="show" position="bottom">
        <nut-date-picker
        v-model="reminderTime"
        type="datetime"
        :min-date="min"
        :three-dimensional="false"
        @confirm="confirm"
        ></nut-date-picker>
    </nut-popup>

    <nut-action-sheet
        v-model:visible="showMember"
        :menu-items="menuItems"
        cancel-txt="取消"
        @choose="choose"
    />
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro';
import { ref } from 'vue'
import { dateFormat } from '../utils'
import { getAction, postAction } from 'src/http';
import { useRouter,useDidShow } from '@tarojs/taro';

const content = ref('')
var now: Date = new Date();
const reminderTimeString = ref(dateFormat(now, 'YYYY-MM-DD HH:mm'))
const member = ref(0)
let memberKeyToString: string[] = ['自己', '他/她', '一起']
const menuItems = [
  {
    name: '自己',
    key: 0
  },
  {
    name: '他/她',
    key: 1
  },
  {
    name: '一起',
    key: 2
  }
]
const showMember = ref(false)
const click = () => {
    showMember.value = true
}
const choose = (item) => {
    member.value = item.key
}

const show = ref(false)
const min = now
const reminderTime = ref(now)
let createOrModify = 0 // 0创建  1修改
let reminderId = 0
let status = 0

// 页面显示时 获取reminder数据
useDidShow(() => {
    const router = useRouter()
    reminderId = Number(router.params.id) as number
    if (reminderId !== 0) {
        Taro.setNavigationBarTitle({title:"修改提醒"})
        reminderByid(reminderId)
        createOrModify = 1
    } else {
        Taro.setNavigationBarTitle({title:"添加提醒"})
        createOrModify = 0
    }
})


const confirm = ({ selectedValue }) => {
  // console.log(selectedValue)
  show.value = false
  reminderTimeString.value = dateFormat(reminderTime.value, 'YYYY-MM-DD HH:mm')
}

const reset = () => {
    now = new Date();
    content.value = ""
    reminderTimeString.value = dateFormat(now, 'YYYY-MM-DD HH:mm')
    reminderTime.value = now
}

const submit = () => {
    if (createOrModify === 0) {
      Taro.requestSubscribeMessage({
        tmplIds:["T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI"],
        entityIds:["T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI"],
        success: function (res) { 
          // console.log(res)
          if (res.T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI === "accept") {
            createReq(content.value, reminderTime.value, member.value)
          }
        }
      })
    } else {
        modifyReq(content.value, reminderTime.value, member.value, status)
    }
}
interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}

async function createReq(content, reminderTime, member) {
  try {
    const result = await postAction('/reminder/v1/reminder/create', {
        reminderTime: Math.floor(reminderTime.getTime()/1000),
        content: content,
        member: member,
    }, {
      loadingTitle: '正在创建...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '创建成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        Taro.navigateBack({delta: 1})
    } else {
        Taro.showToast({
            title: '创建失败！' + result.message,
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

async function reminderByid(reminderId) {
  try {
    const result = await getAction('/reminder/v1/reminder/reminderById', {
        id: reminderId
    }, {
      loadingTitle: '正在获取...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        content.value = result.data.reminder.content
        reminderTimeString.value = dateFormat(result.data.reminder.reminderTime*1000, 'YYYY-MM-DD HH:mm')
        reminderTime.value = new Date(result.data.reminder.reminderTime*1000)
        member.value = result.data.reminder.member
        status = result.data.reminder.status
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

async function modifyReq(content, reminderTime, member, status) {
  try {
    const result = await postAction('/reminder/v1/reminder/modify', {
        reminderTime: Math.floor(reminderTime.getTime()/1000),
        content: content,
        member: member,
        id: reminderId,
        status: status
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
        Taro.navigateBack({delta: 1})
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
</script>
