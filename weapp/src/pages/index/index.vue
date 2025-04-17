<template>
  <nut-grid :column-num="2" style="margin-top: 20vh;">
    <nut-grid-item text="提醒待办" @click="toReminder"><Order /></nut-grid-item>
    <nut-grid-item text="签到" @click="toSignin"><Check /></nut-grid-item>
  </nut-grid>

  <Tabbar></Tabbar>
</template>

<script lang="ts" setup>
import { Order, Check } from '@nutui/icons-vue-taro'
import Taro from '@tarojs/taro';
import Tabbar from "../../components/Tabbar.vue";
import { useDidShow } from '@tarojs/taro'
import { useUserStore } from 'src/stores';
import { postAction } from 'src/http';

const store = useUserStore()

const toReminder = () => {
  Taro.navigateTo({ 
      url: '/pages/reminderIndex/index'
  })
}

const toSignin = () => {
  Taro.navigateTo({ 
      url: '/pages/signin/index/index'
  })
}

useDidShow(() => {
    if (!store.is_login) {
      login()
    }
  })

  const login = () => {
      Taro.login({
          success: function (res) {
              if (res.code) {
                  loginReq({
                      code: res.code,
                  });
              } else {
                  Taro.showToast({
                      title: '登录失败！' + res.errMsg,
                      icon: 'none', // 'error' 'success' 'loading' 'none'
                      duration: 1500
                  })
              }
          }
      })
  }
  
  // 用户登录接口的参数
  interface LoginParams {
      code: string;
  }
  interface ApiResponse {
      success: boolean;
      data?: any,
      message?: string;
  }
  
  // 调用登录接口
  async function loginReq(params: LoginParams) {
      const result = await postAction('/usercenter/v1/user/wxMiniAuth', params, {
        loadingTitle: '正在登录...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
      }, false) as ApiResponse;
  
      if (result.success) {
        store.setLogin({
          is_login: true,
          accessToken: result.data.accessToken,
          accessExpire: result.data.accessExpire,
          refreshAfter: result.data.refreshAfter,
        })
        store.infoReq()
      } 
  }
</script>
