// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/libraryService.proto

package Book

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Book
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId        string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	YearPublished int32  `protobuf:"varint,4,opt,name=year_published,json=yearPublished,proto3" json:"year_published,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetYearPublished() int32 {
	if x != nil {
		return x.YearPublished
	}
	return 0
}

// AddBook
type AddBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	YearPublished int32  `protobuf:"varint,3,opt,name=year_published,json=yearPublished,proto3" json:"year_published,omitempty"`
}

func (x *AddBookRequest) Reset() {
	*x = AddBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookRequest) ProtoMessage() {}

func (x *AddBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookRequest.ProtoReflect.Descriptor instead.
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AddBookRequest) GetYearPublished() int32 {
	if x != nil {
		return x.YearPublished
	}
	return 0
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{2}
}

func (x *AddBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

// SearchBook
type SearchBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchText string `protobuf:"bytes,1,opt,name=search_text,json=searchText,proto3" json:"search_text,omitempty"`
}

func (x *SearchBookRequest) Reset() {
	*x = SearchBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookRequest) ProtoMessage() {}

func (x *SearchBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookRequest.ProtoReflect.Descriptor instead.
func (*SearchBookRequest) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{3}
}

func (x *SearchBookRequest) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

type SearchBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book []*Book `protobuf:"bytes,1,rep,name=book,proto3" json:"book,omitempty"`
}

func (x *SearchBookResponse) Reset() {
	*x = SearchBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookResponse) ProtoMessage() {}

func (x *SearchBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookResponse.ProtoReflect.Descriptor instead.
func (*SearchBookResponse) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{4}
}

func (x *SearchBookResponse) GetBook() []*Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// BorrowBook
type BorrowBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BorrowBookRequest) Reset() {
	*x = BorrowBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRequest) ProtoMessage() {}

func (x *BorrowBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRequest.ProtoReflect.Descriptor instead.
func (*BorrowBookRequest) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{5}
}

func (x *BorrowBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowBookRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BorrowBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BorrowBookResponse) Reset() {
	*x = BorrowBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_libraryService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookResponse) ProtoMessage() {}

func (x *BorrowBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_libraryService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookResponse.ProtoReflect.Descriptor instead.
func (*BorrowBookResponse) Descriptor() ([]byte, []int) {
	return file_protos_libraryService_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowBookResponse) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

var File_protos_libraryService_proto protoreflect.FileDescriptor

var file_protos_libraryService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x74, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e,
	0x79, 0x65, 0x61, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x79, 0x65, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x79, 0x65, 0x61,
	0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x45, 0x0a, 0x11,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xe6, 0x02, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1e, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x21, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x42,
	0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_libraryService_proto_rawDescOnce sync.Once
	file_protos_libraryService_proto_rawDescData = file_protos_libraryService_proto_rawDesc
)

func file_protos_libraryService_proto_rawDescGZIP() []byte {
	file_protos_libraryService_proto_rawDescOnce.Do(func() {
		file_protos_libraryService_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_libraryService_proto_rawDescData)
	})
	return file_protos_libraryService_proto_rawDescData
}

var file_protos_libraryService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_libraryService_proto_goTypes = []any{
	(*Book)(nil),               // 0: LibraryService.Book
	(*AddBookRequest)(nil),     // 1: LibraryService.AddBookRequest
	(*AddBookResponse)(nil),    // 2: LibraryService.AddBookResponse
	(*SearchBookRequest)(nil),  // 3: LibraryService.SearchBookRequest
	(*SearchBookResponse)(nil), // 4: LibraryService.SearchBookResponse
	(*BorrowBookRequest)(nil),  // 5: LibraryService.BorrowBookRequest
	(*BorrowBookResponse)(nil), // 6: LibraryService.BorrowBookResponse
}
var file_protos_libraryService_proto_depIdxs = []int32{
	0, // 0: LibraryService.SearchBookResponse.book:type_name -> LibraryService.Book
	3, // 1: LibraryService.LibraryService.SearchBookByID:input_type -> LibraryService.SearchBookRequest
	3, // 2: LibraryService.LibraryService.SearchBookByTitle:input_type -> LibraryService.SearchBookRequest
	1, // 3: LibraryService.LibraryService.AddBook:input_type -> LibraryService.AddBookRequest
	5, // 4: LibraryService.LibraryService.BorrowBook:input_type -> LibraryService.BorrowBookRequest
	4, // 5: LibraryService.LibraryService.SearchBookByID:output_type -> LibraryService.SearchBookResponse
	4, // 6: LibraryService.LibraryService.SearchBookByTitle:output_type -> LibraryService.SearchBookResponse
	2, // 7: LibraryService.LibraryService.AddBook:output_type -> LibraryService.AddBookResponse
	6, // 8: LibraryService.LibraryService.BorrowBook:output_type -> LibraryService.BorrowBookResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_libraryService_proto_init() }
func file_protos_libraryService_proto_init() {
	if File_protos_libraryService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_libraryService_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SearchBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SearchBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_libraryService_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_libraryService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_libraryService_proto_goTypes,
		DependencyIndexes: file_protos_libraryService_proto_depIdxs,
		MessageInfos:      file_protos_libraryService_proto_msgTypes,
	}.Build()
	File_protos_libraryService_proto = out.File
	file_protos_libraryService_proto_rawDesc = nil
	file_protos_libraryService_proto_goTypes = nil
	file_protos_libraryService_proto_depIdxs = nil
}
