# go- start

## Motivacao
Ao aprender uma nova linguagem de programação vai trazer diversos benefícios
 1. - compreensão geral da programação:
    As linguagem de programação tem sua própria sintaxe e abordagem para resolver problemas. Ao aprender uma linguagem diferente posso expandir meus conhecimento e compreensão geral da programação.
 2. Aumentar a produtividade:
    Algumas linguagens de programação são mais eficientes em determinadas tarefas do que outras. Aprender uma nova linguagem que se encaixa melhor em uma determinada tarefa pode aumentar a produtividade e, consequentemente, do projeto.
  3. Estar atualizado:
    A tecnologia está em constante evolução e novas linguagens de programação surgem a todo momento. Aprender novas linguagens mantém-me atualizado e preparado para as  minhas contribuicao  e a demandas do mercado em constante mudancas.

### pontos de pegada conceitos

1.  Go lang tem como mascote o Gopher,
2.  Linguagem de programacao open source que tem objetivo de torna os devs mais produtivos
3.  simplicidade produtividade
4.  Expressiva concisa limpa eficiente
5.  criada para aproveitar o maximo dos recursos multicore e de Rede
6.  Foi criada pensando no contexto de Hoje , recursos computacional
7.  compilada
8.  Rapida compilacao e ao mesmo tempo trabalha com garbage collection
9.  Rapida, estancialmente tipada. compilada, mas que ao mesmo tempo parece ate uam linguagem dinamicamente typada e interpretada de tao rapida.
10. estaticamente tipada, essa variavel vai ter ser mesmo type ate ela morrer
11. nao necessariamente precisamos declarar todas variaves (java like)
12. Ela compila em apenas um arquivo, tudo vai ta la.
14. ela mesma se auto compila. linguagem reescreu seu propio compilador em Go
15. deploy simples
16. Deteccao de Race Conditions
17. paradigma" usa go way jeito do Go ser
18. Ela tem 25 palavras chaves
19. convencao de boas presticas (Lint por padrao)


### pontos de motivacao para a construcao Go lang

Limitacao que algumas das principais linguagens que a GOOGLE como Python, Java, C++

- Python: problemas com lentidao
- c/c++: problemas complexidade e demorada para compilar
- java complexidade gerada ao longo do tempo / verbosidade da linguagem

- Nasceu na Google  em  2007 o projeto.
- Nasceu nativamente pensando no futuro em Cocorrencia Multithreading
- A versao 1.0 saiu em 2012
genios criadore do go
1.  Robert Griesemar - V8
2.  Rob Pike  - Unix & UTF8
3.  Ken Thompson - Unix & UTF8

### Topico 01

- Workspace padrão do Go
```
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
```
- A bin , onde ficará os compilados do nosso código Go.
- A pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando ao longo de nossos códigos.
- A src, onde escreveremos o código fonte de cada aplicação.
### Topico 02 comamd

1. go build arquivo   // cria exe
2. go run arquivo     // esse comando compilamos e executamos o programa

### Topico 03 Variaveis
```
exemplo em Java
// declaração da  variável idade
int idade;
idade = 15;

// imprime a idade
System.out.println(idade);

```

```
exe,plo em go
package main

import (
    "fmt"
)

func main() {
    var (
        idade  = 15
        altura = 1.78
        nome   = "Guilherme"
    )
    fmt.Println(nome, idade, altura)
}
```
quando so declaramos a variavel e nao colocamos valor
```
----------------------------------------------
| Tipo         | valor atribuído pelo Go     |
|----------    |-------------------------    |
| numérico     | 0                           |
| booleano     | false                       |
| string       | " "                         |
----------------------------------------------
```

### Topico 04 POO
- struct
- formas de utilizar a struct: informando ou não os nomes dos campos.
- entender o que são ponteiros
- forma de utilizar a struct e criar um método



### Topico 05 função receba um número indeterminado de parâmetros.
Funções deste tipo são conhecidas em Go como variadic functions, ou função variádica.
Rest parametrs do javascript parecido

```
package main

import (
    "fmt"
)

func Somando(numeros ...int) int {
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}

func main() {
    fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))
}

```

 - como trabalhar com structs;
 - referências;
 - funções com múltiplos retornos,
 - funções sem retornos;
 - composição;
 - divisão do projeto em pacotes;
 - encapsulamento e interface, uma das maiores peculiaridades da linguagem Go
 - composicao
 
 
