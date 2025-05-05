<template>
    <nut-form>
        <nut-form-item label="封面">
            <nut-uploader
                v-model:file-list="file"
                maximum="1"
                multiple="false"
                :before-xhr-upload="beforeXhrUpload"
                media-type="['image']"
            >
            </nut-uploader>
        </nut-form-item>

        <nut-form-item label="标题" required :rules="[{ required: true, message: '请填写标题' }]">
            <nut-input v-model="lotteryData.name" placeholder="请输入标题" type="text" max-length="50" />
        </nut-form-item>
    </nut-form>

    <nut-form
        v-for="(item, index) in lotteryData.prizes"
            :key="item.key"
    >
        <nut-form-item>
            <nut-radio-group v-model="item.type" direction="horizontal">
                <nut-radio label="1" shape="button">奖品</nut-radio>
                <nut-radio label="2" shape="button">优惠券</nut-radio>
                <nut-radio label="3" shape="button">兑换码</nut-radio>
            </nut-radio-group>
        </nut-form-item>
            <nut-form-item
                required
                :label="item.levelString"
                :rules="[{ required: true, message: '请填写名称'}]"
            >
                <nut-input v-model="item.name" placeholder="请输入名称" type="text" max-length="24" />
            </nut-form-item>
            <nut-form-item
                required
                label="奖品份数"
                :rules="[{ required: true, message: '请输入份数'}]"
            >
                <nut-input v-model="item.count" placeholder="输入份数" type="number"/>
            </nut-form-item>
            <nut-cell title="奖品发放方式" :desc="memberKeyToString[item.grantType]" @click="item.show = true"></nut-cell>
            <nut-popup v-model:visible="item.show" position="bottom">
                <nut-picker :columns="columns" title="奖品发放方式" @confirm="confirm" @cancel="item.show = false" />
            </nut-popup>
        <nut-space style="margin: 10px">
            <nut-button type="primary" size="small" @click="add" v-if="index === lotteryData.prizes.length-1">+ 添加奖项</nut-button>
            <nut-button size="small" @click="remove(index)" v-if="index !== 0">- 删除奖项</nut-button>
        </nut-space>
    </nut-form>

    <nut-form>
        <nut-form-item>
            <nut-radio-group v-model="lotteryData.announceType" direction="horizontal">
                <nut-radio label="1" shape="button">按时间开奖</nut-radio>
                <nut-radio label="2" shape="button">按人数开奖</nut-radio>
                <!-- <nut-radio label="3" shape="button">即抽即中</nut-radio> -->
            </nut-radio-group>
        </nut-form-item>
        <div v-if="lotteryData.announceType === '1'">
            <nut-form-item
                required
                label="开奖时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="lotteryData.announceTimeString" readonly   @click="show = true" />
            </nut-form-item>
        </div>
        <div v-if="lotteryData.announceType === '2'">
            <nut-form-item
                required
                label="开始时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="lotteryData.announceTimeString" readonly   @click="show = true" />
            </nut-form-item>
            <nut-form-item
                required
                label="达到指定人数自动开奖"
                :rules="[{ required: true, message: '请输入人数'}]"
            >
                <nut-input v-model="lotteryData.joinNumber" placeholder="输入人数" type="number"/>
            </nut-form-item>
            <nut-form-item
                required
                label="未达到人数自动开奖时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="lotteryData.awardDeadlineString" readonly   @click="awardDeadlineshow = true" />
            </nut-form-item>
        </div>
        <!-- <div v-if="lotteryData.announceType === '3'">
            <nut-form-item
                required
                label="参与人数上限"
                :rules="[{ required: true, message: '请输入人数'}]"
            >
                <nut-input v-model="lotteryData.joinNumber" placeholder="输入人数" type="number"/>
            </nut-form-item>
            <nut-form-item
                required
                label="开始时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="announceTimeString" readonly   @click="show = true" />
            </nut-form-item>
            <nut-form-item
                required
                label="截止时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="awardDeadlineString" readonly   @click="awardDeadlineshow = true" />
            </nut-form-item>
        </div> -->
        <nut-form-item label="抽奖说明">
            <nut-textarea autosize="{ maxHeight: 200, minHeight: 100 }" max-length="80" limit-show placeholder="请输入抽奖说明" type="text" v-model="lotteryData.introduce" />
        </nut-form-item>
        <nut-collapse>
            <nut-collapse-item name="name1">
                <template #title>更多功能</template>
                <nut-form-item label="设置签到任务">
                    <nut-switch v-model="lotteryData.isClocked"></nut-switch>
                </nut-form-item>
                <nut-form-item label="任务内容">
                    <nut-cell title="click" :desc="val" @click="clockShow = true"></nut-cell>
                    <nut-action-sheet
                        v-model:visible="clockShow"
                        title="任务内容"
                        :menu-items="menuItems"
                        @choose="choose"
                    />
                </nut-form-item>
            </nut-collapse-item>
        </nut-collapse>
    </nut-form>

    <nut-button block type="primary" @click="createLottery">{{lotteryId === 0 ? "发起抽奖" : "更新抽奖"}}</nut-button>

    <nut-popup v-model:visible="show" position="bottom">
        <nut-date-picker
        v-model="lotteryData.announceTime"
        type="datetime"
        :min-date="min"
        :three-dimensional="false"
        @confirm="timeConfirm"
        @cancel="show = false"
        :formatter="formatter"
        ></nut-date-picker>
    </nut-popup>
    <nut-popup v-model:visible="awardDeadlineshow" position="bottom">
        <nut-date-picker
        v-model="lotteryData.awardDeadline"
        type="datetime"
        :min-date="min"
        :three-dimensional="false"
        @confirm="awardDeadlineConfirm"
        @cancel="awardDeadlineshow = false"
        :formatter="formatter"
        ></nut-date-picker>
    </nut-popup>

    <lotteryTabbar></lotteryTabbar>
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro';
import { ref, reactive } from 'vue'
import lotteryTabbar from "src/components/lotteryTabbar.vue"
import { useUserStore } from "src/stores/user";
import { numberToChinese, dateFormat } from 'src/pages/utils'
import { postAction, getAction } from "src/http";
import { useReady, useRouter  } from '@tarojs/taro'


