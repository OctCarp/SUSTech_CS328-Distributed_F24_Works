from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LogLevel(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INFO: _ClassVar[LogLevel]
    WARNING: _ClassVar[LogLevel]
    ERROR: _ClassVar[LogLevel]
    DEBUG: _ClassVar[LogLevel]
INFO: LogLevel
WARNING: LogLevel
ERROR: LogLevel
DEBUG: LogLevel

class LogMessage(_message.Message):
    __slots__ = ("level", "service_name", "message", "timestamp", "trace_id")
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    SERVICE_NAME_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TRACE_ID_FIELD_NUMBER: _ClassVar[int]
    level: LogLevel
    service_name: str
    message: str
    timestamp: int
    trace_id: str
    def __init__(self, level: _Optional[_Union[LogLevel, str]] = ..., service_name: _Optional[str] = ..., message: _Optional[str] = ..., timestamp: _Optional[int] = ..., trace_id: _Optional[str] = ...) -> None: ...

class LogResponse(_message.Message):
    __slots__ = ("success", "message")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    success: bool
    message: str
    def __init__(self, success: bool = ..., message: _Optional[str] = ...) -> None: ...
