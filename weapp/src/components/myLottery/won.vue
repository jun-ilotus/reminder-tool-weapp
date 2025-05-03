<template>
    <nut-cell-group style="margin-top: 5vh;">
        <div v-for="item in AnnouncedData">
            <nut-cell center @click="click(item.id)" :title="item.name" :sub-title="item.prizeName" is-link></nut-cell>
        </div>
    </nut-cell-group>
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { useReady } from '@tarojs/taro'
import { ref } from 'vue'
import { getAction } from "src/http";
import { dateFormat, numberToChinese } from 'src/pages/utils'

const lotteryData = ref([])
const AnnouncedData = ref([])

useReady (() => {
    getUserLotteryList(1)
})

const click = (lotteryId) => {
    Taro.navigateTo({ 
        url: '/pages/lottery/detail/index?id=' + String(lotteryId)
    })
}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getUserLotteryList (isAnnounced) {
    try {
        const result = await getAction('/lottery/v1/lottery/getUserLotteryList', {
            type: 3,
            isAnnounced,
        }, {
        loadingTitle: '正在加载...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;

        if (result.success) {
            const list = result.data.list
            for (let index = 0; index < list.length; index++) {
                let lottery = {
                    key: index,
                    id: list[index].lotteryId,
                    name: list[index].name,
                    lotteryParticipationId: list[index].LotteryParticipationId,
                    level: list[index].level,
                    prizeName: numberToChinese(list[index].level) + '等奖：' + list[index].prizeName,
                    prizeId: list[index].prizeId,
                }
                AnnouncedData.value.push(lottery)
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