// types.go
package digiflazz

// ResponseCode merepresentasikan kode respon dari API Digiflazz.
type ResponseCode string

// Konstanta untuk semua ResponseCode yang diketahui.
const (
	ResponseCodeSuccess               ResponseCode = "00"
	ResponseCodeTimeout               ResponseCode = "01"
	ResponseCodeFailed                ResponseCode = "02"
	ResponseCodePending               ResponseCode = "03"
	ResponseCodePayloadError          ResponseCode = "40"
	ResponseCodeInvalidSignature      ResponseCode = "41"
	ResponseCodeSKUNotFound           ResponseCode = "43"
	ResponseCodeInsufficientBalance   ResponseCode = "44"
	ResponseCodeIPNotRecognized       ResponseCode = "45"
	ResponseCodeDuplicateRefID        ResponseCode = "49"
	ResponseCodeTransactionNotFound   ResponseCode = "50"
	ResponseCodeProductUnavailable    ResponseCode = "53"
	ResponseCodeCutOff                ResponseCode = "58"
	ResponseCodeStockEmpty            ResponseCode = "68"
	ResponseCodeTransactionRefunded   ResponseCode = "74"
	ResponseCodeAccountNotVerified    ResponseCode = "82"
	ResponseCodeLimitReached          ResponseCode = "85"
)

// baseResponse adalah struktur generik untuk membungkus semua respons dari API.
type BaseResponse[T any] struct {
	Data T `json:"data"`
}

// CheckBalanceResponse adalah respons dari endpoint Cek Saldo.
type CheckBalanceResponse struct {
	Deposit float64 `json:"deposit"`
}

// PriceListRequest adalah parameter untuk meminta daftar harga.
type PriceListRequest struct {
	Cmd      string `json:"cmd,omitempty"`      // 'prepaid' atau 'pasca'
	Category string `json:"category,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Type     string `json:"type,omitempty"`
}

// PriceListItem merepresentasikan satu item dalam daftar harga.
type PriceListItem struct {
	ProductName        string  `json:"product_name"`
	Category           string  `json:"category"`
	Brand              string  `json:"brand"`
	Type               string  `json:"type"`
	SellerName         string  `json:"seller_name"`
	Price              float64 `json:"price"`
	BuyerSkuCode       string  `json:"buyer_sku_code"`
	BuyerProductStatus bool    `json:"buyer_product_status"`
	Stock              int     `json:"stock"`
	Desc               string  `json:"desc"`
}

// TransactionRequest adalah parameter dasar untuk transaksi.
type TransactionRequest struct {
	BuyerSkuCode string `json:"buyer_sku_code"`
	CustomerNo   string `json:"customer_no"`
	RefID        string `json:"ref_id"`
	Testing      bool   `json:"testing,omitempty"`
	MaxPrice     int    `json:"max_price,omitempty"`
}

// TransactionResponse adalah respons umum untuk transaksi.
type TransactionResponse struct {
	RefID          string       `json:"ref_id"`
	CustomerNo     string       `json:"customer_no"`
	BuyerSkuCode   string       `json:"buyer_sku_code"`
	Message        string       `json:"message"`
	Status         string       `json:"status"`
	ResponseCode   ResponseCode `json:"rc"`
	SN             string       `json:"sn,omitempty"`
	BuyerLastSaldo float64      `json:"buyer_last_saldo"`
	Price          float64      `json:"price"`
	Tele           string       `json:"tele,omitempty"`
	Wa             string       `json:"wa,omitempty"`
}

// InquiryPascaResponse adalah respons spesifik untuk inquiry pascabayar.
type InquiryPascaResponse struct {
	TransactionResponse
	CustomerName string                 `json:"customer_name"`
	Admin        float64                `json:"admin"`
	SellingPrice float64                `json:"selling_price"`
	Desc         map[string]interface{} `json:"desc"`
}

// BankName adalah tipe untuk nama bank yang didukung.
type BankName string

const (
	BankBCA     BankName = "BCA"
	BankMandiri BankName = "MANDIRI"
	BankBRI     BankName = "BRI"
	BankBNI     BankName = "BNI"
)

// DepositRequest adalah parameter untuk meminta tiket deposit.
type DepositRequest struct {
	Amount    float64  `json:"amount"`
	Bank      BankName `json:"Bank"`
	OwnerName string   `json:"owner_name"`
}

// DepositResponse adalah respons dari permintaan tiket deposit.
type DepositResponse struct {
	ResponseCode ResponseCode `json:"rc"`
	Amount       float64      `json:"amount"`
	Notes        string       `json:"notes"`
}

// InquiryPlnRequest adalah parameter untuk inquiry PLN.
type InquiryPlnRequest struct {
	CustomerNo string `json:"customer_no"`
}

// InquiryPlnResponse adalah respons dari inquiry PLN.
type InquiryPlnResponse struct {
	Message      string       `json:"message"`
	Status       string       `json:"status"`
	ResponseCode ResponseCode `json:"rc"`
	CustomerNo   string       `json:"customer_no"`
	MeterNo      string       `json:"meter_no,omitempty"`
	SubscriberID string       `json:"subscriber_id,omitempty"`
	Name         string       `json:"name,omitempty"`
	SegmentPower string       `json:"segment_power,omitempty"`
}

// WebhookTransactionPayload adalah payload untuk event 'create' & 'update'.
type WebhookTransactionPayload TransactionResponse

// PingEventPayload adalah payload untuk event 'ping'.
type PingEventPayload struct {
	Sed    string `json:"sed"`
	HookID string `json:"hook_id"`
	Hook   struct {
		URL    string `json:"url"`
		Secret string `json:"secret"`
		Type   string `json:"type"`
		Status int    `json:"status"`
	} `json:"hook"`
}