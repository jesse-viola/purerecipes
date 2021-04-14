<template>
    <Overlay :show="open">
        <div class="toast-wrapper">
            <div class="toast-content">
                <slot></slot>
            </div>
        </div>
    </Overlay>
</template>

<script>
import Overlay from '@/components/Overlay.vue'

export default {
    components: { Overlay },
    props: {
        open: Boolean,
    },
}
</script>

<style>
.toast-wrapper {
    width: 200px;
    /* width: 100%;
    height: 100%; */
    display: flex;
    align-items: center;
    justify-content: center;
}
.toast-content {
    background-color: #fff;
    border-radius: 0.25rem;
    padding: 2rem;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}
</style>
