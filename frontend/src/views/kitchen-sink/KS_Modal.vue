<template>
    <h2 class="pb-8">Modal</h2>
    <Button label="Show Modal" @click="modalOpen = true" />
    <Modal :open="modalOpen">
        <Button label="Close" @click="modalOpen = false" />
    </Modal>
</template>

<script>
import { ref } from 'vue'
import Button from '@components/Button.vue'
import Modal from '@components/Modal.vue'

export default {
    components: { Button, Modal },
    setup() {
        const modalOpen = ref(false)

        return {
            modalOpen,
        }
    },
}
</script>

<style></style>
