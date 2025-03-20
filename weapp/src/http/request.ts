import Taro from '@tarojs/taro'

console.log('NODE_ENV', process.env.NODE_ENV)
console.log('TARO_APP_PROXY', process.env.TARO_APP_PROXY)
const baseUrl = process.env.TARO_APP_PROXY

interface RequestParams {
  url: string
  method: 'OPTIONS'|'GET'|'HEAD'|'POST'|'PUT'|'PATCH'|'DELETE'|'TRACE'|'CONNECT'
  data: any
  header?: any
  timeout?: number
  loadingTitle?: string
  [key: string]: any
}
function requestInterceptor (method: RequestParams['method'], data: any) {
//   console.log('method', method)
//   console.log('data', data)
  // 添加公共参数
  return data
}
export function request (params: RequestParams) {
  const { url, method, data, header, args: { timeout = 6000, loadingTitle = '', toastDuration = 1500 } } = params
  Taro.showLoading({
    title: loadingTitle,
    mask: true
  })
  return new Promise(resolve =>{
    Taro.request({
      data: requestInterceptor(method, data),
      url: baseUrl + url,
      method: method,
      timeout: timeout,
      header: {
        'content-type': 'application/json;charset=UTF-8,text/plain,*/*',
        ...header
      },
      success: (res) => { // 接口调用成功的回调函数
        Taro.hideLoading()
        // console.log('success', res)
        if (res.data.code === 200) { // 具体逻辑可参考接口返回的数据结构
          resolve({...res.data, success: true })
        } else {
          resolve({ message: "失败", success: false })
          // showError(res.data.message.message, toastDuration)
        }
      },
      fail: (res) => {
        Taro.hideLoading()
        console.log('fail', res)
        resolve({ message: res, success: false })
        showError('请求失败', toastDuration)
      },
      complete: (res: any) => { // 接口调用结束的回调函数（调用成功、失败都会执行）
        // console.log('complete', res)
      }
    }).catch(e => {
      Taro.hideLoading()
      console.log('catch err', e)
      resolve({ message: e.errMsg, success: false })
      showError(e.errMsg, toastDuration)
    })
  })
}
function showError (message: string, duration = 1500) {
  Taro.showToast({
    title: message,
    icon: 'none', // 'error' 'success' 'loading' 'none'
    duration: duration
  })
}