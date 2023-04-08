// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/pytorch.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fpytorch_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fpytorch_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fpytorch_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fpytorch_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fpytorch_2eproto();
namespace flyteidl {
namespace plugins {
class DistributedPyTorchTrainingTask;
class DistributedPyTorchTrainingTaskDefaultTypeInternal;
extern DistributedPyTorchTrainingTaskDefaultTypeInternal _DistributedPyTorchTrainingTask_default_instance_;
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::DistributedPyTorchTrainingTask* Arena::CreateMaybeMessage<::flyteidl::plugins::DistributedPyTorchTrainingTask>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {

// ===================================================================

class DistributedPyTorchTrainingTask final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.DistributedPyTorchTrainingTask) */ {
 public:
  DistributedPyTorchTrainingTask();
  virtual ~DistributedPyTorchTrainingTask();

  DistributedPyTorchTrainingTask(const DistributedPyTorchTrainingTask& from);

  inline DistributedPyTorchTrainingTask& operator=(const DistributedPyTorchTrainingTask& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DistributedPyTorchTrainingTask(DistributedPyTorchTrainingTask&& from) noexcept
    : DistributedPyTorchTrainingTask() {
    *this = ::std::move(from);
  }

  inline DistributedPyTorchTrainingTask& operator=(DistributedPyTorchTrainingTask&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const DistributedPyTorchTrainingTask& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DistributedPyTorchTrainingTask* internal_default_instance() {
    return reinterpret_cast<const DistributedPyTorchTrainingTask*>(
               &_DistributedPyTorchTrainingTask_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(DistributedPyTorchTrainingTask* other);
  friend void swap(DistributedPyTorchTrainingTask& a, DistributedPyTorchTrainingTask& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DistributedPyTorchTrainingTask* New() const final {
    return CreateMaybeMessage<DistributedPyTorchTrainingTask>(nullptr);
  }

  DistributedPyTorchTrainingTask* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<DistributedPyTorchTrainingTask>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const DistributedPyTorchTrainingTask& from);
  void MergeFrom(const DistributedPyTorchTrainingTask& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(DistributedPyTorchTrainingTask* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string RDZVBackend = 2;
  void clear_rdzvbackend();
  static const int kRDZVBackendFieldNumber = 2;
  const ::std::string& rdzvbackend() const;
  void set_rdzvbackend(const ::std::string& value);
  #if LANG_CXX11
  void set_rdzvbackend(::std::string&& value);
  #endif
  void set_rdzvbackend(const char* value);
  void set_rdzvbackend(const char* value, size_t size);
  ::std::string* mutable_rdzvbackend();
  ::std::string* release_rdzvbackend();
  void set_allocated_rdzvbackend(::std::string* rdzvbackend);

  // int32 workers = 1;
  void clear_workers();
  static const int kWorkersFieldNumber = 1;
  ::google::protobuf::int32 workers() const;
  void set_workers(::google::protobuf::int32 value);

  // int32 minReplicas = 3;
  void clear_minreplicas();
  static const int kMinReplicasFieldNumber = 3;
  ::google::protobuf::int32 minreplicas() const;
  void set_minreplicas(::google::protobuf::int32 value);

  // int32 maxReplicas = 4;
  void clear_maxreplicas();
  static const int kMaxReplicasFieldNumber = 4;
  ::google::protobuf::int32 maxreplicas() const;
  void set_maxreplicas(::google::protobuf::int32 value);

  // int32 nProcPerNode = 5;
  void clear_nprocpernode();
  static const int kNProcPerNodeFieldNumber = 5;
  ::google::protobuf::int32 nprocpernode() const;
  void set_nprocpernode(::google::protobuf::int32 value);

  // int32 maxRestarts = 6;
  void clear_maxrestarts();
  static const int kMaxRestartsFieldNumber = 6;
  ::google::protobuf::int32 maxrestarts() const;
  void set_maxrestarts(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.DistributedPyTorchTrainingTask)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr rdzvbackend_;
  ::google::protobuf::int32 workers_;
  ::google::protobuf::int32 minreplicas_;
  ::google::protobuf::int32 maxreplicas_;
  ::google::protobuf::int32 nprocpernode_;
  ::google::protobuf::int32 maxrestarts_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fpytorch_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// DistributedPyTorchTrainingTask

// int32 workers = 1;
inline void DistributedPyTorchTrainingTask::clear_workers() {
  workers_ = 0;
}
inline ::google::protobuf::int32 DistributedPyTorchTrainingTask::workers() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.workers)
  return workers_;
}
inline void DistributedPyTorchTrainingTask::set_workers(::google::protobuf::int32 value) {
  
  workers_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.workers)
}

// string RDZVBackend = 2;
inline void DistributedPyTorchTrainingTask::clear_rdzvbackend() {
  rdzvbackend_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& DistributedPyTorchTrainingTask::rdzvbackend() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
  return rdzvbackend_.GetNoArena();
}
inline void DistributedPyTorchTrainingTask::set_rdzvbackend(const ::std::string& value) {
  
  rdzvbackend_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
}
#if LANG_CXX11
inline void DistributedPyTorchTrainingTask::set_rdzvbackend(::std::string&& value) {
  
  rdzvbackend_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
}
#endif
inline void DistributedPyTorchTrainingTask::set_rdzvbackend(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  rdzvbackend_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
}
inline void DistributedPyTorchTrainingTask::set_rdzvbackend(const char* value, size_t size) {
  
  rdzvbackend_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
}
inline ::std::string* DistributedPyTorchTrainingTask::mutable_rdzvbackend() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
  return rdzvbackend_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* DistributedPyTorchTrainingTask::release_rdzvbackend() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
  
  return rdzvbackend_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void DistributedPyTorchTrainingTask::set_allocated_rdzvbackend(::std::string* rdzvbackend) {
  if (rdzvbackend != nullptr) {
    
  } else {
    
  }
  rdzvbackend_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), rdzvbackend);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.DistributedPyTorchTrainingTask.RDZVBackend)
}

// int32 minReplicas = 3;
inline void DistributedPyTorchTrainingTask::clear_minreplicas() {
  minreplicas_ = 0;
}
inline ::google::protobuf::int32 DistributedPyTorchTrainingTask::minreplicas() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.minReplicas)
  return minreplicas_;
}
inline void DistributedPyTorchTrainingTask::set_minreplicas(::google::protobuf::int32 value) {
  
  minreplicas_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.minReplicas)
}

