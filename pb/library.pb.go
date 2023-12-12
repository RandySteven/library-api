// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/library.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BorrowRequest struct {
	BookId               uint32   `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BorrowRequest) Reset()         { *m = BorrowRequest{} }
func (m *BorrowRequest) String() string { return proto.CompactTextString(m) }
func (*BorrowRequest) ProtoMessage()    {}
func (*BorrowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{0}
}

func (m *BorrowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BorrowRequest.Unmarshal(m, b)
}
func (m *BorrowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BorrowRequest.Marshal(b, m, deterministic)
}
func (m *BorrowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowRequest.Merge(m, src)
}
func (m *BorrowRequest) XXX_Size() int {
	return xxx_messageInfo_BorrowRequest.Size(m)
}
func (m *BorrowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowRequest proto.InternalMessageInfo

func (m *BorrowRequest) GetBookId() uint32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

type AuthRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{1}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type BookRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity             uint32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	AuthorId             uint32   `protobuf:"varint,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRequest) Reset()         { *m = BookRequest{} }
func (m *BookRequest) String() string { return proto.CompactTextString(m) }
func (*BookRequest) ProtoMessage()    {}
func (*BookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{2}
}

func (m *BookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookRequest.Unmarshal(m, b)
}
func (m *BookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookRequest.Marshal(b, m, deterministic)
}
func (m *BookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRequest.Merge(m, src)
}
func (m *BookRequest) XXX_Size() int {
	return xxx_messageInfo_BookRequest.Size(m)
}
func (m *BookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookRequest proto.InternalMessageInfo

func (m *BookRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BookRequest) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *BookRequest) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *BookRequest) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type ReturnRequest struct {
	BorrowId             uint32   `protobuf:"varint,1,opt,name=borrow_id,json=borrowId,proto3" json:"borrow_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnRequest) Reset()         { *m = ReturnRequest{} }
func (m *ReturnRequest) String() string { return proto.CompactTextString(m) }
func (*ReturnRequest) ProtoMessage()    {}
func (*ReturnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{3}
}

func (m *ReturnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnRequest.Unmarshal(m, b)
}
func (m *ReturnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnRequest.Marshal(b, m, deterministic)
}
func (m *ReturnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnRequest.Merge(m, src)
}
func (m *ReturnRequest) XXX_Size() int {
	return xxx_messageInfo_ReturnRequest.Size(m)
}
func (m *ReturnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnRequest proto.InternalMessageInfo

func (m *ReturnRequest) GetBorrowId() uint32 {
	if m != nil {
		return m.BorrowId
	}
	return 0
}

type UserRegisterRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{4}
}

func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (m *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(m, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

func (m *UserRegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRegisterRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Author struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{5}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Author) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Author) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Author) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type BookResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Quantity             uint32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Cover                string   `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	Author               *Author  `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookResponse) Reset()         { *m = BookResponse{} }
func (m *BookResponse) String() string { return proto.CompactTextString(m) }
func (*BookResponse) ProtoMessage()    {}
func (*BookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{6}
}

func (m *BookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookResponse.Unmarshal(m, b)
}
func (m *BookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookResponse.Marshal(b, m, deterministic)
}
func (m *BookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookResponse.Merge(m, src)
}
func (m *BookResponse) XXX_Size() int {
	return xxx_messageInfo_BookResponse.Size(m)
}
func (m *BookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookResponse proto.InternalMessageInfo

func (m *BookResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BookResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BookResponse) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *BookResponse) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *BookResponse) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *BookResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *BookResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *BookResponse) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type BorrowResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BorrowStatus         string   `protobuf:"bytes,2,opt,name=borrow_status,json=borrowStatus,proto3" json:"borrow_status,omitempty"`
	UserId               uint32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	BookId               uint32   `protobuf:"varint,5,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BookName             string   `protobuf:"bytes,6,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	BookQuantity         uint32   `protobuf:"varint,7,opt,name=book_quantity,json=bookQuantity,proto3" json:"book_quantity,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BorrowResponse) Reset()         { *m = BorrowResponse{} }
func (m *BorrowResponse) String() string { return proto.CompactTextString(m) }
func (*BorrowResponse) ProtoMessage()    {}
func (*BorrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{7}
}

func (m *BorrowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BorrowResponse.Unmarshal(m, b)
}
func (m *BorrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BorrowResponse.Marshal(b, m, deterministic)
}
func (m *BorrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowResponse.Merge(m, src)
}
func (m *BorrowResponse) XXX_Size() int {
	return xxx_messageInfo_BorrowResponse.Size(m)
}
func (m *BorrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowResponse proto.InternalMessageInfo

func (m *BorrowResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BorrowResponse) GetBorrowStatus() string {
	if m != nil {
		return m.BorrowStatus
	}
	return ""
}

func (m *BorrowResponse) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BorrowResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *BorrowResponse) GetBookId() uint32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *BorrowResponse) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

func (m *BorrowResponse) GetBookQuantity() uint32 {
	if m != nil {
		return m.BookQuantity
	}
	return 0
}

func (m *BorrowResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *BorrowResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{8}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UserResponse) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type AuthResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{9}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type BorrowBookResponse struct {
	Message              string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BorrowResponse       *BorrowResponse `protobuf:"bytes,2,opt,name=borrow_response,json=borrowResponse,proto3" json:"borrow_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BorrowBookResponse) Reset()         { *m = BorrowBookResponse{} }
func (m *BorrowBookResponse) String() string { return proto.CompactTextString(m) }
func (*BorrowBookResponse) ProtoMessage()    {}
func (*BorrowBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{10}
}

func (m *BorrowBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BorrowBookResponse.Unmarshal(m, b)
}
func (m *BorrowBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BorrowBookResponse.Marshal(b, m, deterministic)
}
func (m *BorrowBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowBookResponse.Merge(m, src)
}
func (m *BorrowBookResponse) XXX_Size() int {
	return xxx_messageInfo_BorrowBookResponse.Size(m)
}
func (m *BorrowBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowBookResponse proto.InternalMessageInfo

func (m *BorrowBookResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BorrowBookResponse) GetBorrowResponse() *BorrowResponse {
	if m != nil {
		return m.BorrowResponse
	}
	return nil
}

type GetBookResponses struct {
	Message              string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BookResponses        []*BookResponse `protobuf:"bytes,2,rep,name=book_responses,json=bookResponses,proto3" json:"book_responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetBookResponses) Reset()         { *m = GetBookResponses{} }
func (m *GetBookResponses) String() string { return proto.CompactTextString(m) }
func (*GetBookResponses) ProtoMessage()    {}
func (*GetBookResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{11}
}

func (m *GetBookResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookResponses.Unmarshal(m, b)
}
func (m *GetBookResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookResponses.Marshal(b, m, deterministic)
}
func (m *GetBookResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookResponses.Merge(m, src)
}
func (m *GetBookResponses) XXX_Size() int {
	return xxx_messageInfo_GetBookResponses.Size(m)
}
func (m *GetBookResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookResponses.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookResponses proto.InternalMessageInfo

func (m *GetBookResponses) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetBookResponses) GetBookResponses() []*BookResponse {
	if m != nil {
		return m.BookResponses
	}
	return nil
}

type UserResponsesDTO struct {
	Message              string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Users                []*UserResponse `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserResponsesDTO) Reset()         { *m = UserResponsesDTO{} }
func (m *UserResponsesDTO) String() string { return proto.CompactTextString(m) }
func (*UserResponsesDTO) ProtoMessage()    {}
func (*UserResponsesDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{12}
}

func (m *UserResponsesDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponsesDTO.Unmarshal(m, b)
}
func (m *UserResponsesDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponsesDTO.Marshal(b, m, deterministic)
}
func (m *UserResponsesDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponsesDTO.Merge(m, src)
}
func (m *UserResponsesDTO) XXX_Size() int {
	return xxx_messageInfo_UserResponsesDTO.Size(m)
}
func (m *UserResponsesDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponsesDTO.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponsesDTO proto.InternalMessageInfo

func (m *UserResponsesDTO) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserResponsesDTO) GetUsers() []*UserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_594b252d607689b6, []int{13}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BorrowRequest)(nil), "library.BorrowRequest")
	proto.RegisterType((*AuthRequest)(nil), "library.AuthRequest")
	proto.RegisterType((*BookRequest)(nil), "library.BookRequest")
	proto.RegisterType((*ReturnRequest)(nil), "library.ReturnRequest")
	proto.RegisterType((*UserRegisterRequest)(nil), "library.UserRegisterRequest")
	proto.RegisterType((*Author)(nil), "library.Author")
	proto.RegisterType((*BookResponse)(nil), "library.BookResponse")
	proto.RegisterType((*BorrowResponse)(nil), "library.BorrowResponse")
	proto.RegisterType((*UserResponse)(nil), "library.UserResponse")
	proto.RegisterType((*AuthResponse)(nil), "library.AuthResponse")
	proto.RegisterType((*BorrowBookResponse)(nil), "library.BorrowBookResponse")
	proto.RegisterType((*GetBookResponses)(nil), "library.GetBookResponses")
	proto.RegisterType((*UserResponsesDTO)(nil), "library.UserResponsesDTO")
	proto.RegisterType((*Empty)(nil), "library.Empty")
}

