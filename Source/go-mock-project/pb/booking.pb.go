// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: booking.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GrpcBookFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode    string                `protobuf:"bytes,1,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	ContactEmail    string                `protobuf:"bytes,2,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	ContactPhone    string                `protobuf:"bytes,3,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	FlightOptionIds []int64               `protobuf:"varint,4,rep,packed,name=flightOptionIds,proto3" json:"flightOptionIds,omitempty"`
	Passengers      []*GrpcBfReqPassenger `protobuf:"bytes,5,rep,name=passengers,proto3" json:"passengers,omitempty"`
}

func (x *GrpcBookFlightRequest) Reset() {
	*x = GrpcBookFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcBookFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcBookFlightRequest) ProtoMessage() {}

func (x *GrpcBookFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcBookFlightRequest.ProtoReflect.Descriptor instead.
func (*GrpcBookFlightRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcBookFlightRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *GrpcBookFlightRequest) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *GrpcBookFlightRequest) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *GrpcBookFlightRequest) GetFlightOptionIds() []int64 {
	if x != nil {
		return x.FlightOptionIds
	}
	return nil
}

func (x *GrpcBookFlightRequest) GetPassengers() []*GrpcBfReqPassenger {
	if x != nil {
		return x.Passengers
	}
	return nil
}

type GrpcBfReqPassenger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title           string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	LastName        string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName       string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	DateOfBirth     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	TravelDocType   string                 `protobuf:"bytes,5,opt,name=travel_doc_type,json=travelDocType,proto3" json:"travel_doc_type,omitempty"`
	TravelDocNumber string                 `protobuf:"bytes,6,opt,name=travel_doc_number,json=travelDocNumber,proto3" json:"travel_doc_number,omitempty"`
	FfpNumber       string                 `protobuf:"bytes,7,opt,name=ffp_number,json=ffpNumber,proto3" json:"ffp_number,omitempty"`
}

func (x *GrpcBfReqPassenger) Reset() {
	*x = GrpcBfReqPassenger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcBfReqPassenger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcBfReqPassenger) ProtoMessage() {}

func (x *GrpcBfReqPassenger) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcBfReqPassenger.ProtoReflect.Descriptor instead.
func (*GrpcBfReqPassenger) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcBfReqPassenger) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GrpcBfReqPassenger) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcBfReqPassenger) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GrpcBfReqPassenger) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *GrpcBfReqPassenger) GetTravelDocType() string {
	if x != nil {
		return x.TravelDocType
	}
	return ""
}

func (x *GrpcBfReqPassenger) GetTravelDocNumber() string {
	if x != nil {
		return x.TravelDocNumber
	}
	return ""
}

func (x *GrpcBfReqPassenger) GetFfpNumber() string {
	if x != nil {
		return x.FfpNumber
	}
	return ""
}

type GrpcBookFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationCode string                   `protobuf:"bytes,1,opt,name=reservation_code,json=reservationCode,proto3" json:"reservation_code,omitempty"`
	ContactEmail    string                   `protobuf:"bytes,2,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	ContactPhone    string                   `protobuf:"bytes,3,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	CurrencyCode    string                   `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	TotalPrice      float64                  `protobuf:"fixed64,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	FlightOptions   []*GrpcBfResFlightOption `protobuf:"bytes,6,rep,name=flight_options,json=flightOptions,proto3" json:"flight_options,omitempty"`
	Passengers      []*GrpcBfResPassenger    `protobuf:"bytes,7,rep,name=passengers,proto3" json:"passengers,omitempty"`
}

func (x *GrpcBookFlightResponse) Reset() {
	*x = GrpcBookFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcBookFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcBookFlightResponse) ProtoMessage() {}

func (x *GrpcBookFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcBookFlightResponse.ProtoReflect.Descriptor instead.
func (*GrpcBookFlightResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcBookFlightResponse) GetReservationCode() string {
	if x != nil {
		return x.ReservationCode
	}
	return ""
}

func (x *GrpcBookFlightResponse) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *GrpcBookFlightResponse) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *GrpcBookFlightResponse) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *GrpcBookFlightResponse) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *GrpcBookFlightResponse) GetFlightOptions() []*GrpcBfResFlightOption {
	if x != nil {
		return x.FlightOptions
	}
	return nil
}

