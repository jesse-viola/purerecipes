<template>
    <div class="container flex justify-center">
        <div class="flex justify-between content-center p-8">
            <Category
                v-for="item in items"
                :key="item.id"
                :label="item.label"
            ></Category>
        </div>
    </div>
</template>

<script>
import Category from '@/components/Category.vue'

export default {
    components: { Category },
    setup() {
        const items = [
            {
                id: 1,
                label: 'Brunch',
            },
            {
                id: 2,
                label: 'Vegetarian',
            },
            {
                id: 3,
                label: 'Fishy Fish',
            },
            {
                id: 4,
                label: 'Dessert',
            },
            {
                id: 5,
                label: 'Gordon',
            },
        ]

        return {
            items,
        }
    },
}
</script>

<style></style>
