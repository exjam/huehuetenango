import Vue from 'vue';
import Vuex from 'vuex';
import auth from './modules/auth';
import titles from './modules/titles';
import search from './modules/search';
import rpls from './modules/rpls';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
    titles,
    search,
    rpls,
  },
});
