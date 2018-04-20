<template>
  <div v-if="!user">
    <A href="/">Login</A>
  </div>

  <div v-else-if="users === null">
    Loading
  </div>

  <div v-else>
    <table>
      <thead>
      <tr><td>user</td><td>email</td><td>Is administrator</td></tr>
      </thead>
      <tr v-for="user in users">
        <td>{{user.name}}</td>
        <td>{{user.email}}</td>
        <td><input type="checkbox"
                   v-model="user.is_admin"
                   @click="toggleAdmin(user)"
                   :disabled="disabled(user)">
        </td>
      </tr>
    </table>
    <a href="/">Return</a>
  </div>

</template>

<script>
  import { API } from '../api.js'
  export default {
    name: 'adminApp',
    components: {
    },
    data () {
      return {
        users: null,
        loading: true,
        user: null
      }
    },
    computed: {
      noAdmins () {
        return this.users && !this.users.some(u => u.is_admin)
      }
    },
    created () {
      try {
        var cookie = this.$cookie.get('lets_vote.authenticated')
        var u = decodeURIComponent(cookie)
        this.user = JSON.parse(u)
      } catch (e) {
        this.user = null
        return
      }
      API.get('/api/users')
        .then((users) => {
          this.users = users.sort((a, b) => {
            return a.name.localeCompare(b.name)
          })
        })
    },
    destroyed () {

    },
    methods: {
      toggleAdmin (user) {
        API.patch('/api/users/' + user.id, {is_admin: !user.is_admin})
          .then(() => console.log('ok'))
        return false
      },
      disabled (user) {
        if (this.noAdmins && user.id !== this.user.id) {
          return true
        }
        if (user.id === this.user.id && user.is_admin && this.users.filter(u => u.is_admin).length > 1) {
          return true
        }
        return false
      }
    }
  }
</script>

<style>
  #app {
    height: 100%;
  }
</style>
