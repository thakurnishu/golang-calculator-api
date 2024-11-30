package route

type Numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}
type SumNumber struct {
	Nums []float64 `json:"nums"`
}

type ResultResponse struct {
	Result float64 `json:"result"`
}
