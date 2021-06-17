<template>
    <teleport to="body">
        <div
            v-if="show"
            class="fixed top-0 left-0 w-full h-full"
            :class="background"
        >
            <slot></slot>
        </div>
    </teleport>
</template>

<script>
import { computed } from 'vue'

export default {
    props: {
        show: Boolean,
        transparent: Boolean,
    },
    setup(props) {
        const background = computed(() => {
            if (!props.transparent) {
                return ['bg-black', 'bg-opacity-30']
            }
        })

        return {
            background,
        }
    },
}
</script>
