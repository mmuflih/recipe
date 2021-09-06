export function getErrorMessage (eres) {
  if (eres.response.status === 404) {
    return 'Router ' + eres.response.statusText
  }
  return eres.response.data.user_message
}

export function successMessage (res, toastr) {
  if (!(res.data && res.data.data)) {
    toastr.error('Error Server')
    return
  }
  const msg = parseMsg(res.data.data, toastr)
  if (msg !== '') {
    toastr.success(msg)
  }
}

function parseMsg (data, toastr) {
  switch (data) {
    case 'CHK_MAIL':
      return 'Silahkan check inbox email kamu'
    case 'ENF':
      return 'Invalid email'
    default:
      return data
  }
}

export function showErrorMessage (eres, toastr) {
  if (!(eres.response && eres.response.status)) {
    toastr.error(eres)
    return
  }

  if (eres.response.status === 404) {
    const msg = 'Service Router ' + eres.response.statusText
    toastr.error(msg)
    return []
  }
  if (eres.response.status === 500) {
    const msg = eres.response.statusText + ` [${eres.response.data.user_message}]`
    toastr.error(msg)
    return []
  }
  if (eres.response.status === 422) {
    if (eres.response.data) {
      const msg = parseMsg(eres.response.data.user_message, toastr)
      if (msg !== '') {
        toastr.error(msg)
      }
      return
    }
    if (eres.response.data.user_message) {
      const errors = eres.response.data.user_message.split(';')
      for (var i in errors) {
        toastr.error(errors[i])
      }
      if (eres.response.data.errors) {
        return errors
      }
      return []
    }
    const msg = eres.response.statusText
    toastr.error(msg)
    return eres.response.data
  }
  return eres.response.data
}
