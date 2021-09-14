import { grpc } from '@improbable-eng/grpc-web';
import React from 'react';
import './App.css';
import { Input, Output } from "./pbs/base_pb"
import { BaseServiceClient } from './pbs/base_pb_service';
import { BaseService } from "./pbs/base_pb_service"

const input = new Input();
const output = new Output();
input.setInput("input value");
output.setOutput("output value");

function App() {

    const client = new BaseServiceClient('http://localhost:8000')
    let request = new Input();
    const metadata: grpc.Metadata = new grpc.Metadata();

    metadata.append("Access-Control-Allow-Origin", "*")
    // metadata.set("Access-Control-Allow-Origin", "*")
    // metadata.append("hellotest", "hello")
    // metadata.set("hellotest", "hello")
    // metadata.
    // metadata.set('Access-Control-Allow-Origin', '*')

    // client.serviceHost
    let stream = client.ping(request, (resp) => {
      console.log("resp", resp)
    });

  return (
    <div className="App">
      Hello !
    </div>
  );
}

export default App;
