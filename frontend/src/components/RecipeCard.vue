<template>
    <router-link to="/recipe">
        <div
            class="flex flex-col shadow border-gray-400 transition duration-150 ease-in-out hover:shadow-lg cursor-pointer"
        >
            <img src="https://via.placeholder.com/200x150" />
            <div class="flex flex-col p-3 bg-white">
                <div>{{ recipe.title }}</div>
                <div>{{ recipe.rating }}</div>
                <div class="flex justify-end">{{ cookTime }}</div>
            </div>
        </div>
    </router-link>
</template>

<script lang="ts">
import { RecipeCardType } from '@/types/Recipe'
import { computed, defineComponent, PropType } from 'vue'

export default defineComponent({
    props: {
        recipe: {
            type: Object as PropType<RecipeCardType>
        }
    },
    setup(props) {
        const cookTime = computed(() => {
            return `${props.recipe.time}m`
        })

        return {
            cookTime
        }
    },
})
</script>

<style></style>
