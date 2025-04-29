<template>
    <nut-card
        :img-url=url
        title="名称"
        shop-desc="简介"
        delivery="到点开奖"
        is-need-price="false"
    ></nut-card>

    <nut-card
        :img-url=url
        title="名称"
        shop-desc="简介"
        delivery="到点开奖"
        is-need-price="false"
    ></nut-card>

    <nut-card
        :img-url=url
        title="名称"
        shop-desc="简介"
        delivery="到点开奖"
        is-need-price="false"
    ></nut-card>
    <lotteryTabbar></lotteryTabbar>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import lotteryTabbar from "src/components/lotteryTabbar.vue";

const url = ref("http://image.jilotus.cn/lottery.png")

</script>