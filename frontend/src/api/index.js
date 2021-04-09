import axios from 'axios'

const REQUEST_TIMEOUT = 10000

const $api = axios.create({
    baseURL: '/api/v1/',
    headers: { Accept: 'application/json' },
    timeout: REQUEST_TIMEOUT,
})

export default {
    get(url) {
        return $api.get(url)
    },
    post(url, payload) {
        return $api.post(url, payload)
    },
}
