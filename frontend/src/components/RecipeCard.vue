<template>
    <router-link to="/recipe">
        <div
            class="flex flex-col shadow border-gray-400 transition duration-150 ease-in-out hover:shadow-lg cursor-pointer"
        >
            <img src="https://via.placeholder.com/200x150" />
            <div class="flex flex-col p-3 bg-white">
                <div>Title</div>
                <div>Rating</div>
                <div class="flex justify-end">Time</div>
            </div>
        </div>
    </router-link>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    setup() {},
})
</script>

<style></style>
