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
                    <h1 class="h4 text-gray-900 mb-4">Verifikasi Kode!</h1>
                    <div>
                      Silahkan check inbox di alamat email <strong>{{email}}</strong>
                    </div>
                    <br/>
                  </div>
                  <form class="user">
                    <div class="form-group otp-container">
                      <input
                        type="text"
                        class="form-control form-control-otp"
                        id="otp1"
                        aria-describedby="otp1"
                        placeholder="-"
                        ref="otp1"
                        @keyup="keyup1"
                        maxlength="1"
                        v-model="data.otp1"
                      />
                      <input
                        type="text"
                        class="form-control form-control-otp"
                        id="otp2"
                        aria-describedby="otp2"
                        placeholder="-"
                        ref="otp2"
                        @keyup="keyup2"
                        maxlength="1"
                        v-model="data.otp2"
                      />
                      <input
                        type="text"
                        class="form-control form-control-otp"
                        id="otp3"
                        aria-describedby="otp3"
                        placeholder="-"
                        ref="otp3"
                        @keyup="keyup3"
                        maxlength="1"
                        v-model="data.otp3"
                      />
                      <input
                        type="text"
                        class="form-control form-control-otp"
                        id="otp4"
                        aria-describedby="otp4"
                        placeholder="-"
                        ref="otp4"
                        maxlength="1"
                        v-model="data.otp4"
                      />
                    </div>
                    <div class="text-center">
                      <label v-show="!requestButton">Kode kadaluarsa dalam
                        <strong>{{timer}}</strong>
                      </label>
                      <label v-show="requestButton">Kode kadaluarsa,
                        <strong><a class="link" @click="requestOTP">Meminta ulang kode</a></strong>
                      </label>
                    </div>
                    <hr/>
                    <a
                      class="btn btn-primary btn-user btn-block"
                      @click="verifyOTP"
                    >
                      Verifikasi
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
import { decrypt, encrypt } from './lib/crypto'
import { showErrorMessage } from './lib/helpers'
import Loader from './components/Loader.vue'
import differenceInSeconds from 'date-fns/differenceInSeconds'
import parseISO from 'date-fns/parseISO'
import * as DateFNS from 'date-fns'
import formatISO from 'date-fns/formatISO'

export default {
  data () {
    return {
      loading: false,
      data: {},
      email: '',
      encEmail: '',
      account: '',
      encAccount: '',
      timer: '',
      requestButton: false,
      token: ''
    }
  },
  components: {
    Loader
  },
  created () {
    const q = this.$route.query
    if (q && q.code) {
      this.email = decrypt(q.code)
      this.encEmail = q.code
    }
    if (q && q.account) {
      this.account = decrypt(q.account)
      this.encAccount = q.account
    }
    if (q && q.token) {
      this.token = q.token
    }
    if (q && q.exp) {
      const exp = parseISO(decrypt(q.exp))
      const expIn = differenceInSeconds(exp, new Date())
      if (expIn <= 0) {
        this.requestButton = true
        return
      }
      window.setInterval(() => { this.counterTime(exp) }, 1000)
    }
  },
  methods: {
    counterTime (exp) {
      const expIn = differenceInSeconds(exp, new Date())
      if (expIn <= 0) {
        this.requestButton = true
        return
      }
      const min = Math.floor(expIn / 60)
      const sec = expIn % 60
      var sMin = ''
      var sSec = ''
      if (min < 10) {
        sMin = '0' + min
      } else {
        sMin = min
      }
      if (sec < 10) {
        sSec = '0' + sec
      } else {
        sSec = sec
      }
      this.timer = `${sMin}:${sSec}`
    },
    verifyOTP () {
      this.loading = true
      const otp = this.data.otp1 + this.data.otp2 + this.data.otp3 + this.data.otp4
      var data = {
        email: this.email,
        code: otp
      }
      var endpoint = '/api/v1/otp/verify'
      if (this.account === 'true') {
        endpoint = '/api/v1/otp/verify-account'
      }

      HTTP.post(endpoint, data).then(res => {
        const token = this.token
        const account = this.account
        setTimeout(function () {
          let redirectUrl = '#/do-reset-password?code=' + token
          if (account === 'true') {
            redirectUrl = '#/dashboard'
          }
          window.location.href = redirectUrl
        }, 1000)
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    },
    requestOTP () {
      this.loading = true
      HTTP.post('/api/v1/otp/request', { email: this.email }).then(res => {
        this.$toast.success('Silahkan cek inbox email kamu')
        const expired = formatISO(DateFNS.addMinutes(new Date(), 5))
        const exp = encrypt(expired)
        const query = 'code=' + this.encEmail +
          '&account=' + this.encAccount +
          '&exp=' + exp
        console.log(query)
        window.location.href = '#/otp?' + query
        this.$router.go()
        this.loading = false
      }).catch(eres => {
        showErrorMessage(eres, this.$toast)
        this.loading = false
      })
    },
    keyup1 () {
      this.$refs.otp2.focus()
    },
    keyup2 () {
      this.$refs.otp3.focus()
    },
    keyup3 () {
      this.$refs.otp4.focus()
    }
  }
}
</script>
