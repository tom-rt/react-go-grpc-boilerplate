import React from 'react';
import './App.css';
import { Input, Output } from "./pbs/base_pb"
import { BaseServiceClient } from './pbs/base_pb_service';
import { BaseService } from "./pbs/base_pb_service"

const input = new Input();
const output = new Output();
input.setInput("input value");
output.setOutput("output value");

const client = new BaseServiceClient('http://localhost:8000')

function App() {

    let request = new Input();

    let stream = client.ping(request, (resp) => {
      console.log(resp)
    });

  return (
    <div className="App">
      Hello !
    </div>
  );
}

export default App;
