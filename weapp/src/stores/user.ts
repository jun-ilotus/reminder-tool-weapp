import { defineStore } from 'pinia'
import { postAction } from 'src/http'
import Taro from '@tarojs/taro'


const useUserStore = defineStore('user', {
    state: () => ({
        is_login: false,
        accessToken: "",
        accessExpire: 0,
        refreshAfter: 0,
        id: 0,
        mobile: "",
        nickname: "",
        sex: 0,
        avatar: "",
        info: "",
    }),
    actions: {
      setLogin(value) {
        this.is_login = value.is_login
        this.accessToken = value.accessToken
        this.accessExpire = value.accessExpire
        this.refreshAfter = value.refreshAfter
      },
      setUserInfo(value) {
        this.is_login = value.is_login
        this.id = value.id
        this.mobile = value.mobile
        this.nickname = value.nickname
        this.sex = value.sex
        this.avatar = value.avatar
        this.info = value.info
      },
      async infoReq() {
        try {
          const result = await postAction('/usercenter/v1/user/detail', {}, {}, true) as ApiResponse;
            if (result.success) {
                this.is_login = true
                this.id = result.data.userInfo.id
                this.mobile = result.data.userInfo.mobile
                this.nickname = result.data.userInfo.nickname
                this.sex = result.data.userInfo.sex
                this.avatar = result.data.userInfo.avatar
                this.info = result.data.userInfo.info
            } else {
                Taro.showToast({
                    title: '登录失败！' + result.message,
                    icon: 'none', // 'error' 'success' 'loading' 'none'
                    duration: 1500
                })
            }
    
        } catch (error) {
          Taro.showToast({
              title: '请求异常!',
              icon: 'none', // 'error' 'success' 'loading' 'none'
              duration: 1500
          })
        }
      },
    },
    getters: {
        getIsLogin() {
            return this.is_login
        },
        getAccessToken() {
            return this.accessToken
        },
        getNickName() {
            return this.nickname
        },
    },
  })

  interface ApiResponse {
    success: boolean;
    data?: any; // 或者更具体的类型
    message?: any;
  }

  export { useUserStore }