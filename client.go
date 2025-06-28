package digiflazz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseURL = "https://api.digiflazz.com/v1"
)

// Client adalah klien untuk berinteraksi dengan API Digiflazz.
type Client struct {
	username   string
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// NewClient membuat instance baru dari Digiflazz Client.
func NewClient(username, apiKey string) *Client {
	return &Client{
		username:   username,
		apiKey:     apiKey,
		httpClient: &http.Client{},
		baseURL:    apiBaseURL,
	}
}

// _request adalah metode internal untuk melakukan semua request ke API.
func (c *Client) _request(ctx context.Context, endpoint string, body map[string]interface{}, signIdentifier string, responseContainer interface{}) error {
	body["username"] = c.username
	body["sign"] = generateSign(c.username, c.apiKey, signIdentifier)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d - %s", resp.StatusCode, string(respBody))
	}

	// Unmarshal ke wrapper 'data'
	wrappedResp := BaseResponse[json.RawMessage]{}
	if err := json.Unmarshal(respBody, &wrappedResp); err != nil {
		return fmt.Errorf("failed to unmarshal base response: %w", err)
	}

	// Cek response code (rc) di dalam payload 'data'
	// Kita perlu unmarshal sebagian untuk mendapatkan `rc`
	var rcChecker struct {
		ResponseCode ResponseCode `json:"rc"`
	}
	if err := json.Unmarshal(wrappedResp.Data, &rcChecker); err == nil {
		if rcChecker.ResponseCode != "" && rcChecker.ResponseCode != ResponseCodeSuccess {
			apiErr := &APIError{}
			_ = json.Unmarshal(wrappedResp.Data, apiErr)
			return apiErr
		}
	}

	// Unmarshal penuh ke container yang disediakan
	if err := json.Unmarshal(wrappedResp.Data, responseContainer); err != nil {
		return fmt.Errorf("failed to unmarshal final response data: %w", err)
	}

	return nil
}

// CheckBalance mengecek sisa saldo deposit.
func (c *Client) CheckBalance(ctx context.Context) (*CheckBalanceResponse, error) {
	var resp CheckBalanceResponse
	body := map[string]interface{}{"cmd": "deposit"}
	err := c._request(ctx, "/cek-saldo", body, "depo", &resp)
	return &resp, err
}

// RequestDeposit membuat tiket permintaan deposit.
func (c *Client) RequestDeposit(ctx context.Context, req DepositRequest) (*DepositResponse, error) {
	var resp DepositResponse
	body := map[string]interface{}{
		"amount":     req.Amount,
		"Bank":       req.Bank,
		"owner_name": req.OwnerName,
	}
	err := c._request(ctx, "/deposit", body, "deposit", &resp)
	return &resp, err
}

// PriceList mendapatkan daftar harga produk.
func (c *Client) PriceList(ctx context.Context, req PriceListRequest) ([]PriceListItem, error) {
	var resp []PriceListItem
	body := map[string]interface{}{"cmd": "prepaid"}
	if req.Cmd != "" {
		body["cmd"] = req.Cmd
	}
	// ... bisa ditambahkan filter lain seperti brand, category, dll.
	err := c._request(ctx, "/price-list", body, "pricelist", &resp)
	return resp, err
}

// TopUp melakukan pembelian produk prabayar.
func (c *Client) TopUp(ctx context.Context, req TransactionRequest) (*TransactionResponse, error) {
	var resp TransactionResponse
	body := map[string]interface{}{
		"buyer_sku_code": req.BuyerSkuCode,
		"customer_no":    req.CustomerNo,
		"ref_id":         req.RefID,
		"testing":        req.Testing,
	}
	err := c._request(ctx, "/transaction", body, req.RefID, &resp)
	return &resp, err
}

// InquiryPasca melakukan cek tagihan produk pascabayar.
func (c *Client) InquiryPasca(ctx context.Context, req TransactionRequest) (*InquiryPascaResponse, error) {
	var resp InquiryPascaResponse
	body := map[string]interface{}{
		"commands":       "inq-pasca",
		"buyer_sku_code": req.BuyerSkuCode,
		"customer_no":    req.CustomerNo,
		"ref_id":         req.RefID,
	}
	err := c._request(ctx, "/transaction", body, req.RefID, &resp)
	return &resp, err
}

// PayPasca melakukan pembayaran tagihan pascabayar.
func (c *Client) PayPasca(ctx context.Context, req TransactionRequest) (*InquiryPascaResponse, error) {
	var resp InquiryPascaResponse
	body := map[string]interface{}{
		"commands":       "pay-pasca",
		"buyer_sku_code": req.BuyerSkuCode,
		"customer_no":    req.CustomerNo,
		"ref_id":         req.RefID,
	}
	err := c._request(ctx, "/transaction", body, req.RefID, &resp)
	return &resp, err
}

// CheckStatus mengecek status transaksi yang sudah ada.
func (c *Client) CheckStatus(ctx context.Context, req TransactionRequest) (*TransactionResponse, error) {
	var resp TransactionResponse
	body := map[string]interface{}{
		"commands":       "status-pasca",
		"buyer_sku_code": req.BuyerSkuCode,
		"customer_no":    req.CustomerNo,
		"ref_id":         req.RefID,
	}
	err := c._request(ctx, "/transaction", body, req.RefID, &resp)
	return &resp, err
}

// InquiryPln memvalidasi ID Pelanggan PLN.
func (c *Client) InquiryPln(ctx context.Context, req InquiryPlnRequest) (*InquiryPlnResponse, error) {
	var resp InquiryPlnResponse
	body := map[string]interface{}{"customer_no": req.CustomerNo}
	// Signature untuk endpoint ini menggunakan customer_no sebagai identifier
	err := c._request(ctx, "/inquiry-pln", body, req.CustomerNo, &resp)
	return &resp, err
}