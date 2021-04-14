<template>
    <teleport to="body">
        <div v-if="show" class="overlay">
            <slot></slot>
        </div>
    </teleport>
</template>

<script>
export default {
    props: {
        show: Boolean,
    },
}
</script>

<style>
.overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3);
}
</style>
