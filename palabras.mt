#|
  Archivo de prueba para el analizador léxico de Mote.
  Contiene sintaxis válida y casos de error léxico explícitos.
|#

# Comentario de línea: función factorial
fn factorial(n: int): int
    if n <= 1 then
        return 1
    else
        return n * factorial(n - 1)
    end
end

fn main(): int
    # 1. Variables y Literales válidos
    let x: int := 5
    print(factorial(x))

    var total: int := 0
    let nums: [5]int := [1, 2, 3, 4, 5]

    for i in 0..5 do
        total := total + nums[i]
    end
    print(total)


    # Error 1: Cadena sin cerrar antes del salto de línea
    let mensaje: string := "Hola Mote sin cerrar

    # Error 2: Identificador que empieza con dígito
    let 3variable: int := 10

    # Error 3: Literales flotantes mal formados (.5 y 3.)
    let f1: float := .5
    let f2: float := 3.

    # Error 4: Caracteres no reconocidos fuera de alfabetos
    let simboloInvalido := @ + $ ~

    # Error 5: Comentario de bloque sin cerrar al final
    #| Este comentario nunca se cierra con su delimitador correspondiente
