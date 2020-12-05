'use strict';

const Bus = require('./bus');

const unwrap = response => {
  if (response.statusCode != 200) {
    throw new Error(response.body);
  }
  const body = JSON.parse(response.body);
  return body;
}

class JSONBus extends Bus {
  constructor(host, port) {
    super(host, port);
  }

  async request(method, path, rawBody) {
    const response = await super.request(method, path, JSON.stringify(rawBody));
    const body = response.statusCode == 200 ?
      JSON.parse(response.body) :
      response.body;
    return unwrap(response);
  }
}

module.exports = JSONBus;
