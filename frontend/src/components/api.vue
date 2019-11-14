<template>
  <div class="api">
    <h1>{{name}} <br> {{msg}}</h1>
    <div>
      <p>Access to http://localhost:8081</p>
      <button @click="apiPublic">PUBLIC</button>
      <button @click="apiPrivate">PRIVATE</button>
    </div>
    <div>
      <p>Access to http://localhost:8083</p>
      <button @click="apiPing">PING</button>
      <button @click="apiSession">Session</button>
    </div>
    <h2>{{response}}</h2>
  </div>
</template>

<script>
  import firebase from 'firebase'
  import axios from "axios"
  export default {
    name: "api.vue",
    data() {
      return {
        msg: "Hello! You can try JWT authentication!",
        response: "",
        name: firebase.auth().currentUser.email
      }
    },
    methods: {
      apiPublic: function() {
        axios.get("http://localhost:8081/public").then(res => {
          this.response = res.data
        }).catch(err => {
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      },
      apiPrivate: function() {
        axios.get("http://localhost:8081/private", {
          headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
        }).then(res => {
          this.response = res.data
        }).catch(err => {
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      },
      apiPing: function() {
        axios.get("http://localhost:8083/ping").then(res => {
          this.response = res.data
        }).catch(err => {
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      },
      apiSession: function() {
        axios.get("http://localhost:8083/session", {
          headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
        }).then(res => {
          this.response = res.data
        }).catch(err => {
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      }
    }
  }
</script>

<style scoped>
  button {
    margin: 10px;
    padding: 10px;
  }
</style>
