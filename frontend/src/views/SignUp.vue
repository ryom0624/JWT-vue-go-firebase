<template>
  <div class="signup">
    <h2>Sign up</h2>
    <input type="text" placeholder="email" v-model="email">
    <input type="password" placeholder="password" v-model="password">
    <button @click="signUp">Register</button>
    <p>Do you have an account?
      <router-link to="/signin">sign in now!!</router-link>
    </p>
  </div>
</template>

<script>
  import firebase from 'firebase'
  export default {
    name: "Signup.vue",
    data () {
      return {
        email: '',
        password: '',
      }
    },
    methods: {
      signUp: function() {
        firebase.auth().createUserWithEmailAndPassword(this.email, this.password).then(res => {
          res.user.getIdToken().then(idToken => {
            localStorage.setItem("jwt", idToken.toString())
          })
          this.$router.push("/")
        }).catch(function(error) {
          // eslint-disable-next-line no-console
          console.log("Failed to Authenticate. Reason : " + error)
        })
      }
    }
  }
</script>

<style scoped>

</style>
