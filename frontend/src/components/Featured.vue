<template>
    <div class="container flex justify-center gap-12">
        <RecipeCard v-for="recipe in featuredRecipes" :key="recipe.key" :recipe="recipe"/>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onBeforeMount } from 'vue'
import $recipes from '@/api/recipes'
import RecipeCard from '@/components/RecipeCard.vue'
import { RecipeCardType } from '@/types/Recipe'

export default defineComponent({
    components: { RecipeCard },
    setup() {
        const featuredRecipes = ref()

        onBeforeMount(async () => {
            featuredRecipes.value = await $recipes.getFeaturedRecipes()
        })

        return {
            featuredRecipes
        }
    },
})
</script>

<style></style>
