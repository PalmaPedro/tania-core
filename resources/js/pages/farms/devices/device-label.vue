<template lang="pug">
    span
      span(v-if="type == 'TYPE'")
        span.label.label-robot(v-if="device.type == 'ROBOT'")
          translate ROBOT
        span.label.label-temperature-sensor(v-if="device.type == 'SENSOR'")
          translate SENSOR
</template>
<script>
export default {
    name: 'DeviceLabel',
    props: ['type', 'device'],
};
</script>