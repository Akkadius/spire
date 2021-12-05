<template>
  <b-container :class="[ 'bv-example-row', 'p-' + padding ]" v-show="isLoading">
    <b-row class="justify-content-md-center loader">
      <li>L</li>
      <li>O</li>
      <li>A</li>
      <li>D</li>
      <li>I</li>
      <li>N</li>
      <li>G</li>
      <li>.</li>
      <li>.</li>
      <li>.</li>
      <li></li>
      <li>P</li>
      <li>L</li>
      <li>E</li>
      <li>A</li>
      <li>S</li>
      <li>E</li>
      <li></li>
      <li>W</li>
      <li>A</li>
      <li>I</li>
      <li>T</li>
      <li>.</li>
      <li>.</li>
      <li>.</li>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: 'app-loader',
  props: {
    isLoading: {
      default: null,
      required: true
    },
    padding: {
      default: 2,
      required: false
    },
    size: {
      default: 30,
      required: false
    }
  }
}
</script>

<style scoped>
.loader {
  display: flex;
  list-style-type: none;
  padding: 0;
  margin: 0;
  font-weight: bold;
  font-size: 43px;
  color: #222;
  font-family: menuetto, serif;
}

.loader li {
  padding: 0 3px;
  animation: blink 2s linear infinite;
}

@keyframes blink {
  0% {
    color: #ffff00;
    text-shadow: 0px 0px 3px #000, 0px 0px 0px #000, 1px 2px 5px #000, 4px 4px 5px #000;
  }

  50% {
    color: #ffff00;
    text-shadow: 0 0 10px #fff900, 0 0 10px #fff900, 0 0 30px #fff900;
  }

  100% {
    color: #ffff00;
    text-shadow: 0px 0px 3px #000, 0px 0px 0px #000, 1px 2px 5px #000, 4px 4px 5px #000;
  }
}
</style>
