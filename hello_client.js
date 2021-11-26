const {EchoRequest,
    ServerStreamingEchoRequest} = require('./services/hello_pb.js');
const {EchoServiceClient} = require('./services/hello_grpc_web_pb.js');
const {EchoApp} = require('../echoapp.js');
const grpc = {};
grpc.web = require('grpc-web');

var echoService = new EchoServiceClient('http://localhost:8080', null, null);

var echoApp = new EchoApp(
    echoService,
    {
        EchoRequest: EchoRequest,
        ServerStreamingEchoRequest: ServerStreamingEchoRequest
    },
    {
        checkGrpcStatusCode: function(status) {
            if (status.code !== grpc.web.StatusCode.OK) {
                EchoApp.addRightMessage('Error code: '+status.code+' "'+
                    status.details+'"');
            }
        }
    }
);

echoApp.load();