<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <el-row>
      <el-button enabled v-on:click="onsubmit">获取信息</el-button>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  //axios
  methods: {
    onsubmit: function () {
      this.axios({
        method: 'get',
        url: 'http://119.28.77.217:9090/getdata'
      }).then(function (response) {
        console.log(response)
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
