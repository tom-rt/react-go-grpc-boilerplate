// package: base
// file: base.proto

var base_pb = require("./base_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var BaseService = (function () {
  function BaseService() {}
  BaseService.serviceName = "base.BaseService";
  return BaseService;
}());

BaseService.Ping = {
  methodName: "Ping",
  service: BaseService,
  requestStream: false,
  responseStream: false,
  requestType: base_pb.Input,
  responseType: base_pb.Output
};

exports.BaseService = BaseService;

function BaseServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BaseServiceClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BaseService.Ping, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.BaseServiceClient = BaseServiceClient;

