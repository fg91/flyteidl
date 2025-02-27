from flyteidl.core import execution_pb2 as _execution_pb2
from flyteidl.core import identifier_pb2 as _identifier_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Waitable(_message.Message):
    __slots__ = ["phase", "wf_exec_id", "workflow_id"]
    PHASE_FIELD_NUMBER: _ClassVar[int]
    WF_EXEC_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    phase: _execution_pb2.WorkflowExecution.Phase
    wf_exec_id: _identifier_pb2.WorkflowExecutionIdentifier
    workflow_id: str
    def __init__(self, wf_exec_id: _Optional[_Union[_identifier_pb2.WorkflowExecutionIdentifier, _Mapping]] = ..., phase: _Optional[_Union[_execution_pb2.WorkflowExecution.Phase, str]] = ..., workflow_id: _Optional[str] = ...) -> None: ...
