import NProgress from 'nprogress'

import * as types from '../../mutation-types'
import { AddClicked } from '../../helpers/farms/device'
import { calculateNumberOfPages, pageLength } from '../../constants'
import FarmApi from '../../api/farm'
//import moment from 'moment-timezone'

const state = {
    devices: [],
    pages: 0,
    total: 0,
}
  
const getters = {
    getDevices: state => state.devices,
    getNumberOfDevices: state => state.total,
    getDevicesNumberOfPages: state => state.pages,
}

const actions = {
    fetchDevices ({ commit, state }, payload) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        FarmApi
          .ApiFetchDevice(payload.pageId, ({ data }) => {
            commit(types.FETCH_DEVICES, data.data)
            commit(types.SET_PAGES, data.total_rows)
            resolve(data)
          }, error => reject(error.response))
      })
    },
    getDevicesByDomainAndAssetId ({ commit, state }, payload) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        FarmApi
          .ApiFindDevicesByDomainAndAssetId(payload.pageId, payload.domain, payload.assetId, ({ data }) => {
            commit(types.FETCH_DEVICES, data.data)
            commit(types.SET_PAGES, data.total_rows)
            resolve(data)
          }, error => reject(error.response))
      })
    },
    getDevicesByCategoryAndPriorityAndStatus ({ commit, state }, payload) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        let query = '&'
        if (payload.status == 'COMPLETED') {
          query += 'status=COMPLETED'
        } /*else if (payload.status == 'THISWEEK') {
          let due_start = moment().startOf('week').format('YYYY-MM-DD')
          let due_end = moment().endOf('week').format('YYYY-MM-DD')
          query += 'due_start=' + due_start +'&due_end=' + due_end
        }  else if (payload.status == 'THISMONTH') {
          let due_start = moment().startOf('month').format('YYYY-MM-DD')
          let due_end = moment().endOf('month').format('YYYY-MM-DD')
          query += 'due_start=' + due_start +'&due_end=' + due_end
        } else if (payload.status == 'OVERDUE') {
          query += 'is_due=true'
        } else if (payload.status == 'TODAY') {
          let due = moment().format('YYYY-MM-DD')
          query += 'due_date=' + due
        } */else {
          query += 'status=CREATED'
        }
        FarmApi
          .ApiFindDevicesByCategoryAndPriorityAndStatus(payload.pageId, payload.category, payload.priority, query, ({ data }) => {
            commit(types.FETCH_DEVICES, data.data)
            commit(types.SET_PAGES, data.total_rows)
            resolve(data)
          }, error => reject(error.response))
      })
    },
    submitDevice ({ commit, state, getters }, payload) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        if (payload.uid != '') {
          FarmApi
            .ApiUpdateDevice(payload.uid, payload, ({ data }) => {
              payload = data.data
              commit(types.UPDATE_DEVICE, payload)
              resolve(payload)
            }, error => reject(error.response))
        } else {
          FarmApi
            .ApiCreateDevice(payload, ({ data }) => {
              payload = data.data
              commit(types.CREATE_DEVICE, payload)
              resolve(payload)
            }, error => reject(error.response))
        }
      })
    },/*
    setTaskDue ({ commit, state, getters }, taskId) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        FarmApi
          .ApiSetTaskDue(taskId, ({ data }) => {
            resolve(data)
          }, error => reject(error.response))
      })
    },
    setTaskCompleted ({ commit, state, getters }, taskId) {
      NProgress.start()
      return new Promise((resolve, reject) => {
        FarmApi
          .ApiSetTaskCompleted(taskId, ({ data }) => {
            resolve(data)
          }, error => reject(error.response))
      })
    },*/
}

const mutations = {
    [types.CREATE_DEVICE] (state, payload) {
      state.devices.unshift(payload)
      if (state.devices.length > pageLength) {
        state.devices.pop()
      }
      state.pages = calculateNumberOfPages(state.devices.length + 1)
    },
    [types.UPDATE_DEVICE] (state, payload) {
      const devices = state.devices
      state.devices = devices.map(device => (device.uid === payload.uid) ? payload : device)
      state.devices = AddClicked(state.devices)
    },
    [types.FETCH_DEVICES] (state, payload) {
      state.devices = AddClicked(payload)
    },
    [types.SET_PAGES] (state, payload) {
      state.total = payload
      state.pages = calculateNumberOfPages(payload)
    },
  }
  
  export default {
    state, getters, actions, mutations
  }
  