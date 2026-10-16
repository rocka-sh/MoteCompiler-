
print("Este print esta fuera de cualquier funcion")


fn testErrorTipo()
    let variableSinTipo := 10
end

fn testComaParam(a: int, b: string, )
     print("hola")
 end

 fn testComaArreglo()
     let arr: [3]int := [1, 2, ]
 end

 fn testOperadorDuplicado()
     let x: int := 10 + * 2
 end

 record EstructuraSinEnd
     campo1: int
     campo2: string

 fn funcionSinEnd(): int
     print(123)