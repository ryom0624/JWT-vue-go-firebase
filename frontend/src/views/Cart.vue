<template>
  <div class="cart">
    <h1>This is Cart page</h1>
    <h2>{{user}}様のカートの中身はこちらです</h2>
    <p v-show="Err">{{Err}}</p>
    <p v-show="item">{{item}}</p>
    <div v-show="items">
      <ul id="example-1">
        <li v-for="(item, index) in items" v-bind:key="index">
          {{ item }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  import firebase from 'firebase'

  export default {
    name: 'cart',
    data() {
      return {
        Err: '',
        item: '',
        items: [],
        user: firebase.auth().currentUser.email
      }
    },
    created() {
      axios.get("http://localhost:8089/session/items", {
        withCredentials: true,
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      }).then(res => {
        if (Array.isArray(res.data['Products'])) {
          // eslint-disable-next-line no-console
          console.log(res.data["Products"])
          this.items = res.data["Products"]
        } else if (res.data['Products']) {
          // eslint-disable-next-line no-console
          console.log(res.data["Products"])
          this.item = res.data["Products"]
        } else {
          // eslint-disable-next-line no-console
          console.log(res.data["Products"])
          this.item = "There is no product."
        }
      }).catch(err => {
        this.Err = err
      })
    }
  }


</script>
