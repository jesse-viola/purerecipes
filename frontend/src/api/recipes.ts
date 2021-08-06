import $api from './index'
import { RecipeCardType } from '../types/Recipe'

const getFeaturedRecipes = async () => {
    // const url = '/recipes/featured'

    const response: RecipeCardType[] = [
        {
            id: 1,
            title: "Duck Souffle",
            rating: 3.5,
            time: 120
        },
        {
            id: 2,
            title: "Fart Sandwich",
            rating: 0.5,
            time: 2
        },
        {
            id: 3,
            title: "Pad Thai",
            rating: 3,
            time: 20
        }
    ]

    return await response
}

export default {
    getFeaturedRecipes,
}
