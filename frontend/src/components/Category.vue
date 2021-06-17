<template>
    <Button :label="label"></Button>
</template>

<script>
import Button from '@/components/Button.vue'

export default {
    components: { Button },
    props: {
        label: String,
    },
    setup() {},
}
</script>

<style></style>
