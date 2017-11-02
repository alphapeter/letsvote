let idPrefix = Math.random().toString(36).substring(2, 10) + '_'
let id = 1
export const Rpc = {
  getRoots () {
    return this.call('df', [])
  },
  call (method, params) {
    return fetch('api', {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify({
        'jsonrpc': '2.0',
        'method': method,
        'params': params,
        'id': idPrefix + id++
      })
    }).then((response) => {
      return response.json()
    })
  }
}