func (x *GrpcBookFlightResponse) GetPassengers() []*GrpcBfResPassenger {
	if x != nil {
		return x.Passengers
	}
	return nil
}

type GrpcBfResFlightOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OriginAirportCode      string                 `protobuf:"bytes,2,opt,name=origin_airport_code,json=originAirportCode,proto3" json:"origin_airport_code,omitempty"`
	DestinationAirportCode string                 `protobuf:"bytes,3,opt,name=destination_airport_code,json=destinationAirportCode,proto3" json:"destination_airport_code,omitempty"`
	DepartureDateTime      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=departure_date_time,json=departureDateTime,proto3" json:"departure_date_time,omitempty"`
	ArrivalDateTime        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=arrival_date_time,json=arrivalDateTime,proto3" json:"arrival_date_time,omitempty"`
	BookingClass           string                 `protobuf:"bytes,6,opt,name=booking_class,json=bookingClass,proto3" json:"booking_class,omitempty"`
	FlightNumber           string                 `protobuf:"bytes,7,opt,name=flight_number,json=flightNumber,proto3" json:"flight_number,omitempty"`
	FlightDuration         int32                  `protobuf:"varint,8,opt,name=flight_duration,json=flightDuration,proto3" json:"flight_duration,omitempty"`
	CurrencyCode           string                 `protobuf:"bytes,9,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Price                  float64                `protobuf:"fixed64,10,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GrpcBfResFlightOption) Reset() {
	*x = GrpcBfResFlightOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcBfResFlightOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcBfResFlightOption) ProtoMessage() {}

func (x *GrpcBfResFlightOption) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcBfResFlightOption.ProtoReflect.Descriptor instead.
func (*GrpcBfResFlightOption) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcBfResFlightOption) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GrpcBfResFlightOption) GetOriginAirportCode() string {
	if x != nil {
		return x.OriginAirportCode
	}
	return ""
}

func (x *GrpcBfResFlightOption) GetDestinationAirportCode() string {
	if x != nil {
		return x.DestinationAirportCode
	}
	return ""
}

func (x *GrpcBfResFlightOption) GetDepartureDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureDateTime
	}
	return nil
}

func (x *GrpcBfResFlightOption) GetArrivalDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ArrivalDateTime
	}
	return nil
}

func (x *GrpcBfResFlightOption) GetBookingClass() string {
	if x != nil {
		return x.BookingClass
	}
	return ""
}

func (x *GrpcBfResFlightOption) GetFlightNumber() string {
	if x != nil {
		return x.FlightNumber
	}
	return ""
}

func (x *GrpcBfResFlightOption) GetFlightDuration() int32 {
	if x != nil {
		return x.FlightDuration
	}
	return 0
}

func (x *GrpcBfResFlightOption) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *GrpcBfResFlightOption) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GrpcBfResPassenger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title           string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	LastName        string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName       string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	DateOfBirth     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	TravelDocType   string                 `protobuf:"bytes,5,opt,name=travel_doc_type,json=travelDocType,proto3" json:"travel_doc_type,omitempty"`
	TravelDocNumber string                 `protobuf:"bytes,6,opt,name=travel_doc_number,json=travelDocNumber,proto3" json:"travel_doc_number,omitempty"`
	FfpNumber       string                 `protobuf:"bytes,7,opt,name=ffp_number,json=ffpNumber,proto3" json:"ffp_number,omitempty"`
}

func (x *GrpcBfResPassenger) Reset() {
	*x = GrpcBfResPassenger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcBfResPassenger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcBfResPassenger) ProtoMessage() {}

func (x *GrpcBfResPassenger) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcBfResPassenger.ProtoReflect.Descriptor instead.
func (*GrpcBfResPassenger) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcBfResPassenger) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GrpcBfResPassenger) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcBfResPassenger) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GrpcBfResPassenger) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *GrpcBfResPassenger) GetTravelDocType() string {
	if x != nil {
		return x.TravelDocType
	}
	return ""
}

