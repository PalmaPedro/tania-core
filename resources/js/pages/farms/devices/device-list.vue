<template lang="pug">
.device-list.table-responsive
    table.table.m-b-none(v-if="domain == 'AREA' || domain == 'HOME'")
      thead
        tr
          th
          th
            translate Items
          th
            translate Category
          th(v-if="domain != 'AREA'")
      tbody
        tr(v-if="devices.length == 0")
          td(colspan="3")
            translate No Device Created
        tr(v-for="device in devices")
          td
            .checkbox
              label.i-checks
                input(
                  type="checkbox"
                  v-on:change="setDeviceStatus(device.uid, device.status)"
                  :checked="isCompleted(device.status)"
                )
                i
          td
            a(href="#")
              div {{ device.title }}
              MoreDetail(:data="device" :description="device.description")
              small.text-muted(v-if="device.due_date") Due date:
                |
                | {{ device.due_date | moment('timezone', 'Europe/Copenhagen').format('DD/MM/YYYY') }}
                DeviceLabel(:type="'PRIORITY'" :device="device")
                span.text-danger(v-if="device.is_due == true")
                  translate Overdue!
                span.text-success(v-if="isToday(device.due_date)")
                  translate Today
          td
            DeviceLabel(:type="'CATEGORY'" :device="device")
          td(v-if="domain != 'AREA' && domain != 'HOME'")
            a.h3(style="cursor: pointer;" @click="openModal(device)")
              i.fas.fa-edit

    .wrapper(v-else)
      b-list-group
        b-list-group-item(v-if="devices.length == 0")
          translate No Device Created
        b-list-group-item.clearfix(v-for="device in devices" :key="device.uid")
          .row
            .col-sm-1.col-md-1.col-lg-1
              .checkbox
                label.i-checks
                  input(
                    type="checkbox"
                    v-on:change="setDeviceStatus(device.uid, device.status)"
                    :checked="isCompleted(device.status)"
                  )
                  i
            .col-sm-8.col-md-6.col-lg-8
              span.h4.text-dark(v-if="device.category == 'ROBOT' || device.category == 'SENSOR'")
                translate Apply
                |
                |
                u(v-if="device.domain_details.material")
                  | {{ device.domain_details.material.material_name }}
                |
                |
                translate to
                |
                |
                span.identifier-sm(v-if="device.domain_details.crop")
                  | {{ device.domain_details.crop.crop_batch_id }}
                |
                |
                translate on
                |
                |
                span.areatag-sm(v-if="device.domain_details.area")
                  | {{ device.domain_details.area.area_name }}
              span.h4.text-dark(v-else-if="device.category == 'AREA'")
                span.areatag-sm(v-if="device.domain_details.area")
                  |{{ device.domain_details.area.area_name }}
                i.fas.fa-long-arrow-alt-right
                |  {{ device.title }}
              span.h4.text-dark(v-else-if="device.category == 'ROBOT'")
                u(v-if="device.domain_details.robot")
                  | {{ device.domain_details.robot.robot_name }}
                i.fas.fa-long-arrow-alt-right
                |  {{ device.title }}
              span.h4.text-dark(v-else-if="device.category == 'SENSOR'")
                span.identifier-sm(v-if="device.domain_details.sensor")
                  | {{ device.domain_details.sensor.sensor_name }}
                translate on
                //span.areatag-sm(v-if="device.domain_details.area")
                  | {{ device.domain_details.area.area_name }}
                //i.fas.fa-long-arrow-alt-right
                //|  {{ device.title }}
              //span.h4.text-dark(
                v-else-if="device.category == 'SAFETY' || task.category == 'SANITATION'"
              //)
                span.areatag-sm(v-if="task.domain_details.area")
                  | {{ task.domain_details.area.area_name }}
                i.fas.fa-long-arrow-alt-right
                |  {{ task.title }}
              //span.h4.text-dark(v-else) {{ task.title }}
              //MoreDetail(:data="task" :description="task.description")
              //div
                small.text-muted Due date:
                  |
                  | {{ task.due_date | moment('timezone', 'Europe/Copenhagen').format('DD/MM/YYYY') }}
                .status.status-urgent(v-if="task.priority == 'URGENT'")
                  translate URGENT
                span.text-danger(v-if="task.is_due == true")
                  translate Overdue!
            .col-sm-2.col-md-3.col-lg-2
              DeviceLabel(:type="'CATEGORY'" :device="device")
            .col-sm-1.col-md-2.col-lg-1.text-right
              a.h3(v-if="!isCompleted(device.status)" style="cursor: pointer;" @click="openModal(device)")
                i.fas.fa-edit
        Pagination(:pages="pages" @reload="getDevices")

    
</template>

<script>

import { mapActions, mapGetters } from 'vuex';
import moment from 'moment-timezone';
import MoreDetail from '../../../components/more-detail.vue';
import DeviceLabel from './device-label.vue';
import Pagination from '../../../components/pagination.vue';

export default {
  name: 'DevicesList',
  components: {
    MoreDetail,
    Pagination,
    DeviceLabel,
  },
  props: ['asset_id', 'category', 'domain', 'priority', 'reload', 'status'],
  computed: {
    ...mapGetters({
      devices: 'getDevices',
      pages: 'getDevicesNumberOfPages',
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
      //'setDeviceCompleted',
      //'setTaskDue',
      'setDeviceStatus',
    ]),
    getDevices() {
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
    isCompleted(status) {
      return (status === 'COMPLETED');
    },
    //isToday(date) {
    //  return moment(date).tz('Europe/Copenhagen').isSame(moment(), 'day');
    //},
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
</script>

<style lang="scss" scoped>
.wrapper {
    padding: 20px 0;
}

</style>
