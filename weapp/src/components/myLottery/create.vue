<template>
    <nut-cell-group desc="未开奖">
        <div v-for="item in lotteryData">
            <nut-cell center v-if="item.announceType === 1" :title="item.name" :sub-title="item.announceTime">
                <template #desc>
                    <IconFont @click="clickSheet(item)" name="horizontal-n"></IconFont>
                </template>
            </nut-cell>
            <nut-cell center v-if="item.announceType === 2" :title="item.name" :sub-title="item.awardDeadline">
                <template #desc>
                    <IconFont @click="clickSheet(item)" name="horizontal-n"></IconFont>
                </template>
            </nut-cell>
        </div>
    </nut-cell-group>
    <nut-cell-group desc="已开奖" style="margin-top: 5vh;">
        <div v-for="item in AnnouncedData">
            <nut-cell center v-if="item.announceType === 1" :title="item.name" :sub-title="item.announceTime">
                <template #desc>
                    <IconFont @click="clickSheet(item)" name="horizontal-n"></IconFont>
                </template>
            </nut-cell>
            <nut-cell center v-if="item.announceType === 2" :title="item.name" :sub-title="item.awardDeadline">
                <template #desc>
                    <IconFont @click="clickSheet(item)" name="horizontal-n"></IconFont>
                </template>
            </nut-cell>
        </div>
    </nut-cell-group>
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { useReady } from '@tarojs/taro'
import { ref } from 'vue'
import { getAction } from "src/http";
import { dateFormat } from 'src/pages/utils'
import { IconFont } from '@nutui/icons-vue-taro'

const lotteryData = ref([])
const AnnouncedData = ref([])
const actionSheetTitle = ref('')
const menuItems = [
    {
      key: 0,
      name: '详情',
    },
    {
      key: 1,
      name: '修改',
    },
]
const showSheet = ref(false)
const clickLottery = ref<any>()
const clickSheet = (lottery) => {
    showSheet.value = true
    actionSheetTitle.value = lottery.name
    clickLottery.value = lottery
}
const choose = (item) => {
    if (item.key === 0) {  //查看
        Taro.navigateTo({ 
            url: '/pages/lottery/detail/index?id=' + String(clickLottery.value.id)
        })
    } else if (item.key === 1) {  //修改
    }
  }

useReady (() => {
    getUserLotteryList(0)
    getUserLotteryList(1)
})

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getUserLotteryList (isAnnounced) {
    try {
        const result = await getAction('/lottery/v1/lottery/getUserLotteryList', {
            type: 2,
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