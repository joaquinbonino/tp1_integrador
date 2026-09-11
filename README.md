# TP Integrador Clases 1 a 4 — e-commerce y logística
Consigna: https://github.com/TomiCassanelli/arquitectura-de-software/tree/main/TRABAJO_PRACTICO_1_4

Partiendo del monolito en `monolito/` (clientes, productos y pedidos en el
mismo proceso, puerto `8080`), se separaron las responsabilidades en dos
microservicios independientes dentro de `microservicios/`:

```text
Antes:
cliente -> monolito :8080

Después:
cliente -> clientes :8081
        -> pedidos  :8082

pedidos -- evento pedido.confirmado --> RabbitMQ (cola pedidos-confirmados) --> logística (conceptual)
```

## Estructura

```text
tp1_integrador/
├── docker-compose.yml          # RabbitMQ (user/pass en localhost:5672)
├── monolito/                   # Punto de partida (sin modificar)
└── microservicios/
    ├── clientes/                # :8081 — POST /clientes, GET /clientes/:id
    │   ├── controllers/
    │   ├── services/
    │   ├── repositories/
    │   └── models/
    ├── pedidos/                  # :8082 — GET /productos, POST /pedidos
    │   ├── controllers/
    │   ├── services/
    │   ├── repositories/
    │   ├── cache/                # caché en memoria con TTL para GET /productos
    │   ├── messaging/            # publisher de RabbitMQ (pedido.confirmado)
    │   └── models/
    └── eventos/                  # documentación del contrato de eventos
```

Cada microservicio es su propio módulo de Go, con datos en memoria y capas
separadas (controlador → servicio → repositorio).

## Cómo correrlo

1. Levantar RabbitMQ (necesario para `pedidos`, que se conecta al arrancar):

   ```bash
   docker compose up -d
   ```

2. En una terminal, levantar `clientes`:

   ```bash
   cd microservicios/clientes
   go mod tidy
   go run .
   ```

3. En otra terminal, levantar `pedidos`:

   ```bash
   cd microservicios/pedidos
   go mod tidy
   go run .
   ```

## Probarlo

```bash
# Crear un cliente
curl -X POST http://localhost:8081/clientes \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Ana Pérez","email":"ana@mail.com"}'

# Consultar un cliente
curl http://localhost:8081/clientes/C-1

# Listar productos (primera vez pega al repositorio, siguientes usan la caché)
curl http://localhost:8082/productos

# Confirmar un pedido -> publica pedido.confirmado en la cola pedidos-confirmados
curl -X POST http://localhost:8082/pedidos \
  -H "Content-Type: application/json" \
  -d '{"cliente_id":"C-1","producto_id":"P-1"}'
```

Podés ver la cola `pedidos-confirmados` y los mensajes publicados en la UI de
management de RabbitMQ: http://localhost:15672 (user/pass).
