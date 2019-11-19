<template>
  <div class="api">
    <h1>{{name}} <br> {{msg}}</h1>
    <div>
      <p>Access to http://localhost:8081</p>
      <button @click="apiPublic">PUBLIC</button>
      <button @click="apiPrivate">PRIVATE</button>
    </div>
    <div>
      <p>Access to http://localhost:8089</p>
      <button @click="apiPing">PING</button>
      <button @click="getUID">GET UID</button>
    </div>
    <h2>{{response}}</h2>

    <div>
      <input type="text" v-model="product">
      <button @click="addCart(product)">Send</button>
      <p>{{ addProduct }}</p>
    </div>

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
        product: "",
        addProduct: "",
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
        axios.get("http://localhost:8089/ping", {
          withCredentials: true
        }).then(res => {
          this.response = res.data
        }).catch(err => {
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      },
      getUID: function() {
        axios.get("http://localhost:8089/session/getuid", {
          headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`},
          withCredentials: true
        }).then(res => {
          this.response = res.data
        }).catch(err => {
          firebase.auth().signOut().then(() =>{
            localStorage.removeItem("jwt")
            this.$router.push("/signin")
          })
          // eslint-disable-next-line no-console
          console.log("Faild to connect API server REASON:" + err)
        })
      },
      addCart: function(product) {
        axios.post("http://localhost:8089/session/cart", {
          "product": product
        }, {
          withCredentials: true,
          headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
        }).then(res => {
          // eslint-disable-next-line no-console
          console.log(product)
          this.addProduct = res.data
        }).catch(err => {
          firebase.auth().signOut().then(() =>{
            localStorage.removeItem("jwt")
            this.$router.push("/signin")
          })
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
