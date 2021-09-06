import {
  createRouter,
  createWebHashHistory
} from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Admin',
    component: () =>
      import('../Admin.vue'),
    meta: {
      requireAuth: true
    },
    children: [
      {
        path: '/ingredients',
        name: 'Ingredient',
        component: () =>
          import('../views/ingredient/Index.vue'),
        children: [
          {
            path: '',
            name: 'IngredientList',
            component: () =>
              import('../views/ingredient/List.vue')
          },
          {
            path: 'add',
            name: 'IngredientAdd',
            component: () =>
              import('../views/ingredient/ListAdd.vue')
          }
        ]
      },
      {
        path: '/categories',
        name: 'Category',
        component: () =>
          import('../views/category/Index.vue'),
        children: [
          {
            path: '',
            name: 'CategoryList',
            component: () =>
              import('../views/category/List.vue')
          },
          {
            path: 'add',
            name: 'CategoryAdd',
            component: () =>
              import('../views/category/ListAdd.vue')
          }
        ]
      },
      {
        path: '/recipes',
        name: 'Recipe',
        component: () =>
          import('../views/recipe/Index.vue'),
        children: [
          {
            path: '',
            name: 'RecipeList',
            component: () =>
              import('../views/recipe/List.vue')
          },
          {
            path: 'add',
            name: 'RecipeAdd',
            component: () =>
              import('../views/recipe/ListAdd.vue')
          }
        ]
      }
    ]
  }
]

var router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
