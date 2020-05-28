<template>
  <v-container>
    <v-btn @click="loadModel">Load model...</v-btn>
  </v-container>
</template>

<script>
import * as tf from "@tensorflow/tfjs";
import * as tmImage from "@teachablemachine/image";

export default {
  name: "Perzon",

  data: () => ({
    model: null,
    publicPath: process.env.BASE_URL,
  }),
  methods: {
    loadModel() {
      //   const uploadModel = document.getElementById("upload-model");
      //   const uploadWeights = document.getElementById("upload-weights");
      //   const uploadMetadata = document.getElementById("upload-metadata");
      // file_mode = new File()
      this.model = tmImage
        // .loadFromFiles(
        //   uploadModel.files[0],
        //   uploadWeights.files[0],
        //   uploadMetadata.files[0]
        // )
        .loadFromFiles("model.json", "metadata.json", "weights.bin")
        .then((mod) => {
          this.model = mod;
          console.log("loaded model " + mod.name);
          console.log("loaded model " + mod.getTotalClasses());
        })
        .catch(function(error) {
          console.log("loadLayersModel error");
          console.log(error);
        });

      //   tmImage
      //     .load(modelURL, metadataURL)
      //     .then((mod) => {
      //       this.model = mod;
      //       console.log("loaded model " + mod.name);
      //     })
      //     .catch(function(error) {
      //       console.log("loadLayersModel error");
      //       console.log(error);
      //     });

      //   this.model = tf
      //     .loadLayersModel(
      //       "https://storage.googleapis.com/tfjs-models/tfjs/iris_v1/model.json"
      //     )
      //     .then((mod) => {
      //       console.log("loaded model " + mod.name);
      //     })
      //     .catch(function(error) {
      //       console.log("loadLayersModel error");
      //       console.log(error);
      //     });
    },
  },
  mounted() {
    tf.setBackend("webgl");
    console.log(tf.getBackend());
  },
};
</script>
