<template>
  <v-container>
    <v-btn @click="loadModel">Load model...</v-btn>
  </v-container>
</template>

<script>
import * as tf from "@tensorflow/tfjs";
import * as tmImage from "@teachablemachine/image";

// import axios from "axios";

export default {
  name: "Perzon",

  data: () => ({
    model_json: "",
    metadata_json: "",
    model: null,
  }),
  methods: {
    loadModel() {
      //
      // Create model
      //
      tmImage
        .load("http://localhost/api/v1/model", "http://localhost/api/v1/metadata")
        .then((mod) => {
          this.model = mod;
          console.log("Created model " + mod.name);
        })
        .catch(function(error) {
          console.log("tmImage load error");
          console.log(error);
        });

      // //
      // // Load Model file
      // //
      // axios
      //   .get("/api/v1/model")
      //   .then((response) => {
      //     console.log("Model: ");
      //     console.log(response.data);
      //     this.model_json = response.data;
      //     //
      //     // Load Metadata file
      //     //
      //     axios
      //       .get("/api/v1/metadata")
      //       .then((response) => {
      //         console.log("Metadata: ");
      //         console.log(response.data);
      //         this.metadata_json = response.data;

      //         //
      //         // Create model
      //         //
      //         tmImage
      //           .load(this.model_json, this.metadata_json)
      //           .then((mod) => {
      //             this.model = mod;
      //             console.log("Created model " + mod.name);
      //           })
      //           .catch(function(error) {
      //             console.log("tmImage load error");
      //             console.log(error);
      //           });
      //       })
      //       .catch(function(error) {
      //         console.log("axios get /api/v1/metadata error");
      //         console.log(error);
      //       });
      //   })
      //   .catch(function(error) {
      //     console.log("axios get /api/v1/model error");
      //     console.log(error);
      //   });
    },
  },
  mounted() {
    tf.setBackend("webgl");
    console.log(tf.getBackend());
  },
};
</script>
