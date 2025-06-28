// types.go
package digiflazz

// ResponseCode merepresentasikan kode respon dari API Digiflazz.
type ResponseCode string

// Konstanta lengkap untuk semua ResponseCode yang diketahui dari dokumentasi Digiflazz.
const (
	// Sukses & Pending
	ResponseCodeSuccess ResponseCode = "00" // Transaksi Sukses
	ResponseCodePending ResponseCode = "03" // Transaksi Pending

	// Kegagalan Umum Transaksi
	ResponseCodeTimeout             ResponseCode = "01" // Timeout
	ResponseCodeTransactionFailed   ResponseCode = "02" // Transaksi Gagal
	ResponseCodeTransactionRefunded ResponseCode = "74" // Transaksi Refund

	// Kegagalan Validasi & Sistem Internal (4x series)
	ResponseCodePayloadError          ResponseCode = "40" // Payload Error
	ResponseCodeInvalidSignature      ResponseCode = "41" // Signature tidak valid
	ResponseCodeAPIBuyerProcessFailed ResponseCode = "42" // Gagal memproses API Buyer
	ResponseCodeSKUNotFound           ResponseCode = "43" // SKU tidak di temukan atau Non-Aktif
	ResponseCodeInsufficientBalance   ResponseCode = "44" // Saldo tidak cukup
	ResponseCodeIPNotRecognized       ResponseCode = "45" // IP Anda tidak kami kenali
	ResponseCodeTransactionExistsOtherBuyer ResponseCode = "47" // Transaksi sudah terjadi di buyer lain
	ResponseCodeDuplicateRefID        ResponseCode = "49" // Ref ID tidak unik

	// Kegagalan Terkait Produk & Tujuan (5x series)
	ResponseCodeTransactionNotFound   ResponseCode = "50" // Transaksi Tidak Ditemukan
	ResponseCodeDestinationBlocked    ResponseCode = "51" // Nomor Tujuan Diblokir
	ResponseCodePrefixInvalid         ResponseCode = "52" // Prefix Tidak Sesuai Dengan Operator
	ResponseCodeProductSellerUnavailable ResponseCode = "53" // Produk Seller Sedang Tidak Tersedia
	ResponseCodeDestinationInvalid    ResponseCode = "54" // Nomor Tujuan Salah
	ResponseCodeProductIssue          ResponseCode = "55" // Produk Sedang Gangguan
	ResponseCodeSellerBalanceLimit    ResponseCode = "56" // Limit saldo seller
	ResponseCodeInvalidDigitCount     ResponseCode = "57" // Jumlah Digit Kurang Atau Lebih
	ResponseCodeCutOff                ResponseCode = "58" // Sedang Cut Off
	ResponseCodeDestinationOutOfArea  ResponseCode = "59" // Tujuan di Luar Wilayah/Cluster

	// Kegagalan Terkait Tagihan & Deposit (6x series)
	ResponseCodeBillNotAvailable      ResponseCode = "60" // Tagihan belum tersedia
	ResponseCodeNoDepositHistory      ResponseCode = "61" // Belum pernah melakukan deposit
	ResponseCodeSellerIssue           ResponseCode = "62" // Seller sedang mengalami gangguan
	ResponseCodeMultiTransactionUnsupported ResponseCode = "63" // Tidak support transaksi multi
	ResponseCodeDepositTicketFailed   ResponseCode = "64" // Tarik tiket gagal
	ResponseCodeMultiTransactionLimit ResponseCode = "65" // Limit transaksi multi
	ResponseCodeCutOffSellerSystem    ResponseCode = "66" // Cut Off (Perbaikan Sistem Seller)
	ResponseCodeSellerNotVerified     ResponseCode = "67" // Seller belum ter-verfikasi
	ResponseCodeStockEmpty            ResponseCode = "68" // Stok habis
	ResponseCodePriceExceedsLimit     ResponseCode = "69" // Harga seller lebih besar dari ketentuan harga Buyer

	// Kegagalan Lanjutan & Spesifik (7x series)
	ResponseCodeBillerTimeout      ResponseCode = "70" // Timeout Dari Biller
	ResponseCodeProductUnstable    ResponseCode = "71" // Produk Sedang Tidak Stabil
	ResponseCodeUnregPackageRequired ResponseCode = "72" // Lakukan Unreg Paket Dahulu
	ResponseCodeKwhExceedsLimit    ResponseCode = "73" // Kwh Melebihi Batas

	// Kegagalan Terkait Akun & Limitasi (8x series)
	ResponseCodeAccountBlockedBySeller  ResponseCode = "80" // Akun Anda telah diblokir oleh Seller
	ResponseCodeSellerBlockedByYou    ResponseCode = "81" // Seller ini telah diblokir oleh Anda
	ResponseCodeAccountNotVerified    ResponseCode = "82" // Akun Anda belum ter-verfikasi
	ResponseCodePriceListLimitReached ResponseCode = "83" // Anda telah mencapai limitasi pengecekan pricelist
	ResponseCodeInvalidAmount         ResponseCode = "84" // Nominal tidak valid
	ResponseCodeTransactionLimitReached ResponseCode = "85" // Anda telah mencapai limitasi transaksi
	ResponseCodePlnCheckLimitReached  ResponseCode = "86" // Anda telah mencapai limitasi pengecekan nomor PLN

	// Lain-lain
	ResponseCodeRouterIssue ResponseCode = "99" // DF Router Issue (Pending)
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