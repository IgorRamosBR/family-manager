import { CategoryModel } from "src/models/category";
import { Auth } from "./auth";


const URL = `${process.env.REACT_APP_API_SERVER_URL}/categories`;

async function getCategories(): Promise<CategoryModel[]> {
    return fetch(URL, {
        method: 'GET',
        headers: {
            'Authorization': Auth.getToken(),
        }
    })
        .then(response => response.json())
        .then(response => { return response as CategoryModel[] })
        .then(cat => cat.sort((a, b) => a.priority >= b.priority ? 1 : -1))
}

async function createCategory(category: CategoryModel) {
    const response = await fetch(URL, {
        method: 'POST',
        body: JSON.stringify(category),
        headers: {
            'Authorization': Auth.getToken(),
            'Content-Type': 'application/json'
        }
    })

    if (response.status < 200 || response.status > 299) {
        throw new Error('Error to create the category')
    }
}

async function updateCategoryOrder(categories: CategoryModel[]) {
    const response = await fetch(`${URL}/order`, {
        method: 'PUT',
        body: JSON.stringify(categories),
        headers: {
            'Authorization': Auth.getToken(),
            'Content-Type': 'application/json'
        }
    })

    if (response.status < 200 || response.status > 299) {
        throw new Error('Error to update category order')
    }
}

export const CategoryApi = {
    getCategories,
    createCategory,
    updateCategoryOrder,
}