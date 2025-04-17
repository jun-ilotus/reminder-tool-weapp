<template>
    <nut-grid :border="false" column-num=2>
        <nut-grid-item>
            我的积分：
        </nut-grid-item>
        <nut-grid-item>
            {{ store.points }}
        </nut-grid-item>
    </nut-grid>
    <nut-divider style="margin-top: 5vh;"> 积分记录 </nut-divider>
    <nut-form disabled v-for="pointsRecode in pointsRecodeList">
        <nut-form-item :label="pointsRecode.content">
            <nut-input input-align="right" v-model="pointsRecode.points"  />
        </nut-form-item>
    </nut-form>
</template>

<script lang="ts" setup>
import { useDidShow } from '@tarojs/taro'
import { useUserStore } from 'src/stores';
import { getAction } from 'src/http';
import { ref } from 'vue';
import Taro from '@tarojs/taro';

const store = useUserStore()
const pointsRecodeList = ref<any[]>([])

// 页面显示时
useDidShow(() => {
    store.infoReq()  // 刷新一下积分值
    getPointsRecodeList()
})


interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getPointsRecodeList() {
  try {
    const result = await getAction('/usercenter/v1/pointsRecode/list', {}, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;
    if (result.success) {
        pointsRecodeList.value = []
        const list = result.data.pointsRecode
        for (let index = 0; index < list.length; index++) {
          let task = {
            key: index,
            id: list[index].id,
            content: list[index].content,
            points: list[index].points > 0 ? `+${list[index].points}` : `${list[index].points}`,
          }
          pointsRecodeList.value.push(task)
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