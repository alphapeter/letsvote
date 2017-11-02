export const API = {
  getPolls () {
    return this.get('polls')
  },
  get (url) {
    return fetch('api/polls', {
      credentials: 'include',
      method: 'GET'
    }).then((response) => {
      return response.json()
    })
  },
  post (url, payload) {
    return fetch('polls', {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify(payload)
    }).then((response) => {
      return response.json()
    })
  },
  put (url, payload) {
    return fetch('polls', {
      credentials: 'include',
      method: 'PUT',
      body: JSON.stringify(payload)
    }).then((response) => {
      return response.json()
    })
  }
}
