import { request } from './request'
import { useUserStore } from 'src/stores'

export function getAction (url: string, parameter = {}, args = {}, isToken: false) {
    if (isToken) {
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