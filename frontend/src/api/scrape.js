import $api from './index'

const postScrapeSite = async (scrapeApi, urlToScrape) => {
    const response = await $api.post(scrapeApi, urlToScrape)
    console.log(response)
}

export default {
    postScrapeSite
}