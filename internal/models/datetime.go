package models

type DateTime struct {

	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	// datetime without a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Required. Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	// month.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	// may choose to allow the value "24:00:00" for scenarios like business
	// closing time.
	Hours int32 `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	// Required. Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,5,opt,name=minutes,proto3" json:"minutes,omitempty"`
	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	// API may allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,6,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	// 999,999,999.
	Nanos int32 `protobuf:"varint,7,opt,name=nanos,proto3" json:"nanos,omitempty"`
	// Optional. Specifies either the UTC offset or the time zone of the DateTime.
	// Choose carefully between them, considering that time zone data may change
	// in the future (for example, a country modifies their DST start/end dates,
	// and future DateTimes in the affected range had already been stored).
	// If omitted, the DateTime is considered to be in local time.
	//
	// Types that are assignable to TimeOffset:
	//	*DateTime_UtcOffset
	//	*DateTime_TimeZone
	// TimeOffset isDateTime_TimeOffset `protobuf_oneof:"time_offset"`
	// contains filtered or unexported fields
}
