<template>
  <div class="uk-overflow-auto uk-padding-small uk-padding-remove-horizontal">
    <table class="uk-table uk-table-striped uk-table-small" v-if="functions && functions.length">
      <caption>Functions</caption>
      <thead>
        <tr>
          <th class="uk-table-expand">Name</th>
          <th class="uk-table-shrink uk-text-right">RPL</th>
        </tr>
      </thead>
      <tbody>
        <RplImportRow v-for="item in functions" :item="item" :key="item.name" />
      </tbody>
    </table>
    <table class="uk-table uk-table-striped uk-table-small" v-if="data && data.length">
      <caption>Data</caption>
      <thead>
        <tr>
          <th class="uk-table-expand">Name</th>
          <th class="uk-table-shrink uk-text-right">RPL</th>
        </tr>
      </thead>
      <tbody>
        <RplImportRow v-for="item in data" :item="item" :key="item.name" />
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import RplImportRow from './RplImportRow.vue';

export default {
  components: { RplImportRow },
  computed: {
    ...mapGetters(['importData', 'importFunctions']),
  	data() {
      if (!this.importData) {
        return [];
      }

    	return this.importData.map(item => Object.freeze(item));
    },
  	functions() {
      if (!this.importFunctions) {
        return [];
      }

    	return this.importFunctions.map(item => Object.freeze(item));
    },
  },
};
</script>

<style>
</style>
