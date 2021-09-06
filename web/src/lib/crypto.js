import CryptoJS from 'crypto-js'

export function encrypt (text) {
  const cipher = CryptoJS.DES.encrypt(text, '****')
  return encodeURIComponent(cipher.toString())
}

export function decrypt (text) {
  text = decodeURIComponent(text)
  const cipher = CryptoJS.DES.decrypt(text, '****')
  return CryptoJS.enc.Utf8.stringify(cipher).toString()
}
