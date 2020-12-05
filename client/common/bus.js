'use strict';

const http = require('http');

const request = ({ options, body }) => new Promise((resolve, reject) => {
  const req = http.request(options, res => {
    const response = {
      statusCode: res.statusCode,
      headers: res.headers,
      body: '',
    };
    res.setEncoding('utf8');
    res.on('data', c => response.body += c);
    res.on('end', () => resolve(response));
  });

  req.on('error', reject);

  if (body) req.write(body);
  req.end();
});

const makeRequest = (host, port, method, path, body) => {
  const options = {
    method,
    host,
    port,
    path,
  };
  return { body, options };
};

class Bus {
  constructor(host, port) {
    this.host = host;
    this.port = port;
  }

  async request(method, path, body) {
    const req = makeRequest(this.host, this.port, method, path, body);
    return request(req);
  }
}


module.exports = Bus;
