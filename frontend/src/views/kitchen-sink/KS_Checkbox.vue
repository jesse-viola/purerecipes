<template>
    <h2 class="pb-8">Checkbox</h2>
    <div class="flex flex-col">
        <div>
            <Checkbox id="demo-checkbox" />
            <label for="demo-checkbox">Basic</label>
        </div>
        <div>
            <Checkbox id="demo-checkbox-disabled" disabled />
            <label for="demo-checkbox-disabled">Disabled</label>
        </div>
    </div>
</template>

<script>
import Checkbox from '@/components/Checkbox.vue'

export default {
    components: { Checkbox },
    setup() {},
}
</script>

<style></style>
