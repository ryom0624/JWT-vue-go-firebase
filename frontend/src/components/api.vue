<template>
  <div class="api">
    <h1>{{msg}}</h1>
    <button @click="apiPublic">PUBLIC</button>
    <button @click="apiPrivate">PRIVATE</button>
    <p>{{response}}</p>
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
    }
  }
</script>

<style scoped>
  button {
    margin: 10px;
    padding: 10px;
  }
</style>
