<template>
  <div class="signin">
    <h2>Sign in</h2>
    <input type="text" placeholder="email" v-model="email">
    <input type="password" placeholder="Password" v-model="password">
    <button @click="signIn">Signin</button>
    <p>You don't have an account?
      <router-link to="/signup">create account now!!</router-link>
    </p>
  </div>
</template>

<script>
  import firebase from 'firebase'
  export default {
    name: "SignIn.vue",
    data () {
      return {
        email: '',
        password: '',
      }
    },
    methods: {
      signIn: function() {
        firebase.auth().signInWithEmailAndPassword(this.email, this.password).then(res => {
          res.user.getIdToken().then(idToken => {
            localStorage.setItem('jwt', idToken.toString())
          })
          this.$router.push('/')
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