func (x *GrpcBfResPassenger) GetTravelDocNumber() string {
	if x != nil {
		return x.TravelDocNumber
	}
	return ""
}

func (x *GrpcBfResPassenger) GetFfpNumber() string {
	if x != nil {
		return x.FfpNumber
	}
	return ""
}

type GrpcViewBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationCode string `protobuf:"bytes,1,opt,name=reservation_code,json=reservationCode,proto3" json:"reservation_code,omitempty"`
	LastName        string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	TravelDocNumber string `protobuf:"bytes,3,opt,name=travel_doc_number,json=travelDocNumber,proto3" json:"travel_doc_number,omitempty"`
}

func (x *GrpcViewBookingRequest) Reset() {
	*x = GrpcViewBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcViewBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcViewBookingRequest) ProtoMessage() {}

func (x *GrpcViewBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcViewBookingRequest.ProtoReflect.Descriptor instead.
func (*GrpcViewBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *GrpcViewBookingRequest) GetReservationCode() string {
	if x != nil {
		return x.ReservationCode
	}
	return ""
}

func (x *GrpcViewBookingRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcViewBookingRequest) GetTravelDocNumber() string {
	if x != nil {
		return x.TravelDocNumber
	}
	return ""
}

type GrpcCancelBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationCode string `protobuf:"bytes,1,opt,name=reservation_code,json=reservationCode,proto3" json:"reservation_code,omitempty"`
	LastName        string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	TravelDocNumber string `protobuf:"bytes,3,opt,name=travel_doc_number,json=travelDocNumber,proto3" json:"travel_doc_number,omitempty"`
}

func (x *GrpcCancelBookingRequest) Reset() {
	*x = GrpcCancelBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcCancelBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcCancelBookingRequest) ProtoMessage() {}

func (x *GrpcCancelBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcCancelBookingRequest.ProtoReflect.Descriptor instead.
func (*GrpcCancelBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *GrpcCancelBookingRequest) GetReservationCode() string {
	if x != nil {
		return x.ReservationCode
	}
	return ""
}

func (x *GrpcCancelBookingRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcCancelBookingRequest) GetTravelDocNumber() string {
	if x != nil {
		return x.TravelDocNumber
	}
	return ""
}

type GrpcCancelBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationCode string                 `protobuf:"bytes,1,opt,name=reservation_code,json=reservationCode,proto3" json:"reservation_code,omitempty"`
	CancelDate      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=cancel_date,json=cancelDate,proto3" json:"cancel_date,omitempty"`
}

func (x *GrpcCancelBookingResponse) Reset() {
	*x = GrpcCancelBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcCancelBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcCancelBookingResponse) ProtoMessage() {}

