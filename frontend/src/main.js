import Vue from 'vue'
import App from './App.vue'
import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false

const firebaseConfig = {
  apiKey: process.env.VUE_APP_API_KEY,
  authDomain: process.env.VUE_APP_DOMAIN,
  databaseURL: process.env.VUE_APP_DATABASE_URL,
  projectId: process.env.VUE_APP_PROJECTID,
  storageBucket: process.env.VUE_APP_STORAGE_BUCKET,
  messagingSenderId: process.env.VUE_APP_MESSAGEING_SENDERID,
  appId: process.env.VUE_APP_APP_ID
};
firebase.initializeApp(firebaseConfig);

let app

// eslint-disable-next-line no-unused-vars
firebase.auth().onAuthStateChanged(user => {
  if (!app) {
    // eslint-disable-next-line no-console
    console.log(firebase.auth().currentUser)
    new Vue({
      router,
      render: h => h(App)
    }).$mount('#app')
  }
})
