<template>
  <v-app>
    <v-app-bar
      app
      color="black"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Logo"
          class="shrink mr-2"
          contain
          src="dkt.png"
          transition="scale-transition"
          width="120"
        />
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/dkt64/perzon"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-content>
      <Perzon/>
    </v-content>
  </v-app>
</template>

<script>
import Perzon from './components/Perzon';

export default {
  name: 'App',

  components: {
    Perzon,
  },

  data: () => ({
    //
  }),
};
</script>
