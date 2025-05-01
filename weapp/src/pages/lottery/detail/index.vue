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
        <div style="display: flex; justify-content: center; margin-top: 15%; height: 20vh;">

            <nut-animate v-if="lotteryData.isParticipated === 0" type="ripple" loop duration="1000" 
            style="border-radius: 50%; width: 30vw; height: 30vw;">
                <nut-button @click="participation()" type="primary" style="width: 30vw; height: 30vw; border-radius: 50%;">参与抽奖</nut-button>
            </nut-animate>
            <nut-button v-else plain type="primary" style="width: 30vw; height: 30vw; border-radius: 50%;">待开奖</nut-button>
        </div>

        <!-- 参与者 -->
        <div style="display: flex; justify-content: center;">
            <div style="color:#ccc; font-size: small;">
                已有 {{ participationsData.count }} 人参与
            </div>
        </div>
        <div style="display: flex; justify-content: center; margin-top: 3vh; margin-bottom: 5vh;">
            <nut-avatar-group max-count="7" max-content="..." span="2" size="small">
                <nut-avatar v-for="info in participationsData.list" size="small">
                    <img :src="info.avatar" />
                </nut-avatar>
            </nut-avatar-group>
        </div>
    </nut-form>
    
    <!-- 已开奖 -->
    <nut-form v-else>
        <div style="margin-left: 5vw; margin-top: 3vw; color:#ccc; font-size: small; margin-bottom: 3vw;">
            中奖名单
        </div>
        <div v-for="win in winData" style=" margin: 2vh 5vw 3vh 3vw; border: 1px solid #ccc;  border-radius: 8px;">
            <div style="text-align: center;  line-height: 30px; font-size: small; background-color: #F8F8FF; width: 100%; height: 30px; border-radius: 8px;">
                {{ win.prize.name }}，{{ win.winnerCount }} 人中奖
            </div>
            <div style="display: flex; justify-content: center; margin-top: 3vh; margin-bottom: 3vh;">
                <nut-grid :column-num="4" :border="false" gutter="10">
                    <nut-grid-item v-for="user in win.users" :text="user.nickname">
                        <nut-avatar>
                            <img :src="user.avatar" />
                        </nut-avatar>
                    </nut-grid-item>
                </nut-grid>
            </div>
        </div>
    </nut-form>

</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useReady, useReachBottom, useRouter  } from '@tarojs/taro'
import Taro from '@tarojs/taro'
import { numberToChinese, dateFormat } from 'src/pages/utils'
import { getAction, postAction } from "src/http";

import { My } from '@nutui/icons-vue-taro'


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
const participationsData = ref({
    count: 0,
    list: []
})

const winData = ref([{
    prize: {},
    winnerCount: 0,
    users: [],
}])

let lotteryId = 0
useReady(() => {
    const router = useRouter()
    lotteryId = Number(router.params.id) as number
    getlottery(lotteryId)
})

const participation = () => {
    lotteryParticipation()
}

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
            if (lotteryData.value.isAnnounced === 0) {  // 未开奖
                GetParticipations(0, 10)
            } else if (lotteryData.value.isAnnounced === 1) {  // 已开奖
                GetParticipationsWin()
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

async function lotteryParticipation () {
    try {
        const result = await postAction('/lottery/v1/lottery/participation', {
            lotteryId: lotteryData.value.id
        }, {
        loadingTitle: '正在参与...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;
        if (result.success) {
            lotteryData.value.isParticipated = 1
        } else {
            Taro.showToast({
                title: '参与失败！' + result.message,
                icon: 'error', // 'error' 'success' 'loading' 'none'
                duration: 1500
            })
        }
    } catch (error) {
        
    }
}

async function GetParticipations (lastId, pageSize) {
    try {
        const result = await getAction('/lottery/v1/lottery/participations', {
            lotteryId: lotteryData.value.id,
            lastId,
            pageSize,
        }, {
        loadingTitle: '正在加载...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;
        let data = result.data
        let participations = data.list
        if (result.success) {
            participationsData.value.count = data.count
            for (let index = 0; index < participations.length; index++) {
                let participation = {
                    key: index,
                    nickname: participations[index].nickname,
                    avatar: participations[index].avatar,
                }
                participationsData.value.list.push(participation)
            }
            console.log(participationsData.value)
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

async function GetParticipationsWin () {
    try {
        const result = await getAction('/lottery/v1/lottery/getLotteryWinnersList', {
            lotteryId: lotteryData.value.id,
        }, {
        loadingTitle: '正在加载...', // 请求时显示的加载提示
        toastDuration: 1500 // 错误提示的显示时长
        }, true) as ApiResponse;
        let list = result.data.list
        if (result.success) {
            winData.value = []

            for (let i = 0; i < list.length; i ++ ) {
                let won = {
                    prize: {
                        id: list[i].prize.id,
                        lotteryId: list[i].prize.lottery_id,
                        type: list[i].prize.type,
                        name: list[i].prize.name,
                        level: list[i].prize.level,
                        count: list[i].prize.count,
                        grantType: list[i].prize.grant_type,
                    },
                    winnerCount: list[i].winnerCount,
                    users: [],
                }
                for (let j = 0; j < list[i].users.length; j ++ ) {
                    let user = {
                        key: j,
                        nickname: list[i].users[j].nickname,
                        avatar: list[i].users[j].avatar,
                    }
                    won.users.push(user)
                }
                winData.value.push(won)
            }
            console.log(winData.value)
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