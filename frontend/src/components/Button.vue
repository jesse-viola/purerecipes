<template>
    <button
        class="rounded-full px-4 py-2 bg-green-500 hover:bg-green-600 text-white focus:outline-none focus:ring-4 focus:ring-green-500 focus:ring-opacity-50"
    >
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
