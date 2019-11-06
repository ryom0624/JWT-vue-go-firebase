<template>
  <div class="api">
    <api></api>
    <p>Do you want to sign out? click here!</p>
    <button @click="signOut">SignOut</button>
  </div>
</template>

<script>
import api from '../components/api.vue'
import firebase from 'firebase'

export default {
  name: "API.vue",
  components: {
    api,
  },
  methods : {
    signOut: function() {
      firebase.auth().signOut().then(() =>{
        localStorage.removeItem("jwt")
        this.$router.push("/signin")
      })
    }
  }
}
</script>

<style scoped>

</style>
