package bersamabilling

import "encoding/xml"

type CreatePaymentCodeRequest struct {
	XMLName         xml.Name
	URLCheckout     string `xml:"-"`
	Type            string `xml:"type"`
	BookingID       string `xml:"bookingid"`
	ClientID        string `xml:"clientid"`
	CustomerName    string `xml:"customer_name"`
	Amount          int64  `xml:"amount"`
	ProductID       string `xml:"productid"`
	Interval        int    `xml:"interval"`
	Username        string `xml:"username"`
	BookingDatetime string `xml:"booking_datetime"`
	Signature       string `xml:"signature"`
}

type StatusInquiryPaymentRequest struct {
	XMLName      xml.Name
	URLGetStatus string `xml:"-"`
	Type         string `xml:"type"`
	Item         []ItemRequest
}

type ItemRequest struct {
	XMLName         xml.Name
	BookingID       string `xml:"-"`
	VaID            string `xml:"vaid"`
	BookingDatetime string `xml:"booking_datetime"`
	Username        string `xml:"username"`
	Signature       string `xml:"signature"`
}
