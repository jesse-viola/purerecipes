import $api from './index'

const getRecipes = async () => {
    const url = 'http://taco-randomizer.herokuapp.com/random/'
    const { data } = await $api.get(url)
    const arr = []
    for (const key in data) {
        const value = data[key]
        arr.push({
            name: value.name,
            recipe: value.recipe.replace(/\n/g, '<br/>'),
        })
    }
    console.log(arr)
    return arr
}

export default {
    getRecipes,
}
