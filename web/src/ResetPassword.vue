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
                    <h1 class="h4 text-gray-900 mb-4">Pemulihan Kata Sandi!</h1>
                  </div>
                  <form class="user">
                    <div class="form-group">
                      <input
                        type="password"
                        class="form-control form-control-user"
                        id="password"
                        aria-describedby="emailHelp"
                        placeholder="Password baru"
                        v-model="data.password"
                      />
                    </div>
                    <div class="form-group">
                      <input
                        type="password"
                        class="form-control form-control-user"
                        id="password_confirmed"
                        aria-describedby="emailHelp"
                        placeholder="Masukkan ulang password baru"
                        v-model="data.password_confirmation"
                      />
                    </div>
                    <hr/>
                    <a
                      class="btn btn-primary btn-user btn-block"
                      @click="reset"
                    >
                      Pulihkan
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
import { HTTP } from './lib/httplib'
import { showErrorMessage } from './lib/helpers'
import Loader from './components/Loader.vue'

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
  created () {
    const q = this.$route.query
    if (q && q.code) {
      this.data.token = q.code
    }
  },
  methods: {
    reset () {
      this.loading = true
      HTTP.put('/api/v1/users/reset-password', this.data).then(res => {
        this.$toast.success('Pemulihan kata sandi berhasil')
        window.location.href = '#/login'
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    }
  }
}
</script>
