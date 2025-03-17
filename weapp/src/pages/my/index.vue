<template>
    <div v-if="!store.getIsLogin">
        <nut-grid :border="false" column-num=2>
            <nut-grid-item>
                <nut-avatar size="large">
                    <img
                    :src="logoUrl"
                    />
                </nut-avatar>
            </nut-grid-item>
            <nut-grid-item>
                <nut-button type="success" @click="getInfo">登录</nut-button>
            </nut-grid-item>
        </nut-grid>
    </div>
    <div v-else>
        <nut-grid :border="false" column-num=2>
            <nut-grid-item>
                <nut-avatar size="large">
                    <img
                    :src="store.avatar"
                    />
                </nut-avatar>
            </nut-grid-item>
            <nut-grid-item>
                {{store.getNickName || '加载中...'}}
            </nut-grid-item>
        </nut-grid>
    </div>
    <Tabbar></Tabbar>
</template>
  
<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { postAction } from "../../http/index";
import Tabbar from "../../components/Tabbar.vue";
import { useUserStore } from "../../stores/user";
import logoUrl from "../../assets/images/anonymous.png";

const store = useUserStore()

const getInfo = () => {
    Taro.getUserProfile({
        desc: '用于完善用户资料',
        success: (res) => {
            if (res.iv) {
                login(res.iv, res.encryptedData)
            }
        }
    })
}

const login = (iv, encryptedData) => {
    Taro.login({
        success: function (res) {
            if (res.code) {
                loginReq({
                    code: res.code,
                    iv: iv,
                    encryptedData: encryptedData,
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
    iv: string;
    encryptedData: string;
}

// 调用登录接口
async function loginReq(params: LoginParams) {
  try {
    const result = await postAction('/usercenter/v1/user/wxMiniAuth', params, {
      loadingTitle: '正在登录...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    });

    if (result.success) {
    //   console.log('登录成功', result.data);
      store.setLogin({
        accessToken: result.data.accessToken,
        accessExpire: result.data.accessExpire,
        refreshAfter: result.data.refreshAfter,
      })
      store.infoReq()
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
}



</script>