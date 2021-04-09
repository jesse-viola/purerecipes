import $api from './index'

const login = async (payload) => {
    const url = '/login'
    // const url = 'http://localhost:8000/api/v1/login'
    // await $api.post(url, payload)
    return await $api.get(url)
}

export default {
    login,
}
