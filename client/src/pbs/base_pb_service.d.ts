// package: base
// file: base.proto

import * as base_pb from "./base_pb";
import {grpc} from "@improbable-eng/grpc-web";

type BaseServicePing = {
  readonly methodName: string;
  readonly service: typeof BaseService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof base_pb.Input;
  readonly responseType: typeof base_pb.Output;
};

export class BaseService {
  static readonly serviceName: string;
  static readonly Ping: BaseServicePing;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class BaseServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  ping(
    requestMessage: base_pb.Input,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: base_pb.Output|null) => void
  ): UnaryResponse;
  ping(
    requestMessage: base_pb.Input,
    callback: (error: ServiceError|null, responseMessage: base_pb.Output|null) => void
  ): UnaryResponse;
}

