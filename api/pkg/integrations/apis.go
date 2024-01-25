package integrations

// Integrations URLs
var urls = map[string]string{
	"bcb_exchange":  "https://www.bcb.gov.br/api/servico/sitebcb/indicadorCambio",
	"bcb_rates": "https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoMoedaDia(moeda=@moeda,dataCotacao=@dataCotacao)",
	"exchange_rates":  "https://exchange-rates.abstractapi.com/v1/live/",
	"fixer":  "https://api.apilayer.com/fixer/latest",
	"open_exchanges_rates":  "https://openexchangerates.org/api/latest.json",
}


func GetUrl(index string) string {
	return urls[index]
}
