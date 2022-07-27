package data

import (
	"encoding/json"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/helper"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/invoice"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type VAController interface {
	AllVA(ctx *gin.Context)
	CreateInvoice(ctx *gin.Context)
	CallbackTransaksi(ctx *gin.Context)
	HistoryTransaksi(ctx *gin.Context)
	DetailHistoryTransaksi(ctx *gin.Context)
	AllTransaksi(ctx *gin.Context)
}

type vaController struct {
	xenService *client.API
	txService service.TransaksiService
	detailTxService service.DetailTransaksiService
}

func (v vaController) AllTransaksi(ctx *gin.Context) {
	var transaksis []entity.Transaksi = v.txService.AllTransaksi()
	resp := helper.BuildResponse(true, "Get Data Berhasil", transaksis)
	ctx.JSON(http.StatusOK, resp)
}

func (v vaController) DetailHistoryTransaksi(ctx *gin.Context) {
	transaksi_id, err := strconv.ParseUint(ctx.Param("transaksi_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter transaksi_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		res := v.detailTxService.DetailTransaksi(uint32(transaksi_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (v vaController) HistoryTransaksi(ctx *gin.Context) {
	kustomer_id, err := strconv.ParseUint(ctx.Param("customer_id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter kustomer_id yang ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}else {
		res := v.txService.TransaksiByCustomer(uint32(kustomer_id))
		resp := helper.BuildResponse(true, "Get Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (v vaController) CallbackTransaksi(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {

	}
	var dataForCallback map[string]string
	json.Unmarshal([]byte(jsonData), &dataForCallback)

	externalId := dataForCallback["external_id"]
	externalSplit := strings.Split(externalId, "_")

	id_tx, _ := strconv.ParseUint(externalSplit[2], 10, 64)

	if dataForCallback["status"] == "PAID" {
		var upTransaksi dto.UpdateTransaksi
		var upDetailTransaksi dto.UpdateDetailTransaksi

		upTransaksi.PaymentMethods = dataForCallback["payment_channel"]
		upTransaksi.Status = "1"

		upDetailTransaksi.PaymentMethods = dataForCallback["payment_channel"]
		upDetailTransaksi.Status = "1"

		_ = v.txService.UpdateTransaksi(uint32(id_tx), upTransaksi)
		_ = v.detailTxService.UpdateDetailTransaksi(uint32(id_tx), upDetailTransaksi)
	} else if dataForCallback["status"] == "EXPIRED" {
		var upTransaksi dto.UpdateTransaksi
		var upDetailTransaksi dto.UpdateDetailTransaksi

		//upTransaksi.PaymentMethods = dataForCallback["payment_channel"]
		upTransaksi.Status = "0"

		//upDetailTransaksi.PaymentMethods = dataForCallback["payment_channel"]
		upDetailTransaksi.Status = "0"

		_ = v.txService.UpdateTransaksi(uint32(id_tx), upTransaksi)
		_ = v.detailTxService.UpdateDetailTransaksi(uint32(id_tx), upDetailTransaksi)
	}
}

func (v vaController) CreateInvoice(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {

	}

	xendit.Opt.SecretKey = "xnd_development_kenbzs6xGEMYzUGQv8Sd84NV4nFT33R4arYuOndlqvoS310PwMYmQbZmlxSpSpG"
	var dataForTransaksi map[string]string
	json.Unmarshal([]byte(jsonData), &dataForTransaksi)

	var dataTotal map[string]int
	json.Unmarshal([]byte(jsonData), &dataTotal)

	id_customer, _ := strconv.ParseUint(dataForTransaksi["id"], 10, 64)
	newTransaksi := entity.Transaksi{
		Customer:   	uint32(id_customer),
		TotalPrice: 	uint64(dataTotal["totalHarga"]),
		Status:         "2",
	}

	tglServ := dataForTransaksi["tgl"]
	splits := strings.Split(tglServ, "T")
	concateTGL := splits[0]+" "+splits[1]+":00"
	newTGL := carbon.Parse(concateTGL).Time

	alamatUser := dataForTransaksi["alamat"]

	res := v.txService.CreateTransaksi(newTransaksi)

	var dataProduk map[string][]map[string]string
	json.Unmarshal([]byte(jsonData), &dataProduk)

	var items []xendit.InvoiceItem
	var fees []xendit.InvoiceFee

	for _, value := range dataProduk["cart"]{
		idSeller, _ := strconv.ParseInt(value["id_seller"],10, 64)
		idProduk, _ := strconv.ParseInt(value["id_produk"],10, 64)
		quantity, _ := strconv.ParseFloat(value["quantity"],10)
		hargaNormal, _ := strconv.ParseFloat(value["harga_normal"],10)
		dp, _ := strconv.ParseFloat(value["dp"],10)
		hargaxQuantity := hargaNormal*quantity
		var decimalPercent = 0.01
		desimalDP := dp * decimalPercent
		hargaxdpxQuantity := hargaxQuantity*desimalDP

		PotFee := -(hargaxQuantity - (hargaxQuantity*desimalDP))
		//hargaxdp := hargaNormal*decimalPercent

		//convertHargaNormal := hargaNormal - hargaxdp

		newDetailTX := entity.Detail_Transaksi{
			IdTransaksi:    res.Id_Transaksi,
			Customer:       uint32(id_customer),
			Alamat: 		alamatUser,
			Seller:         uint32(idSeller),
			Produk:         uint32(idProduk),
			Pax:            uint64(quantity),
			DownPayment:    uint64(hargaxdpxQuantity),
			TanggalRes:     newTGL,
			PaymentMethods: "",
			Status:         "2",
		}

		_ = v.detailTxService.CreateDetailTransaksi(newDetailTX)

		item := xendit.InvoiceItem{
			Name: value["nama_produk"],
			Price: hargaNormal,
			Quantity: int(quantity),
		}

		items = append(items, item)

		fee := xendit.InvoiceFee{
			Type:         "DP" + strconv.Itoa(int(res.Id_Transaksi)) + strconv.Itoa(int(idProduk)),
			Value:        PotFee,
		}

		fees = append(fees, fee)
	}

	cus := xendit.InvoiceCustomer{
		Email:        dataForTransaksi["email"],
	}

	NotificationType := []string{"email"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID: "demo_123124123_" + strconv.Itoa(int(res.Id_Transaksi)),
		Amount: float64(dataTotal["totalHarga"]),
		PayerEmail: dataForTransaksi["email"],
		Description: "Wedding",
		InvoiceDuration: 60,
		Items: items,
		Customer: cus,
		CustomerNotificationPreference: customerNotificationPreference,
		Fees: fees,
	}

	resp, errCreateInvoice := invoice.Create(&data)

	if errCreateInvoice != nil {
		response := helper.BuildErrorResponse("Failed To Checkout", errCreateInvoice.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		resss := helper.BuildResponse(true, "Checkout Berhasil", resp)
		ctx.JSON(http.StatusOK, resss)
	}

	//fmt.Printf("created invoice: %+v\n", resp)
}

func (v vaController) AllVA(ctx *gin.Context) {
	availableBanks, _ := v.xenService.VirtualAccount.GetAvailableBanks()
	res := helper.BuildResponse(true, "Get Data Berhasil", availableBanks)
	ctx.JSON(http.StatusOK, res)
}

func NewVAController(xenServ *client.API, txServ service.TransaksiService, detailTxServ service.DetailTransaksiService) VAController {
	return &vaController{
		xenService: xenServ,
		txService: txServ,
		detailTxService: detailTxServ,
	}
}
