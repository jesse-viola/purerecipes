// TODO cleanup later this is for backend hookup
<template>
    <div class="container mx-auto flex h-full justify-center">
        <div class="w-1/3">
            <div class="bg-white shadow-md rounded-md p-8 mt-16">
                <div class="mb-4">
                   <label
                        class="block text-grey-darker text-sm font-bold mb-2"
                        for="scrapeUrl"
                    >
                        Enter Recipe Website URL
                    </label>
                    <InputText v-model="scrapeUrl" />
                </div>
                    <button
                        class="bg-blue-500 hover:bg-blue-dark text-white font-bold py-2 px-4 rounded"
                        type="button"
                        @click="onSubmitClicked(scrapeUrl)"
                    >
                        Submit
                    </button>
                </div>
            </div>
        </div>
</template>

<script>
import { ref } from 'vue'
import InputText from '@components/InputText.vue'
import { api } from '@api/constants'
import $scrape from '@api/scrape'

export default {
    components: { InputText },
    setup() {
        const scrapeUrl = ref('')
        
        return {
            scrapeUrl
        }
    },
    methods: {
        onSubmitClicked
    },
}

function onSubmitClicked(value) {
    // TODO don't do this hardcode
    const url = api.SCRAPE
    var request = {
        url: value
    }
    $scrape.postScrapeSite(url, request)
    console.log(url + " " + value)
}
</script>

<style></style> 