const store = useUserStore();
const memberKeyToString: string[] = ['', '快递邮寄', '让中奖者联系我', '中奖者填写信息']
const columns = ref([
  { text: '快递邮寄', value: 1 },
  { text: '让中奖者联系我', value: 2 },
  { text: '中奖者填写信息', value: 3 },
])
const menuItems = [
  {
    name: 'A',
    value: 0,
  },
  {
    name: 'B',
    value: 1,
  },
  {
    name: 'C',
    value: 2,
  }
]
const choose = (item) => {
    lotteryData.value.clockTaskId = item.value
}
const confirm = ({ selectedValue, selectedOptions }) => {
  for (let i = 0; i < lotteryData.value.prizes.length; i ++) {
        if (lotteryData.value.prizes[i].show) {
            lotteryData.value.prizes[i].grantType = selectedValue[0]
            lotteryData.value.prizes[i].show = false
        }
  }
}
const clockShow = ref(false)
const show = ref(false)
const awardDeadlineshow = ref(false)
var now: Date = new Date()
const min = now

const lotteryId = ref(0)
useReady(() => {
    const router = useRouter()
    lotteryId.value = Number(router.params.id) as number
    if (lotteryId.value !== 0) {
        getlottery(lotteryId.value)
    }
})

const timeConfirm = ({ selectedValue }) => {
    show.value = false
    lotteryData.value.announceTimeString = dateFormat(lotteryData.value.announceTime, 'YYYY年MM月DD日 HH:mm')
}
const awardDeadlineConfirm = ({ selectedValue }) => {
    awardDeadlineshow.value = false
    lotteryData.value.awardDeadlineString = dateFormat(lotteryData.value.awardDeadline, 'YYYY年MM月DD日 HH:mm')
}

const file = reactive([
{
    name: 'lottery.png',
    url: 'http://image.jilotus.cn/lottery.png',
    status: 'success',
    message: '上传成功',
    type: 'image',
  },
])
const lotteryData = ref({
    url: "http://image.jilotus.cn/lottery.png",
    name: "",
    announceType: "1",
    announceTime: now,
    joinNumber: 0,
    introduce: "",
    isClocked: false,
    clockTaskId: 0,
    awardDeadline: now,
    announceTimeString: dateFormat(now, 'YYYY年MM月DD日 HH:mm'),
    awardDeadlineString: dateFormat(now, 'YYYY年MM月DD日 HH:mm'),
    prizes: [
        {
            id: 0,
            key: Date.now(),
            show: false,
            levelString: '一等奖',
            level: 1,
            type: "1",
            name: "",
            count: 0,
            grantType: 1,
        }
    ],
})

const add = () => {
    const k = lotteryData.value.prizes.length
    lotteryData.value.prizes.push({
        id: 0,
        key: Date.now(),
        show: false,
        levelString: numberToChinese(k+1) + '等奖',
        level: 1,
        type: "1",
        name: "",
        count: 0,
        grantType: 1,
    })
}

