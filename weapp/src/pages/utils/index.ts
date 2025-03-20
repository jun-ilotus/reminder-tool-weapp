// 日期格式化
export function dateFormat (value: number|string|Date = Date.now(), format = 'YYYY-MM-DD HH:mm:ss'): string {
    if (typeof value === 'number' || typeof value === 'string') {
      var date = new Date(value)
    } else {
      var date = value
    }
    let showTime = format
    if (showTime.includes('SSS')) {
      const S = String(date.getMilliseconds())
      showTime = showTime.replace('SSS', S.padStart(3, '0'))
    }
    if (showTime.includes('YY')) {
      const Y = String(date.getFullYear())
      showTime = showTime.includes('YYYY') ? showTime.replace('YYYY', Y) : showTime.replace('YY', Y.slice(2, 4))
    }
    if (showTime.includes('M')) {
      const M = String(date.getMonth() + 1)
      showTime = showTime.includes('MM') ? showTime.replace('MM', M.padStart(2, '0')) : showTime.replace('M', M)
    }
    if (showTime.includes('D')) {
      const D = String(date.getDate())
      showTime = showTime.includes('DD') ? showTime.replace('DD', D.padStart(2, '0')) : showTime.replace('D', D)
    }
    if (showTime.includes('H')) {
      const H = String(date.getHours())
      showTime = showTime.includes('HH') ? showTime.replace('HH', H.padStart(2, '0')) : showTime.replace('H', H)
    }
    if (showTime.includes('m')) {
      const m = String(date.getMinutes())
      showTime = showTime.includes('mm') ? showTime.replace('mm', m.padStart(2, '0')) : showTime.replace('m', m)
    }
    if (showTime.includes('s')) {
      const s = String(date.getSeconds())
      showTime = showTime.includes('ss') ? showTime.replace('ss', s.padStart(2, '0')) : showTime.replace('s', s)
    }
    return showTime
  }