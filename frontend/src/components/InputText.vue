<template>
    <label v-if="label" class="block text-grey-darker text-sm font-bold mb-2">
        {{ label }}
    </label>
    <input
        class="full-width border border-black px-2 leading-8"
        :value="modelValue"
        @input="onInput"
    />
</template>

<script>
export default {
    emits: ['update:modelValue'],
    props: {
        modelValue: null,
        label: String,
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
