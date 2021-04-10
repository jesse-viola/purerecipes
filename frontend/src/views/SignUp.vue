<template>
    <div class="container mx-auto flex h-full justify-center">
        <div class="w-1/3">
            <div class="bg-white shadow-md rounded-md p-8 mt-16">
                <form>
                    <div class="flex flex-col mb-4">
                        <InputText v-model="email" label="Email" />
                    </div>
                    <div class="flex flex-col mb-8">
                        <InputText
                            type="password"
                            v-model="password"
                            label="Password"
                        />
                    </div>
                    <div class="flex justify-center">
                        <Button @click="signUp()">Sign Up</Button>
                    </div>
                </form>
                <p class="flex justify-center text-sm mt-4">
                    Already have an account?
                    <router-link to="/login" class="ml-2 font-bold"
                        >Sign In</router-link
                    >
                </p>
            </div>
        </div>
    </div>
</template>

<script>
import { ref } from 'vue'
import $auth from '@/api/auth'
import InputText from '@/components/InputText.vue'
import Button from '@/components/Button.vue'

export default {
    components: { InputText, Button },
    setup() {
        const email = ref('')
        const password = ref('')

        const signUp = async () => {
            const res = await $auth.login({
                email: email.value,
                password: password.value,
            })
        }

        return {
            email,
            password,
            signUp,
        }
    },
}
</script>

<style></style>
