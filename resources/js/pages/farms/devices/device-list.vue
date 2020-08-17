<template lang="pug">
    
</template>
<script>
export default {
    
}
</script>
import { mapActions, mapGetters } from 'vuex';
import moment from 'moment-timezone';
import MoreDetail from '../../../components/more-detail.vue';
import DeviceLabel from './device-label.vue';
import Pagination from '../../../components/pagination.vue';

export default {
  name: 'DeviceList',
  components: {
    MoreDetail,
    Pagination,
    DeviceLabel,
  },
  props: ['asset_id', 'category', 'domain', 'priority', 'reload', 'status'],
  computed: {
    ...mapGetters({
      devices: 'getDevices',
      pages: 'getDeviceNumberOfPages',
    }),
  },
  created() {
    this.getDevices();
  },
  mounted() {
    this.$watch('reload', () => {
      this.getDevices();
    }, {});
    this.$watch('category', () => {
      this.getDevices();
    }, {});
    this.$watch('priority', () => {
      this.getDevices();
    }, {});
    this.$watch('status', () => {
      this.getDevices();
    }, {});
  },
  methods: {
    ...mapActions([
      'getDevicesByDomainAndAssetId',
      'getDevicesByCategoryAndPriorityAndStatus',
      'getDevices',
      'fetchDevices',
      'setTaskStatus',
    ]),
    getDdevices() {
      let pageId = 1;
      if (typeof this.$route.query.page !== 'undefined') {
        pageId = parseInt(this.$route.query.page, 10);
      }
      if (this.domain) {
        this.getDevicesByDomainAndAssetId({ domain: this.domain, assetId: this.asset_id, pageId });
      } else if (this.category !== '' || this.priority !== '' || this.status !== '') {
        const status = (this.status === 'INCOMPLETE') ? '' : this.status;
        this.getDevicesByCategoryAndPriorityAndStatus({
          category: this.category,
          priority: this.priority,
          status,
          pageId,
        });
      } else {
        this.fetchDevices({ pageId });
      }
    },
    openModal(data) {
      this.data = data;
      this.$emit('openModal', this.data);
    },
    setDeviceStatus(deviceId) {
      this.setDeviceCompleted(deviceId)
        .then(() => { this.getDevices(); })
        .catch(({ data }) => {
          this.message = data;
          return this.message;
        });
    },
  },
};

<style lang="scss" scoped>
.wrapper {
  padding: 20px 0;
}

</style>