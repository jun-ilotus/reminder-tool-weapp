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
                <nut-radio label="3" shape="button">即抽即中</nut-radio>
            </nut-radio-group>
        </nut-form-item>
        <div v-if="lotteryData.announceType === '1'">
            <nut-form-item
                required
                label="开奖时间"
                :rules="[{ required: true, message: '请填写时间'}]"
            >
                <nut-input v-model="announceTimeString" readonly   @click="show = true" />
            </nut-form-item>
        </div>
        <div v-if="lotteryData.announceType === '2'">
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
                <nut-input v-model="announceTimeString" readonly   @click="show = true" />
            </nut-form-item>
        </div>
        <div v-if="lotteryData.announceType === '3'">
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
        </div>
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

    <nut-button block type="primary" @click="createLottery">发起抽奖</nut-button>

    <nut-popup v-model:visible="show" position="bottom">
        <nut-date-picker
        v-model="reminderTime"
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
        v-model="awardDeadlineTime"
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
import { postAction } from "src/http";


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
const reminderTime = ref(now)
const awardDeadlineTime = ref(now)
const announceTimeString = ref(dateFormat(now, 'YYYY-MM-DD HH:mm'))
const awardDeadlineString = ref(dateFormat(now, 'YYYY-MM-DD HH:mm'))

const timeConfirm = ({ selectedValue }) => {
    show.value = false
    announceTimeString.value = dateFormat(reminderTime.value, 'YYYY-MM-DD HH:mm')
    lotteryData.value.announceTime = Math.floor(reminderTime.value/1000)
}
const awardDeadlineConfirm = ({ selectedValue }) => {
    awardDeadlineshow.value = false
    awardDeadlineString.value = dateFormat(awardDeadlineTime.value, 'YYYY-MM-DD HH:mm')
    lotteryData.value.awardDeadline = Math.floor(awardDeadlineTime.value/1000)
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
    announceTime: Math.floor(now/1000),
    joinNumber: 0,
    introduce: "",
    isClocked: false,
    clockTaskId: 0,
    awardDeadline: Math.floor(now/1000),
    prizes: [
        {
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
        name: lotteryData.value.name,
        thumb: lotteryData.value.url,
        publishTime: 0,
        announceType: Number(lotteryData.value.announceType),
        announceTime: lotteryData.value.announceTime,
        joinNumber: Number(lotteryData.value.joinNumber),
        introduce: lotteryData.value.introduce,
        awardDeadline: lotteryData.value.awardDeadline,
        isClocked: lotteryData.value.isClocked ? 1:0,
        clockTaskId: lotteryData.value.clockTaskId,
        prizes: [],
    }
    for (let i = 0; i < lotteryData.value.prizes.length; i ++) {
        data.prizes.push({
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
</script>