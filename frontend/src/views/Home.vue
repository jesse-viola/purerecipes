<template>
    <SearchBar></SearchBar>
    <Categories></Categories>
    <Featured></Featured>
</template>

<script>
import SearchBar from '@/components/SearchBar.vue'
import Categories from '@/components/Categories.vue'
import Featured from '@/components/Featured.vue'

export default {
    components: { SearchBar, Categories, Featured },
    setup() {},
}
</script>