// int32 maxReplicas = 4;
inline void DistributedPyTorchTrainingTask::clear_maxreplicas() {
  maxreplicas_ = 0;
}
inline ::google::protobuf::int32 DistributedPyTorchTrainingTask::maxreplicas() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.maxReplicas)
  return maxreplicas_;
}
inline void DistributedPyTorchTrainingTask::set_maxreplicas(::google::protobuf::int32 value) {
  
  maxreplicas_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.maxReplicas)
}

// int32 nProcPerNode = 5;
inline void DistributedPyTorchTrainingTask::clear_nprocpernode() {
  nprocpernode_ = 0;
}
inline ::google::protobuf::int32 DistributedPyTorchTrainingTask::nprocpernode() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.nProcPerNode)
  return nprocpernode_;
}
inline void DistributedPyTorchTrainingTask::set_nprocpernode(::google::protobuf::int32 value) {
  
  nprocpernode_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.nProcPerNode)
}

// int32 maxRestarts = 6;
inline void DistributedPyTorchTrainingTask::clear_maxrestarts() {
  maxrestarts_ = 0;
}
inline ::google::protobuf::int32 DistributedPyTorchTrainingTask::maxrestarts() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedPyTorchTrainingTask.maxRestarts)
  return maxrestarts_;
}
inline void DistributedPyTorchTrainingTask::set_maxrestarts(::google::protobuf::int32 value) {
  
  maxrestarts_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedPyTorchTrainingTask.maxRestarts)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace plugins
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fpytorch_2eproto
