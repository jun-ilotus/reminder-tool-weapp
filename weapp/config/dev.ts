import type { UserConfigExport } from "@tarojs/cli";
export default {
   logger: {
    quiet: false,
    stats: true
  },
  env: { // 设置环境变量
    NODE_ENV: '"development"', // JSON.stringify('development')
    TARO_APP_PROXY: '"https://remind.jilotus.cn"'
  },
  mini: {},
  h5: {}
} satisfies UserConfigExport<'webpack5'>
