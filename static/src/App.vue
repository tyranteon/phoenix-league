<template>
    <v-app :dark="nightmode">
        <v-navigation-drawer
                fixed
                clipped
                v-model="drawer"
                app
        >
            <v-list dense>
                <v-list-tile
                        v-for="(item, i) in items"
                        :key="i"
                        :to="item.url"
                        ripple
                >
                    <v-list-tile-action style="margin-right: -16px">
                        <font-awesome-icon :icon="item.icon" fixed-width size="lg"
                                           style="opacity: .7;"></font-awesome-icon>
                    </v-list-tile-action>
                    <v-list-tile-content>
                        <v-list-tile-title v-text="item.title"></v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>
        <v-toolbar fixed app clipped-left>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-toolbar-title v-text="title"></v-toolbar-title>
            <v-spacer></v-spacer>
            <login></login>
            <v-btn icon @click.stop="nightmode = !nightmode">
                <font-awesome-icon icon="moon"></font-awesome-icon>
            </v-btn>
        </v-toolbar>
        <v-content>
            <v-container fluid>
                <v-slide-y-transition mode="out-in">
                    <router-view></router-view>
                </v-slide-y-transition>
            </v-container>
        </v-content>
        <v-footer :fixed="false" app>
            <span>&copy; 2017</span>
        </v-footer>
    </v-app>
</template>

<script>
    import Login from "./Login";

    export default {
        components: {Login},
        data() {
            return {
                drawer: true,
                nightmode: false,
                items: [
                    {icon: 'play', title: 'Play', url: '/play'},
                    {icon: 'trophy', title: 'Leaderboards', url: '/leaderboards'},
                    {icon: 'home', title: 'Profile', url: '/profile'},
                ],
                title: 'Phoenix League',
            }
        },
    }
</script>

<style scoped>

</style>