const remove = (index) => {
    for (let i = 0; i < lotteryData.value.prizes.length; i ++) {
        if (i > index) {
            lotteryData.value.prizes[i - 1] = lotteryData.value.prizes[i]
            lotteryData.value.prizes[i - 1].level = i
            lotteryData.value.prizes[i - 1].levelString = numberToChinese(i) + '等奖'
        }
    }
    lotteryData.value.prizes.pop()
}


const beforeXhrUpload = (taroUploadFile, options) => {
      const uploadTask = taroUploadFile({
        url: process.env.TARO_APP_PROXY + '/upload/v1/upload/add',
        filePath: options.taroFilePath,
        fileType: options.fileType,
        header: {
          'Authorization': 'Bearer ' + store.accessToken,
          ...options.headers
        }, //
        formData: options.formData,
        name: 'file',
        success(response: { errMsg; statusCode; data }) {
          if (options.xhrState == response.statusCode) {
            options.onSuccess?.(response, options);
            lotteryData.value.url = JSON.parse(response.data).data.url
          } else {
            options.onFailure?.(response, options);
          }
        },
        fail(e) {
          options.onFailure?.(e, options);
        }
      });
      options.onStart?.(options);
      uploadTask.progress((res) => {
        options.onProgress?.(res, options);
        // console.log('上传进度', res.progress);
        // console.log('已经上传的数据长度', res.totalBytesSent);
        // console.log('预期需要上传的数据总长度', res.totalBytesExpectedToSend);
      });
      // uploadTask.abort(); // 取消上传任务
    };
const formatter = (type, option) => {
  switch (type) {
    case 'year':
      option.text += '年'
      break
    case 'month':
      option.text += '月'
      break
    case 'day':
      option.text += '日'
      break
    default:
      option.text += ''
  }
  return option
}


interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function createLottery() {
    const data = {
        id: lotteryId.value,
        name: lotteryData.value.name,
        thumb: lotteryData.value.url,
        publishTime: 0,
        announceType: Number(lotteryData.value.announceType),
        announceTime: Math.floor(lotteryData.value.announceTime/1000),
        joinNumber: Number(lotteryData.value.joinNumber),
        introduce: lotteryData.value.introduce,
        awardDeadline: Math.floor(lotteryData.value.awardDeadline/1000),
        isClocked: lotteryData.value.isClocked ? 1:0,
        clockTaskId: lotteryData.value.clockTaskId,
        prizes: [],
    }
    for (let i = 0; i < lotteryData.value.prizes.length; i ++) {
        data.prizes.push({
            id: lotteryData.value.prizes[i].id,
            lotteryId: lotteryData.value.prizes[i].lotteryId,
            type: Number(lotteryData.value.prizes[i].type),
            name: lotteryData.value.prizes[i].name,
            count: Number(lotteryData.value.prizes[i].count),
            thumb: "",
            level: lotteryData.value.prizes[i].level,
            grantType: lotteryData.value.prizes[i].grantType,
        })
    }
  try {
    const result = await postAction('/lottery/v1/lottery/createLottery', data, {
      loadingTitle: '正在创建...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '创建成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        // Taro.navigateBack({delta: 1})
    } else {
        Taro.showToast({
            title: result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    console.log(error)
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
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
            lotteryData.value.awardDeadline = data.awardDeadline*1000,
            lotteryData.value.awardDeadlineString = dateFormat(data.awardDeadline*1000, 'YYYY年MM月DD日 HH:mm'),
            lotteryData.value.isSelected = data.isSelected,
            lotteryData.value.announceType = String(data.announceType),
            lotteryData.value.announceTime = data.announceTime*1000,
            lotteryData.value.announceTimeString = dateFormat(data.announceTime*1000, 'YYYY年MM月DD日 HH:mm'),
            lotteryData.value.isAnnounced = data.isAnnounced,
            lotteryData.value.isClocked = data.isClocked === 1 ? true : false,
            lotteryData.value.clockTaskId = data.clockTaskId,
            lotteryData.value.isParticipated = result.data.isParticipated
            const prizes = result.data.prizes
            lotteryData.value.prizes = []
            for (let index = 0; index < prizes.length; index++) {
                let prize = {
                    key: index,
                    id: prizes[index].id,
                    lotteryId: prizes[index].lotteryId,
                    levelString: numberToChinese(prizes[index].level) + '等奖',
                    level: prizes[index].level,
                    type: String(prizes[index].type),
                    name: prizes[index].name,
                    count: prizes[index].count,
                    grantType: prizes[index].grantType,
                }
                lotteryData.value.prizes.push(prize)
            }
            // console.log(lotteryData.value)
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