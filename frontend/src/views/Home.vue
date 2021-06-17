<template>
    <SearchBar></SearchBar>
    <Categories></Categories>
    <Featured></Featured>
</template>

<script>
import { ref, computed, onMounted, onBeforeMount } from 'vue'
import $recipes from '../api/recipes'
import SearchBar from '@/components/SearchBar.vue'

export default {
    components: { SearchBar },
    setup() {},
}
</script>

<style>
.recipe {
    text-align: left;
}
</style>
