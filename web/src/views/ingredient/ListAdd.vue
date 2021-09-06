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
          <h6 class="m-0 font-weight-bold text-primary">Tambah Bahan</h6>
        </div>
        <div class="card-body">
          <div class="form-group">
            <label>Nama</label>
            <input type="text" class="form-control" v-model="data.name">
          </div>
        </div>
        <div class="card-footer">
          <div class="btn-toolbar" role="toolbar" aria-label="Toolbar with button groups">
            <div class="button-group mr-2">
              <a href="#/ingredients" class="btn btn-info btn-icon-split">
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
  name: 'IngredientAdd',
  components: { Loader },
  data () {
    return {
      loading: false,
      data: {}
    }
  },
  created () {
    const query = this.$route.query
    if (query && query.slug) {
      HTTP.get(`/v1/ingredients/slug/${query.slug}`).then(res => {
        this.data = res.data.data
      })
    }
  },
  methods: {
    save () {
      this.loading = true
      HTTP.post('/v1/ingredients', this.data).then(res => {
        this.loading = false
        this.$toast.success('Berhasil menambahkan ingredient').goAway()
        this.$router.go({ path: '/ingredients' })
      }).catch(eres => {
        this.loading = false
        this.$toast.error(eres.response.data.user_message).goAway()
      })
    },
    edit () {
      this.loading = true
      HTTP.put('/v1/ingredients/' + this.data.id, this.data).then(res => {
        this.loading = false
        this.$toast.success('Berhasil mengubah kategori').goAway()
        this.$router.go({ path: '/categories' })
      }).catch(eres => {
        this.loading = false
        this.$toast.error(eres.response.data.user_message).goAway()
      })
    }
  }
}
</script>
