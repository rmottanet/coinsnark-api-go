package models

import (
    "time"
)


type ConversionRequest struct {
    From   string  `json:"from"`
    To     string  `json:"to"`
    Amount float64 `json:"amount"`
}


type ConversionResponse struct {
    ApiResponse
    Conversion        map[string]string `json:"conversion"`
}


func NewConversionResponse(conversion map[string]string, timestamp time.Time) *ConversionResponse {
    response := &ConversionResponse{
        Conversion: conversion,
    }
    response.ApiResponse = *NewApiResponse()
    response.Timestamp = timestamp.Format(time.RFC3339)

    return response
}
