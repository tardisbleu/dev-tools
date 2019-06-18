<template>
  <div class="container">
    <a @click="generatePDF">Générer les PDF</a>
  </div>
</template>

<script>
export default {
  methods: {
    generatePDF: function() {
      var self = this;
      window.backend.PdfGenerator.generatePDF(10).then(result => { });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.6em;
  border-color: #6b778d;
  background-color: #6b778d;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.5em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>