func init() {
	proto.RegisterFile("proto/library.proto", fileDescriptor_594b252d607689b6)
}

var fileDescriptor_594b252d607689b6 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x85, 0x28, 0x91, 0x12, 0x87, 0x92, 0xec, 0xd2, 0x76, 0xad, 0xca, 0x28, 0xe0, 0xb2, 0x87,
	0x0a, 0x68, 0x61, 0x03, 0x6a, 0x0f, 0x0d, 0x10, 0x24, 0x91, 0x9d, 0x40, 0x10, 0x10, 0x38, 0x08,
	0x9d, 0x1c, 0x92, 0x8b, 0x40, 0x89, 0x0b, 0x9b, 0x91, 0xc4, 0xa5, 0x77, 0x97, 0x36, 0x7c, 0xc9,
	0x21, 0xb7, 0x5c, 0x72, 0xf2, 0x3d, 0x7f, 0x2e, 0x3f, 0x24, 0xd8, 0x2f, 0x8a, 0xa4, 0x2c, 0x9e,
	0xec, 0x79, 0xc3, 0x9d, 0x99, 0x7d, 0xef, 0x0d, 0x29, 0xd8, 0x4b, 0x08, 0x66, 0xf8, 0x74, 0x19,
	0xcd, 0x48, 0x40, 0xee, 0x4f, 0x44, 0xe4, 0x36, 0x55, 0xe8, 0x0d, 0xa0, 0x73, 0x86, 0x09, 0xc1,
	0x77, 0x3e, 0xba, 0x49, 0x11, 0x65, 0xee, 0x21, 0x34, 0x67, 0x18, 0x2f, 0xa6, 0x51, 0xd8, 0xab,
	0x1d, 0xd7, 0x06, 0x1d, 0xdf, 0xe2, 0xe1, 0x24, 0xf4, 0x9e, 0x83, 0x33, 0x4a, 0xd9, 0xb5, 0x7e,
	0x6e, 0x1f, 0x4c, 0xb4, 0x0a, 0xa2, 0xa5, 0x78, 0xca, 0xf6, 0x65, 0xe0, 0xf6, 0xa1, 0x95, 0x04,
	0x94, 0xde, 0x61, 0x12, 0xf6, 0x0c, 0x91, 0xc8, 0x62, 0xef, 0xa1, 0x06, 0xce, 0x19, 0xc6, 0x8b,
	0x5c, 0x05, 0x16, 0xb1, 0x25, 0xd2, 0x15, 0x44, 0xe0, 0x1e, 0x83, 0x13, 0x22, 0x3a, 0x27, 0x51,
	0xc2, 0x22, 0x1c, 0xab, 0x22, 0x79, 0x88, 0xf7, 0xb8, 0x49, 0x83, 0x98, 0x45, 0xec, 0xbe, 0x57,
	0x17, 0x23, 0x66, 0x31, 0xaf, 0x39, 0xc7, 0xb7, 0x88, 0xf4, 0x1a, 0xb2, 0xa6, 0x08, 0xdc, 0x23,
	0xb0, 0x83, 0x94, 0x5d, 0x63, 0xc2, 0x6f, 0x65, 0xca, 0x23, 0x12, 0x98, 0x84, 0xde, 0x3f, 0xd0,
	0xf1, 0x11, 0x4b, 0x49, 0xac, 0xe7, 0x3a, 0x02, 0x7b, 0x26, 0x28, 0x59, 0x73, 0xd0, 0x92, 0xc0,
	0x24, 0xf4, 0x3e, 0xc3, 0xde, 0x7b, 0x8a, 0x88, 0x8f, 0xae, 0x22, 0xca, 0xf8, 0x5f, 0x79, 0xc6,
	0x85, 0x46, 0x1c, 0xac, 0xf4, 0x55, 0xc4, 0xff, 0x6b, 0x86, 0x8c, 0x3c, 0x43, 0x7f, 0x40, 0x3b,
	0xb9, 0xc6, 0x31, 0x9a, 0xc6, 0xe9, 0x6a, 0x86, 0x88, 0xb8, 0x81, 0xed, 0x3b, 0x02, 0xbb, 0x10,
	0x50, 0x81, 0xc4, 0x46, 0x89, 0xc4, 0xaf, 0x35, 0xb0, 0x46, 0x62, 0x74, 0xb7, 0x0b, 0x46, 0x36,
	0xa0, 0x11, 0x85, 0xd9, 0x0c, 0x46, 0x6e, 0x86, 0xdf, 0x01, 0xe6, 0x04, 0x05, 0x0c, 0x85, 0xd3,
	0x80, 0xa9, 0x5e, 0xb6, 0x42, 0x46, 0x8c, 0xa7, 0xd3, 0x24, 0xd4, 0x69, 0xd9, 0xcb, 0x56, 0x88,
	0x4c, 0x87, 0x68, 0x89, 0x54, 0xda, 0x94, 0x69, 0x85, 0x8c, 0x98, 0xf7, 0xcd, 0x80, 0xb6, 0x14,
	0x94, 0x26, 0x38, 0xa6, 0x68, 0x63, 0xa2, 0x4c, 0x61, 0xa3, 0x42, 0xe1, 0x7a, 0xb5, 0xc2, 0x8d,
	0x6d, 0x0a, 0x9b, 0x79, 0x85, 0xff, 0x02, 0x4b, 0x0a, 0xda, 0xb3, 0x8e, 0x6b, 0x03, 0x67, 0xb8,
	0x73, 0xa2, 0xfd, 0x2e, 0xc9, 0xf2, 0x55, 0xba, 0x44, 0x48, 0xb3, 0x9a, 0x90, 0x56, 0x35, 0x21,
	0x76, 0x99, 0x90, 0x07, 0x03, 0xba, 0x7a, 0x9b, 0xb6, 0x50, 0xf2, 0x27, 0x74, 0x94, 0xb9, 0x28,
	0x0b, 0x58, 0x4a, 0x15, 0x35, 0x6d, 0x09, 0x5e, 0x0a, 0x8c, 0xef, 0x60, 0x4a, 0x91, 0x70, 0xab,
	0x34, 0xb8, 0xc5, 0xc3, 0x49, 0xc8, 0xad, 0x29, 0x12, 0x42, 0x67, 0x65, 0x0d, 0x0e, 0x5c, 0x70,
	0xad, 0x73, 0x9b, 0x6b, 0xe6, 0x37, 0x57, 0x1a, 0x1a, 0x2f, 0xe4, 0x29, 0x4b, 0x9e, 0xe2, 0x80,
	0x38, 0x25, 0x06, 0xc2, 0x8b, 0x69, 0x46, 0x78, 0x53, 0x9c, 0x6d, 0x73, 0xf0, 0xad, 0x26, 0xbd,
	0xc8, 0x5a, 0xab, 0x9a, 0x35, 0xbb, 0xc4, 0x9a, 0xf7, 0xa3, 0x06, 0x6d, 0xb9, 0x34, 0x5b, 0x48,
	0x79, 0xcc, 0xb9, 0xd9, 0xf6, 0xd4, 0xab, 0xb6, 0xa7, 0x51, 0xbd, 0x3d, 0x66, 0x71, 0x7b, 0x4a,
	0xf7, 0xb0, 0xaa, 0xef, 0xd1, 0xac, 0x56, 0xbf, 0x55, 0x56, 0xff, 0x19, 0xb4, 0xe5, 0x0b, 0x52,
	0xdd, 0xb2, 0x07, 0xcd, 0x15, 0xa2, 0x34, 0xb8, 0xd2, 0xaf, 0x05, 0x1d, 0x8a, 0xbd, 0xc0, 0x0b,
	0x14, 0x67, 0x7b, 0xc1, 0x03, 0x2f, 0x01, 0x57, 0x9a, 0xa7, 0xb0, 0x53, 0xdb, 0xab, 0xbc, 0x80,
	0x1d, 0x65, 0x25, 0xa2, 0x1e, 0x16, 0xf5, 0x9c, 0xe1, 0x61, 0x66, 0xfe, 0xa2, 0x19, 0xfd, 0xee,
	0xac, 0x10, 0x7b, 0x9f, 0x60, 0x77, 0x8c, 0x58, 0xbe, 0x1d, 0xad, 0xe8, 0xf7, 0x14, 0xba, 0xc2,
	0x29, 0xba, 0x1b, 0xf7, 0x6e, 0x7d, 0xe0, 0x0c, 0x0f, 0x72, 0xed, 0xd6, 0x95, 0x7c, 0x61, 0xab,
	0xac, 0xae, 0xf7, 0x01, 0x76, 0xf3, 0x1e, 0xa0, 0x2f, 0xdf, 0xbd, 0xa9, 0xe8, 0xf5, 0x37, 0x98,
	0xdc, 0xd7, 0x9b, 0x2d, 0xf2, 0x35, 0x7c, 0xf9, 0x8c, 0xd7, 0x04, 0xf3, 0xd5, 0x2a, 0x61, 0xf7,
	0xc3, 0xef, 0x35, 0xf8, 0x65, 0x4d, 0xe1, 0x25, 0x22, 0xb7, 0xd1, 0x1c, 0xb9, 0x63, 0xd8, 0x3d,
	0x17, 0x12, 0xaf, 0x53, 0xee, 0xaf, 0x1b, 0x14, 0x89, 0xf7, 0x78, 0xff, 0xa8, 0x84, 0x17, 0xa4,
	0x18, 0x01, 0xf8, 0x88, 0x91, 0x34, 0x2e, 0x95, 0x28, 0x7c, 0x3e, 0x2a, 0x4b, 0x0c, 0xcf, 0xe5,
	0x47, 0x54, 0x8f, 0xf6, 0x1f, 0x98, 0xaf, 0xf1, 0x55, 0x14, 0xbb, 0xfb, 0x85, 0xf7, 0x95, 0x2e,
	0x75, 0x50, 0x42, 0x55, 0x91, 0x2f, 0xea, 0x43, 0xaa, 0xab, 0xfc, 0x0f, 0xce, 0x18, 0xb1, 0xd1,
	0x72, 0xc9, 0x41, 0xea, 0x76, 0xb3, 0x53, 0x82, 0x95, 0xfe, 0x6f, 0x59, 0xbc, 0x21, 0xf6, 0x13,
	0x00, 0x4d, 0x0d, 0x5e, 0xe4, 0x86, 0xc8, 0x7d, 0xa6, 0xfb, 0x8f, 0xcb, 0x3b, 0x1c, 0x83, 0xc3,
	0xb5, 0xd8, 0x98, 0x81, 0x83, 0x55, 0x33, 0x94, 0x4d, 0x70, 0x66, 0x7d, 0x6c, 0x9c, 0x9c, 0x26,
	0xb3, 0x99, 0x25, 0x7e, 0x99, 0xfc, 0xfb, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x36, 0x54, 0x57,
	0xb0, 0x08, 0x00, 0x00,
}
