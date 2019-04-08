<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <ul>
      <form @submit.prevent="submit">
        <input type="text" name="name" v-model="user.username">
        <input type="password" name="password" v-model="user.password">
        <input type="submit" value="提交">
      </form>
    </ul>
  </div>
</template>

<script>
import {http} from '../common/http.js'
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      user: {
        username: 'test',
        password: 'test'
      }
    }
  },
  methods: {
    submit: function (event) {
      var qs = JSON.stringify(this.user)
      http.post('/login', qs).then((response) => {
        // success callback
        alert('success')
      }, (response) => {
        // error callback
        alert('error')
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
