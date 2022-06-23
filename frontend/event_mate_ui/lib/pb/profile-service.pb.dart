///
//  Generated code. Do not modify.
//  source: profile-service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetUserEventRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'GetUserEventRequest',
      createEmptyInstance: create)
    ..a<$core.int>(
        1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'userId',
        $pb.PbFieldType.O3,
        protoName: 'userId')
    ..hasRequiredFields = false;

  GetUserEventRequest._() : super();
  factory GetUserEventRequest({
    $core.int? userId,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    return _result;
  }
  factory GetUserEventRequest.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory GetUserEventRequest.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  GetUserEventRequest clone() => GetUserEventRequest()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  GetUserEventRequest copyWith(void Function(GetUserEventRequest) updates) =>
      super.copyWith((message) => updates(message as GetUserEventRequest))
          as GetUserEventRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetUserEventRequest create() => GetUserEventRequest._();
  GetUserEventRequest createEmptyInstance() => create();
  static $pb.PbList<GetUserEventRequest> createRepeated() =>
      $pb.PbList<GetUserEventRequest>();
  @$core.pragma('dart2js:noInline')
  static GetUserEventRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<GetUserEventRequest>(create);
  static GetUserEventRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get userId => $_getIZ(0);
  @$pb.TagNumber(1)
  set userId($core.int v) {
    $_setSignedInt32(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);
}

class User extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'User',
      createEmptyInstance: create)
    ..a<$core.int>(
        1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'UserId',
        $pb.PbFieldType.O3,
        protoName: 'UserId')
    ..aOS(
        2,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'Name',
        protoName: 'Name')
    ..aOS(
        3,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'LastName',
        protoName: 'LastName')
    ..aOS(
        4,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'About',
        protoName: 'About')
    ..aOS(
        5,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'Photo',
        protoName: 'Photo')
    ..pc<Event>(
        6,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'AttandedEvents',
        $pb.PbFieldType.PM,
        protoName: 'AttandedEvents',
        subBuilder: Event.create)
    ..aOM<UserProfileAdress>(
        7,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'Adress',
        protoName: 'Adress',
        subBuilder: UserProfileAdress.create)
    ..hasRequiredFields = false;

  User._() : super();
  factory User({
    $core.int? userId,
    $core.String? name,
    $core.String? lastName,
    $core.String? about,
    $core.String? photo,
    $core.Iterable<Event>? attandedEvents,
    UserProfileAdress? adress,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (name != null) {
      _result.name = name;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    if (about != null) {
      _result.about = about;
    }
    if (photo != null) {
      _result.photo = photo;
    }
    if (attandedEvents != null) {
      _result.attandedEvents.addAll(attandedEvents);
    }
    if (adress != null) {
      _result.adress = adress;
    }
    return _result;
  }
  factory User.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory User.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  User clone() => User()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  User copyWith(void Function(User) updates) =>
      super.copyWith((message) => updates(message as User))
          as User; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static User create() => User._();
  User createEmptyInstance() => create();
  static $pb.PbList<User> createRepeated() => $pb.PbList<User>();
  @$core.pragma('dart2js:noInline')
  static User getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<User>(create);
  static User? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get userId => $_getIZ(0);
  @$pb.TagNumber(1)
  set userId($core.int v) {
    $_setSignedInt32(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lastName => $_getSZ(2);
  @$pb.TagNumber(3)
  set lastName($core.String v) {
    $_setString(2, v);
  }

  @$pb.TagNumber(3)
  $core.bool hasLastName() => $_has(2);
  @$pb.TagNumber(3)
  void clearLastName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get about => $_getSZ(3);
  @$pb.TagNumber(4)
  set about($core.String v) {
    $_setString(3, v);
  }

  @$pb.TagNumber(4)
  $core.bool hasAbout() => $_has(3);
  @$pb.TagNumber(4)
  void clearAbout() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get photo => $_getSZ(4);
  @$pb.TagNumber(5)
  set photo($core.String v) {
    $_setString(4, v);
  }

  @$pb.TagNumber(5)
  $core.bool hasPhoto() => $_has(4);
  @$pb.TagNumber(5)
  void clearPhoto() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<Event> get attandedEvents => $_getList(5);

  @$pb.TagNumber(7)
  UserProfileAdress get adress => $_getN(6);
  @$pb.TagNumber(7)
  set adress(UserProfileAdress v) {
    setField(7, v);
  }

  @$pb.TagNumber(7)
  $core.bool hasAdress() => $_has(6);
  @$pb.TagNumber(7)
  void clearAdress() => clearField(7);
  @$pb.TagNumber(7)
  UserProfileAdress ensureAdress() => $_ensure(6);
}

class UserProfileAdress extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'UserProfileAdress',
      createEmptyInstance: create)
    ..aOS(
        1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'UserId',
        protoName: 'UserId')
    ..aOS(
        2,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'City',
        protoName: 'City')
    ..hasRequiredFields = false;

  UserProfileAdress._() : super();
  factory UserProfileAdress({
    $core.String? userId,
    $core.String? city,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (city != null) {
      _result.city = city;
    }
    return _result;
  }
  factory UserProfileAdress.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory UserProfileAdress.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserProfileAdress clone() => UserProfileAdress()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserProfileAdress copyWith(void Function(UserProfileAdress) updates) =>
      super.copyWith((message) => updates(message as UserProfileAdress))
          as UserProfileAdress; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserProfileAdress create() => UserProfileAdress._();
  UserProfileAdress createEmptyInstance() => create();
  static $pb.PbList<UserProfileAdress> createRepeated() =>
      $pb.PbList<UserProfileAdress>();
  @$core.pragma('dart2js:noInline')
  static UserProfileAdress getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserProfileAdress>(create);
  static UserProfileAdress? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get city => $_getSZ(1);
  @$pb.TagNumber(2)
  set city($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasCity() => $_has(1);
  @$pb.TagNumber(2)
  void clearCity() => clearField(2);
}

class Event extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'Event',
      createEmptyInstance: create)
    ..aOS(
        1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'Name',
        protoName: 'Name')
    ..aOS(
        2,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'CoverPhoto',
        protoName: 'CoverPhoto')
    ..hasRequiredFields = false;

  Event._() : super();
  factory Event({
    $core.String? name,
    $core.String? coverPhoto,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (coverPhoto != null) {
      _result.coverPhoto = coverPhoto;
    }
    return _result;
  }
  factory Event.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory Event.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Event clone() => Event()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Event copyWith(void Function(Event) updates) =>
      super.copyWith((message) => updates(message as Event))
          as Event; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Event create() => Event._();
  Event createEmptyInstance() => create();
  static $pb.PbList<Event> createRepeated() => $pb.PbList<Event>();
  @$core.pragma('dart2js:noInline')
  static Event getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Event>(create);
  static Event? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get coverPhoto => $_getSZ(1);
  @$pb.TagNumber(2)
  set coverPhoto($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasCoverPhoto() => $_has(1);
  @$pb.TagNumber(2)
  void clearCoverPhoto() => clearField(2);
}

class GetUserEventResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'GetUserEventResponse',
      createEmptyInstance: create)
    ..aOM<User>(
        1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'user',
        subBuilder: User.create)
    ..hasRequiredFields = false;

  GetUserEventResponse._() : super();
  factory GetUserEventResponse({
    User? user,
  }) {
    final _result = create();
    if (user != null) {
      _result.user = user;
    }
    return _result;
  }
  factory GetUserEventResponse.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory GetUserEventResponse.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  GetUserEventResponse clone() =>
      GetUserEventResponse()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  GetUserEventResponse copyWith(void Function(GetUserEventResponse) updates) =>
      super.copyWith((message) => updates(message as GetUserEventResponse))
          as GetUserEventResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetUserEventResponse create() => GetUserEventResponse._();
  GetUserEventResponse createEmptyInstance() => create();
  static $pb.PbList<GetUserEventResponse> createRepeated() =>
      $pb.PbList<GetUserEventResponse>();
  @$core.pragma('dart2js:noInline')
  static GetUserEventResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<GetUserEventResponse>(create);
  static GetUserEventResponse? _defaultInstance;

  @$pb.TagNumber(1)
  User get user => $_getN(0);
  @$pb.TagNumber(1)
  set user(User v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearUser() => clearField(1);
  @$pb.TagNumber(1)
  User ensureUser() => $_ensure(0);
}
