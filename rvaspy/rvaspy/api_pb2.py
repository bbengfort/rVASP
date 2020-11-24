# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='api.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\tapi.proto\x12\x02pb\"&\n\x05\x45rror\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\"T\n\x08Identity\x12\x16\n\x0ewallet_address\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x0f\n\x07ivms101\x18\x03 \x01(\t\x12\x10\n\x08provider\x18\x04 \x01(\t\"\xa8\x01\n\x0bTransaction\x12\x0f\n\x07\x61\x63\x63ount\x18\x01 \x01(\t\x12 \n\noriginator\x18\x02 \x01(\x0b\x32\x0c.pb.Identity\x12!\n\x0b\x62\x65neficiary\x18\x03 \x01(\x0b\x32\x0c.pb.Identity\x12\x0e\n\x06\x61mount\x18\x04 \x01(\x02\x12\r\n\x05\x64\x65\x62it\x18\x05 \x01(\x08\x12\x11\n\tcompleted\x18\x06 \x01(\x08\x12\x11\n\ttimestamp\x18\x07 \x01(\t\"G\n\x0fTransferRequest\x12\x0f\n\x07\x61\x63\x63ount\x18\x01 \x01(\t\x12\x13\n\x0b\x62\x65neficiary\x18\x02 \x01(\t\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x02\"O\n\rTransferReply\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.pb.Error\x12$\n\x0btransaction\x18\x02 \x01(\x0b\x32\x0f.pb.Transaction\"Z\n\x0e\x41\x63\x63ountRequest\x12\x0f\n\x07\x61\x63\x63ount\x18\x01 \x01(\t\x12\x17\n\x0fno_transactions\x18\x02 \x01(\x08\x12\x0c\n\x04page\x18\x03 \x01(\r\x12\x10\n\x08per_page\x18\x04 \x01(\r\"\xb9\x01\n\x0c\x41\x63\x63ountReply\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.pb.Error\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x16\n\x0ewallet_address\x18\x04 \x01(\t\x12\x0f\n\x07\x62\x61lance\x18\x05 \x01(\x02\x12\x11\n\tcompleted\x18\x06 \x01(\x04\x12\x0f\n\x07pending\x18\x07 \x01(\x04\x12%\n\x0ctransactions\x18\x08 \x03(\x0b\x32\x0f.pb.Transaction\"\x97\x01\n\x07\x43ommand\x12\x15\n\x04type\x18\x01 \x01(\x0e\x32\x07.pb.RPC\x12\n\n\x02id\x18\x02 \x01(\x04\x12\x0e\n\x06\x63lient\x18\x03 \x01(\t\x12\'\n\x08transfer\x18\x0b \x01(\x0b\x32\x13.pb.TransferRequestH\x00\x12%\n\x07\x61\x63\x63ount\x18\x0c \x01(\x0b\x32\x12.pb.AccountRequestH\x00\x42\t\n\x07request\"\xcb\x01\n\x07Message\x12\x15\n\x04type\x18\x01 \x01(\x0e\x32\x07.pb.RPC\x12\n\n\x02id\x18\x02 \x01(\x04\x12\x0e\n\x06update\x18\x03 \x01(\t\x12\x11\n\ttimestamp\x18\x04 \x01(\t\x12%\n\x08\x63\x61tegory\x18\x05 \x01(\x0e\x32\x13.pb.MessageCategory\x12%\n\x08transfer\x18\x0b \x01(\x0b\x32\x11.pb.TransferReplyH\x00\x12#\n\x07\x61\x63\x63ount\x18\x0c \x01(\x0b\x32\x10.pb.AccountReplyH\x00\x42\x07\n\x05reply*+\n\x03RPC\x12\t\n\x05NORPC\x10\x00\x12\x0c\n\x08TRANSFER\x10\x01\x12\x0b\n\x07\x41\x43\x43OUNT\x10\x02*S\n\x0fMessageCategory\x12\n\n\x06LEDGER\x10\x00\x12\x0b\n\x07TRISADS\x10\x01\x12\x0c\n\x08TRISAP2P\x10\x02\x12\x0e\n\nBLOCKCHAIN\x10\x03\x12\t\n\x05\x45RROR\x10\x04\x32\x38\n\tTRISADemo\x12+\n\x0bLiveUpdates\x12\x0b.pb.Command\x1a\x0b.pb.Message(\x01\x30\x01\x32}\n\x10TRISAIntegration\x12\x32\n\x08Transfer\x12\x13.pb.TransferRequest\x1a\x11.pb.TransferReply\x12\x35\n\rAccountStatus\x12\x12.pb.AccountRequest\x1a\x10.pb.AccountReplyb\x06proto3'
)

_RPC = _descriptor.EnumDescriptor(
  name='RPC',
  full_name='pb.RPC',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NORPC', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TRANSFER', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ACCOUNT', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1108,
  serialized_end=1151,
)
_sym_db.RegisterEnumDescriptor(_RPC)

