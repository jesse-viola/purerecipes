<template>
    <div class="container mx-auto flex h-full justify-center">
        <div class="w-1/3">
            <div class="bg-white shadow-md rounded-md p-8 mt-16">
                <form>
                    <div class="flex flex-col mb-4">
                        <InputText v-model="email" label="Email" />
                    </div>
                    <div class="flex flex-col mb-4">
                        <InputText
                            type="password"
                            v-model="password"
                            label="Password"
                        />
                    </div>
                    <div class="flex items-center justify-between mb-8">
                        <Checkbox label="Remember me" />
                        <router-link to="/forgot" class="ml-2 font-bold text-sm"
                            >Forgot?</router-link
                        >
                    </div>
                    <div class="flex justify-center">
                        <Button @click="logIn()">Log In</Button>
                    </div>
                </form>
                <p class="flex justify-center text-sm mt-4">
                    Are you new here?
                    <router-link to="/signUp" class="ml-2 font-bold"
                        >Sign Up</router-link
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
import Checkbox from '@/components/Checkbox.vue'

export default {
    components: { InputText, Button, Checkbox },
    setup() {
        const email = ref('')
        const password = ref('')

        const logIn = async () => {
            const res = await $auth.login({
                email: email.value,
                password: password.value,
            })

            alert(res.data)
        }

        return {
            email,
            password,
            logIn,
        }
    },
}
</script>

<style></style>
