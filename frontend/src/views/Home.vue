<template>
    <SearchBar></SearchBar>
    <Categories></Categories>
    <Featured></Featured>
</template>

<script>
import Categories from '@/components/Categories.vue'

export default {
    components: { Categories },
    setup() {},
}
</script>

<style>
.recipe {
    text-align: left;
}
</style>
