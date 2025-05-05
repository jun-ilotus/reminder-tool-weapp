<template>
    
    <div v-for="item in data" style="padding-left: 5vw; padding-right: 5vw; padding-top: 3vw;">
        <nut-card
            :img-url="item.thumb"
            :title="item.name"
            is-need-price="false"
        >
        <template #price>
            <nut-tag color="#FA2400" plain><div style="font-size: x-small;">{{announceTypeString[item.announceType]}} </div> </nut-tag>
        </template>
        <template #origin><div> </div></template>
        <template #shop-tag >
            <div v-if="item.announceType === 1">
                <div style="font-size: x-small;">开奖时间</div>
                <AtCountdown 
                    :isCard="false"
                    :format="{day:'天', hours: '时', minutes: '分', seconds: '秒' }"
                    isShowDay
                    :seconds="item.announceTime"
                    style="font-size: x-small;"
                />
            </div>
            <div v-if="item.announceType === 2">
                <div v-if="item.announceTime > 0">
                    <div style="font-size: x-small;">开始时间</div>
                        <AtCountdown 
                            :isCard="false"
                            :format="{day:'天', hours: '时', minutes: '分', seconds: '秒' }"
                            isShowDay
                            :seconds="item.announceTime"
                            style="font-size: x-small;"
                        />
                </div>
                <div v-else>
                    <div style="font-size: x-small;">结束时间</div>
                        <AtCountdown 
                            :isCard="false"
                            :format="{day:'天', hours: '时', minutes: '分', seconds: '秒' }"
                            isShowDay
                            :seconds="item.awardDeadline"
                            style="font-size: x-small;"
                        />
                </div>
            </div>
        </template>
        <template #footer >
            <nut-button size="mini"  type="primary" @click="lotteryDetail(item.id)" >参与抽奖</nut-button>
        </template>
    </nut-card>  
    </div>

    <nut-divider dashed v-if="none"> 暂时没有了 </nut-divider>
    
    <lotteryTabbar></lotteryTabbar>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useReady, useReachBottom } from '@tarojs/taro'
import Taro from '@tarojs/taro'
import lotteryTabbar from "src/components/lotteryTabbar.vue";
import { dateFormat, getDays, getHours, getMinutes, getSeconds } from 'src/pages/utils'
import { getAction } from "src/http";
import { AtCountdown } from 'taro-ui-vue3'
import "taro-ui-vue3/dist/style/components/countdown.scss";


const announceTypeString: string[] = ['', '按时间开奖', '按人数开奖', '即抽即中']
const announceTimeString: string[] = ['', '开奖时间 ', '结束时间 ', '开始时间 ']

var now: Date = new Date()
const data = ref([])
const none = ref(false)

useReady (() => {
    getLotteryList(0, 10, 0)
})

useReachBottom(() => {
    getLotteryList(data.value[data.value.length-1].id, 10, 0)
})

const lotteryDetail = (id) => {
    Taro.navigateTo({ 
      url: '/pages/lottery/detail/index?id=' + String(id)
    })
}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getLotteryList (lastId, pageSize, isSelected) {
    try {
        const result = await getAction('/lottery/v1/lottery/lotteryList', {
            lastId,
            pageSize,
            isSelected,
        }, {
        loadingTitle: '正在加载...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;

        if (result.success) {
            if (result.data.list === null || result.data.list.length < 10) {
                none.value = true
            }
            const list = result.data.list
            for (let index = 0; index < list.length; index++) {
                let lottery = {
                    key: index,
                    id: list[index].id,
                    userId: list[index].userId,
                    name: list[index].name,
                    thumb: list[index].thumb,
                    joinNumber: list[index].joinNumber,
                    introduce: list[index].introduce,
                    awardDeadline: (list[index].awardDeadline*1000 - Date.now())/1000,
                    isSelected: list[index].isSelected,
                    announceType: list[index].announceType,
                    announceTime: (list[index].announceTime*1000 - Date.now())/1000,
                    isAnnounced: list[index].isAnnounced,
                    isClocked: list[index].isClocked,
                    clockTaskId: list[index].clockTaskId,
                }
                data.value.push(lottery)
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