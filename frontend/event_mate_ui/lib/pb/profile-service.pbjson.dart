///
//  Generated code. Do not modify.
//  source: profile-service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use getUserEventRequestDescriptor instead')
const GetUserEventRequest$json = const {
  '1': 'GetUserEventRequest',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 5, '10': 'userId'},
  ],
};

/// Descriptor for `GetUserEventRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserEventRequestDescriptor = $convert.base64Decode('ChNHZXRVc2VyRXZlbnRSZXF1ZXN0EhYKBnVzZXJJZBgBIAEoBVIGdXNlcklk');
@$core.Deprecated('Use userDescriptor instead')
const User$json = const {
  '1': 'User',
  '2': const [
    const {'1': 'UserId', '3': 1, '4': 1, '5': 5, '10': 'UserId'},
    const {'1': 'Name', '3': 2, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'LastName', '3': 3, '4': 1, '5': 9, '10': 'LastName'},
    const {'1': 'About', '3': 4, '4': 1, '5': 9, '10': 'About'},
    const {'1': 'Photo', '3': 5, '4': 1, '5': 9, '10': 'Photo'},
    const {'1': 'AttandedEvents', '3': 6, '4': 3, '5': 11, '6': '.Event', '10': 'AttandedEvents'},
    const {'1': 'Adress', '3': 7, '4': 1, '5': 11, '6': '.UserProfileAdress', '10': 'Adress'},
  ],
};

/// Descriptor for `User`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDescriptor = $convert.base64Decode('CgRVc2VyEhYKBlVzZXJJZBgBIAEoBVIGVXNlcklkEhIKBE5hbWUYAiABKAlSBE5hbWUSGgoITGFzdE5hbWUYAyABKAlSCExhc3ROYW1lEhQKBUFib3V0GAQgASgJUgVBYm91dBIUCgVQaG90bxgFIAEoCVIFUGhvdG8SLgoOQXR0YW5kZWRFdmVudHMYBiADKAsyBi5FdmVudFIOQXR0YW5kZWRFdmVudHMSKgoGQWRyZXNzGAcgASgLMhIuVXNlclByb2ZpbGVBZHJlc3NSBkFkcmVzcw==');
@$core.Deprecated('Use userProfileAdressDescriptor instead')
const UserProfileAdress$json = const {
  '1': 'UserProfileAdress',
  '2': const [
    const {'1': 'UserId', '3': 1, '4': 1, '5': 9, '10': 'UserId'},
    const {'1': 'City', '3': 2, '4': 1, '5': 9, '10': 'City'},
  ],
};

/// Descriptor for `UserProfileAdress`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userProfileAdressDescriptor = $convert.base64Decode('ChFVc2VyUHJvZmlsZUFkcmVzcxIWCgZVc2VySWQYASABKAlSBlVzZXJJZBISCgRDaXR5GAIgASgJUgRDaXR5');
@$core.Deprecated('Use eventDescriptor instead')
const Event$json = const {
  '1': 'Event',
  '2': const [
    const {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    const {'1': 'CoverPhoto', '3': 2, '4': 1, '5': 9, '10': 'CoverPhoto'},
  ],
};

/// Descriptor for `Event`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List eventDescriptor = $convert.base64Decode('CgVFdmVudBISCgROYW1lGAEgASgJUgROYW1lEh4KCkNvdmVyUGhvdG8YAiABKAlSCkNvdmVyUGhvdG8=');
@$core.Deprecated('Use getUserEventResponseDescriptor instead')
const GetUserEventResponse$json = const {
  '1': 'GetUserEventResponse',
  '2': const [
    const {'1': 'user', '3': 1, '4': 1, '5': 11, '6': '.User', '10': 'user'},
  ],
};

/// Descriptor for `GetUserEventResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserEventResponseDescriptor = $convert.base64Decode('ChRHZXRVc2VyRXZlbnRSZXNwb25zZRIZCgR1c2VyGAEgASgLMgUuVXNlclIEdXNlcg==');
