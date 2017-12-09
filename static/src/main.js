import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import VueRouter from 'vue-router'
import 'vuetify/dist/vuetify.css'
import Play from './Play'
import Leaderboards from './Leaderboards'
import Profile from './Profile'
import FontAwesomeIcon from '@fortawesome/vue-fontawesome'
import fontawesome from '@fortawesome/fontawesome'
import faSolid from '@fortawesome/fontawesome-free-solid'
import VueNativeSock from 'vue-native-websocket'


import store from './store'

Vue.use(VueNativeSock, "ws://localhost:8081/api/socket", {
    format: "json",
    store: store,
});

Vue.use(Vuetify);
Vue.use(VueRouter);

Vue.component(FontAwesomeIcon.name, FontAwesomeIcon);

fontawesome.library.add(faSolid);

const router = new VueRouter({
    routes: [
        {path: '/play', component: Play},
        {path: '/leaderboards', component: Leaderboards},
        {path: '/profile', component: Profile},
    ]
});

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
});
