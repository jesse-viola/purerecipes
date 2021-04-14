<template>
    <h2 class="pb-8">Toast</h2>
    <Button label="Show Modal" @click="modalOpen = true" />
    <Toast :open="modalOpen">
        <Button label="Close" @click="modalOpen = false" />
    </Toast>
</template>

<script>
import { ref } from 'vue'
import Button from '@/components/Button.vue'
import Toast from '@/components/Toast.vue'

export default {
    components: { Button, Toast },
    setup() {
        const modalOpen = ref(false)

        return {
            modalOpen,
        }
    },
}
</script>

<style></style>
