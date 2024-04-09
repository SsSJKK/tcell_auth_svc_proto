from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ChangePasswordRequest(_message.Message):
    __slots__ = ("old_password", "new_password", "app_id")
    OLD_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    NEW_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    old_password: str
    new_password: str
    app_id: int
    def __init__(self, old_password: _Optional[str] = ..., new_password: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

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
    __slots__ = ("access_token", "action", "app_id")
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    APP_ID_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    action: str
    app_id: int
    def __init__(self, access_token: _Optional[str] = ..., action: _Optional[str] = ..., app_id: _Optional[int] = ...) -> None: ...

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
