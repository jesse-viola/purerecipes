<template>
    <teleport to="body">
        <div v-if="open" class="overlay">
            <div class="modal-content">
                <slot></slot>
            </div>
        </div>
    </teleport>
</template>

<script>
export default {
    props: {
        open: Boolean,
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
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-content {
    background-color: #fff;
    border-radius: 0.25rem;
    padding: 2rem;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}
</style>
