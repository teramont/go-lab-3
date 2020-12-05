'use strict';

const unwrap = response => {
  if (response.message != undefined) {
    throw new Error(response.message);
  }
  return response;
}

class Storage {
  constructor(bus) {
    this.bus = bus;
  }

  async listMachines() {
    const response = await this.bus.request('GET', '/machines');
    return unwrap(response);
  }

  async connect(machineName, diskId) {
    const body = { machineName, diskId };
    const response = await this.bus.request('PATCH', '/machines', body);
    return unwrap(response);
  }
}

module.exports = Storage;
