package dto

type ShortenQueryParam struct {
	Original string `query:"url" validate:"required,url"`
}

type RedirectParam struct {
	Code string `params:"code" validate:"required,len=8"`
}
