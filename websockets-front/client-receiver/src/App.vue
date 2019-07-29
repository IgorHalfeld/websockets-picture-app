<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-bold primary--text">Nerdzão</span>
        <span class="font-weight-light primary--text">Week</span>
      </v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container justify-center align-center fill-height column class="text-center">
        <h1 class="font-weight-bold display-2 primary--text" v-if="hasSomeImage">Waiting for photos 😘</h1>
        <v-layout column justify-center align-center v-else>
          <v-img
            v-for="(image, index) in images"
            :key="`image-${index}`"
            :src="image"
            class="grey lighten-2 mt-3 mb-3"
            max-width="700" />
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    images: [],
  }),
  computed: {
    hasSomeImage() {
      return !this.images.length;
    },
  },
  created() {
    if (!window.WebSocket) {
      this.hasError = {
        status: true,
        message: 'Your browser doesn\'t supports WebSockets :c',
      };
      return;
    }

    const URI = process.env.VUE_APP_WS_URI || 'ws://localhost:3000/ws';
    this.wsConnection = new WebSocket(URI);
    this.wsConnection.addEventListener('message', event => {
      this.images.push(event.data);
      this.$forceUpdate();
      setTimeout(() => {
        window.scroll({
          top: document.documentElement.scrollHeight,
          behavior: 'smooth',
        });
      }, 500);
    });
  },
};
</script>

<style>
.wrapper {
  overflow: hidden;
}
</style>