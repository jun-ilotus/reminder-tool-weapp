<template>
    <nut-navbar title="事项" @click-right="addRecode">
        <template #right>
            <Uploader></Uploader>
        </template>
    </nut-navbar>

    <nut-swipe-group lock>
        <div v-for="item in itemsList">
            <nut-swipe :ref="(el) => handleSwipeRef(item.key, el)" :name="item.key">
                <nut-cell :title="item.content" is-link center @click.stop="toLeft(item.key)"></nut-cell>
                <template #right>
                    <nut-button shape="square" style="height: 100%" type="success" @click="detail(item.id)">查看</nut-button>
                    <nut-button shape="square" style="height: 100%" type="info" @click="modify(item.id)">修改</nut-button>
                </template>
            </nut-swipe>
        </div>
    </nut-swipe-group>

    <Tabbar></Tabbar>
</template>

<script lang="ts" setup>
import Taro from "@tarojs/taro";
import Tabbar from "../../components/Tabbar.vue";
import { Uploader } from '@nutui/icons-vue-taro'
import { useDidShow } from "@tarojs/taro";
import { getAction } from "src/http";
import { ref } from 'vue'

const itemsList = ref<any[]>([])
// 存储所有 swipe 实例的 Map
const swipeInstances = ref(new Map());

const addRecode = () => {
    Taro.navigateTo({ 
      url: '/pages/addRecode/index?id=0'
    })
}
const modify = (id) => {
    Taro.navigateTo({ 
      url: '/pages/addRecode/index?id=' + String(id)
    })
}

const detail = (id) => {
    Taro.navigateTo({ 
      url: '/pages/calendar/index?id=' + String(id)
    })
}

// ref 绑定处理函数
const handleSwipeRef = (key, el) => {
  if (el) {
    swipeInstances.value.set(key, el);
  } else {
    swipeInstances.value.delete(key);
  }
};

const toLeft = (key) => {
    const instance = swipeInstances.value.get(key);
    instance?.open('left');
}

useDidShow(() => {
  getItemsList()
})

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function getItemsList() {
  try {
    const result = await getAction('/recode/v1/items/list', {
        lastId: 0,
        page: 0,
        pageSize: 0,
    }, {
      loadingTitle: '正在加载...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        itemsList.value = []
        const list = result.data.list
        for (let index = 0; index < list.length; index++) {
          let items = {
            key: index,
            id: list[index].id,
            userId: list[index].userId,
            content: list[index].content,
            member: list[index].member,
          }
          itemsList.value.push(items)
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