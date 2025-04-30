<template>
    <nut-form>
        <nut-swiper
            :auto-play="3000"
            style="margin: 2vw; height: 30vh;"
        >
            <nut-swiper-item style="height: 100%">
                <img :src="lotteryData.thumb" alt="" style="height: 100%; width: 100%" draggable="false" />
            </nut-swiper-item>
        </nut-swiper>
        <div style="margin-top: 5vw;">
            <div v-for="prize in lotteryData.prizes" style="margin-left: 5vw; margin-top: 2vw;">
                <text>{{ prize.levelString }}：{{ prize.name }}</text>
                <text style="color:#ccc; margin-left: 2vw; font-size: x-small;">× {{ prize.count }}份</text>
            </div>
        </div>
        <div style="margin-left: 5vw; margin-top: 3vw; color:#ccc; font-size: x-small; margin-bottom: 3vw;">
            {{ announceTimeString[lotteryData.announceType] }}：{{ lotteryData.announceTimeString }}
        </div>
    </nut-form>
    <nut-form>
        <div style="margin-left: 5vw; margin-top: 3vw; color:#ccc; font-size: small; margin-bottom: 3vw;">
            抽奖说明
        </div>
        <div style="margin-left: 5vw; margin-top: 3vw; margin-bottom: 3vw;">
            {{ lotteryData.introduce }}
        </div>
    </nut-form>
    <!-- 待开奖 -->
    <nut-form v-if="lotteryData.isAnnounced === 0">
        <div style="display: flex; margin-left: 36%; margin-top: 15%; height: 30vh;">

            <nut-animate v-if="lotteryData.isParticipated === 0" type="ripple" loop duration="1000" 
            style="border-radius: 50%; width: 90vw; height: 90vw;">
                <nut-button type="primary" style="width: 30vw; height: 30vw; border-radius: 50%;">参与抽奖</nut-button>
            </nut-animate>
            <nut-button v-else plain type="primary" style="width: 30vw; height: 30vw; border-radius: 50%;">待开奖</nut-button>

        </div>
    </nut-form>
    
    <!-- 已开奖 -->
    <nut-form v-else>
        
    </nut-form>

</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useReady, Taro, useReachBottom, useRouter  } from '@tarojs/taro'
import { numberToChinese, dateFormat } from 'src/pages/utils'
import { getAction } from "src/http";


const announceTimeString: string[] = ['', '开奖时间 ', '结束时间 ', '开始时间 ']
const lotteryData = ref({
    thumb: "http://image.jilotus.cn/lottery.png",
    id: 0,
    userId: 0,
    name: "",
    announceType: 1,
    announceTime: Math.floor(Date.now()/1000),
    announceTimeString: "",
    joinNumber: 0,
    introduce: "",
    isClocked: false,
    clockTaskId: 0,
    awardDeadline: Math.floor(Date.now()/1000),
    isSelected: 0,
    isAnnounced: 0,
    isParticipated: 0,
    prizes: [
        {
            key: Date.now(),
            id: 0,
            levelString: '一等奖',
            level: 1,
            type: 1,
            name: "",
            count: 0,
            grantType: 1,
        }
    ],
})

let lotteryId = 0
useReady(() => {
    const router = useRouter()
    lotteryId = Number(router.params.id) as number
    getlottery(lotteryId)
})

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getlottery (id) {
    try {
        const result = await getAction('/lottery/v1/lottery/lotteryDetail', {
            id,
        }, {
        loadingTitle: '正在加载...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;
        let data = result.data.lottery
        if (result.success) {
            lotteryData.value.id = data.id,
            lotteryData.value.userId = data.userId,
            lotteryData.value.name = data.name,
            lotteryData.value.thumb = data.thumb,
            lotteryData.value.joinNumber = data.joinNumber,
            lotteryData.value.introduce = data.introduce,
            lotteryData.value.awardDeadline = dateFormat(data.awardDeadline*1000, 'YYYY-MM-DD HH:mm'),
            lotteryData.value.isSelected = data.isSelected,
            lotteryData.value.announceType = data.announceType,
            lotteryData.value.announceTime = (data.announceTime*1000 - Date.now())/1000,
            lotteryData.value.announceTimeString = dateFormat(data.announceTime*1000, 'YYYY年MM月DD日 HH:mm'),
            lotteryData.value.isAnnounced = data.isAnnounced,
            lotteryData.value.isClocked = data.isClocked,
            lotteryData.value.clockTaskId = data.clockTaskId,
            lotteryData.value.isParticipated = result.data.isParticipated
            const prizes = result.data.prizes
            lotteryData.value.prizes = []
            for (let index = 0; index < prizes.length; index++) {
                let prize = {
                    key: index,
                    id: prizes[index].id,
                    levelString: numberToChinese(prizes[index].level) + '等奖',
                    level: prizes[index].level,
                    type: prizes[index].type,
                    name: prizes[index].name,
                    count: prizes[index].count,
                    grantType: prizes[index].grantType,
                }
                lotteryData.value.prizes.push(prize)
            }
            console.log(lotteryData.value)
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