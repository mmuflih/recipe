<template>
<div class="container">
  <!-- Outer Row -->
  <div class="row justify-content-center">
    <div class="col-xl-10 col-lg-12 col-md-9">
      <div class="card o-hidden border-0 shadow-lg my-5">
        <div class="card-body p-0">
          <!-- Nested Row within Card Body -->
          <loader v-show="loading" />
          <div class="row">
            <div class="col-lg-6 d-none d-lg-block bg-login-image"></div>
            <div class="col-lg-6">
              <div class="p-5">
                <div class="text-center">
                  <h1 class="h4 text-gray-900 mb-4">Selamat Datang Kembali!</h1>
                </div>
                <form class="user">
                  <div class="form-group">
                    <input type="email" class="form-control form-control-user" id="exampleInputEmail" aria-describedby="emailHelp" placeholder="Masukkan alamat email" v-model="data.email" />
                  </div>
                  <div class="form-group">
                    <input type="password" class="form-control form-control-user" id="exampleInputPassword" placeholder="Masukkan password" v-model="data.password" />
                  </div>
                  <hr />
                  <a class="btn btn-success btn-user btn-block" @click="login">
                    Masuk
                  </a>
                  <div class="col-md-12" style="text-align: center; padding: 10px 0">atau</div>
                  <a href="#/register" class="btn btn-primary btn-user btn-block">
                    Daftar
                  </a>
                  <div class="col-md-12" style="text-align: center; padding: 10px 0">
                    <a href="#/reset-password">Pulihkan kata sandi</a>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>
<script>
import {
  HTTP
} from '../src/lib/httplib'
import {
  setCookie,
  showErrorMessage
} from '../src/lib/helpers'
import Loader from './components/Loader.vue'
export default {
  data() {
    return {
      loading: false,
      data: {}
    }
  },
  components: {
    Loader
  },
  methods: {
    login() {
      this.loading = true
      HTTP.post('/api/v1/users/login', this.data).then(res => {
        setCookie(res.data)
        setTimeout(function() {
          window.location.href = '#/dashboard'
        }, 1000)
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    }
  }
}
</script>
