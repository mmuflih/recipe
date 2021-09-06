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
                        type="email"
                        class="form-control form-control-user"
                        id="exampleInputEmail"
                        aria-describedby="emailHelp"
                        placeholder="Masukkan alamat email"
                        v-model="data.email"
                      />
                    </div>
                    <hr/>
                    <a
                      class="btn btn-primary btn-user btn-block"
                      @click="reset"
                    >
                      Pemulihan
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
import { encrypt } from './lib/crypto'
import formatISO from 'date-fns/formatISO'
import * as DateFNS from 'date-fns'

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
    reset () {
      this.loading = true
      HTTP.post('/api/v1/otp/request-reset', this.data).then(res => {
        this.$toast.success('Silahkan check inbox email anda')
        const expired = formatISO(DateFNS.addMinutes(new Date(), 5))
        const query = encrypt(this.data.email) +
          '&account=' + encrypt('false') +
          '&token=' + res.data.data.reset_token +
          '&exp=' + encrypt(expired)
        window.location.href = '#/otp?code=' + query
        this.loading = false
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    }
  }
}
</script>
