package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func NuevoMedicine(name string, manufacturer string, manufacturedate time.Time, shelflife int) *Medicine {
	if name == "" {
		fmt.Println("el nombre no puede estar vacio")
	}
	if manufacturer == "" {
		fmt.Println("el fabricante no puede estar vacio")
	}
	if shelflife <= 0 {
		fmt.Println("la vida util debe ser positiva")
	}
	return &Medicine{
		name:            name,
		manufacturer:    manufacturer,
		manufactureDate: manufacturedate,
		shelfLife:       shelflife,
	}
}
func (c Medicine) GetName() string {
	return c.name
}
func (c Medicine) GetManufacturer() string {
	return c.manufacturer
}
func (c Medicine) GetManufactureDate() time.Time {
	return c.manufactureDate
}
func (c Medicine) GetShelLife() int {
	return c.shelfLife
}
func (c *Medicine) SetName(nuevoname string) {
	nuevoname = c.name
}
func (c *Medicine) SetManufacturer(nuevomanufacturer string) {
	nuevomanufacturer = c.manufacturer
}
func (c *Medicine) SetShelfLife(nuevoshelflive int) {
	nuevoshelflive = c.shelfLife
}

type Tablet struct {
	Medicine
	dosePertablet  float64
	isPrescription bool
}
type Syrup struct {
	Medicine
	volumen float64
	flavor  string
}

func NuevoTablet(name string, manufacturer string, manufacturedate time.Time, shelflife int, dosePertable float64, isPrescription bool) *Tablet {
	if dosePertable <= 0 {
		fmt.Println("la dosis por tabla debe ser mayor a 0")
	}
	medicine := NuevoMedicine(name, manufacturer, manufacturedate, shelflife)
	return &Tablet{
		Medicine:       *medicine,
		dosePertablet:  dosePertable,
		isPrescription: isPrescription,
	}
}
func NuevoSyrup(name string, manufacturer string, manufacturedate time.Time, shelflife int, volumen float64, flavor string) *Syrup {
	if volumen <= 0 {
		fmt.Println("el volumen debe ser mayor a 0")
	}
	medicine := NuevoMedicine(name, manufacturer, manufacturedate, shelflife)
	return &Syrup{
		Medicine: *medicine,
		volumen:  volumen,
		flavor:   flavor,
	}
}
func (c Medicine) CalcularCaducidad() time.Time {
	return c.manufactureDate.AddDate(0, c.shelfLife, 0)
}
func (t Tablet) MostrarDetalles() {
	fmt.Println("TABLET")
	fmt.Println("Nombre:", t.name)
	fmt.Println("Fabricante:", t.manufacturer)
	fmt.Println("Fecha de fabricacion:", t.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida util:", t.shelfLife, "meses")
	fmt.Println("Dosis por tableta:", t.dosePertablet, "mg")
	fmt.Println("Requiere receta:", t.isPrescription)
	fmt.Println("Fecha de caducidad:", t.CalcularCaducidad().Format("2006-01-02"))
	fmt.Println()
}
func (s Syrup) MostrarDetalles() {
	fmt.Println("SYRUP")
	fmt.Println("Nombre:", s.name)
	fmt.Println("Fabricante:", s.manufacturer)
	fmt.Println("Fecha de fabricacion:", s.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida util:", s.shelfLife, "meses")
	fmt.Println("Volumen:", s.volumen, "ml")
	fmt.Println("Sabor:", s.flavor)
	fmt.Println("Fecha de caducidad:", s.CalcularCaducidad().Format("2006-01-02"))
	fmt.Println()
}
func main() {
	inventario := []Tablet{}
	inventario2 := []Syrup{}

	fecha1 := time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)
	fecha2 := time.Date(2025, 3, 5, 0, 0, 0, 0, time.UTC)

	tablet1 := NuevoTablet("pepe", "jorgitoelguayako", fecha1, 24, 500, true)
	tablet2 := NuevoTablet("lololo", "homelander", fecha2, 24, 500, true)

	syrup1 := NuevoSyrup("galletasMaria", "pabloemilio", fecha1, 12, 120, "Cereza")
	syrup2 := NuevoSyrup("galletasMaria", "pabloemilio", fecha2, 12, 120, "Cereza")

}
