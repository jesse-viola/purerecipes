<template>
    <SearchBar></SearchBar>
    <Categories></Categories>
    <Featured></Featured>
</template>

<script>
import SearchBar from '@/components/SearchBar.vue'
import Categories from '@/components/Categories.vue'

export default {
    components: { SearchBar, Categories },
    setup() {},
}
</script>

<style>
.recipe {
    text-align: left;
}
</style>
