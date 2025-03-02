# Run program

O output estará no último byte(o 255º byte). Para rodar o programa, basta fazer a build e passar o arquivo como primeiro argumento. 
Execute na raiz do projeto:
```go build && ./neanderVM program.mem;```

Para rodar o programa basta rodar a função passando os bytes do arquivo mem:
```resultados, _ := neander.RunProgram(bytes, false, true)```

Para imprimir o programa original, use a função passando os bytes do arquivo mem:
```neander.PrintProgram(bytes, false, false)```
