<template lang="pug">
.container-fluid.bottom-space
  .row
    .col
      modal(v-if="showCropModal" @close="showCropModal = false")
        DeviceForm(:asset="'Device'" :data="data")
      //modal(v-if="showModal" @close="showModal = false")
      //  DeviceForm(:asset="'General'" :data="data")

      h3.title-page
        translate Devices

  .row
    .col-xs-12.col-sm-12.col-md-8.col-lg-9
      BtnAddNew(:title="$gettext('Add Device')" v-on:click.native="openModal()")

      .cards-wrapper
        DevicesList(:category="selected_type" :status="status" @openModal="openModal")

    
</template>
<script>
import { mapActions } from 'vuex';
import Modal from '../../components/modal.vue';
import DeviceList from '../farms/devices/device-list.vue';
import DeviceForm from '../farms/devices/device-form.vue';
//import CropTaskForm from '../farms/devices/crop-device-form.vue';
import { DeviceDomainCategories } from '../../stores/helpers/farms/device';
import BtnAddNew from '../../components/common/btn-add-new.vue';

export default {
  name: 'Devices',
  components: {
    //CropTaskForm,
    Modal,
    DeviceForm,
    DeviceList,
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
      //showCropModal: false,
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

.cards-wrappe {
    margin-top: 20px;
}

form {
    padding-top: 30px;
}

.list-group-item:hover {
    cursor: pointer;
}

</style>