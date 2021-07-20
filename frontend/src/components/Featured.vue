<template>
    <div class="container flex justify-center">
        <RecipeCard />
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import RecipeCard from '@/components/RecipeCard.vue'

export default defineComponent({
    components: { RecipeCard },
    setup() {},
})
</script>

<style></style>
