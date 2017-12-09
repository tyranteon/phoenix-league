<template>
    <span v-if="!isLoggedIn">
        <a href="https://steamcommunity.com/openid/login?openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.realm=http%3A%2F%2Fwww.phoenix-league.net%2F&openid.return_to=http%3A%2F%2Fwww.phoenix-league.net%2Fsignin">
            <img src="/img/signin-steam.png" style="cursor: pointer !important;">
        </a>
    </span>
    <span v-else>
        Logged in!
    </span>
</template>

<script>
    export default {
        name: "login",

        computed: {
            isLoggedIn() {
                return this.$store.state.loggedIn;
            }
        }
    }
</script>

<style scoped>

</style>