RPC = enum_type_wrapper.EnumTypeWrapper(_RPC)
_MESSAGECATEGORY = _descriptor.EnumDescriptor(
  name='MessageCategory',
  full_name='pb.MessageCategory',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='LEDGER', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TRISADS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TRISAP2P', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='BLOCKCHAIN', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ERROR', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1153,
  serialized_end=1236,
)
_sym_db.RegisterEnumDescriptor(_MESSAGECATEGORY)

MessageCategory = enum_type_wrapper.EnumTypeWrapper(_MESSAGECATEGORY)
NORPC = 0
TRANSFER = 1
ACCOUNT = 2
LEDGER = 0
TRISADS = 1
TRISAP2P = 2
BLOCKCHAIN = 3
ERROR = 4



_ERROR = _descriptor.Descriptor(
  name='Error',
  full_name='pb.Error',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='pb.Error.code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='pb.Error.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=17,
  serialized_end=55,
)


_IDENTITY = _descriptor.Descriptor(
  name='Identity',
  full_name='pb.Identity',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='wallet_address', full_name='pb.Identity.wallet_address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='email', full_name='pb.Identity.email', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ivms101', full_name='pb.Identity.ivms101', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='provider', full_name='pb.Identity.provider', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=141,
)


_TRANSACTION = _descriptor.Descriptor(
  name='Transaction',
  full_name='pb.Transaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account', full_name='pb.Transaction.account', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='originator', full_name='pb.Transaction.originator', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='beneficiary', full_name='pb.Transaction.beneficiary', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='amount', full_name='pb.Transaction.amount', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='debit', full_name='pb.Transaction.debit', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='completed', full_name='pb.Transaction.completed', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='pb.Transaction.timestamp', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=144,
  serialized_end=312,
)


_TRANSFERREQUEST = _descriptor.Descriptor(
  name='TransferRequest',
  full_name='pb.TransferRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account', full_name='pb.TransferRequest.account', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='beneficiary', full_name='pb.TransferRequest.beneficiary', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='amount', full_name='pb.TransferRequest.amount', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=314,
  serialized_end=385,
)


_TRANSFERREPLY = _descriptor.Descriptor(
  name='TransferReply',
  full_name='pb.TransferReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='pb.TransferReply.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='transaction', full_name='pb.TransferReply.transaction', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=387,
  serialized_end=466,
)


_ACCOUNTREQUEST = _descriptor.Descriptor(
  name='AccountRequest',
  full_name='pb.AccountRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account', full_name='pb.AccountRequest.account', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='no_transactions', full_name='pb.AccountRequest.no_transactions', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='page', full_name='pb.AccountRequest.page', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='per_page', full_name='pb.AccountRequest.per_page', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=468,
  serialized_end=558,
)


_ACCOUNTREPLY = _descriptor.Descriptor(
  name='AccountReply',
  full_name='pb.AccountReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='pb.AccountReply.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='pb.AccountReply.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='email', full_name='pb.AccountReply.email', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wallet_address', full_name='pb.AccountReply.wallet_address', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='balance', full_name='pb.AccountReply.balance', index=4,
      number=5, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='completed', full_name='pb.AccountReply.completed', index=5,
      number=6, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pending', full_name='pb.AccountReply.pending', index=6,
      number=7, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='transactions', full_name='pb.AccountReply.transactions', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=561,
  serialized_end=746,
)


_COMMAND = _descriptor.Descriptor(
  name='Command',
  full_name='pb.Command',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='pb.Command.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.Command.id', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='client', full_name='pb.Command.client', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='transfer', full_name='pb.Command.transfer', index=3,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='account', full_name='pb.Command.account', index=4,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='request', full_name='pb.Command.request',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=749,
  serialized_end=900,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='pb.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='pb.Message.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.Message.id', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='update', full_name='pb.Message.update', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='pb.Message.timestamp', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='category', full_name='pb.Message.category', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='transfer', full_name='pb.Message.transfer', index=5,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='account', full_name='pb.Message.account', index=6,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='reply', full_name='pb.Message.reply',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=903,
  serialized_end=1106,
)

_TRANSACTION.fields_by_name['originator'].message_type = _IDENTITY
_TRANSACTION.fields_by_name['beneficiary'].message_type = _IDENTITY
_TRANSFERREPLY.fields_by_name['error'].message_type = _ERROR
_TRANSFERREPLY.fields_by_name['transaction'].message_type = _TRANSACTION
_ACCOUNTREPLY.fields_by_name['error'].message_type = _ERROR
_ACCOUNTREPLY.fields_by_name['transactions'].message_type = _TRANSACTION
_COMMAND.fields_by_name['type'].enum_type = _RPC
_COMMAND.fields_by_name['transfer'].message_type = _TRANSFERREQUEST
_COMMAND.fields_by_name['account'].message_type = _ACCOUNTREQUEST
_COMMAND.oneofs_by_name['request'].fields.append(
  _COMMAND.fields_by_name['transfer'])
_COMMAND.fields_by_name['transfer'].containing_oneof = _COMMAND.oneofs_by_name['request']
_COMMAND.oneofs_by_name['request'].fields.append(
  _COMMAND.fields_by_name['account'])
