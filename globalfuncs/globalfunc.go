package globalfuncs

type ValorParaPositivar struct {
	ValorVariavel float64
}

func (c *ValorParaPositivar) TornaPositivo(v float64) float64 {
	if v >= 0 {
		return v
	} else {
		return v * -1
	}
}
