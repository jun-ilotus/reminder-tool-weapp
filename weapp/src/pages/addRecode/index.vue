<template>
    <nut-form>
        <nut-textarea v-model="content" limit-show :max-length="20" placeholder="事项名" />
        <nut-cell title="是否共有">
            <template #link>
                <nut-switch v-model="memberSwitch" />
            </template>
        </nut-cell>
        <nut-space v-if="createOrModify==0" style="margin-left: 20vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button size="small" @click="reset" style="margin-right: 20vw;">重置</nut-button>
            <nut-button type="primary" size="small" @click="submit">提交</nut-button>
        </nut-space>
        <nut-space v-else style="margin-left: 20vw; margin-top: 10px; margin-bottom: 10px;">
            <nut-button type="danger" size="small" @click="deleteItem" style="margin-right: 20vw;">删除事项</nut-button>
            <nut-button type="primary" size="small" @click="submit">修改</nut-button>
        </nut-space>
    </nut-form>

    <nut-dialog v-model:visible="deleteDialog"  @ok="okDelete" >确定要删除该事项吗</nut-dialog>
</template>

<script lang="ts" setup>
import Taro from '@tarojs/taro'
import { useDidShow, useRouter } from '@tarojs/taro'
import { ref } from 'vue'
import { getAction, postAction } from "src/http";

const content = ref('')
const memberSwitch = ref(false)
let itemsId = 0
let createOrModify = 0  // 0添加  1修改
const deleteDialog = ref(false)

// 页面显示时 获取reminder数据
useDidShow(() => {
    const router = useRouter()
    itemsId = Number(router.params.id) as number
    if (itemsId !== 0) {
        Taro.setNavigationBarTitle({title:"修改事项"})
        itemsByid(itemsId)
        createOrModify = 1
    } else {
        Taro.setNavigationBarTitle({title:"添加事项"})
        createOrModify = 0
    }
})

const reset = () => {
    content.value = ""
    memberSwitch.value = false
}

const submit = () => {
    let member = memberSwitch.value ? 1 : 0
    if (createOrModify === 0) {
        createReq(content.value, member)
    } else {
        modifyReq(itemsId, content.value, member)
    }
}

const deleteItem = () => {
    deleteDialog.value = true
}
const okDelete = () => {
    deleteItemReq(itemsId)
}

interface ApiResponse {
    success: boolean;
    data?: any,
    message?: string;
}
async function createReq(content, member) {
  try {
    const result = await postAction('/recode/v1/items/create', {
        content: content,
        member: member,
    }, {
      loadingTitle: '正在创建...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '创建成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        Taro.navigateBack({delta: 1})
    } else {
        Taro.showToast({
            title: '创建失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}

async function itemsByid(itemsId) {
  try {
    const result = await getAction('/recode/v1/items/itemsById', {
        id: itemsId
    }, {
      loadingTitle: '正在获取...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        content.value = result.data.items.content
        if (result.data.items.member === 0) {
            memberSwitch.value = false
        } else {
            memberSwitch.value = true
        }
    } else {
        Taro.showToast({
            title: '获取失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}

async function modifyReq(itemsId, content, member) {
  try {
    const result = await postAction('/recode/v1/items/modify', {
        content: content,
        member: member,
        id: itemsId,
    }, {
      loadingTitle: '正在修改...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '修改成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        Taro.navigateBack({delta: 1})
    } else {
        Taro.showToast({
            title: '修改失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}

async function deleteItemReq(itemsId) {
  try {
    const result = await postAction('/recode/v1/items/delete', {
        id: itemsId,
    }, {
      loadingTitle: '正在删除...', // 请求时显示的加载提示
      toastDuration: 1500 // 错误提示的显示时长
    }, true) as ApiResponse;

    if (result.success) {
        Taro.showToast({
            title: '删除成功！',
            icon: 'success', // 'error' 'success' 'loading' 'none'
            duration: 1000
        })
        Taro.navigateBack({delta: 1})
    } else {
        Taro.showToast({
            title: '删除失败！' + result.message,
            icon: 'error', // 'error' 'success' 'loading' 'none'
            duration: 1500
        })
    }
  } catch (error) {
    Taro.showToast({
        title: '请求异常!',
        icon: 'error', // 'error' 'success' 'loading' 'none'
        duration: 1500
    })
  }
}
</script>
