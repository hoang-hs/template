package enums

import "github.com/thoas/go-funk"

type InvoicePublishStatus string
type InvoicePublishType string
type InvoiceDataCurrencyCode string
type InvoiceDataPaymentMethodName string
type InvoiceDataPaymentMethodNames []InvoiceDataPaymentMethodName
type InvoiceDataItemType string

const (
	InvoicePublishStatusNew       InvoicePublishStatus = "NEW"
	InvoicePublishStatusPublished InvoicePublishStatus = "PUBLISHED"
	InvoicePublishStatusError     InvoicePublishStatus = "ERROR"

	InvoicePublishTypePOS    InvoicePublishType = "POS"
	InvoicePublishTypeNormal InvoicePublishType = "NORMAL"

	InvoiceDataCurrencyCodeVND InvoiceDataCurrencyCode = "VND"

	InvoiceDataPaymentMethodNameCash     InvoiceDataPaymentMethodName = "Tiền mặt"
	InvoiceDataPaymentMethodNameTransfer InvoiceDataPaymentMethodName = "Chuyển khoản"
	InvoiceDataPaymentMethodNameCard     InvoiceDataPaymentMethodName = "Thẻ"
	InvoiceDataPaymentMethodNamePoint    InvoiceDataPaymentMethodName = "Điểm"
	InvoiceDataPaymentMethodNameVoucher  InvoiceDataPaymentMethodName = "Voucher"

	InvoiceDataItemTypeHHDV        InvoiceDataItemType = "HHDV"
	InvoiceDataItemTypePromotion   InvoiceDataItemType = "Khuyến mãi"
	InvoiceDataItemTypeDiscount    InvoiceDataItemType = "Chiết khấu"
	InvoiceDataItemTypeNoteExplain InvoiceDataItemType = "Ghi chú / diễn giải"
)

var (
	validInvoicePublishStatuses = []InvoicePublishStatus{
		InvoicePublishStatusNew,
		InvoicePublishStatusPublished,
		InvoicePublishStatusError,
	}

	validInvoicePublishTypes = []InvoicePublishType{
		InvoicePublishTypePOS,
		InvoicePublishTypeNormal,
	}

	ValidInvoiceDataCurrencyCodes = []InvoiceDataCurrencyCode{
		InvoiceDataCurrencyCodeVND,
	}

	ValidInvoiceDataPaymentMethodNames = []InvoiceDataPaymentMethodName{
		InvoiceDataPaymentMethodNameCash,
		InvoiceDataPaymentMethodNameTransfer,
		InvoiceDataPaymentMethodNameCard,
		InvoiceDataPaymentMethodNamePoint,
		InvoiceDataPaymentMethodNameVoucher,
	}

	ValidInvoiceDataItemTypes = []InvoiceDataItemType{
		InvoiceDataItemTypeHHDV,
		InvoiceDataItemTypePromotion,
		InvoiceDataItemTypeDiscount,
		InvoiceDataItemTypeNoteExplain,
	}
)

func (e InvoicePublishStatus) IsValid() bool {
	return funk.Contains(validInvoicePublishStatuses, e)
}

func (e InvoicePublishType) IsValid() bool {
	return funk.Contains(validInvoicePublishTypes, e)
}

func (e InvoiceDataCurrencyCode) IsValid() bool {
	return funk.Contains(ValidInvoiceDataCurrencyCodes, e)
}

func (e InvoiceDataPaymentMethodNames) IsValid() bool {
	return funk.Every(ValidInvoiceDataPaymentMethodNames, ConvertTypesToInterfaces[InvoiceDataPaymentMethodName](e)...)
}

func (e InvoiceDataItemType) IsValid() bool {
	return funk.Contains(ValidInvoiceDataItemTypes, e)
}

// ConvertTypesToInterfaces Todo consider
func ConvertTypesToInterfaces[T any](tt []T) []interface{} {
	var ii []interface{}
	for _, t := range tt {
		ii = append(ii, t)
	}
	return ii
}
