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
          <h6 class="m-0 font-weight-bold text-primary">Resep</h6>
          <div class="dropdown no-arrow">
            <a href="#/recipes/add" class="link text-success" data-toggle="tooltip"
              title="Tambah Resep">
              <i class="fa fa-plus-square"></i>
            </a>
          </div>
        </div>
        <div class="card-body">
          <div class="scroll">
            <table class="table table-bordered table-stripped">
              <thead>
                <tr>
                  <th>No</th>
                  <th>Kategori</th>
                  <th>Nama</th>
                  <th>Slug</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(v, k) in data.data" :key="k">
                  <td>{{ k + 1 }}</td>
                  <td>{{ v.category_name }} </td>
                  <td>{{ v.name }} </td>
                  <td>{{ v.slug }} </td>
                  <td>
                    <a class="btn btn-outline-primary" :href="`#/recipes/add?slug=${v.slug}`">Edit</a>
                    <button class="btn btn-outline-danger" @click="deleteItem(v)">Delete</button>
                  </td>
                </tr>
              </tbody>
            </table>
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
  name: 'Recipe',
  components: { Loader },
  data () {
    return {
      loading: false,
      data: {}
    }
  },
  created () {
    this.load()
  },
  methods: {
    load () {
      HTTP.get('/v1/recipes').then(res => {
        this.data = res.data
      })
    },
    deleteItem (v) {
      const c = confirm(`Yakin akan mengkahpus "${v.name}" ?`)
      if (c) {
        HTTP.delete('/v1/recipes/' + v.id).then(res => {
          this.$toast.success('berhasil menghapus data')
          this.load()
        }).catch(eres => {
          this.$toast.error('gagal mengapush data')
        })
      }
    }
  }
}
</script>
