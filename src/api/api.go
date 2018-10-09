package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

//Start levamta el servicio Api
func Start() {
	router := gin.Default()
	router.GET("/clima", clima)
	router.GET("/mover_galaxia", ejecutarMoverGalaxia)
	router.GET("/periodos_de_lluvia", periodosDeLluvia)
	router.GET("/periodos_optimos", periodosOptimos)
	router.GET("/periodos_sequia", periodosSequia)
	router.GET("/periodos_despejado", periodosDespejado)
	router.GET("/crear_esquema", crearEsquema)
	router.Run(":9990")
}

func clima(c *gin.Context) {
	dia := c.Query("dia")
	i, err := strconv.Atoi(dia)
	if err != nil {
		c.String(400, "400-Bad-Request: Asegurese de ingresar un numero entero")
		return
	}
	var periodo models.Periodo
	periodo = periodo.ObtenerPeriodo(i)
	if periodo == (models.Periodo{}) {
		c.JSON(200, gin.H{
			"dia":   dia,
			"clima": "No se cuenta con este dato, no esta dentro los proximos 10 anios",
		})
		return
	}

	c.JSON(200, gin.H{
		"dia":   dia,
		"clima": periodo.Clima,
	})
	return
}

func moverGalaxia(c *gin.Context) {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	res := galaxia.MoverGalaxia()
	c.JSON(200, gin.H{
		"resultado": res,
	})
	return
}

func ejecutarMoverGalaxia(c *gin.Context) {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	res := galaxia.MoverGalaxia()
	c.JSON(200, gin.H{
		"resultado": res,
	})
	return
}

func periodosDeLluvia(c *gin.Context) {
	var periodo models.Periodo
	cantidad, err := periodo.CantidadPeriodosClima("Lluvia")
	if err != nil {
		c.String(400, "400-Bad-Request")
		return
	}
	dia, err := periodo.PicoMaximoLluvia()
	c.JSON(200, gin.H{
		"cantidad":    cantidad,
		"pico_maximo": dia,
	})
	return
}

func periodosOptimos(c *gin.Context) {
	periodosClima(c, "Optimo")
	return
}

func periodosSequia(c *gin.Context) {
	periodosClima(c, "Sequia")
	return
}

func periodosDespejado(c *gin.Context) {
	periodosClima(c, "Despejado")
	return
}

func periodosClima(c *gin.Context, clima string) {
	var periodo models.Periodo
	cantidad, err := periodo.CantidadPeriodosClima(clima)
	if err != nil {
		c.String(400, "400-Bad-Request")
		return
	}
	c.JSON(200, gin.H{
		"cantidad": cantidad,
	})
	return
}

func crearEsquema(c *gin.Context) {
	var periodo models.Periodo
	err := periodo.CrearEsquema()
	if err != nil {
		c.String(400, "400-Bad-Request")
		return
	}
	c.JSON(200, gin.H{
		"res": "El esquema se ha creado correctamente",
	})
	return
}
