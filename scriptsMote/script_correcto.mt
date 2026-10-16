# =========================================================
# SCRIPT DE PRUEBA VALIDO - LENGUAJE MOTE
# Este script cubre el 100% de la gramática del lenguaje.
# =========================================================

# 1. Declaraciones globales de variables (con y sin asignacion)
let version: string := "1.0.0"
var ejecuciones: int := 0
var flagGlobal: bool

# 2. Declaraciones de registros (Records)
record Punto
    x: float
    y: float
end

record Jugador
    nombre: string
    posicion: Punto
    vidas: int
    activo: bool
    inventario: [3]int
end

# 3. Funciones auxiliares
# Procedimiento sin tipo de retorno (void)
fn saludar(nombre: string)
    print("Hola " + nombre)
end

# Funcion con multiples parametros y tipo de retorno float
fn distancia(p1: Punto, p2: Punto): float
    let dx: float := p2.x - p1.x
    let dy: float := p2.y - p1.y
    return (dx * dx + dy * dy)
end

# Funcion con condicionales if / elif / else
fn verificarEstado(j: Jugador): bool
    if j.vidas > 0 and j.activo then
        return true
    elif j.vidas = 0 then
        return false
    else
        return false
    end
end

# 4. Funcion principal con todas las estructuras
fn main(): int
    # Variables locales de distintos tipos
    let lim: int := 10
    var acumulador: int := 0
    let notas: [5]int := [10, 8, 9, 7, 10]
    let factor: float := -2.5

    # Llamada a procedimiento
    saludar("Mote")

    # Bucle while con expresiones booleanas y logicas (and, not)
    while acumulador < lim and not false do
        acumulador := acumulador + 1
    end

    # Bucle for con rango (..)
    for i in 0..5 do
        # Operadores relacionales (=, <>) y aritmeticos (*, %)
        if i = 2 or i <> 4 then
            notas[i] := (notas[i] * 2) % 10
        end
    end

    # Instanciacion y acceso a campos de record con '.'
    let pOrigen: Punto
    pOrigen.x := 0.0
    pOrigen.y := 0.0

    let pDestino: Punto
    pDestino.x := 3.0
    pDestino.y := 4.0

    # Asignacion de resultado de llamada a funcion
    let d: float := distancia(pOrigen, pDestino)
    print(d)

    # Comparaciones compuestas (<=, >=)
    if d <= 5.0 and d >= 1.0 then
        print("Distancia en rango correcto")
    end

    return 0
end
