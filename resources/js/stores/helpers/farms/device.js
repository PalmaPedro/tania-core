export const DeviceDomainCategories = [
    { key: 'ROBOT', label: 'Robot' },
    { key: 'SENSOR', label: 'Sensor' }
]

export function AddClicked(data) {
    for (var i = 0; i < data.length; i++) {
      data[i].clicked = false
    }
    return data
  }