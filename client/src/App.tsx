import { grpc } from '@improbable-eng/grpc-web';
import './App.css';
import { Input, Output } from "./pbs/base_pb"
import { BaseServiceClient, ServiceError } from './pbs/base_pb_service';
import React from 'react'

class App extends React.Component {
  constructor(props: any) {
    super(props)
    this.state = {}
  }

  async pingServer() {
    const client = new BaseServiceClient('http://localhost:8000')
    const request = new Input();
    const metadata: grpc.Metadata = new grpc.Metadata();

    metadata.append("Access-Control-Allow-Origin", "*")
    request.setInput("Ping")
    
    await client.ping(request, (error: ServiceError | null, res: Output | null) => {
      if (error) {
        console.error(`Error ${error.code}: ${error.message}`)
      } else {
        const output: string | undefined = res?.getOutput()
        console.log("Server responded:", output)  
      }
    });
  }

  componentDidMount() { 
    this.pingServer()
  }

  componentWillUnmount() {  }

  render() {
    return (
      <div className="App">
        Ping rpc example !
      </div>
    );
  }

}

export default App;
