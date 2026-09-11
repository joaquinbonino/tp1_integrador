package models

// Pedido representa un pedido confirmado por un cliente.
type Pedido struct {
	ID         string `json:"id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
	Estado     string `json:"estado"`
}
