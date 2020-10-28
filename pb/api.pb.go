// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the RPC the command is wrapping in the bidirectional stream.
type RPC int32

const (
	RPC_NORPC    RPC = 0
	RPC_TRANSFER RPC = 1
	RPC_ACCOUNT  RPC = 2
)

var RPC_name = map[int32]string{
	0: "NORPC",
	1: "TRANSFER",
	2: "ACCOUNT",
}

var RPC_value = map[string]int32{
	"NORPC":    0,
	"TRANSFER": 1,
	"ACCOUNT":  2,
}

func (x RPC) String() string {
	return proto.EnumName(RPC_name, int32(x))
}

func (RPC) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

// Allows for standardized error handling for demo purposes.
type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Identity maps a wallet address, email, and VASP provider and is used to store
// originator and beneficiary data as well as the KYC information that is collected
// during the TRISA protocol exchange in JSON format.
type Identity struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Ivms101              string   `protobuf:"bytes,3,opt,name=ivms101,proto3" json:"ivms101,omitempty"`
	Provider             string   `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *Identity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Identity) GetIvms101() string {
	if m != nil {
		return m.Ivms101
	}
	return ""
}

func (m *Identity) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

// Describes a simple transaction between an originator and beneficiary and includes
// identity information that was exchanged during the TRISA protocol.
type Transaction struct {
	Account              string    `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Originator           *Identity `protobuf:"bytes,2,opt,name=originator,proto3" json:"originator,omitempty"`
	Beneficiary          *Identity `protobuf:"bytes,3,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Amount               float32   `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Debit                bool      `protobuf:"varint,5,opt,name=debit,proto3" json:"debit,omitempty"`
	Completed            bool      `protobuf:"varint,6,opt,name=completed,proto3" json:"completed,omitempty"`
	Timestamp            string    `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Transaction) GetOriginator() *Identity {
	if m != nil {
		return m.Originator
	}
	return nil
}

func (m *Transaction) GetBeneficiary() *Identity {
	if m != nil {
		return m.Beneficiary
	}
	return nil
}

func (m *Transaction) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetDebit() bool {
	if m != nil {
		return m.Debit
	}
	return false
}

func (m *Transaction) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Transaction) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

// Initiates a transfer from the specified account to the specified wallet address or
// email address for a known wallet at some other rVASP.
type TransferRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Beneficiary          string   `protobuf:"bytes,2,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *TransferRequest) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *TransferRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// The transfer reply will contain the details of the transaction initiated or completed
// or an error if there are insufficient funds or the account or beneficiary could not
// be looked up. Errors encountered during the TRISA protocol may also be returned.
type TransferReply struct {
	Error                *Error       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Transaction          *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransferReply) Reset()         { *m = TransferReply{} }
func (m *TransferReply) String() string { return proto.CompactTextString(m) }
func (*TransferReply) ProtoMessage()    {}
func (*TransferReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *TransferReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferReply.Unmarshal(m, b)
}
func (m *TransferReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferReply.Marshal(b, m, deterministic)
}
func (m *TransferReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferReply.Merge(m, src)
}
func (m *TransferReply) XXX_Size() int {
	return xxx_messageInfo_TransferReply.Size(m)
}
func (m *TransferReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferReply.DiscardUnknown(m)
}

var xxx_messageInfo_TransferReply proto.InternalMessageInfo

func (m *TransferReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TransferReply) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// Account request is used to fetch the status information of the account as well as
// all the transactions associated with the account (unless otherwise requested).
// TODO: implement transaction pagination.
type AccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	NoTransactions       bool     `protobuf:"varint,2,opt,name=no_transactions,json=noTransactions,proto3" json:"no_transactions,omitempty"`
	Page                 uint32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountRequest) GetNoTransactions() bool {
	if m != nil {
		return m.NoTransactions
	}
	return false
}

func (m *AccountRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AccountRequest) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

