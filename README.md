# Space Invaders CLI

Un juego clásico de Space Invaders implementado como una aplicación de línea de comandos (CLI) en Go, usando el framework Cobra para la gestión de comandos y BubbleTea para crear una interfaz de usuario de terminal (TUI) atractiva.

## Características

- Interfaz de usuario de terminal gráfica y colorida
- Movimiento y disparo de la nave jugador
- Formaciones de enemigos que se mueven y disparan
- Sistema de puntuación y vidas
- Detección de colisiones
- Tres niveles de dificultad
- Pantalla de fin de juego con opción de reinicio
- Mensaje de despedida al salir

## Requisitos

- Go 1.18 o superior
- Las siguientes dependencias:
    - github.com/spf13/cobra
    - github.com/charmbracelet/bubbletea
    - github.com/charmbracelet/lipgloss

## Instalación

1. Clona el repositorio:

```bash
git clone https://github.com/yourusername/space-invaders.git
cd space-invaders
```

2. Instala las dependencias:

```bash
go mod tidy
```

3. Compila el juego:

```bash
go build -o space-invaders
```

## Uso

Para iniciar el juego con la dificultad por defecto (Normal):

```bash
./space-invaders
```

### Opciones de dificultad

- Modo Fácil (más vidas, enemigos más lentos, menos filas de enemigos):
  ```bash
  ./space-invaders -e
  ```
  o
  ```bash
  ./space-invaders --easy
  ```

- Modo Difícil (menos vidas, enemigos más rápidos, más filas de enemigos):
  ```bash
  ./space-invaders -d
  ```
  o
  ```bash
  ./space-invaders --hard
  ```

## Controles

- **Movimiento**: Teclas de flecha izquierda/derecha o A/D
- **Disparar**: Barra espaciadora, tecla W o flecha arriba
- **Salir**: Tecla Q o Ctrl+C
- **Reiniciar** (después de Game Over): Tecla R

## Estructura del proyecto

```
space-invaders/
├── cmd/
│   └── cmd.go   (Configuración de Cobra CLI)
├── internal/
│   └── game/
│         ├── game.go   (Lógica principal del juego)
│         ├── player.go (Jugador y sus acciones)
│         ├── enemy.go  (Enemigos y sus acciones)
│         ├── bullet.go (Proyectiles)
│         └── ui.go     (Interfaz con BubbleTea)
├── main.go       (Punto de entrada)
└── go.mod        (Dependencias)
```

## Mecánicas del juego

- El jugador controla una nave en la parte inferior de la pantalla.
- Los enemigos se mueven de lado a lado y descienden cuando llegan al borde.
- El jugador debe destruir a todos los enemigos antes de que lleguen abajo.
- Si los enemigos llegan a la parte inferior o el jugador pierde todas sus vidas, el juego termina.
- Cuando todos los enemigos son eliminados, aparece una nueva formación.
- La puntuación aumenta con cada enemigo destruido.

## Personalización

Puedes personalizar varios aspectos del juego modificando las constantes en `game/game.go`:

- Velocidad del jugador
- Velocidad de los enemigos
- Velocidad de los proyectiles
- Número de vidas
- Dimensiones de la pantalla

## Licencia

[MIT](LICENSE)

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue para discutir los cambios que te gustaría hacer.

## Agradecimientos

- Inspirado en el juego clásico Space Invaders
- Desarrollado con [Cobra](https://github.com/spf13/cobra) y [BubbleTea](https://github.com/charmbracelet/bubbletea)
