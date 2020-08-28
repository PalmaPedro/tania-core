<template lang="pug">
  .devices-form
    .modal-header
      h4.font-bold(v-if="device.uid")
        translate Update Device
      h4.font-bold(v-else-if="asset != 'General'")
        | {{ asset }}:
        |
        translate Add New Device on
        |
        |
        span.areatag {{ data.name }}
      h4.font-bold(v-else)
        translate Add New Device
    .modal-body
      b-form(@submit.prevent="validateBeforeSubmit")
        .form-row
          .col-xs-12.col-sm-12.col-md-8.col-lg-7
        .form-group
          label#label-category(for="category")
            translate Device Category
          select.form-control#category(
            v-validate="'required'"
            :class="{'input': true, 'text-danger': errors.has('category') }"
            v-model="device.category"
            name="category"
          )
            option(value="")
              translate Please select category
            //option(v-if="asset_name == 'ROBOT'" value="ROBOT")
              translate Robot
            //option(v-if="asset_name == 'SENSOR'" value="SENSOR")
              translate Sensor
            //option(v-if="asset_name == 'GENERAL'" value="GENERAL")
              translate General
            option(v-for="category in options.deviceCategories" :value="category.key")
              | {{ category.label }}
            //option(v-if="asset_name == 'GENERAL'" value="INVENTORY")
              translate Inventory
          span.help-block.text-danger(v-show="errors.has('category')")
            | {{ errors.first('category') }}
        .form-group
          label#label-title(for="title")
            translate Title
          input.form-control#title(
            type="text"
            v-validate="'required|max:100'"
            :class="{'input': true, 'text-danger': errors.has('title') }"
            v-model="device.title"
            name="title"
          )
          span.help-block.text-danger(v-show="errors.has('title')") {{ errors.first('title') }}
        .form-group
          label#label-description(for="description")
            translate Description
          textarea.form-control#description(
            type="text"
            :class="{'input': true, 'text-danger': errors.has('description') }"
            v-model="device.description"
            name="description"
            rows="3"
          )
          span.help-block.text-danger(v-show="errors.has('description')")
            | {{ errors.first('description') }}
        .form-group
          BtnCancel(v-on:click.native="$parent.$emit('close')")
          BtnSave(customClass="float-right")
</template>

<script>
import { mapActions } from 'vuex';
//import Datepicker from 'vuejs-datepicker';
import { StubDevice } from '../../../stores/stubs';
import { DeviceDomainCategories } from '../../../stores/helpers/farms/device';
import BtnCancel from '../../../components/common/btn-cancel.vue';
import BtnSave from '../../../components/common/btn-save.vue';

export default {
  name: 'FarmDevicesForm',
  components: {
    //Datepicker,
    BtnCancel,
    BtnSave,
  },
  props: {
    data: {
      type: Object,
      default: () => {},
    },
    asset: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      device: Object.assign({}, StubDevice),
      asset_name: '',
      options: {
        deviceCategories: Array.from(DeviceDomainCategories),
      },
    };
  },
  mounted() {
    if (typeof this.data.domain !== 'undefined') {
      this.device.uid = this.data.uid;
      //this.task.due_date = this.data.due_date;
      this.device.priority = this.data.priority;
      this.device.category = this.data.category;
      this.device.title = this.data.title;
      this.device.description = this.data.description;
      this.asset_name = this.data.domain;
    } else {
      this.asset_name = this.asset.toUpperCase();
    }
  },
  methods: {
    ...mapActions([
      //'openPicker',
      'submitDevice',
    ]),
    validateBeforeSubmit() {
      this.$validator.validateAll().then((result) => {
        if (result) {
          this.submit();
        }
      });
    },
    //openPicker() {
    //  this.$refs.openCal.showCalendar();
    //},
    submit() {
      if (typeof this.data.domain !== 'undefined') {
        this.device.domain = this.data.domain;
      } else {
        this.device.asset_id = this.data.uid;
        this.device.domain = this.asset.toUpperCase();
      }
      this.submitDevice(this.device)
        .then(() => this.$parent.$emit('close'))
        .catch(() => this.$toasted.error('Error submiting a device'));
    },
  },
};
</script>

<style lang="scss" scoped>
i.fas.fa-check,
i.fas.fa-times {
  width: 30px;
}
</style>