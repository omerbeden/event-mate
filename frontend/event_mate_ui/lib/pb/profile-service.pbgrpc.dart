///
//  Generated code. Do not modify.
//  source: profile-service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'profile-service.pb.dart' as $0;
export 'profile-service.pb.dart';

class ProfileServiceClient extends $grpc.Client {
  static final _$getUser =
      $grpc.ClientMethod<$0.GetUserEventRequest, $0.GetUserEventResponse>(
          '/ProfileService/GetUser',
          ($0.GetUserEventRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetUserEventResponse.fromBuffer(value));

  ProfileServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.GetUserEventResponse> getUser(
      $0.GetUserEventRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getUser, request, options: options);
  }
}

abstract class ProfileServiceBase extends $grpc.Service {
  $core.String get $name => 'ProfileService';

  ProfileServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.GetUserEventRequest, $0.GetUserEventResponse>(
            'GetUser',
            getUser_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.GetUserEventRequest.fromBuffer(value),
            ($0.GetUserEventResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.GetUserEventResponse> getUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.GetUserEventRequest> request) async {
    return getUser(call, await request);
  }

  $async.Future<$0.GetUserEventResponse> getUser(
      $grpc.ServiceCall call, $0.GetUserEventRequest request);
}
