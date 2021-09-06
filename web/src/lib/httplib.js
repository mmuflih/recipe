import axios from 'axios'

export const HTTP = axios.create({
  baseURL: process.env.VUE_APP_API_HOST,
  headers: {
    'Content-Type': 'application/json'
  }
})

export function objectToQuery (data) {
  return Object.keys(data).map(function (k) {
    return encodeURIComponent(k) + '=' + encodeURIComponent(data[k])
  }).join('&')
}
