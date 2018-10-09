package models

import (
	"fmt"
	"math"
	"strconv"
)

//Galaxia es la estructura de un Galaxia
type Galaxia struct {
	dias     int
	planetas []Planeta
	sol      Planeta
	anios    int
}

//Iniciar carga todos los planetas en la galaxia
func (g Galaxia) Iniciar() Galaxia {
	ferengi := Planeta{Radio: 500, VelocidadGradosDia: 1, SentidoHorario: true, Name: "ferengi", X: 500, Y: 0, Grados: 0}
	betasoide := Planeta{Radio: 2000, VelocidadGradosDia: 3, SentidoHorario: true, Name: "betasoide", X: 200, Y: 0, Grados: 0}
	vulcano := Planeta{Radio: 1000, VelocidadGradosDia: 5, SentidoHorario: false, Name: "vulcano", X: 500, Y: 0, Grados: 0}
	g.planetas = append(g.planetas, ferengi)
	g.planetas = append(g.planetas, betasoide)
	g.planetas = append(g.planetas, vulcano)
	g.sol = Planeta{X: 0, Y: 0}
	g.dias = 50
	g.anios = 1
	return g
}

//MoverGalaxia es la funcion que calcula el clima de los proximos N Dias
func (g Galaxia) MoverGalaxia() string {
	cantDias := g.dias * g.anios
	var ultimoPeriodo Periodo
	ultimoPeriodo.EliminarPeriodos() //Limpiamos la BD
	for i := 1; i <= cantDias; i++ {
		var periodo Periodo
		fmt.Printf("DIA: %v \n", i)
		periodo = g.ObtenerClima(periodo, i)
		if periodo.Clima != ultimoPeriodo.Clima {
			fmt.Printf("primer: %+v ultimo: %+v \n", periodo, ultimoPeriodo)
			if ultimoPeriodo != (Periodo{}) {
				ultimoPeriodo.Guardar()
			}
			ultimoPeriodo = periodo
			ultimoPeriodo.Inicio = i
		}
		ultimoPeriodo.Fin = i
		//fmt.Printf(" %+v \n", periodo)
	}
	ultimoPeriodo.Guardar()
	cant, err := ultimoPeriodo.CantidadPeriodos()
	if err != nil {
		return err.Error()
	}
	return "Se calcularon " + strconv.Itoa(cant) + " nuevos periodos"
}

//ObtenerClima funcion para obtener el clima de determinado dia
func (g Galaxia) ObtenerClima(p Periodo, dia int) Periodo {
	g = g.ActualizarGalaxia(dia)
	p = g.Alineacion(p, dia)
	if p.Clima == "" {
		fmt.Print("entro a triangulacion\n")
		p = g.Triagulacion(p, dia)
	}
	return p
}

//Alineacion funcion que calcula si los planetas dentro la galaxia estan alineados
func (g Galaxia) Alineacion(p Periodo, dias int) Periodo {
	p.Clima = g.CondicionesClima()
	return p
}

//Triagulacion funcion que calcula si los planetas estan formando una triangulacion
func (g Galaxia) Triagulacion(p Periodo, dias int) Periodo {
	if g.IsSunIn() {
		p.Clima = "Lluvia"
		perimetro := g.GetPerimetro()
		if perimetro > p.Perimetro {
			p.Pico = dias
			p.Perimetro = perimetro
		}
	} else {
		p.Clima = "Despejado"
	}
	return p
}

//ActualizarGalaxia actualzia la ubicacion del los 3 planetas de la BD
func (g Galaxia) ActualizarGalaxia(dias int) Galaxia {
	for i := range g.planetas {
		g.planetas[i] = g.planetas[i].ActualizarUbicacion(dias)
	}
	return g
}

//IsSunIn dato 3 planetas calcula si el sol se encuentra en el medio
func (g Galaxia) IsSunIn() bool {
	var res = false
	a := g.planetas[0]
	b := g.planetas[1]
	c := g.planetas[2]

	var areaPlanetas = g.GetArea(a, b, c)
	var areaA = g.GetArea(a, b, g.sol)
	var areaB = g.GetArea(a, c, g.sol)
	var areaC = g.GetArea(b, c, g.sol)
	//	fmt.Printf("SIN REDONDEAR-ABC: %v areaPlanetas: %v \n", (areaA + areaB + areaC), areaPlanetas)
	var areaABC = math.Round((areaA+areaB+areaC)*100) / 100
	areaPlanetas = math.Round((areaPlanetas)*100) / 100
	//fmt.Printf("areaABC: %v areaPlanetas: %v \n", areaABC, areaPlanetas)
	if areaABC == areaPlanetas {
		res = true
		fmt.Print("ESTA DENTRO\n")
	} else {
		fmt.Print("esta fuera\n")
	}

	return res
}

//GetArea obtiene el area de 3 puntos
func (g Galaxia) GetArea(a Planeta, b Planeta, c Planeta) float64 {
	var area = (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2
	/*fmt.Print("Planetas\n")
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
	fmt.Printf("Area de los planetas %+v\n", area)
	fmt.Print("Fin Planetas\n")*/
	return math.Abs(area)
}

//CondicionesClima devuelve una cadena si las condiciones eran optimas o sequia
func (g Galaxia) CondicionesClima() string {
	res := ""
	var calc Calculadora
	if calc.Alineados(g.planetas[0], g.planetas[1], g.planetas[2]) {
		if calc.Alineados(g.planetas[0], g.planetas[1], g.sol) {
			res = "Sequia"
		} else {
			res = "Optimo"
		}
	}
	return res
}

// Distancia entre 2 puntos
func (g Galaxia) Distancia(a, b Planeta) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

//GetPerimetro obtener perimetro de un triangulo
func (g Galaxia) GetPerimetro() float64 {
	a := g.planetas[0]
	b := g.planetas[1]
	c := g.planetas[2]
	return (g.Distancia(a, b) + g.Distancia(b, c) + g.Distancia(c, a))
}
