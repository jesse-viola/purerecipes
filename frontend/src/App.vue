<template>
    <Header />
    <router-view />
</template>

<script>
import Header from '@components/Header.vue'

const APP_VERSION = import.meta.env.VITE_APP_VERSION

export default {
    components: { Header },
    setup() {
        console.log('VERSION: ', APP_VERSION)
    },
}
</script>

<style></style>