// Returns the account information and balance as well as transactions ordered from
// most to least recent. An error is returned if the account cannot be found.
type AccountReply struct {
	Error                *Error         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string         `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	WalletAddress        string         `protobuf:"bytes,4,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	Balance              float32        `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Completed            uint64         `protobuf:"varint,6,opt,name=completed,proto3" json:"completed,omitempty"`
	Pending              uint64         `protobuf:"varint,7,opt,name=pending,proto3" json:"pending,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,8,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AccountReply) Reset()         { *m = AccountReply{} }
func (m *AccountReply) String() string { return proto.CompactTextString(m) }
func (*AccountReply) ProtoMessage()    {}
func (*AccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *AccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountReply.Unmarshal(m, b)
}
func (m *AccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountReply.Marshal(b, m, deterministic)
}
func (m *AccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountReply.Merge(m, src)
}
func (m *AccountReply) XXX_Size() int {
	return xxx_messageInfo_AccountReply.Size(m)
}
func (m *AccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccountReply proto.InternalMessageInfo

func (m *AccountReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountReply) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *AccountReply) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountReply) GetCompleted() uint64 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *AccountReply) GetPending() uint64 {
	if m != nil {
		return m.Pending
	}
	return 0
}

func (m *AccountReply) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// A wrapper for the TransferRequet and AccountRequest RPCs to be sent via streaming.
type Command struct {
	Type   RPC    `protobuf:"varint,1,opt,name=type,proto3,enum=pb.RPC" json:"type,omitempty"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Client string `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// only one of these fields can be set, and the field that is set should
	// match the RPC type described above.
	//
	// Types that are valid to be assigned to Request:
	//	*Command_Transfer
	//	*Command_Account
	Request              isCommand_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetType() RPC {
	if m != nil {
		return m.Type
	}
	return RPC_NORPC
}

func (m *Command) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Command) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type isCommand_Request interface {
	isCommand_Request()
}

type Command_Transfer struct {
	Transfer *TransferRequest `protobuf:"bytes,11,opt,name=transfer,proto3,oneof"`
}

type Command_Account struct {
	Account *AccountRequest `protobuf:"bytes,12,opt,name=account,proto3,oneof"`
}

func (*Command_Transfer) isCommand_Request() {}

func (*Command_Account) isCommand_Request() {}

func (m *Command) GetRequest() isCommand_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Command) GetTransfer() *TransferRequest {
	if x, ok := m.GetRequest().(*Command_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *Command) GetAccount() *AccountRequest {
	if x, ok := m.GetRequest().(*Command_Account); ok {
		return x.Account
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Command) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Command_OneofMarshaler, _Command_OneofUnmarshaler, _Command_OneofSizer, []interface{}{
		(*Command_Transfer)(nil),
		(*Command_Account)(nil),
	}
}

func _Command_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Command)
	// request
	switch x := m.Request.(type) {
	case *Command_Transfer:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *Command_Account:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Account); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Command.Request has unexpected type %T", x)
	}
	return nil
}

func _Command_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Command)
	switch tag {
	case 11: // request.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Command_Transfer{msg}
		return true, err
	case 12: // request.account
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AccountRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Command_Account{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Command_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Command)
	// request
	switch x := m.Request.(type) {
	case *Command_Transfer:
		s := proto.Size(x.Transfer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_Account:
		s := proto.Size(x.Account)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Message is either a wrapper for a TransferReply or AccountReply RPCs or it is a live
// update message sent from the rVASP to show the communication interactions of the
// InterVASP protocol. If it is a wrapper, then type will be > 0 and the ID will match
// the id of the command request sent by the client. Otherwise both of these fields will
// be zero and the update string will be populated.
type Message struct {
	Type      RPC    `protobuf:"varint,1,opt,name=type,proto3,enum=pb.RPC" json:"type,omitempty"`
	Id        uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Update    string `protobuf:"bytes,3,opt,name=update,proto3" json:"update,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// if type and id are greater than zero, one of these fields will be set, matching
	// the RPC type described above.
	//
	// Types that are valid to be assigned to Reply:
	//	*Message_Transfer
	//	*Message_Account
	Reply                isMessage_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() RPC {
	if m != nil {
		return m.Type
	}
	return RPC_NORPC
}

func (m *Message) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetUpdate() string {
	if m != nil {
		return m.Update
	}
	return ""
}

func (m *Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type isMessage_Reply interface {
	isMessage_Reply()
}

type Message_Transfer struct {
	Transfer *TransferReply `protobuf:"bytes,11,opt,name=transfer,proto3,oneof"`
}

type Message_Account struct {
	Account *AccountReply `protobuf:"bytes,12,opt,name=account,proto3,oneof"`
}

func (*Message_Transfer) isMessage_Reply() {}

func (*Message_Account) isMessage_Reply() {}

func (m *Message) GetReply() isMessage_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *Message) GetTransfer() *TransferReply {
	if x, ok := m.GetReply().(*Message_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *Message) GetAccount() *AccountReply {
	if x, ok := m.GetReply().(*Message_Account); ok {
		return x.Account
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Transfer)(nil),
		(*Message_Account)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// reply
	switch x := m.Reply.(type) {
	case *Message_Transfer:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *Message_Account:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Account); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Reply has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 11: // reply.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferReply)
		err := b.DecodeMessage(msg)
		m.Reply = &Message_Transfer{msg}
		return true, err
	case 12: // reply.account
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AccountReply)
		err := b.DecodeMessage(msg)
		m.Reply = &Message_Account{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// reply
	switch x := m.Reply.(type) {
	case *Message_Transfer:
		s := proto.Size(x.Transfer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Account:
		s := proto.Size(x.Account)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("pb.RPC", RPC_name, RPC_value)
	proto.RegisterType((*Error)(nil), "pb.Error")
	proto.RegisterType((*Identity)(nil), "pb.Identity")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*TransferRequest)(nil), "pb.TransferRequest")
	proto.RegisterType((*TransferReply)(nil), "pb.TransferReply")
	proto.RegisterType((*AccountRequest)(nil), "pb.AccountRequest")
	proto.RegisterType((*AccountReply)(nil), "pb.AccountReply")
	proto.RegisterType((*Command)(nil), "pb.Command")
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0x64, 0x39, 0x92, 0x8e, 0x6c, 0xc7, 0xe3, 0x86, 0x41, 0xf3, 0x06, 0xcc, 0x10, 0x30,
	0xcc, 0x58, 0x02, 0x2f, 0x76, 0x10, 0x60, 0xb7, 0x9e, 0x9b, 0x22, 0x01, 0xda, 0xc4, 0x60, 0x9c,
	0xeb, 0x80, 0x96, 0x18, 0x83, 0x80, 0xfe, 0x4a, 0xd1, 0x29, 0x7c, 0x91, 0x9b, 0xde, 0xf4, 0xa5,
	0xfa, 0x18, 0x7d, 0x8f, 0xbe, 0x42, 0x41, 0x4a, 0xb2, 0xa5, 0xd8, 0xfd, 0xbb, 0xd3, 0x77, 0x78,
	0xc8, 0xf3, 0x7d, 0xe7, 0x7c, 0xa4, 0xc0, 0x26, 0x29, 0x1b, 0xa6, 0x3c, 0x11, 0x09, 0xd2, 0xd3,
	0x85, 0x77, 0x0e, 0xcd, 0x0b, 0xce, 0x13, 0x8e, 0x10, 0x18, 0x7e, 0x12, 0x50, 0x57, 0xeb, 0x6b,
	0x83, 0x26, 0x56, 0xdf, 0xc8, 0x05, 0x33, 0xa2, 0x59, 0x46, 0x96, 0xd4, 0xd5, 0xfb, 0xda, 0xc0,
	0xc6, 0x25, 0xf4, 0x9e, 0xc0, 0xba, 0x0a, 0x68, 0x2c, 0x98, 0x58, 0xa3, 0xbf, 0xa0, 0xf3, 0x96,
	0x84, 0x21, 0x15, 0xf7, 0x24, 0x08, 0x38, 0xcd, 0x32, 0x75, 0x86, 0x8d, 0xdb, 0x79, 0x74, 0x92,
	0x07, 0xd1, 0x2f, 0xd0, 0xa4, 0x11, 0x61, 0x61, 0x71, 0x54, 0x0e, 0x64, 0x09, 0xf6, 0x18, 0x65,
	0xa3, 0xd3, 0x91, 0xdb, 0xc8, 0x4b, 0x14, 0x10, 0xf5, 0xc0, 0x4a, 0x79, 0xf2, 0xc8, 0x02, 0xca,
	0x5d, 0x43, 0x2d, 0x6d, 0xb0, 0xf7, 0x49, 0x03, 0x67, 0xce, 0x49, 0x9c, 0x11, 0x5f, 0xb0, 0x24,
	0x96, 0xa7, 0x10, 0xdf, 0x4f, 0x56, 0xb1, 0x28, 0x6a, 0x97, 0x10, 0x9d, 0x00, 0x24, 0x9c, 0x2d,
	0x59, 0x4c, 0x44, 0xc2, 0x55, 0x69, 0x67, 0xdc, 0x1a, 0xa6, 0x8b, 0x61, 0x49, 0x1f, 0x57, 0xd6,
	0xd1, 0x10, 0x9c, 0x05, 0x8d, 0xe9, 0x03, 0xf3, 0x19, 0xe1, 0x6b, 0xc5, 0xe8, 0x79, 0x7a, 0x35,
	0x01, 0xfd, 0x0a, 0x87, 0x24, 0x52, 0x65, 0x25, 0x43, 0x1d, 0x17, 0x48, 0x6a, 0x0d, 0xe8, 0x82,
	0x09, 0xb7, 0xd9, 0xd7, 0x06, 0x16, 0xce, 0x01, 0xfa, 0x03, 0x6c, 0x3f, 0x89, 0xd2, 0x90, 0x0a,
	0x1a, 0xb8, 0x87, 0x6a, 0x65, 0x1b, 0x90, 0xab, 0x82, 0x45, 0x34, 0x13, 0x24, 0x4a, 0x5d, 0x53,
	0xa9, 0xd8, 0x06, 0x3c, 0x0a, 0x47, 0x4a, 0xf0, 0x03, 0xe5, 0x98, 0xbe, 0x59, 0xd1, 0x4c, 0x7c,
	0x45, 0x74, 0xbf, 0x2e, 0x23, 0x6f, 0xf8, 0x17, 0x88, 0x37, 0xaa, 0xc4, 0x3d, 0x1f, 0xda, 0xdb,
	0x32, 0x69, 0xb8, 0x46, 0x7f, 0x42, 0x93, 0x4a, 0x7f, 0xa8, 0x12, 0xce, 0xd8, 0x96, 0xbd, 0x50,
	0x86, 0xc1, 0x79, 0x1c, 0x8d, 0xc0, 0x11, 0xdb, 0x49, 0x14, 0x1d, 0x3e, 0x92, 0x69, 0x95, 0x01,
	0xe1, 0x6a, 0x8e, 0xf7, 0x4e, 0x83, 0xce, 0x24, 0xa7, 0xfa, 0x6d, 0x2d, 0x7f, 0xc3, 0x51, 0x9c,
	0xdc, 0x57, 0xb6, 0x67, 0xaa, 0x86, 0x85, 0x3b, 0x71, 0x52, 0xa9, 0x90, 0x49, 0x03, 0xa7, 0xd2,
	0xa9, 0x52, 0x50, 0x1b, 0xab, 0x6f, 0xf4, 0x1b, 0x58, 0x29, 0xe5, 0xf7, 0x2a, 0x6e, 0xa8, 0xb8,
	0x99, 0x52, 0x3e, 0x93, 0x0e, 0x7e, 0xaf, 0x43, 0x6b, 0x43, 0xe2, 0xbb, 0x94, 0x22, 0x30, 0x62,
	0x12, 0x95, 0x57, 0x41, 0x7d, 0x6f, 0x4d, 0xdd, 0xa8, 0x9a, 0x7a, 0xf7, 0x46, 0x18, 0xfb, 0x6e,
	0x84, 0x0b, 0xe6, 0x82, 0x84, 0x24, 0xf6, 0xa9, 0xf2, 0x89, 0x8e, 0x4b, 0xb8, 0xeb, 0x14, 0xa3,
	0xea, 0x14, 0x17, 0xcc, 0x94, 0xc6, 0x01, 0x8b, 0x97, 0xca, 0x27, 0x06, 0x2e, 0x21, 0x3a, 0x83,
	0x56, 0xad, 0x53, 0x56, 0xbf, 0xb1, 0x6f, 0x1a, 0xb5, 0x24, 0xef, 0x83, 0x06, 0xe6, 0x34, 0x89,
	0x22, 0x12, 0x07, 0xe8, 0x77, 0x30, 0xc4, 0x3a, 0xcd, 0x5f, 0x81, 0xce, 0xd8, 0x94, 0x1b, 0xf1,
	0x6c, 0x8a, 0x55, 0x10, 0x75, 0x40, 0x67, 0x81, 0x92, 0x6f, 0x60, 0x9d, 0x05, 0xd2, 0x44, 0x7e,
	0xc8, 0x68, 0x61, 0x22, 0x1b, 0x17, 0x08, 0x8d, 0xc0, 0x12, 0x85, 0x89, 0x5c, 0x47, 0x35, 0xf3,
	0xe7, 0x0d, 0x83, 0xad, 0x7f, 0x2f, 0x0f, 0xf0, 0x26, 0x0d, 0x0d, 0xb7, 0xf3, 0x6f, 0xa9, 0x1d,
	0x48, 0xee, 0xa8, 0x9b, 0xe4, 0xf2, 0x60, 0xe3, 0x8a, 0xff, 0x6d, 0x30, 0x79, 0x1e, 0xf5, 0x3e,
	0x6a, 0x60, 0xbe, 0xce, 0x9f, 0xa5, 0x1f, 0xa6, 0xbf, 0x4a, 0x03, 0x22, 0x68, 0x49, 0x3f, 0x47,
	0xf5, 0x8b, 0x68, 0x3c, 0xbb, 0x88, 0xe8, 0xdf, 0x1d, 0x71, 0x3f, 0xd5, 0xc5, 0xa5, 0xe1, 0xba,
	0x26, 0xed, 0xe4, 0xb9, 0xb4, 0x6e, 0x4d, 0x5a, 0x9e, 0xbe, 0x11, 0x66, 0x42, 0x93, 0xcb, 0xd8,
	0x3f, 0xc7, 0xd0, 0xc0, 0xb3, 0x29, 0xb2, 0xa1, 0x79, 0x7d, 0x83, 0x67, 0xd3, 0xee, 0x01, 0x6a,
	0x81, 0x35, 0xc7, 0x93, 0xeb, 0xdb, 0x97, 0x17, 0xb8, 0xab, 0x21, 0x07, 0xcc, 0xc9, 0x74, 0x7a,
	0x73, 0x77, 0x3d, 0xef, 0xea, 0xe3, 0xff, 0xc0, 0x9e, 0xe3, 0xab, 0xdb, 0xc9, 0x0b, 0x1a, 0x25,
	0xe8, 0x18, 0x9c, 0x57, 0xec, 0x91, 0xde, 0x29, 0x35, 0x19, 0x72, 0x64, 0xb9, 0x62, 0xbe, 0x3d,
	0x05, 0x8a, 0x6e, 0x0d, 0xb4, 0x53, 0x6d, 0xfc, 0x04, 0x5d, 0xb5, 0xf3, 0x2a, 0x16, 0x74, 0xc9,
	0x89, 0x7a, 0x4d, 0xc7, 0x60, 0x95, 0x72, 0xd0, 0xbe, 0xc9, 0xf5, 0x76, 0x15, 0xa3, 0x73, 0x68,
	0x17, 0x92, 0x6e, 0x05, 0x11, 0xab, 0x0c, 0xed, 0x19, 0x60, 0x6f, 0x47, 0xf9, 0xe2, 0x50, 0xfd,
	0x89, 0xce, 0x3e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x61, 0x6a, 0x15, 0x96, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TRISADemoClient is the client API for TRISADemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TRISADemoClient interface {
	LiveUpdates(ctx context.Context, opts ...grpc.CallOption) (TRISADemo_LiveUpdatesClient, error)
}

type tRISADemoClient struct {
	cc *grpc.ClientConn
}

func NewTRISADemoClient(cc *grpc.ClientConn) TRISADemoClient {
	return &tRISADemoClient{cc}
}

func (c *tRISADemoClient) LiveUpdates(ctx context.Context, opts ...grpc.CallOption) (TRISADemo_LiveUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TRISADemo_serviceDesc.Streams[0], "/pb.TRISADemo/LiveUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &tRISADemoLiveUpdatesClient{stream}
	return x, nil
}

type TRISADemo_LiveUpdatesClient interface {
	Send(*Command) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type tRISADemoLiveUpdatesClient struct {
	grpc.ClientStream
}

func (x *tRISADemoLiveUpdatesClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tRISADemoLiveUpdatesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TRISADemoServer is the server API for TRISADemo service.
type TRISADemoServer interface {
	LiveUpdates(TRISADemo_LiveUpdatesServer) error
}

func RegisterTRISADemoServer(s *grpc.Server, srv TRISADemoServer) {
	s.RegisterService(&_TRISADemo_serviceDesc, srv)
}

func _TRISADemo_LiveUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TRISADemoServer).LiveUpdates(&tRISADemoLiveUpdatesServer{stream})
}

type TRISADemo_LiveUpdatesServer interface {
	Send(*Message) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type tRISADemoLiveUpdatesServer struct {
	grpc.ServerStream
}

func (x *tRISADemoLiveUpdatesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tRISADemoLiveUpdatesServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TRISADemo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TRISADemo",
	HandlerType: (*TRISADemoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LiveUpdates",
			Handler:       _TRISADemo_LiveUpdates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

// TRISAIntegrationClient is the client API for TRISAIntegration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TRISAIntegrationClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error)
	AccountStatus(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
}

type tRISAIntegrationClient struct {
	cc *grpc.ClientConn
}

func NewTRISAIntegrationClient(cc *grpc.ClientConn) TRISAIntegrationClient {
	return &tRISAIntegrationClient{cc}
}

func (c *tRISAIntegrationClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error) {
	out := new(TransferReply)
	err := c.cc.Invoke(ctx, "/pb.TRISAIntegration/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tRISAIntegrationClient) AccountStatus(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/pb.TRISAIntegration/AccountStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TRISAIntegrationServer is the server API for TRISAIntegration service.
type TRISAIntegrationServer interface {
	Transfer(context.Context, *TransferRequest) (*TransferReply, error)
	AccountStatus(context.Context, *AccountRequest) (*AccountReply, error)
}

func RegisterTRISAIntegrationServer(s *grpc.Server, srv TRISAIntegrationServer) {
	s.RegisterService(&_TRISAIntegration_serviceDesc, srv)
}

func _TRISAIntegration_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TRISAIntegrationServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TRISAIntegration/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TRISAIntegrationServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TRISAIntegration_AccountStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TRISAIntegrationServer).AccountStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TRISAIntegration/AccountStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TRISAIntegrationServer).AccountStatus(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TRISAIntegration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TRISAIntegration",
	HandlerType: (*TRISAIntegrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TRISAIntegration_Transfer_Handler,
		},
		{
			MethodName: "AccountStatus",
			Handler:    _TRISAIntegration_AccountStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
