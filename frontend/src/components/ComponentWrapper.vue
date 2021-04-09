<template>
    <div class="py-8">
        <h2 class="pb-8">{{ title }}</h2>
        <slot></slot>
    </div>
</template>

<script>
export default {
    props: {
        title: String,
    },
    setup() {},
}
</script>

<style></style>
