<template>
    <nut-cell-group desc="未开奖">
        <div v-for="item in lotteryData">
            <nut-cell center @click="click(item.id)" v-if="item.announceType === 1" :title="item.name" :sub-title="item.announceTime" is-link></nut-cell>
            <nut-cell center @click="click(item.id)" v-if="item.announceType === 2" :title="item.name" :sub-title="item.awardDeadline" is-link></nut-cell>
        </div>
    </nut-cell-group>
    <nut-cell-group desc="已开奖" style="margin-top: 5vh;">
        <div v-for="item in AnnouncedData">
            <nut-cell center @click="click(item.id)" v-if="item.announceType === 1" :title="item.name" :sub-title="item.announceTime" is-link></nut-cell>
            <nut-cell center @click="click(item.id)" v-if="item.announceType === 2" :title="item.name" :sub-title="item.awardDeadline" is-link></nut-cell>
        </div>
    </nut-cell-group>
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { useReady } from '@tarojs/taro'
import { ref } from 'vue'
import { getAction } from "src/http";
import { dateFormat } from 'src/pages/utils'

const lotteryData = ref([])
const AnnouncedData = ref([])

useReady (() => {
    getUserLotteryList(0)
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
            type: 1,
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
                    userId: list[index].userId,
                    name: list[index].name,
                    thumb: list[index].thumb,
                    joinNumber: list[index].joinNumber,
                    awardDeadline: '截止时间：' + dateFormat(list[index].awardDeadline*1000, 'YYYY-MM-DD HH:mm'),
                    announceType: list[index].announceType,
                    announceTime: '开奖时间：' + dateFormat(list[index].announceTime*1000, 'YYYY-MM-DD HH:mm'),
                    isAnnounced: list[index].isAnnounced,
                }
                if (isAnnounced === 0) {
                    lotteryData.value.push(lottery)
                } else {
                    AnnouncedData.value.push(lottery)
                }
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