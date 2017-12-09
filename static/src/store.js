import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex);

export default new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    state: {
        loggedIn: false,
    },
    mutations: {
        SOCKET_ONOPEN() {

        },

        login(state) {
            state.loggedIn = true
        }
    }
})