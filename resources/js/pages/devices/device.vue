<template lang="pug">
.container-fluid.bottom-space
  .row
    .col
      //modal(v-if="showCropModal" @close="showCropModal = false")
        CropTaskForm(:asset="'Crop'" :data="data")
      modal(v-if="showModal" @close="showModal = false")
        DeviceForm(:asset="'General'" :data="data")

      h3.title-page
        translate Devices

  .row
    .col-xs-12.col-sm-12.col-md-8.col-lg-9
      BtnAddNew(:title="$gettext('Add Device')" v-on:click.native="openModal()")

      .cards-wrapper
        DevicesList(:category="selected_category" :priority="selected_priority" :status="status" @openModal="openModal")

    .col-xs-12.col-sm-12.col-md-4.col-lg-3
      b-form
        .form-group
          label.control-label
            translate Category

        .form-group
          select.form-control(:class="{'input': true, 'text-danger': errors.has('category') }" name="category" @change="categoryChange($event.target.value)" v-model="selected_category")
            option(value="")
              translate All
            //option(value="ROBOT")
              translate Robot
            //option(value="RESERVOIR")
              //translate Reservoir
            //option(value="SENSOR")
              translate Sensor
            option(v-for="category in options.deviceCategories" :value="category.key") {{ category.label }}
        
      //.wrapper
        b-list-group
          b-list-group-item(
            v-bind:class="{ active: isActive('COMPLETED') }"
            @click="statusSelected('COMPLETED')"
          )
            h5.text-info
              translate Completed
          b-list-group-item(
            v-bind:class="{ active: isActive('INCOMPLETE') }"
            @click="statusSelected('INCOMPLETE')"
          )
            h5.text-muted
              translate Incomplete
          b-list-group-item(
            v-bind:class="{ active: isActive('OVERDUE') }"
            @click="statusSelected('OVERDUE')"
          )
            h5.text-danger
              translate Overdue
          b-list-group-item(
            v-bind:class="{ active: isActive('TODAY') }"
            @click="statusSelected('TODAY')"
          )
            h5.text-success
              translate Today
          b-list-group-item(
            v-bind:class="{ active: isActive('THISWEEK') }"
            @click="statusSelected('THISWEEK')"
          )
            h5.text-lt
              translate This Week
          b-list-group-item(
            v-bind:class="{ active: isActive('THISMONTH') }"
            @click="statusSelected('THISMONTH')"
          )
            h5.text-lt
              translate This Month
</template>

<script>
import { mapActions } from 'vuex';
import Modal from '../../components/modal.vue';
import DevicesList from '../farms/devices/device-list.vue';
import DeviceForm from '../farms/devices/device-form.vue';
//import CropTaskForm from '../farms/tasks/crop-task-form.vue';
import { DeviceDomainCategories } from '../../stores/helpers/farms/device';
import BtnAddNew from '../../components/common/btn-add-new.vue';

export default {
  name: 'Devices',
  components: {
    //CropTaskForm,
    Modal,
    DeviceForm,
    DevicesList,
    BtnAddNew,
  },
  data() {
    return {
      data: {},
      options: {
        deviceCategories: Array.from(DeviceDomainCategories),
      },
      selected_category: '',
      selected_priority: '',
      showCropModal: false,
      showModal: false,
      status: 'INCOMPLETE',
    };
  },
  methods: {
    ...mapActions([
    ]),
    closeModal() {
      this.showModal = false;
    },
    categoryChange(type) {
      this.selected_category = type;
    },
    openModal(data) {
      if (data) {
        this.data = data;
        if (data.domain === 'CROP') {
          this.showCropModal = true;
        }
      } else {
        this.data = {};
      }
      if (!this.showCropModal) {
        this.showModal = true;
      }
    },
    priorityChange(type) {
      this.selected_priority = type;
    },
    statusSelected(status) {
      this.status = status;
    },
    isActive(status) {
      return this.status === status;
    },
  },
};
</script>

<style lang="scss" scoped>
h3.title-page {
  margin: 20px 0 30px 0;
}

.bottom-space {
  padding-bottom: 60px;
}

.cards-wrapper {
  margin-top: 20px;
}

form {
  padding-top: 30px;
}

.list-group-item:hover {
  cursor: pointer;
}
</style>