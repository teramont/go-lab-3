'use strict';

const Bus = require('./common/json-bus');
const Storage = require('./storage/storage');

(async () => {
  const bus = new Bus('localhost', 8080);
  const storage = new Storage(bus);

  const machines = await storage.listMachines();
  console.table(machines);

  const status = await storage.connect('server-1', 1);
  console.dir(status);
})();

