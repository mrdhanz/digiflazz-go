# Digiflazz Go SDK

SDK Klien Go yang modern, tangguh, dan idiomatis untuk berinteraksi dengan [API Buyer Digiflazz](https://developer.digiflazz.com/api/buyer).

**Disclaimer:** Pustaka ini dikembangkan secara independen dan tidak berafiliasi dengan Digiflazz.

## Fitur

- ✅ **Idiomatis**: Ditulis dengan mengikuti praktik terbaik Go, termasuk penggunaan `context.Context` dan penanganan error yang jelas.
- 🛡️ **Tipe Aman**: Semua request dan response dimodelkan dengan `struct` Go.
- 📦 **Tanpa Dependensi**: Hanya menggunakan pustaka standar Go.
- 🚨 **Penanganan Error Jelas**: Menyediakan `APIError` kustom untuk error dari sisi API.
- 🌐 **Cakupan Penuh**: Mendukung semua endpoint API Buyer, termasuk deposit dan verifikasi webhook.

## Instalasi

```bash
go get github.com/mrdhanz/digiflazz-go
```

## Penggunaan

### Inisialisasi Klien

```go
import (
    "context"
    "fmt"
    "log"
    "github.com/mrdhanz/digiflazz-go"
)

func main() {
    username := "USERNAME_ANDA"
    apiKey := "API_KEY_ANDA"

    client := digiflazz.NewClient(username, apiKey)

    // Contoh: Cek Saldo
    ctx := context.Background()
    saldo, err := client.CheckBalance(ctx)
    if err != nil {
        log.Fatalf("Gagal mengecek saldo: %v", err)
    }

    fmt.Printf("Saldo Anda saat ini: %.2f\n", saldo.Deposit)
}
```

### Penanganan Error

Library ini membedakan antara error jaringan dan error dari API Digiflazz. Anda dapat menggunakan *type assertion* untuk menangani `APIError` secara spesifik.

```go
    // Contoh: Transaksi dengan SKU yang salah
    _, err := client.TopUp(ctx, digiflazz.TransactionRequest{
        BuyerSkuCode: "sku-tidak-valid",
        CustomerNo:   "081234567890",
        RefID:        "unique-ref-id-123",
    })
    if err != nil {
        var apiErr *digiflazz.APIError
        if errors.As(err, &apiErr) {
            // Ini adalah error yang dikembalikan oleh API
            fmt.Printf("Error dari API: %s\n", apiErr.Message)
            fmt.Printf("Response Code: %s\n", apiErr.ResponseCode)
            if apiErr.ResponseCode == digiflazz.ResponseCodeSKUNotFound {
                fmt.Println("Aksi: Harap periksa kembali SKU produk.")
            }
        } else {
            // Ini adalah error lain (jaringan, parsing, dll)
            log.Fatalf("Error tak terduga: %v", err)
        }
    }
```

### Penanganan Webhook

Gunakan fungsi `VerifyWebhookSignature` di dalam HTTP handler Anda untuk memvalidasi request yang masuk.

**Contoh dengan `net/http`:**

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    webhookSecret := "SECRET_KEY_ANDA"

    // Baca raw body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Cannot read body", http.StatusBadRequest)
        return
    }

    // Verifikasi signature
    signature := r.Header.Get("X-Hub-Signature")
    if err := digiflazz.VerifyWebhookSignature(webhookSecret, body, signature); err != nil {
        log.Printf("Webhook verification failed: %v", err)
        http.Error(w, "Invalid signature", http.StatusForbidden)
        return
    }

    // Signature valid, proses payload...
    var payload digiflazz.WebhookTransactionPayload
    if err := json.Unmarshal(body.Data, &payload); err != nil {
        http.Error(w, "Cannot parse payload", http.StatusBadRequest)
        return
    }
    
    fmt.Printf("Webhook diterima untuk ref_id: %s, status: %s\n", payload.RefID, payload.Status)

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Webhook diterima")
}

func main() {
    http.HandleFunc("/webhook", webhookHandler)
    log.Println("Server webhook berjalan di :8080")
    http.ListenAndServe(":8080", nil)
}
```

## Referensi API (Contoh)

- `client.CheckBalance(ctx)`
- `client.RequestDeposit(ctx, req)`
- `client.PriceList(ctx, req)`
- `client.TopUp(ctx, req)`
- `client.InquiryPasca(ctx, req)`
- `client.PayPasca(ctx, req)`
- `client.CheckStatus(ctx, req)`
- `client.InquiryPln(ctx, req)`

## Lisensi

[MIT](LICENSE)