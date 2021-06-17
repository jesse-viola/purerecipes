<template>
    <SearchBar></SearchBar>
    <Categories></Categories>
    <Featured></Featured>
</template>

<script>
import { ref, computed, onMounted, onBeforeMount } from 'vue'
import $recipes from '../api/recipes'

export default {
    components: {},
    setup() {},
}
</script>

<style>
.recipe {
    text-align: left;
}
</style>
