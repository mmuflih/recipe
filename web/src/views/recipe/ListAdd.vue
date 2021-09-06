<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card border-success shadow h-100 py-2">
      <loader v-show="loading"/>
        <div
          class="
            card-header
            py-3
            d-flex
            flex-row
            align-items-center
            justify-content-between
          "
        >
          <h6 class="m-0 font-weight-bold text-primary">Tambah Resep</h6>
        </div>
        <div class="card-body">
          <div class="form-group">
            <label>Nama</label>
            <input type="text" class="form-control" v-model="data.name">
          </div>
          <div class="form-group">
            <label>Kategori</label>
            <select v-model="data.category_id" class="form-control">
              <option v-for="(v,k) in categories" :key="k" :value="v.id">{{v.name}}</option>
            </select>
          </div>
          <div class="form-group">
            <label>Bahan</label>
            <table class="table table-bordered">
              <thead>
                <tr>
                  <th>Bahan</th>
                  <th width="300">Keterangan</th>
                  <th width="200"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(v, k) in data.ingredients" :key="k">
                  <td>{{ v.item.name }}</td>
                  <td>{{ v.notes }}</td>
                  <td>
                    <button class="btn-outline-danger btn" @click="removeItem(k)">Hapus</button>
                  </td>
                </tr>
                <tr>
                  <td>
                    <select v-model="ingredient.item" class="form-control">
                      <option v-for="(v,k) in ingredients" :key="k" :value="v">{{v.name}}</option>
                    </select>
                  </td>
                  <td>
                    <input type="text" class="form-control" v-model="ingredient.notes" placeholder="contoh: 1 sendok makan">
                  </td>
                  <td align="center">
                    <button class="btn btn-outline-success" @click="addIngredient">Tambah bahan</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="card-footer">
          <div class="btn-toolbar" role="toolbar" aria-label="Toolbar with button groups">
            <div class="button-group mr-2">
              <a href="#/recipes" class="btn btn-info btn-icon-split">
                <span class="icon text-white-50">
                  <i class="fa fa-arrow-left"></i>
                </span>
                <span class="text">Batal</span>
              </a>
            </div>
            <div class="button-group">
              <button v-if="data.id" class="btn-success btn" @click="edit">Edit</button>
              <button v-if="!data.id" class="btn-success btn" @click="save">Simpan</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Loader from '../../components/Loader.vue'
import { HTTP } from '../../lib/httplib'
export default {
  name: 'RecipeAdd',
  components: { Loader },
  data () {
    return {
      loading: false,
      data: {
        ingredients: []
      },
      categories: [],
      ingredients: [],
      ingredient: {}
    }
  },
  created () {
    this.categoryList()
    this.ingredientList()
    const query = this.$route.query
    if (query && query.slug) {
      HTTP.get(`/v1/recipes/slug/${query.slug}`).then(res => {
        this.data = res.data.data
      })
    }
  },
  methods: {
    save () {
      this.loading = true
      HTTP.post('/v1/recipes', this.data).then(res => {
        this.loading = false
        this.$toast.success('Berhasil menambahkan resep').goAway()
        this.$router.go({ path: '/ingredients' })
      }).catch(eres => {
        this.loading = false
        this.$toast.error(eres.response.data.user_message).goAway()
      })
    },
    edit () {
      this.loading = true
      HTTP.put('/v1/recipes/' + this.data.id, this.data).then(res => {
        this.loading = false
        this.$toast.success('Berhasil mengubah kategori').goAway()
        this.$router.go({ path: '/categories' })
      }).catch(eres => {
        this.loading = false
        this.$toast.error(eres.response.data.user_message).goAway()
      })
    },
    categoryList () {
      HTTP.get('/v1/categories').then(res => {
        this.categories = res.data.data
      })
    },
    ingredientList () {
      HTTP.get('/v1/ingredients').then(res => {
        this.ingredients = res.data.data
      })
    },
    addIngredient () {
      this.data.ingredients.push({
        ingredient_id: this.ingredient.item.id,
        notes: this.ingredient.notes,
        item: this.ingredient.item
      })
      this.ingredient = {}
    },
    removeItem (k) {
      this.data.ingredients.splice(k, 1)
    }
  }
}
</script>
