<template>
    <div class="flex">
        <div class="h-screen bg-white w-40">
            <nav>
                <router-link
                    v-for="link in links"
                    :key="link.path"
                    :to="`/kitchen-sink/${link.path}`"
                    class="block px-4 py-2 hover:bg-gray-200"
                    >{{ link.label }}</router-link
                >
            </nav>
        </div>
        <div class="p-8">
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
export default {
    setup() {
        const links = [
            { path: 'input-text', label: 'InputText' },
            { path: 'checkbox', label: 'Checkbox' },
            { path: 'button', label: 'Button' },
            { path: 'modal', label: 'Modal' },
            { path: 'toast', label: 'Toast' },
        ]

        return {
            links,
        }
    },
}
</script>

<style></style>
