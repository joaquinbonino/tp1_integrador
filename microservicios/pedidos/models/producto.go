package models

// Producto representa un producto del catálogo.
type Producto struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Stock  int    `json:"stock"`
}
