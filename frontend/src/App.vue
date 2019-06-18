<template>
  <div id="app">
    <PdfGenerator/>
  </div>
</template>

<script>
import PdfGenerator from "./components/PdfGenerator.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    PdfGenerator
  }
};
</script>
