package router

type PruebaRoute struct {
	nombre string
}

func (p *PruebaRoute) visualizarNombre() string {
	return p.nombre
}
