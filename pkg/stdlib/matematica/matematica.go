package matematica

import "math"

// Absoluto retorna o valor absoluto
func Absoluto(x float64) float64 { return math.Abs(x) }

// Teto arredonda para cima
func Teto(x float64) float64 { return math.Ceil(x) }

// Piso arredonda para baixo
func Piso(x float64) float64 { return math.Floor(x) }

// Maximo retorna o maior valor
func Maximo(a, b float64) float64 { return math.Max(a, b) }

// Minimo retorna o menor valor
func Minimo(a, b float64) float64 { return math.Min(a, b) }

// Potencia eleva a base ao expoente
func Potencia(base, exp float64) float64 { return math.Pow(base, exp) }

// Raiz retorna a raiz quadrada
func Raiz(x float64) float64 { return math.Sqrt(x) }
