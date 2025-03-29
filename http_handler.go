package main

import (
	"net/http"
	"time"

	frt "github.com/tiendat-go/common-service/utils/format"
	jsn "github.com/tiendat-go/common-service/utils/json"
	pbCore "github.com/tiendat-go/proto-service/gen/core/v1"
	pbCrypto "github.com/tiendat-go/proto-service/gen/crypto/v1"
)

type handler struct {
	grpcClient *GrpcClient
}

func NewHandler(grpcClient *GrpcClient) *handler {
	return &handler{grpcClient}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/crypto", h.HandleGetCrypto)
	mux.HandleFunc("GET /api/crypto/v1/uiKlines", h.HandleGetKlines)
	mux.HandleFunc("GET /api/core/v1/sayHello/{name}", h.HandleSayHello)
}

func (h *handler) HandleGetCrypto(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	symbol := frt.GetString(query.Get("symbol"), "BTCUSDT")
	endTime := frt.GetInt64(query.Get("endTime"), time.Now().UnixMilli())
	limit := frt.GetInt(query.Get("limit"), 1000)
	interval := frt.GetString(query.Get("interval"), "1d")

	klines, err := h.grpcClient.Crypto.GetKlinesBySymbol(r.Context(), &pbCrypto.GetKlinesBySymbolRequest{
		Symbol:   symbol,
		EndTime:  endTime,
		Limit:    int32(limit),
		Interval: interval,
	})
	if err != nil {
		jsn.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsn.WriteJSON(w, http.StatusOK, klines)
}

func (h *handler) HandleGetKlines(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	symbol := frt.GetString(query.Get("symbol"), "BTCUSDT")
	endTime := frt.GetInt64(query.Get("endTime"), time.Now().UnixMilli())
	limit := frt.GetInt(query.Get("limit"), 1000)
	interval := frt.GetString(query.Get("interval"), "1d")

	klines, err := h.grpcClient.Crypto.GetKlinesBySymbol(r.Context(), &pbCrypto.GetKlinesBySymbolRequest{
		Symbol:   symbol,
		EndTime:  endTime,
		Limit:    int32(limit),
		Interval: interval,
	})
	if err != nil {
		jsn.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsn.WriteJSON(w, http.StatusOK, klines)
}

func (h *handler) HandleSayHello(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	response, err := h.grpcClient.Core.SayHello(r.Context(), &pbCore.SayHelloRequest{Name: name})
	if err != nil {
		jsn.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsn.WriteJSON(w, http.StatusOK, response)
}
