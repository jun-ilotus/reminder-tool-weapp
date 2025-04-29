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




// 定义数字与汉字的映射关系
const numMap = [
  '零', '一', '二', '三', '四', '五', '六', '七', '八', '九',
];
const unitMap = ['', '十', '百', '千', '万', '亿', '兆'];
const chineseNumMap: { [key: string]: number } = {
  零: 0,
  一: 1,
  二: 2,
  三: 3,
  四: 4,
  五: 5,
  六: 6,
  七: 7,
  八: 8,
  九: 9,
  十: 10,
  百: 100,
  千: 1000,
  万: 10000,
  亿: 100000000,
  兆: 1000000000000,
};

// 数字转汉字
export function numberToChinese(num: number): string {
  if (num === 0) return '零';
  let result = '';
  let index = 0;
  while (num > 0) {
    const digit = num % 10; // 获取当前位的数字
    if (digit !== 0) {
      result = numMap[digit] + unitMap[index] + result;
    } else if (result[0] !== '零') {
      result = '零' + result;
    }
    num = Math.floor(num / 10);
    index++;
    if (index === 4 && num > 0) {
      result = '万' + result;
      index = 0;
    } else if (index === 8 && num > 0) {
      result = '亿' + result;
      index = 0;
    } else if (index === 12 && num > 0) {
      result = '兆' + result;
      index = 0;
    }
  }
  return result;
}

// 汉字转数字
export function chineseToNumber(chinese: string): number {
  let result = 0;
  let temp = 0;
  let unit = 1;
  for (let i = chinese.length - 1; i >= 0; i--) {
    const char = chinese[i];
    const value = chineseNumMap[char];
    if (value === undefined) {
      throw new Error('Invalid Chinese number');
    }
    if (value < 10) {
      temp = value * unit + temp;
    } else {
      unit = value;
      result = temp * unit + result;
      temp = 0;
    }
  }
  return result + temp;
}