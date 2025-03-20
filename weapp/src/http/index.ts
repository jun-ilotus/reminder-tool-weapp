import { request } from './request'
import { useUserStore } from 'src/stores'
import { useTabbarStore } from 'src/stores'
import Taro from '@tarojs/taro'

export function getAction (url: string, parameter = {}, args = {}, isToken: boolean) {
    if (isToken) {
        if (!useUserStore().getIsLogin) {
            useTabbarStore().setTabbarSelectedName("my")
            Taro.redirectTo({ url: '/pages/my/index' })
            return new Promise(resolve =>
                resolve({ message: '请先登录！', success: false })
            )
        }
        return request({
            url: url,
            method: 'GET',
            data: parameter,
            args: args,
            header: {
                'Authorization': 'Bearer ' + useUserStore().accessToken,
            }
        })
      } else {
            return request({
                url: url,
                method: 'GET',
                data: parameter,
                args: args
        })
    }
}
export function postAction (url: string, parameter = {}, args = {}, isToken: boolean) {
  if (isToken) {
    if (!useUserStore().getIsLogin) {
        useTabbarStore().setTabbarSelectedName("my")
        Taro.redirectTo({ url: '/pages/my/index' })
        return new Promise(resolve =>
            resolve({ message: '请先登录！', success: false })
        )
    }
    return request({
        url: url,
        method: 'POST',
        data: parameter,
        args: args,
        header: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + useUserStore().accessToken,
        }
    })
  } else {
        return request({
            url: url,
            method: 'POST',
            data: parameter,
            args: args,
            header: {
                'Content-Type': 'application/json'
            }
        })
    }
}