_COMMAND.fields_by_name['account'].containing_oneof = _COMMAND.oneofs_by_name['request']
_MESSAGE.fields_by_name['type'].enum_type = _RPC
_MESSAGE.fields_by_name['category'].enum_type = _MESSAGECATEGORY
_MESSAGE.fields_by_name['transfer'].message_type = _TRANSFERREPLY
_MESSAGE.fields_by_name['account'].message_type = _ACCOUNTREPLY
_MESSAGE.oneofs_by_name['reply'].fields.append(
  _MESSAGE.fields_by_name['transfer'])
_MESSAGE.fields_by_name['transfer'].containing_oneof = _MESSAGE.oneofs_by_name['reply']
_MESSAGE.oneofs_by_name['reply'].fields.append(
  _MESSAGE.fields_by_name['account'])
_MESSAGE.fields_by_name['account'].containing_oneof = _MESSAGE.oneofs_by_name['reply']
DESCRIPTOR.message_types_by_name['Error'] = _ERROR
DESCRIPTOR.message_types_by_name['Identity'] = _IDENTITY
DESCRIPTOR.message_types_by_name['Transaction'] = _TRANSACTION
DESCRIPTOR.message_types_by_name['TransferRequest'] = _TRANSFERREQUEST
DESCRIPTOR.message_types_by_name['TransferReply'] = _TRANSFERREPLY
DESCRIPTOR.message_types_by_name['AccountRequest'] = _ACCOUNTREQUEST
DESCRIPTOR.message_types_by_name['AccountReply'] = _ACCOUNTREPLY
DESCRIPTOR.message_types_by_name['Command'] = _COMMAND
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.enum_types_by_name['RPC'] = _RPC
DESCRIPTOR.enum_types_by_name['MessageCategory'] = _MESSAGECATEGORY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Error = _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), {
  'DESCRIPTOR' : _ERROR,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.Error)
  })
_sym_db.RegisterMessage(Error)

Identity = _reflection.GeneratedProtocolMessageType('Identity', (_message.Message,), {
  'DESCRIPTOR' : _IDENTITY,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.Identity)
  })
_sym_db.RegisterMessage(Identity)

Transaction = _reflection.GeneratedProtocolMessageType('Transaction', (_message.Message,), {
  'DESCRIPTOR' : _TRANSACTION,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.Transaction)
  })
_sym_db.RegisterMessage(Transaction)

TransferRequest = _reflection.GeneratedProtocolMessageType('TransferRequest', (_message.Message,), {
  'DESCRIPTOR' : _TRANSFERREQUEST,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.TransferRequest)
  })
_sym_db.RegisterMessage(TransferRequest)

TransferReply = _reflection.GeneratedProtocolMessageType('TransferReply', (_message.Message,), {
  'DESCRIPTOR' : _TRANSFERREPLY,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.TransferReply)
  })
_sym_db.RegisterMessage(TransferReply)

AccountRequest = _reflection.GeneratedProtocolMessageType('AccountRequest', (_message.Message,), {
  'DESCRIPTOR' : _ACCOUNTREQUEST,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.AccountRequest)
  })
_sym_db.RegisterMessage(AccountRequest)

AccountReply = _reflection.GeneratedProtocolMessageType('AccountReply', (_message.Message,), {
  'DESCRIPTOR' : _ACCOUNTREPLY,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.AccountReply)
  })
_sym_db.RegisterMessage(AccountReply)

Command = _reflection.GeneratedProtocolMessageType('Command', (_message.Message,), {
  'DESCRIPTOR' : _COMMAND,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.Command)
  })
_sym_db.RegisterMessage(Command)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGE,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:pb.Message)
  })
_sym_db.RegisterMessage(Message)



_TRISADEMO = _descriptor.ServiceDescriptor(
  name='TRISADemo',
  full_name='pb.TRISADemo',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1238,
  serialized_end=1294,
  methods=[
  _descriptor.MethodDescriptor(
    name='LiveUpdates',
    full_name='pb.TRISADemo.LiveUpdates',
    index=0,
    containing_service=None,
    input_type=_COMMAND,
    output_type=_MESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TRISADEMO)

DESCRIPTOR.services_by_name['TRISADemo'] = _TRISADEMO


_TRISAINTEGRATION = _descriptor.ServiceDescriptor(
  name='TRISAIntegration',
  full_name='pb.TRISAIntegration',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1296,
  serialized_end=1421,
  methods=[
  _descriptor.MethodDescriptor(
    name='Transfer',
    full_name='pb.TRISAIntegration.Transfer',
    index=0,
    containing_service=None,
    input_type=_TRANSFERREQUEST,
    output_type=_TRANSFERREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='AccountStatus',
    full_name='pb.TRISAIntegration.AccountStatus',
    index=1,
    containing_service=None,
    input_type=_ACCOUNTREQUEST,
    output_type=_ACCOUNTREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TRISAINTEGRATION)

DESCRIPTOR.services_by_name['TRISAIntegration'] = _TRISAINTEGRATION

# @@protoc_insertion_point(module_scope)
