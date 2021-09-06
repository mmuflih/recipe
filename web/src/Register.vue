<template>
  <div class="container">
    <!-- Outer Row -->
    <div class="row justify-content-center">
      <div class="col-xl-10 col-lg-12 col-md-9">
        <div class="card o-hidden border-0 shadow-lg my-5">
          <div class="card-body p-0">
            <!-- Nested Row within Card Body -->
            <loader v-show="loading"/>
            <div class="row">
              <div class="col-lg-6 d-none d-lg-block bg-login-image"></div>
              <div class="col-lg-6">
                <div class="p-5">
                  <div class="text-center">
                    <h1 class="h4 text-gray-900 mb-4">Membuat akun baru!</h1>
                  </div>
                  <form class="user">
                    <div class="form-group">
                      <input
                        type="text"
                        class="form-control form-control-user"
                        id="fullName"
                        aria-describedby="fullName"
                        placeholder="Nama lengkap"
                        v-model="data.full_name"
                      />
                    </div>
                    <div class="form-group">
                      <input
                        type="email"
                        class="form-control form-control-user"
                        id="inputEmail"
                        aria-describedby="inputEmail"
                        placeholder="Alamat Email"
                        v-model="data.email"
                      />
                    </div>
                    <div class="form-group">
                      <input
                        type="password"
                        class="form-control form-control-user"
                        id="exampleInputPassword"
                        placeholder="Password"
                        v-model="data.password"
                      />
                    </div>
                    <div class="form-group">
                      <input
                        type="password"
                        class="form-control form-control-user"
                        id="exampleInputPassword2"
                        placeholder="Masukkal kembali password"
                        v-model="data.password_confirmation"
                      />
                    </div>
                    <hr>
                    <a
                      class="btn btn-success btn-user btn-block"
                      @click="register"
                    >
                      Mendaftar
                    </a>
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
import { HTTPAUTH } from './lib/httplib'
import { showErrorMessage } from './lib/helpers'
import Loader from './components/Loader.vue'
import { encrypt } from './lib/crypto'
import * as DateFNS from 'date-fns'
import formatISO from 'date-fns/formatISO'

export default {
  data () {
    return {
      loading: false,
      data: {}
    }
  },
  components: {
    Loader
  },
  methods: {
    async register () {
      this.loading = true
      await HTTPAUTH.post('/api/v1/users/register', this.data).then(res => {
        const expired = formatISO(DateFNS.addMinutes(new Date(), 5))
        const query = encrypt(this.data.email) +
          '&account=' + encrypt('true') +
          '&exp=' + encrypt(expired)
        this.$toast.success('Silahkan cek inbox email kamu')
        window.location.href = '#/otp?code=' + query
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    }
  }
}
</script>
