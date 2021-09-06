import { createStore } from 'vuex'
import { HTTPAUTH } from '../lib/httplib'

export default createStore({
  state: {
    qrData: {},
    loading: false,
    auth: {},
    isDelete: {},
    close: {},
    qrExpired: false,
    message: {}
  },
  mutations: {
    loading (state, value) {
      state.loading = value
    },
    qrExpired (state, value) {
      state.qrExpired = value
    },
    SOCKET_message (state, data) {
      state.message = data
      state.loading = false
      console.log('message', data, state.loading)
    },
    SOCKET_qr: (state, data) => {
      state.qrData = data
      state.loading = false
      console.log('qr', data, state.loading)
    },
    SOCKET_authenticated: (state, data) => {
      state.auth = data
      state.loading = false
      HTTPAUTH.put(`/api/v1/devices/${data.id}/auth`, data)
      console.log('auth', data, state.loading)
    },
    SOCKET_isdelete: (state, data) => {
      state.isDelete = data
      state.loading = false
      console.log('isdelete', data, state.loading)
    },
    SOCKET_close: (state, data) => {
      state.close = data
      state.loading = false
      console.log('close', data, state.loading)
    }
  },
  actions: {
  },
  modules: {
  }
})
