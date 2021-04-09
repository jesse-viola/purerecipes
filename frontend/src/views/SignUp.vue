<template>
    <div class="container mx-auto flex h-full justify-center">
        <div class="w-1/3">
            <div class="bg-white shadow-md rounded-md p-8 mt-16">
                <div class="mb-4">
                    <label
                        class="block text-grey-darker text-sm font-bold mb-2"
                        for="username"
                    >
                        Username
                    </label>
                    <InputText v-model="username" />
                </div>
                <div class="mb-6">
                    <label
                        class="block text-grey-darker text-sm font-bold mb-2"
                        for="password"
                    >
                        Password
                    </label>
                    <InputText type="password" v-model="password" />
                </div>
                <div class="flex items-center justify-between">
                    <button
                        @click="signIn()"
                        class="bg-blue-500 hover:bg-blue-dark text-white font-bold py-2 px-4 rounded"
                        type="button"
                    >
                        Sign In
                    </button>
                    <a
                        class="inline-block align-baseline font-bold text-sm text-blue hover:text-blue-darker"
                        href="#"
                    >
                        Forgot Password?
                    </a>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref } from 'vue'
import $auth from '@/api/auth'
import InputText from '@/components/InputText.vue'

export default {
    components: { InputText },
    setup() {
        const username = ref('')
        const password = ref('')

        const signIn = async () => {
            const res = await $auth.login({
                username: username.value,
                password: password.value,
            })

            alert(res.data)
        }

        return {
            username,
            password,
            signIn,
        }
    },
}
</script>

<style></style>
