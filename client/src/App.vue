<template>
  <v-app>
    <span>
      <a :href="openidURL">
        <img
          alt="sign in through Steam"
          src="/img/signin-steam.png"
          style="cursor: pointer !important;"
        />
      </a>
      <router-view />
    </span>
  </v-app>
</template>

<script>
export default {
  name: "App",
  components: {},
  data() {
    let host = encodeURIComponent("http://localhost:8080");

    return {
      openidURL: `https://steamcommunity.com/openid/login?openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.realm=${host}%2F&openid.return_to=${host}%2Fsignin`
    };
  }
};
</script>
