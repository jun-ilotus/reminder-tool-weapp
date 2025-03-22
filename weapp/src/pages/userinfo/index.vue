<template>
    <nut-form>
        <nut-form-item label="昵称">
            <nut-input v-model="nickname" type="nickname"></nut-input>
        </nut-form-item>
        <nut-form-item label="头像">
            <nut-avatar> <img :src="avatar" /> </nut-avatar>
            <nut-button  style="float: right; margin-right: 5vw;"
                type="success" open-type="chooseAvatar" @chooseavatar="onChooseAvatar" size="small">选择</nut-button>
        </nut-form-item>
        <nut-space style="margin-left: 40vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button type="primary" size="small" @click="modify">修改</nut-button>
        </nut-space>
    </nut-form>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue';
import { useUserStore } from 'src/stores';
import Taro from '@tarojs/taro';
import { postAction } from 'src/http';

const store = useUserStore()
const nickname = ref(store.nickname)
const avatar = ref(store.avatar)
const baseUrl = process.env.TARO_APP_PROXY
const doAvatar = ref(false)

const onChooseAvatar = (e) => {
    avatar.value = e.detail.avatarUrl
    doAvatar.value = true
}

interface ApiResponse1 {
  code: number;
  msg: string;
  data: {
    url: string;
  };
}
const modify = () => {
    if (doAvatar.value) {
        Taro.uploadFile({
            url: baseUrl + '/upload/v1/upload/add', //仅为示例，非真实的接口地址
            filePath: avatar.value,
            name: 'file',
            formData:{
               
            },
            header: {
                'Authorization': 'Bearer ' + store.accessToken,
                // 'content-type': "multipart/json"
            },
            success: function (result){
                if (result.statusCode == 200) {
                    const parsedData: ApiResponse = JSON.parse(result.data);
                    avatar.value = parsedData.data.url
                    modifyReq()
                } else {
                    Taro.showToast({
                        title: '上传失败！',
                        icon: 'none', // 'error' 'success' 'loading' 'none'
                        duration: 1500
                    })
                }
            }
        })
    } else {
        modifyReq()
    }

}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function modifyReq() {
  try {
    const result = await postAction('/usercenter/v1/user/modifyInfo', {
        avatar: avatar.value, 
        nickname: nickname.value,
    }, {
      loadingTitle: '正在修改...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        store.infoReq()
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