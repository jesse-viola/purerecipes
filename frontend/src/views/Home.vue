<template>
    <h1>Home</h1>
    <div class="recipe" v-for="recipe in recipes" :key="recipe.name">
        <p v-html="recipe.recipe"></p>
    </div>
</template>

<script>
import { ref, computed, onMounted, onBeforeMount } from 'vue'
import $recipes from '../api/recipes'

export default {
    components: {},
    setup() {
        const recipes = ref([])

        onBeforeMount(async () => {
            recipes.value = await $recipes.getRecipes()
        })

        return {
            recipes,
        }
    },
}
</script>

<style>
.recipe {
    text-align: left;
}
</style>
