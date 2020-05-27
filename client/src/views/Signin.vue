<template>
  <div>Signing in ...</div>
</template>

<script>
import axios from "axios";

export default {
  name: "Signin",
  data() {
    return {
      iframeURL: ""
    };
  },
  created() {
    axios
      .get("http://localhost:3000/auth/steam/return" + window.location.search, {
        withCredentials: true
      })
      .then(() => {
        this.$router.push({ name: "home" });
      });
  }
};
</script>

<style scoped></style>
