from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SparkApplication(_message.Message):
    __slots__ = []
    class Type(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    JAVA: SparkApplication.Type
    PYTHON: SparkApplication.Type
    R: SparkApplication.Type
    SCALA: SparkApplication.Type
    def __init__(self) -> None: ...

class SparkJob(_message.Message):
    __slots__ = ["applicationType", "databricksConf", "databricksInstance", "databricksToken", "executorPath", "hadoopConf", "mainApplicationFile", "mainClass", "sparkConf"]
    class HadoopConfEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    class SparkConfEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    APPLICATIONTYPE_FIELD_NUMBER: _ClassVar[int]
    DATABRICKSCONF_FIELD_NUMBER: _ClassVar[int]
    DATABRICKSINSTANCE_FIELD_NUMBER: _ClassVar[int]
    DATABRICKSTOKEN_FIELD_NUMBER: _ClassVar[int]
    EXECUTORPATH_FIELD_NUMBER: _ClassVar[int]
    HADOOPCONF_FIELD_NUMBER: _ClassVar[int]
    MAINAPPLICATIONFILE_FIELD_NUMBER: _ClassVar[int]
    MAINCLASS_FIELD_NUMBER: _ClassVar[int]
    SPARKCONF_FIELD_NUMBER: _ClassVar[int]
    applicationType: SparkApplication.Type
    databricksConf: _struct_pb2.Struct
    databricksInstance: str
    databricksToken: str
    executorPath: str
    hadoopConf: _containers.ScalarMap[str, str]
    mainApplicationFile: str
    mainClass: str
    sparkConf: _containers.ScalarMap[str, str]
    def __init__(self, applicationType: _Optional[_Union[SparkApplication.Type, str]] = ..., mainApplicationFile: _Optional[str] = ..., mainClass: _Optional[str] = ..., sparkConf: _Optional[_Mapping[str, str]] = ..., hadoopConf: _Optional[_Mapping[str, str]] = ..., executorPath: _Optional[str] = ..., databricksConf: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., databricksToken: _Optional[str] = ..., databricksInstance: _Optional[str] = ...) -> None: ...
