<template>
    <div class="flex items-center">
        <input type="checkbox" />
        <label v-if="label" class="ml-2 text-sm">
            {{ label }}
        </label>
    </div>
</template>

<script>
export default {
    props: {
        label: String,
    },
}
</script>

<style></style>
