// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passwordservice.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// how data is stored in mongo
type UserInformation struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Auth *Auth `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	// repeated Login logins = 3;
	Passwords            string   `protobuf:"bytes,4,opt,name=passwords,proto3" json:"passwords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserInformation) Reset()         { *m = UserInformation{} }
func (m *UserInformation) String() string { return proto.CompactTextString(m) }
func (*UserInformation) ProtoMessage()    {}
func (*UserInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{0}
}
func (m *UserInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInformation.Unmarshal(m, b)
}
func (m *UserInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInformation.Marshal(b, m, deterministic)
}
func (dst *UserInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInformation.Merge(dst, src)
}
func (m *UserInformation) XXX_Size() int {
	return xxx_messageInfo_UserInformation.Size(m)
}
func (m *UserInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInformation.DiscardUnknown(m)
}

var xxx_messageInfo_UserInformation proto.InternalMessageInfo

func (m *UserInformation) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserInformation) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UserInformation) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

// user information
type User struct {
	First                string   `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Last                 string   `protobuf:"bytes,2,opt,name=last,proto3" json:"last,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

func (m *User) GetLast() string {
	if m != nil {
		return m.Last
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// authentication info for server
type Auth struct {
	Dn                   string          `protobuf:"bytes,1,opt,name=dn,proto3" json:"dn,omitempty"`
	authpassword         string          `protobuf:"bytes,2,opt,name=authpassword,proto3" json:"authpassword,omitempty"`
	AccessToken          string          `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	FailedLogins         int32           `protobuf:"varint,4,opt,name=failedLogins,proto3" json:"failedLogins,omitempty"`
	AuthQuestions        []*AuthQuestion `protobuf:"bytes,5,rep,name=authQuestions,proto3" json:"authQuestions,omitempty"`
	KnownIps             []string        `protobuf:"bytes,6,rep,name=knownIps,proto3" json:"knownIps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" bson:"-"`
	XXX_unrecognized     []byte          `json:"-" bson:"-"`
	XXX_sizecache        int32           `json:"-" bson:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{2}
}
func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (dst *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(dst, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetDn() string {
	if m != nil {
		return m.Dn
	}
	return ""
}

func (m *Auth) Getauthpassword() string {
	if m != nil {
		return m.authpassword
	}
	return ""
}

func (m *Auth) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Auth) GetFailedLogins() int32 {
	if m != nil {
		return m.FailedLogins
	}
	return 0
}

func (m *Auth) GetAuthQuestions() []*AuthQuestion {
	if m != nil {
		return m.AuthQuestions
	}
	return nil
}

func (m *Auth) GetKnownIps() []string {
	if m != nil {
		return m.KnownIps
	}
	return nil
}

// basic authentication questions for user
type AuthQuestion struct {
	Q                    string   `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
	A                    string   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *AuthQuestion) Reset()         { *m = AuthQuestion{} }
func (m *AuthQuestion) String() string { return proto.CompactTextString(m) }
func (*AuthQuestion) ProtoMessage()    {}
func (*AuthQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{3}
}
func (m *AuthQuestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthQuestion.Unmarshal(m, b)
}
func (m *AuthQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthQuestion.Marshal(b, m, deterministic)
}
func (dst *AuthQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthQuestion.Merge(dst, src)
}
func (m *AuthQuestion) XXX_Size() int {
	return xxx_messageInfo_AuthQuestion.Size(m)
}
func (m *AuthQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_AuthQuestion proto.InternalMessageInfo

func (m *AuthQuestion) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *AuthQuestion) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

// login information fetched from IP stack
type Login struct {
	Timestamp            int64     `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Location             *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{4}
}
func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (dst *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(dst, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Login) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CountryCode          string   `protobuf:"bytes,3,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	CountryName          string   `protobuf:"bytes,4,opt,name=countryName,proto3" json:"countryName,omitempty"`
	RegionCode           string   `protobuf:"bytes,5,opt,name=regionCode,proto3" json:"regionCode,omitempty"`
	RegionName           string   `protobuf:"bytes,6,opt,name=regionName,proto3" json:"regionName,omitempty"`
	City                 string   `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	Zip                  string   `protobuf:"bytes,8,opt,name=zip,proto3" json:"zip,omitempty"`
	Latitude             float32  `protobuf:"fixed32,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,10,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{5}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Location) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Location) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *Location) GetCountryName() string {
	if m != nil {
		return m.CountryName
	}
	return ""
}

func (m *Location) GetRegionCode() string {
	if m != nil {
		return m.RegionCode
	}
	return ""
}

func (m *Location) GetRegionName() string {
	if m != nil {
		return m.RegionName
	}
	return ""
}

func (m *Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Location) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type CrudResponse struct {
	Passwords            string   `protobuf:"bytes,1,opt,name=passwords,proto3" json:"passwords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CrudResponse) Reset()         { *m = CrudResponse{} }
func (m *CrudResponse) String() string { return proto.CompactTextString(m) }
func (*CrudResponse) ProtoMessage()    {}
func (*CrudResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{6}
}
func (m *CrudResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrudResponse.Unmarshal(m, b)
}
func (m *CrudResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrudResponse.Marshal(b, m, deterministic)
}
func (dst *CrudResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrudResponse.Merge(dst, src)
}
func (m *CrudResponse) XXX_Size() int {
	return xxx_messageInfo_CrudResponse.Size(m)
}
func (m *CrudResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CrudResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CrudResponse proto.InternalMessageInfo

func (m *CrudResponse) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

type CrudRequest struct {
	NewPasswords         string   `protobuf:"bytes,1,opt,name=newPasswords,proto3" json:"newPasswords,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CrudRequest) Reset()         { *m = CrudRequest{} }
func (m *CrudRequest) String() string { return proto.CompactTextString(m) }
func (*CrudRequest) ProtoMessage()    {}
func (*CrudRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{7}
}
func (m *CrudRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrudRequest.Unmarshal(m, b)
}
func (m *CrudRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrudRequest.Marshal(b, m, deterministic)
}
func (dst *CrudRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrudRequest.Merge(dst, src)
}
func (m *CrudRequest) XXX_Size() int {
	return xxx_messageInfo_CrudRequest.Size(m)
}
func (m *CrudRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CrudRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CrudRequest proto.InternalMessageInfo

func (m *CrudRequest) GetNewPasswords() string {
	if m != nil {
		return m.NewPasswords
	}
	return ""
}

func (m *CrudRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type ChallengeRequestBody struct {
	UserQuestionResponse *AuthQuestion `protobuf:"bytes,1,opt,name=userQuestionResponse,proto3" json:"userQuestionResponse,omitempty"`
	Location             *Location     `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" bson:"-"`
	XXX_unrecognized     []byte        `json:"-" bson:"-"`
	XXX_sizecache        int32         `json:"-" bson:"-"`
}

func (m *ChallengeRequestBody) Reset()         { *m = ChallengeRequestBody{} }
func (m *ChallengeRequestBody) String() string { return proto.CompactTextString(m) }
func (*ChallengeRequestBody) ProtoMessage()    {}
func (*ChallengeRequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{8}
}
func (m *ChallengeRequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeRequestBody.Unmarshal(m, b)
}
func (m *ChallengeRequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeRequestBody.Marshal(b, m, deterministic)
}
func (dst *ChallengeRequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeRequestBody.Merge(dst, src)
}
func (m *ChallengeRequestBody) XXX_Size() int {
	return xxx_messageInfo_ChallengeRequestBody.Size(m)
}
func (m *ChallengeRequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeRequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeRequestBody proto.InternalMessageInfo

func (m *ChallengeRequestBody) GetUserQuestionResponse() *AuthQuestion {
	if m != nil {
		return m.UserQuestionResponse
	}
	return nil
}

func (m *ChallengeRequestBody) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type ChallengeRequest struct {
	User                 string                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Body                 *ChallengeRequestBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" bson:"-"`
	XXX_unrecognized     []byte                `json:"-" bson:"-"`
	XXX_sizecache        int32                 `json:"-" bson:"-"`
}

func (m *ChallengeRequest) Reset()         { *m = ChallengeRequest{} }
func (m *ChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*ChallengeRequest) ProtoMessage()    {}
func (*ChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{9}
}
func (m *ChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeRequest.Unmarshal(m, b)
}
func (m *ChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeRequest.Marshal(b, m, deterministic)
}
func (dst *ChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeRequest.Merge(dst, src)
}
func (m *ChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_ChallengeRequest.Size(m)
}
func (m *ChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeRequest proto.InternalMessageInfo

func (m *ChallengeRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ChallengeRequest) GetBody() *ChallengeRequestBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type ChallengeResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserQuestion         string   `protobuf:"bytes,2,opt,name=userQuestion,proto3" json:"userQuestion,omitempty"`
	Logins               []*Login `protobuf:"bytes,3,rep,name=logins,proto3" json:"logins,omitempty"`
	User                 *User    `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Challenge            string   `protobuf:"bytes,5,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *ChallengeResponse) Reset()         { *m = ChallengeResponse{} }
func (m *ChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*ChallengeResponse) ProtoMessage()    {}
func (*ChallengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{10}
}
func (m *ChallengeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeResponse.Unmarshal(m, b)
}
func (m *ChallengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeResponse.Marshal(b, m, deterministic)
}
func (dst *ChallengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeResponse.Merge(dst, src)
}
func (m *ChallengeResponse) XXX_Size() int {
	return xxx_messageInfo_ChallengeResponse.Size(m)
}
func (m *ChallengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeResponse proto.InternalMessageInfo

func (m *ChallengeResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ChallengeResponse) GetUserQuestion() string {
	if m != nil {
		return m.UserQuestion
	}
	return ""
}

func (m *ChallengeResponse) GetLogins() []*Login {
	if m != nil {
		return m.Logins
	}
	return nil
}

func (m *ChallengeResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ChallengeResponse) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type DBEntry struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Auth                 *Auth    `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Logins               []*Login `protobuf:"bytes,3,rep,name=logins,proto3" json:"logins,omitempty"`
	Passwords            string   `protobuf:"bytes,4,opt,name=passwords,proto3" json:"passwords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DBEntry) Reset()         { *m = DBEntry{} }
func (m *DBEntry) String() string { return proto.CompactTextString(m) }
func (*DBEntry) ProtoMessage()    {}
func (*DBEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_passwordservice_2a4914db1d39fefc, []int{11}
}
func (m *DBEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBEntry.Unmarshal(m, b)
}
func (m *DBEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBEntry.Marshal(b, m, deterministic)
}
func (dst *DBEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBEntry.Merge(dst, src)
}
func (m *DBEntry) XXX_Size() int {
	return xxx_messageInfo_DBEntry.Size(m)
}
func (m *DBEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DBEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DBEntry proto.InternalMessageInfo

func (m *DBEntry) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *DBEntry) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DBEntry) GetLogins() []*Login {
	if m != nil {
		return m.Logins
	}
	return nil
}

func (m *DBEntry) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInformation)(nil), "protobuf.UserInformation")
	proto.RegisterType((*User)(nil), "protobuf.User")
	proto.RegisterType((*Auth)(nil), "protobuf.Auth")
	proto.RegisterType((*AuthQuestion)(nil), "protobuf.AuthQuestion")
	proto.RegisterType((*Login)(nil), "protobuf.Login")
	proto.RegisterType((*Location)(nil), "protobuf.Location")
	proto.RegisterType((*CrudResponse)(nil), "protobuf.CrudResponse")
	proto.RegisterType((*CrudRequest)(nil), "protobuf.CrudRequest")
	proto.RegisterType((*ChallengeRequestBody)(nil), "protobuf.ChallengeRequestBody")
	proto.RegisterType((*ChallengeRequest)(nil), "protobuf.ChallengeRequest")
	proto.RegisterType((*ChallengeResponse)(nil), "protobuf.ChallengeResponse")
	proto.RegisterType((*DBEntry)(nil), "protobuf.DBEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PasswordserviceClient is the client API for Passwordservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PasswordserviceClient interface {
	// reads passwords
	ReadPasswords(ctx context.Context, in *CrudRequest, opts ...grpc.CallOption) (*CrudResponse, error)
	// update passwords
	UpdatePasswords(ctx context.Context, in *CrudRequest, opts ...grpc.CallOption) (*CrudResponse, error)
	// get challenge token
	GenerateChallenge(ctx context.Context, in *ChallengeRequest, opts ...grpc.CallOption) (*ChallengeResponse, error)
}

type passwordserviceClient struct {
	cc *grpc.ClientConn
}

func NewPasswordserviceClient(cc *grpc.ClientConn) PasswordserviceClient {
	return &passwordserviceClient{cc}
}

func (c *passwordserviceClient) ReadPasswords(ctx context.Context, in *CrudRequest, opts ...grpc.CallOption) (*CrudResponse, error) {
	out := new(CrudResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Passwordservice/ReadPasswords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordserviceClient) UpdatePasswords(ctx context.Context, in *CrudRequest, opts ...grpc.CallOption) (*CrudResponse, error) {
	out := new(CrudResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Passwordservice/UpdatePasswords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordserviceClient) GenerateChallenge(ctx context.Context, in *ChallengeRequest, opts ...grpc.CallOption) (*ChallengeResponse, error) {
	out := new(ChallengeResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Passwordservice/GenerateChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordserviceServer is the server API for Passwordservice service.
type PasswordserviceServer interface {
	// reads passwords
	ReadPasswords(context.Context, *CrudRequest) (*CrudResponse, error)
	// update passwords
	UpdatePasswords(context.Context, *CrudRequest) (*CrudResponse, error)
	// get challenge token
	GenerateChallenge(context.Context, *ChallengeRequest) (*ChallengeResponse, error)
}

func RegisterPasswordserviceServer(s *grpc.Server, srv PasswordserviceServer) {
	s.RegisterService(&_Passwordservice_serviceDesc, srv)
}

func _Passwordservice_ReadPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordserviceServer).ReadPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Passwordservice/ReadPasswords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordserviceServer).ReadPasswords(ctx, req.(*CrudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwordservice_UpdatePasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordserviceServer).UpdatePasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Passwordservice/UpdatePasswords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordserviceServer).UpdatePasswords(ctx, req.(*CrudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwordservice_GenerateChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordserviceServer).GenerateChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Passwordservice/GenerateChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordserviceServer).GenerateChallenge(ctx, req.(*ChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Passwordservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Passwordservice",
	HandlerType: (*PasswordserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadPasswords",
			Handler:    _Passwordservice_ReadPasswords_Handler,
		},
		{
			MethodName: "UpdatePasswords",
			Handler:    _Passwordservice_UpdatePasswords_Handler,
		},
		{
			MethodName: "GenerateChallenge",
			Handler:    _Passwordservice_GenerateChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passwordservice.proto",
}

func init() {
	proto.RegisterFile("passwordservice.proto", fileDescriptor_passwordservice_2a4914db1d39fefc)
}

var fileDescriptor_passwordservice_2a4914db1d39fefc = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0x13, 0x3b,
	0x18, 0xd5, 0x24, 0x93, 0x34, 0xf9, 0x92, 0x36, 0xad, 0xd5, 0x56, 0xa3, 0xdc, 0xaa, 0x8a, 0x66,
	0x73, 0xa3, 0xab, 0xab, 0x44, 0x0a, 0x3b, 0xc4, 0x86, 0x96, 0x82, 0x8a, 0x2a, 0x54, 0xac, 0x76,
	0xc3, 0xce, 0xcd, 0x38, 0xa9, 0xd5, 0x89, 0x3d, 0x1d, 0x7b, 0x88, 0xc2, 0x92, 0x47, 0x28, 0x1b,
	0xf6, 0xbc, 0x04, 0xef, 0xc1, 0x8e, 0x35, 0x0f, 0x82, 0xec, 0xf1, 0xfc, 0x45, 0x6d, 0x05, 0x82,
	0x55, 0xec, 0xf3, 0x9d, 0x39, 0xfe, 0x7e, 0x8e, 0x1d, 0xd8, 0x8b, 0x88, 0x94, 0x4b, 0x11, 0x07,
	0x92, 0xc6, 0xef, 0xd9, 0x94, 0x8e, 0xa2, 0x58, 0x28, 0x81, 0x5a, 0xe6, 0xe7, 0x2a, 0x99, 0xf5,
	0x0f, 0xe6, 0x42, 0xcc, 0x43, 0x3a, 0x26, 0x11, 0x1b, 0x13, 0xce, 0x85, 0x22, 0x8a, 0x09, 0x2e,
	0x53, 0x9e, 0xbf, 0x84, 0xde, 0xa5, 0xa4, 0xf1, 0x29, 0x9f, 0x89, 0x78, 0x61, 0x22, 0xc8, 0x07,
	0x37, 0x91, 0x34, 0xf6, 0x9c, 0x81, 0x33, 0xec, 0x4c, 0xb6, 0x46, 0x99, 0xd2, 0x48, 0x13, 0xb1,
	0x89, 0x69, 0x0e, 0x49, 0xd4, 0xb5, 0x57, 0x5b, 0xe7, 0x3c, 0x4f, 0xd4, 0x35, 0x36, 0x31, 0x74,
	0x00, 0xed, 0x3c, 0x37, 0xcf, 0x1d, 0x38, 0xc3, 0x36, 0x2e, 0x00, 0xff, 0x25, 0xb8, 0x5a, 0x0f,
	0xed, 0x42, 0x63, 0xc6, 0x62, 0xa9, 0xcc, 0x71, 0x6d, 0x9c, 0x6e, 0x10, 0x02, 0x37, 0x24, 0x52,
	0x19, 0xfd, 0x36, 0x36, 0x6b, 0xcd, 0xa4, 0x0b, 0xc2, 0x42, 0xaf, 0x9e, 0x32, 0xcd, 0xc6, 0xff,
	0xee, 0x80, 0xab, 0x0f, 0x45, 0x5b, 0x50, 0x0b, 0xb8, 0x55, 0xa9, 0x05, 0xba, 0x8c, 0xae, 0x4e,
	0xe3, 0xdc, 0x9e, 0x68, 0xa5, 0x2a, 0x18, 0x1a, 0x40, 0x87, 0x4c, 0xa7, 0x54, 0xca, 0x0b, 0x71,
	0x43, 0xb9, 0x15, 0x2e, 0x43, 0x5a, 0x65, 0x46, 0x58, 0x48, 0x83, 0x33, 0x31, 0x67, 0x3c, 0xad,
	0xa3, 0x81, 0x2b, 0x18, 0x7a, 0x06, 0x9b, 0x5a, 0xf5, 0x6d, 0x42, 0xa5, 0x69, 0xad, 0xd7, 0x18,
	0xd4, 0x87, 0x9d, 0xc9, 0x7e, 0xb5, 0x2b, 0x59, 0x18, 0x57, 0xc9, 0xa8, 0x0f, 0xad, 0x1b, 0x2e,
	0x96, 0xfc, 0x34, 0x92, 0x5e, 0x73, 0x50, 0x1f, 0xb6, 0x71, 0xbe, 0xf7, 0xff, 0x83, 0x6e, 0xf9,
	0x53, 0xd4, 0x05, 0xe7, 0xd6, 0x96, 0xe8, 0xdc, 0xea, 0x1d, 0xb1, 0x65, 0x39, 0xc4, 0xbf, 0x84,
	0x86, 0xc9, 0x47, 0xf7, 0x5d, 0xb1, 0x05, 0x95, 0x8a, 0x2c, 0x22, 0x43, 0xae, 0xe3, 0x02, 0x40,
	0x23, 0x68, 0x85, 0x62, 0x6a, 0x26, 0x6d, 0xa7, 0x87, 0x8a, 0x3c, 0xcf, 0x6c, 0x04, 0xe7, 0x1c,
	0xff, 0xae, 0x06, 0xad, 0x0c, 0xd6, 0x3d, 0x66, 0x51, 0xd6, 0x63, 0x16, 0xe9, 0x31, 0xa9, 0x55,
	0x44, 0xb3, 0x31, 0xe9, 0xb5, 0xee, 0xe9, 0x54, 0x24, 0x5c, 0xc5, 0xab, 0x63, 0x11, 0xd0, 0xac,
	0xa7, 0x25, 0xa8, 0xc4, 0x78, 0x43, 0x16, 0xd4, 0x5a, 0xa3, 0x0c, 0xa1, 0x43, 0x80, 0x98, 0xce,
	0x99, 0xe0, 0x46, 0xa2, 0x61, 0x08, 0x25, 0xa4, 0x88, 0x1b, 0x81, 0x66, 0x39, 0x6e, 0xbe, 0x47,
	0xe0, 0x4e, 0x99, 0x5a, 0x79, 0x1b, 0x69, 0x5e, 0x7a, 0x8d, 0xb6, 0xa1, 0xfe, 0x81, 0x45, 0x5e,
	0xcb, 0x40, 0x7a, 0xa9, 0x3b, 0x1f, 0x12, 0xc5, 0x54, 0x12, 0x50, 0xaf, 0x3d, 0x70, 0x86, 0x35,
	0x9c, 0xef, 0x75, 0x13, 0x43, 0xc1, 0xe7, 0x69, 0x10, 0x4c, 0xb0, 0x00, 0xfc, 0xff, 0xa1, 0x7b,
	0x1c, 0x27, 0x01, 0xa6, 0x32, 0x12, 0x5c, 0xd2, 0xaa, 0xd5, 0x9d, 0x75, 0xab, 0x9f, 0x40, 0x27,
	0x65, 0xdf, 0xea, 0x39, 0x6a, 0x4b, 0x71, 0xba, 0x3c, 0x5f, 0xe3, 0x57, 0x30, 0x5d, 0x80, 0xb9,
	0x83, 0xb6, 0xb1, 0x7a, 0xed, 0xdf, 0x39, 0xb0, 0x7b, 0x7c, 0x4d, 0xc2, 0x90, 0xf2, 0x39, 0xb5,
	0x62, 0x47, 0x22, 0x58, 0xa1, 0xd7, 0xb0, 0xab, 0x09, 0xb9, 0xc1, 0x6c, 0x56, 0xf6, 0x02, 0x3f,
	0x64, 0xc3, 0x7b, 0xbf, 0xf9, 0x6d, 0x7b, 0xbc, 0x83, 0xed, 0xf5, 0x9c, 0xf2, 0xe4, 0x9d, 0x22,
	0x79, 0x34, 0x01, 0xf7, 0x4a, 0x04, 0x2b, 0xab, 0x79, 0x58, 0x68, 0xde, 0x57, 0x11, 0x36, 0x5c,
	0xff, 0xab, 0x03, 0x3b, 0xa5, 0xb0, 0xcd, 0x50, 0x3f, 0x03, 0x71, 0x2c, 0x32, 0xf9, 0x74, 0xa3,
	0x9b, 0x5a, 0xae, 0x27, 0xbb, 0xed, 0x65, 0x0c, 0xfd, 0x0b, 0xcd, 0x30, 0xbd, 0xc5, 0x75, 0x73,
	0x41, 0x7b, 0xe5, 0xca, 0xe6, 0x8c, 0x63, 0x1b, 0xce, 0x5f, 0x40, 0xf7, 0x91, 0x17, 0xf0, 0x00,
	0xda, 0xd3, 0x2c, 0x37, 0xeb, 0xd0, 0x02, 0xf0, 0x3f, 0x3b, 0xb0, 0xf1, 0xe2, 0xe8, 0x44, 0x1b,
	0xfa, 0xaf, 0xbd, 0xa7, 0xbf, 0x9c, 0xfe, 0xa3, 0x0f, 0xef, 0xe4, 0x4b, 0x0d, 0x7a, 0xe7, 0xd5,
	0xff, 0x0c, 0x84, 0x61, 0x13, 0x53, 0x12, 0x14, 0xfe, 0xdb, 0x2b, 0x0d, 0xa8, 0xb0, 0x6e, 0x7f,
	0x7f, 0x1d, 0x4e, 0x67, 0xe2, 0xa3, 0x8f, 0xdf, 0x7e, 0x7c, 0xaa, 0x75, 0x11, 0x8c, 0xf3, 0x73,
	0xd0, 0x05, 0xf4, 0x2e, 0xa3, 0x80, 0x28, 0xfa, 0xa7, 0xaa, 0x7e, 0x59, 0x75, 0x06, 0x3b, 0xaf,
	0x28, 0xa7, 0x31, 0x51, 0x34, 0xb7, 0x06, 0xea, 0x3f, 0x6c, 0xa7, 0xfe, 0x3f, 0xf7, 0xc6, 0xec,
	0x09, 0x9e, 0x39, 0x01, 0xf9, 0x30, 0xce, 0x47, 0xf7, 0xd4, 0x78, 0xef, 0xaa, 0x69, 0xbe, 0x7a,
	0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xfe, 0x6b, 0x1b, 0x5f, 0x07, 0x00, 0x00,
}
