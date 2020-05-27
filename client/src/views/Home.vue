<template>
  <div>Home Page!</div>
</template>

<script>
export default {
  name: "Home",
  mounted() {
    let socket = new WebSocket("ws://localhost:3000/");
    socket.onerror = ev => {
      console.error(ev);
    };
    socket.onopen = () => {
      console.log("Connected!");
    };
    socket.onmessage = ev => {
      console.log("Received Message: ", ev.data);
    };
  }
};
</script>

<style scoped></style>
