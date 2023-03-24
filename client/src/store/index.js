import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    user: null,
    sessionId:0,
  },
  getters: {
    userAuthorized: (state) => {
      return state.user.is_authorized
    },
    getRedirectUrl: (state) => {
      return `${state.user.redirect_url}?code=${state.user.code}`
    },
    getUserInfo: (state) => {
      return state.user.user
    }
  },
  mutations: {
    setUser (state,user) {
      state.user = user
    },
    setSession (state,id) {
      state.sessionId = id
    },

  },
  actions: {
  },
  modules: {
  }
})
