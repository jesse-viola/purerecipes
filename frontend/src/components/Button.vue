<template>
    <button class="px-4 py-2 bg-green-500 hover:bg-green-600 text-white">
        <slot>
            <span>{{ label }}</span>
        </slot>
    </button>
</template>

<script>
export default {
    props: {
        label: {
            type: String,
        },
    },
}
</script>

<style></style>
