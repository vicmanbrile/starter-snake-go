## Mision

1. No dejar que el Battlesnake retroseda a su propio cuello en el proximo movimiento


__Input__
```json
[
    {x: number, y: number},
    {x: number, y: number}
]
```

```js
var posiblesMovimientos: {
    "up": true,
    "down" : true,
    "left": true,
    "right": true
}
```
Vamos a iniciar de [{x: 0, y: 0}, {x: 0, y: -1}]

__Movimiento__

- La x en eje x+ 
- La y en eje y-


__SOLUCION EN TYPESCRIPT__


```js

var cuello = {x: 2, y: 0}; // 0
var cabeza = {x: 1, y: 0}; // 1


function PosicionX(cuelloX, cabezaX) {
    if (cuelloX > cabezaX) {
        posiblesMovimientos.rigth = false
    } else {
        posiblesMovimientos.left = false
    }
}

function PosicionY(cuelloY, cabezaY){
    if(cuelloY > cabezaY){
        posiblesMovimientos.up = false
    } else{
        posiblesMovimientos.down = false
    }
}

PosicionX(cuello.x, cabeza.y)
PosicionY(cuello.x, cabeza.y)



```