<template>
    <input
        class="border border-black px-2 leading-8"
        :value="modelValue"
        @input="onInput"
    />
</template>

<script>
export default {
    emits: ['update:modelValue'],
    props: {
        modelValue: null,
    },
    setup(_, { emit }) {
        const onInput = (e) => {
            emit('update:modelValue', e.target.value)
        }

        return {
            onInput,
        }
    },
}
</script>

<style></style>
