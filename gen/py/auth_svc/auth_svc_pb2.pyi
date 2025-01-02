from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TOTPHasLinkedAccountRequest(_message.Message):
    __slots__ = ("access_token", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class TOTPHasLinkedAccountResponse(_message.Message):
    __slots__ = ("has_linked",)
    HAS_LINKED_FIELD_NUMBER: _ClassVar[int]
    has_linked: bool
    def __init__(self, has_linked: bool = ...) -> None: ...

class TOTPUnlinkAccountRequest(_message.Message):
    __slots__ = ("access_token", "code", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    code: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., code: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class TOTPUnlinkAccountResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class TOTPLinkAccountRequest(_message.Message):
    __slots__ = ("access_token", "code", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    code: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., code: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class TOTPLinkAccountResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class TOTPLinkAccountTmpRequest(_message.Message):
    __slots__ = ("access_token", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class TOTPLinkAccountTmpResponse(_message.Message):
    __slots__ = ("tmp_code",)
    TMP_CODE_FIELD_NUMBER: _ClassVar[int]
    tmp_code: str
    def __init__(self, tmp_code: _Optional[str] = ...) -> None: ...

class CheckTOTPRequest(_message.Message):
    __slots__ = ("access_token", "code", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    code: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., code: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class CheckTOTPResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class AddCustomerToGroupRequest(_message.Message):
    __slots__ = ("app_id", "customer_id", "group_id")
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    customer_id: int
    group_id: int
    def __init__(self, app_id: _Optional[int] = ..., customer_id: _Optional[int] = ..., group_id: _Optional[int] = ...) -> None: ...

class AddCustomerToGroupResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class GetGroupRequest(_message.Message):
    __slots__ = ("app_id",)
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    def __init__(self, app_id: _Optional[int] = ...) -> None: ...

class GetGroupResponse(_message.Message):
    __slots__ = ("groups",)
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[Group]
    def __init__(self, groups: _Optional[_Iterable[_Union[Group, _Mapping]]] = ...) -> None: ...

class AddGroupRequest(_message.Message):
    __slots__ = ("app_id", "name")
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    name: str
    def __init__(self, app_id: _Optional[int] = ..., name: _Optional[str] = ...) -> None: ...

class AddGroupResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class GetActionRequest(_message.Message):
    __slots__ = ("app_id",)
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    def __init__(self, app_id: _Optional[int] = ...) -> None: ...

class GetActionResponse(_message.Message):
    __slots__ = ("actions",)
    ACTIONS_FIELD_NUMBER: _ClassVar[int]
    actions: _containers.RepeatedCompositeFieldContainer[Action]
    def __init__(self, actions: _Optional[_Iterable[_Union[Action, _Mapping]]] = ...) -> None: ...

class AddPermissionCustomerRequest(_message.Message):
    __slots__ = ("app_id", "customer_id", "action_id", "space")
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_ID_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    customer_id: int
    action_id: int
    space: str
    def __init__(self, app_id: _Optional[int] = ..., customer_id: _Optional[int] = ..., action_id: _Optional[int] = ..., space: _Optional[str] = ...) -> None: ...

class AddPermissionCustomerResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class AddPermissionGroupRequest(_message.Message):
    __slots__ = ("app_id", "group_id", "action_id", "space")
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_ID_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    app_id: int
    group_id: int
    action_id: int
    space: str
    def __init__(self, app_id: _Optional[int] = ..., group_id: _Optional[int] = ..., action_id: _Optional[int] = ..., space: _Optional[str] = ...) -> None: ...

class AddPermissionGroupResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class CheckVersionRequest(_message.Message):
    __slots__ = ("version", "app_id", "access_token")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    version: str
    app_id: int
    access_token: str
    def __init__(self, version: _Optional[str] = ..., app_id: _Optional[int] = ..., access_token: _Optional[str] = ...) -> None: ...

class CheckVersionResponse(_message.Message):
    __slots__ = ("need_update", "update_url")
    NEED_UPDATE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_URL_FIELD_NUMBER: _ClassVar[int]
    need_update: bool
    update_url: str
    def __init__(self, need_update: bool = ..., update_url: _Optional[str] = ...) -> None: ...

class HasPasswordRequest(_message.Message):
    __slots__ = ("login", "app_id")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    login: str
    app_id: int
    def __init__(self, login: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class HasPasswordResponse(_message.Message):
    __slots__ = ("has_password",)
    HAS_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    has_password: bool
    def __init__(self, has_password: bool = ...) -> None: ...

class ChangePasswordRequest(_message.Message):
    __slots__ = ("login", "old_password", "new_password", "app_id")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    OLD_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    NEW_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    login: str
    old_password: str
    new_password: str
    app_id: int
    def __init__(self, login: _Optional[str] = ..., old_password: _Optional[str] = ..., new_password: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class ChangePasswordResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class SetPasswordRequest(_message.Message):
    __slots__ = ("login", "password", "app_id")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    login: str
    password: str
    app_id: int
    def __init__(self, login: _Optional[str] = ..., password: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class SetPasswordResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class ResetPasswordRequest(_message.Message):
    __slots__ = ("login", "app_id")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    login: str
    app_id: int
    def __init__(self, login: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class ResetPasswordResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class SignOutRequest(_message.Message):
    __slots__ = ("access_token", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class SignOutResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class SignUpRequest(_message.Message):
    __slots__ = ("login", "password", "app_id", "email")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    login: str
    password: str
    app_id: int
    email: str
    def __init__(self, login: _Optional[str] = ..., password: _Optional[str] = ..., app_id: _Optional[int] = ..., email: _Optional[str] = ...) -> None: ...

class TakePinRequest(_message.Message):
    __slots__ = ("phone_number", "app_id")
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    phone_number: str
    app_id: int
    def __init__(self, phone_number: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class SignUpResponse(_message.Message):
    __slots__ = ("user_id",)
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    def __init__(self, user_id: _Optional[int] = ...) -> None: ...

class RefreshTokenRequest(_message.Message):
    __slots__ = ("refresh_token", "app_id")
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    refresh_token: str
    app_id: int
    def __init__(self, refresh_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class RefreshTokenResponse(_message.Message):
    __slots__ = ("access_token", "refresh_token")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    refresh_token: str
    def __init__(self, access_token: _Optional[str] = ..., refresh_token: _Optional[str] = ...) -> None: ...

class CheckTokenRequest(_message.Message):
    __slots__ = ("access_token", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class CheckTokenResponse(_message.Message):
    __slots__ = ("user_id", "login", "email", "phone_number")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    login: str
    email: str
    phone_number: str
    def __init__(self, user_id: _Optional[int] = ..., login: _Optional[str] = ..., email: _Optional[str] = ..., phone_number: _Optional[str] = ...) -> None: ...

class CheckPermissonRequest(_message.Message):
    __slots__ = ("access_token", "action", "app_id", "space")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    action: str
    app_id: int
    space: str
    def __init__(self, access_token: _Optional[str] = ..., action: _Optional[str] = ..., app_id: _Optional[int] = ..., space: _Optional[str] = ...) -> None: ...

class CheckPermissonResponse(_message.Message):
    __slots__ = ("user_id", "login", "email", "phone_number")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    login: str
    email: str
    phone_number: str
    def __init__(self, user_id: _Optional[int] = ..., login: _Optional[str] = ..., email: _Optional[str] = ..., phone_number: _Optional[str] = ...) -> None: ...

class TakePinResponse(_message.Message):
    __slots__ = ("user_id",)
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: int
    def __init__(self, user_id: _Optional[int] = ...) -> None: ...

class SignInRequest(_message.Message):
    __slots__ = ("login", "password", "app_id", "auth_type", "device_info", "software_info")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    AUTH_TYPE_FIELD_NUMBER: _ClassVar[int]
    DEVICE_INFO_FIELD_NUMBER: _ClassVar[int]
    SOFTWARE_INFO_FIELD_NUMBER: _ClassVar[int]
    login: str
    password: str
    app_id: int
    auth_type: str
    device_info: str
    software_info: str
    def __init__(self, login: _Optional[str] = ..., password: _Optional[str] = ..., app_id: _Optional[int] = ..., auth_type: _Optional[str] = ..., device_info: _Optional[str] = ..., software_info: _Optional[str] = ...) -> None: ...

class SignInResponse(_message.Message):
    __slots__ = ("access_token", "refresh_token")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    refresh_token: str
    def __init__(self, access_token: _Optional[str] = ..., refresh_token: _Optional[str] = ...) -> None: ...

class Group(_message.Message):
    __slots__ = ("id", "name", "app_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    app_id: int
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

class Customer(_message.Message):
    __slots__ = ("id", "login", "password", "app_id", "phone_number", "email")
    ID_FIELD_NUMBER: _ClassVar[int]
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    id: int
    login: str
    password: str
    app_id: int
    phone_number: str
    email: str
    def __init__(self, id: _Optional[int] = ..., login: _Optional[str] = ..., password: _Optional[str] = ..., app_id: _Optional[int] = ..., phone_number: _Optional[str] = ..., email: _Optional[str] = ...) -> None: ...

class Action(_message.Message):
    __slots__ = ("id", "action", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: int
    action: str
    name: str
    def __init__(self, id: _Optional[int] = ..., action: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class PermissionsGroup(_message.Message):
    __slots__ = ("id", "app_id", "group_id", "action_id", "action", "space")
    ID_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    id: int
    app_id: int
    group_id: int
    action_id: int
    action: Action
    space: str
    def __init__(self, id: _Optional[int] = ..., app_id: _Optional[int] = ..., group_id: _Optional[int] = ..., action_id: _Optional[int] = ..., action: _Optional[_Union[Action, _Mapping]] = ..., space: _Optional[str] = ...) -> None: ...

class PermissionsCustomer(_message.Message):
    __slots__ = ("id", "app_id", "customer_id", "action_id", "action", "space")
    ID_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_ID_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    id: int
    app_id: int
    customer_id: int
    action_id: int
    action: Action
    space: str
    def __init__(self, id: _Optional[int] = ..., app_id: _Optional[int] = ..., customer_id: _Optional[int] = ..., action_id: _Optional[int] = ..., action: _Optional[_Union[Action, _Mapping]] = ..., space: _Optional[str] = ...) -> None: ...

class CustomerInfoResponse(_message.Message):
    __slots__ = ("customer", "group_ids", "permissions_customers", "permissions_groups")
    CUSTOMER_FIELD_NUMBER: _ClassVar[int]
    GROUP_IDS_FIELD_NUMBER: _ClassVar[int]
    PERMISSIONS_CUSTOMERS_FIELD_NUMBER: _ClassVar[int]
    PERMISSIONS_GROUPS_FIELD_NUMBER: _ClassVar[int]
    customer: Customer
    group_ids: _containers.RepeatedScalarFieldContainer[int]
    permissions_customers: _containers.RepeatedCompositeFieldContainer[PermissionsCustomer]
    permissions_groups: _containers.RepeatedCompositeFieldContainer[PermissionsGroup]
    def __init__(self, customer: _Optional[_Union[Customer, _Mapping]] = ..., group_ids: _Optional[_Iterable[int]] = ..., permissions_customers: _Optional[_Iterable[_Union[PermissionsCustomer, _Mapping]]] = ..., permissions_groups: _Optional[_Iterable[_Union[PermissionsGroup, _Mapping]]] = ...) -> None: ...

class CustomerInfoRequest(_message.Message):
    __slots__ = ("access_token", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...
