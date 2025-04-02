<template>
    <div v-if="store.intimateId==0">
        <nut-grid :border="false" column-num=2>
            <nut-grid-item>
                <nut-avatar size="large">
                    <img
                    :src="logoUrl"
                    />
                </nut-avatar>
            </nut-grid-item>
            <nut-grid-item>
                友人A
            </nut-grid-item>
        </nut-grid>
        <nut-button block type="success"  openType='share'>绑定好友</nut-button>
    </div>
    <div v-else>
        <nut-grid :border="false" column-num=2>
            <nut-grid-item>
                <nut-avatar size="large">
                    <img
                    :src="store.intimateAvatar"
                    />
                </nut-avatar>
            </nut-grid-item>
            <nut-grid-item>
                {{store.intimateNickname || '加载中...'}}
            </nut-grid-item>
        </nut-grid>
        <nut-button block type="danger" @click="cancelBind">解除绑定</nut-button>
    </div>

    <nut-dialog title="绑定亲密好友" :visible="bindDialog" @ok="okBind" @cancel="onCancel">{{ diglogContent }}</nut-dialog>
    <nut-dialog v-model:visible="cancelBindDialog"  @ok="okCancel" >确定要取消绑定吗</nut-dialog>
    <nut-notify :visible="showNotify" msg="绑定成功" type="success"/>
</template>
  
<script lang="ts" setup>
import { useUserStore } from "../../stores/user";
import logoUrl from "../../assets/images/anonymous.png";
import { useShareAppMessage, useReady, useRouter } from "@tarojs/taro";
import { ref } from "vue";
import Taro from "@tarojs/taro";
import { postAction } from "src/http";

const store = useUserStore()
const bindDialog = ref(false)
const diglogContent = ref('')
const accessToken = ref('')
const showNotify = ref(false)
const cancelBindDialog = ref(false)

const onCancel = () => {
    bindDialog.value = false
}

useShareAppMessage(() => {
    // if (res.from === 'menu') {
    //   // 来自页面内转发按钮
    //   console.log(res.target)
    // }
    return {
      title: '绑定亲密好友',
      path: '/pages/intimate/index?accessToken=' + store.accessToken + '&nickname=' + store.nickname,
    }
  })


// 页面显示时 获取reminder数据
useReady(() => {
    login()
    const router = useRouter()
    const intimateAccessToken = router.params.accessToken
    diglogContent.value = `你的好友，${router.params.nickname}，请求与你绑定亲密好友关系。`
    if (intimateAccessToken !== undefined) {
        console.log(diglogContent)
        bindDialog.value = true
        accessToken.value = router.params.accessToken as string
    }
})

const okBind = () => {
    bindReq()
    bindDialog.value = false
}

const cancelBind = () => {
    cancelBindDialog.value = true
}

const okCancel = () => {
    cancelBindReq()
}


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

// 调用绑定接口
async function bindReq() {
  try {
    const result = await postAction('/usercenter/v1/user/bindIntimate', {
        intimateAccessToken: accessToken.value
    }, {
      loadingTitle: '正在绑定...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
      store.infoReq()
      showNotify.value = true
    } else {
        Taro.showToast({
            title: '绑定失败！' + result.message,
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
// 调用取消绑定接口
async function cancelBindReq() {
  try {
    const result = await postAction('/usercenter/v1/user/cancelBindIntimate', {
    }, {
      loadingTitle: '正在解除...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
      store.infoReq()
    } else {
        Taro.showToast({
            title: '解除失败！' + result.message,
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