func (x *GrpcCancelBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcCancelBookingResponse.ProtoReflect.Descriptor instead.
func (*GrpcCancelBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{7}
}

func (x *GrpcCancelBookingResponse) GetReservationCode() string {
	if x != nil {
		return x.ReservationCode
	}
	return ""
}

func (x *GrpcCancelBookingResponse) GetCancelDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelDate
	}
	return nil
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x15, 0x47, 0x72, 0x70, 0x63, 0x42, 0x6f, 0x6f,
	0x6b, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x42, 0x66, 0x52, 0x65, 0x71, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x22, 0x99,
	0x02, 0x0a, 0x12, 0x47, 0x72, 0x70, 0x63, 0x42, 0x66, 0x52, 0x65, 0x71, 0x50, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x44, 0x6f, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x44, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x66, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x66, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xcd, 0x02, 0x0a, 0x16, 0x47,
	0x72, 0x70, 0x63, 0x42, 0x6f, 0x6f, 0x6b, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x42, 0x66, 0x52, 0x65, 0x73, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x42, 0x66, 0x52, 0x65, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x0a,
	0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x22, 0xd3, 0x03, 0x0a, 0x15, 0x47,
	0x72, 0x70, 0x63, 0x42, 0x66, 0x52, 0x65, 0x73, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4a,
	0x0a, 0x13, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x61, 0x72,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x99, 0x02, 0x0a, 0x12, 0x47, 0x72, 0x70, 0x63, 0x42, 0x66, 0x52, 0x65, 0x73, 0x50, 0x61,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x44, 0x6f, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x44, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x66, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x66, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a,
	0x16, 0x47, 0x72, 0x70, 0x63, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x44, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8e, 0x01, 0x0a, 0x18,
	0x47, 0x72, 0x70, 0x63, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x44, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x83, 0x01, 0x0a,
	0x19, 0x47, 0x72, 0x70, 0x63, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x44, 0x61,
	0x74, 0x65, 0x32, 0xed, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x72,
	0x70, 0x63, 0x12, 0x45, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x42, 0x6f, 0x6f, 0x6b, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x42, 0x6f, 0x6f, 0x6b, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x56, 0x69, 0x65,
	0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x42, 0x6f,
	0x6f, 0x6b, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x6e, 0x68, 0x6c, 0x76, 0x31, 0x31, 0x2f, 0x67, 0x6f,
	0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_booking_proto_goTypes = []interface{}{
	(*GrpcBookFlightRequest)(nil),     // 0: pb.GrpcBookFlightRequest
	(*GrpcBfReqPassenger)(nil),        // 1: pb.GrpcBfReqPassenger
	(*GrpcBookFlightResponse)(nil),    // 2: pb.GrpcBookFlightResponse
	(*GrpcBfResFlightOption)(nil),     // 3: pb.GrpcBfResFlightOption
	(*GrpcBfResPassenger)(nil),        // 4: pb.GrpcBfResPassenger
	(*GrpcViewBookingRequest)(nil),    // 5: pb.GrpcViewBookingRequest
	(*GrpcCancelBookingRequest)(nil),  // 6: pb.GrpcCancelBookingRequest
	(*GrpcCancelBookingResponse)(nil), // 7: pb.GrpcCancelBookingResponse
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_booking_proto_depIdxs = []int32{
	1,  // 0: pb.GrpcBookFlightRequest.passengers:type_name -> pb.GrpcBfReqPassenger
	8,  // 1: pb.GrpcBfReqPassenger.date_of_birth:type_name -> google.protobuf.Timestamp
	3,  // 2: pb.GrpcBookFlightResponse.flight_options:type_name -> pb.GrpcBfResFlightOption
	4,  // 3: pb.GrpcBookFlightResponse.passengers:type_name -> pb.GrpcBfResPassenger
	8,  // 4: pb.GrpcBfResFlightOption.departure_date_time:type_name -> google.protobuf.Timestamp
	8,  // 5: pb.GrpcBfResFlightOption.arrival_date_time:type_name -> google.protobuf.Timestamp
	8,  // 6: pb.GrpcBfResPassenger.date_of_birth:type_name -> google.protobuf.Timestamp
	8,  // 7: pb.GrpcCancelBookingResponse.cancel_date:type_name -> google.protobuf.Timestamp
	0,  // 8: pb.BookingGrpc.BookFlight:input_type -> pb.GrpcBookFlightRequest
	5,  // 9: pb.BookingGrpc.ViewBooking:input_type -> pb.GrpcViewBookingRequest
	6,  // 10: pb.BookingGrpc.CancelBooking:input_type -> pb.GrpcCancelBookingRequest
	2,  // 11: pb.BookingGrpc.BookFlight:output_type -> pb.GrpcBookFlightResponse
	2,  // 12: pb.BookingGrpc.ViewBooking:output_type -> pb.GrpcBookFlightResponse
	7,  // 13: pb.BookingGrpc.CancelBooking:output_type -> pb.GrpcCancelBookingResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcBookFlightRequest); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcBfReqPassenger); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcBookFlightResponse); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcBfResFlightOption); i {
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
		file_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcBfResPassenger); i {
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
		file_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcViewBookingRequest); i {
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
		file_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcCancelBookingRequest); i {
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
		file_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcCancelBookingResponse); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
