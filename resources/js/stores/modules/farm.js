import FarmArea from './farms/area';
import FarmCrop from './farms/crop';
import FarmFarm from './farms/farm';
import FarmInventories from './farms/inventories';
import FarmReservoir from './farms/reservoir';
import Task from './farms/task';
import Device from './farms/device';

const state = Object.assign({},
  FarmArea.state,
  FarmCrop.state,
  FarmFarm.state,
  FarmInventories.state,
  FarmReservoir.state,
  Task.state,
  Device.state);

const getters = Object.assign({},
  FarmArea.getters,
  FarmCrop.getters,
  FarmFarm.getters,
  FarmInventories.getters,
  FarmReservoir.getters,
  Task.getters,
  Device.getters);

const actions = Object.assign({},
  FarmArea.actions,
  FarmCrop.actions,
  FarmFarm.actions,
  FarmInventories.actions,
  FarmReservoir.actions,
  Task.actions,
  Device.actions);

const mutations = Object.assign({},
  FarmArea.mutations,
  FarmCrop.mutations,
  FarmFarm.mutations,
  FarmInventories.mutations,
  FarmReservoir.mutations,
  Task.mutations,
  Device.mutations);

export default {
  state, getters, actions, mutations,
};
