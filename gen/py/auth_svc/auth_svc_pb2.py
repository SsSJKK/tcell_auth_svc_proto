# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth_svc/auth_svc.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17\x61uth_svc/auth_svc.proto\x12\x04\x61uth\"6\n\x0eSignOutRequest\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x02 \x01(\x05\"\"\n\x0fSignOutResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"O\n\rSignUpRequest\x12\r\n\x05login\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x03 \x01(\x05\x12\r\n\x05\x65mail\x18\x04 \x01(\t\"6\n\x0eTakePinRequest\x12\x14\n\x0cphone_number\x18\x01 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x02 \x01(\x05\"!\n\x0eSignUpResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\"<\n\x13RefreshTokenRequest\x12\x15\n\rrefresh_token\x18\x01 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x02 \x01(\x05\"C\n\x14RefreshTokenResponse\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t\x12\x15\n\rrefresh_token\x18\x02 \x01(\t\"9\n\x11\x43heckTokenRequest\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x02 \x01(\x05\"Y\n\x12\x43heckTokenResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\r\n\x05login\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x14\n\x0cphone_number\x18\x04 \x01(\t\"M\n\x15\x43heckPermissonRequest\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t\x12\x0e\n\x06\x61\x63tion\x18\x02 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x03 \x01(\x05\"]\n\x16\x43heckPermissonResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\r\n\x05login\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x14\n\x0cphone_number\x18\x04 \x01(\t\"\"\n\x0fTakePinResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\"\x7f\n\rSignInRequest\x12\r\n\x05login\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0e\n\x06\x61pp_id\x18\x03 \x01(\x05\x12\x11\n\tauth_type\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65vice_info\x18\x05 \x01(\t\x12\x15\n\rsoftware_info\x18\x06 \x01(\t\"=\n\x0eSignInResponse\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t\x12\x15\n\rrefresh_token\x18\x02 \x01(\t2\xb5\x03\n\x04\x41uth\x12\x33\n\x06SignUp\x12\x13.auth.SignUpRequest\x1a\x14.auth.SignUpResponse\x12\x33\n\x06SignIn\x12\x13.auth.SignInRequest\x1a\x14.auth.SignInResponse\x12\x45\n\x0cRefreshToken\x12\x19.auth.RefreshTokenRequest\x1a\x1a.auth.RefreshTokenResponse\x12?\n\nCheckToken\x12\x17.auth.CheckTokenRequest\x1a\x18.auth.CheckTokenResponse\x12K\n\x0e\x43heckPermisson\x12\x1b.auth.CheckPermissonRequest\x1a\x1c.auth.CheckPermissonResponse\x12\x36\n\x07TakePin\x12\x14.auth.TakePinRequest\x1a\x15.auth.TakePinResponse\x12\x36\n\x07SignOut\x12\x14.auth.SignOutRequest\x1a\x15.auth.SignOutResponseB\x16Z\x14sssjkk.auth.v1;auth1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'auth_svc.auth_svc_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\024sssjkk.auth.v1;auth1'
  _globals['_SIGNOUTREQUEST']._serialized_start=33
  _globals['_SIGNOUTREQUEST']._serialized_end=87
  _globals['_SIGNOUTRESPONSE']._serialized_start=89
  _globals['_SIGNOUTRESPONSE']._serialized_end=123
  _globals['_SIGNUPREQUEST']._serialized_start=125
  _globals['_SIGNUPREQUEST']._serialized_end=204
  _globals['_TAKEPINREQUEST']._serialized_start=206
  _globals['_TAKEPINREQUEST']._serialized_end=260
  _globals['_SIGNUPRESPONSE']._serialized_start=262
  _globals['_SIGNUPRESPONSE']._serialized_end=295
  _globals['_REFRESHTOKENREQUEST']._serialized_start=297
  _globals['_REFRESHTOKENREQUEST']._serialized_end=357
  _globals['_REFRESHTOKENRESPONSE']._serialized_start=359
  _globals['_REFRESHTOKENRESPONSE']._serialized_end=426
  _globals['_CHECKTOKENREQUEST']._serialized_start=428
  _globals['_CHECKTOKENREQUEST']._serialized_end=485
  _globals['_CHECKTOKENRESPONSE']._serialized_start=487
  _globals['_CHECKTOKENRESPONSE']._serialized_end=576
  _globals['_CHECKPERMISSONREQUEST']._serialized_start=578
  _globals['_CHECKPERMISSONREQUEST']._serialized_end=655
  _globals['_CHECKPERMISSONRESPONSE']._serialized_start=657
  _globals['_CHECKPERMISSONRESPONSE']._serialized_end=750
  _globals['_TAKEPINRESPONSE']._serialized_start=752
  _globals['_TAKEPINRESPONSE']._serialized_end=786
  _globals['_SIGNINREQUEST']._serialized_start=788
  _globals['_SIGNINREQUEST']._serialized_end=915
  _globals['_SIGNINRESPONSE']._serialized_start=917
  _globals['_SIGNINRESPONSE']._serialized_end=978
  _globals['_AUTH']._serialized_start=981
  _globals['_AUTH']._serialized_end=1418
# @@protoc_insertion_point(module_scope)