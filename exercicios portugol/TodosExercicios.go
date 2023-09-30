// exercicio 01 

programa {
  funcao inicio() {

//declarando e armazenando a variavel do primeiro medicamento.
    cadeia medicamento1 
    escreva("Digite o nome do primeiro medicamento: ")
    leia(medicamento1)

//declarando e armazenando a variavel do segundo medicamento.
    cadeia medicamento2
    escreva("Digite o nome do segundo medicamento: ")
    leia(medicamento2)

//declarando e armazenando a variavel do valor do primeiro medicamento.
    real valor_medicamento1
    escreva("Digite o valor do primeiro medicamento: ")
    leia(valor_medicamento1)

//declarando e armazenando a variavel do valor do segundo medicamento.
    real valor_medicamento2
    escreva("Digite o valor do segundo medicamento: ")
    leia (valor_medicamento2)

//escrevendo o nome e valor de cada um dos medicamentos. 
    escreva("O valor do " + medicamento1 + " é "+ valor_medicamento1)
    escreva("O valor do " + medicamento2 + " é "+ valor_medicamento2)
  }
}

// exercicio 02

programa {
  funcao inicio() {

//criando a variavel e lendo o nome do funcionário:
    cadeia nome 
    escreva ("Digite o nome do funcionário: ")
    leia (nome)
//criando a variavel e lendo a quantidade de dependentes: 
    real dependentes
    escreva ("Digite o número de dependentes: ")
    leia (dependentes)
//criando a variavel e lendo o cargo ocupado: 
    cadeia cargo 
    escreva ("Digite o cargo ocupado: ")
    leia (cargo)
//saida das informações: 
    escreva ("O funcionário " + nome + " possui " + dependentes + " dependentes. E oculpa o cargo de " + cargo + ".") 
  }
}

//exercicio 03

programa {
  funcao inicio() {

    real a1, b1, d2, b2, area1, area2, area_total

escreva ("Digite o valor de 'a' do primeiro retângulo: ")
leia (a1)
escreva ("Digite o valor de 'b' do primeiro retângulo: ")
leia (b1)

escreva("Digite o valor de 'd' do segundo retângulo: ")
leia(d2)
escreva("Digite o valor de 'b' do segundo retângulo: ")
leia(b2)


// Calcular as áreas dos retângulos
area1 = a1 * b1
area2 = d2 * b2

// Calcular a área total
area_total = area1 + area2

// Exibir os resultados
escreva("A área do primeiro retângulo (A1) é: ",area1,"\n")
escreva("A área do segundo retângulo (A2) é: ",area2,"\n")
escreva("A área total das duas figuras (AT) é: ",area_total,"\n")
  }
}


// exercicio 04 

programa {

  inclua biblioteca Matematica --> mat

  funcao inicio() {

    real altura
    escreva ("Digite o valor da altura: ")
    leia (altura)

    real raio
    escreva ("Digite o valor do raio: ")
    leia (raio)

    real area_base = 3.14 * raio * raio
    real volume_base = area_base * altura

    escreva("O valor da area da base é ", area_base, ". E o seu volume é ", volume_base,".")
  }
}

// exercicio 05 

programa {

	inclua biblioteca Matematica --> mat
	
  funcao inicio() {

real a, b, c, delta, x1, x2

    escreva("Informe o valor de a: ")
    leia(a)

    escreva("Informe o valor de b: ")
    leia(b)

    escreva("Informe o valor de c: ")
    leia(c)

    delta = (b * b) - (4 * a * c)

    se (delta < 0) 
    {
        escreva ("A equação não possui raízes reais.")
    }

    
    senao se (delta == 0) 
    {
        x1 = -b / (2 * a)
        escreva ("A equação possui uma raiz real:", x1)
    }

    senao
    {
    	  x1 = (-b + mat.raiz(delta, 2)) / (2 * a)
        x2 = (-b - mat.raiz(delta, 2)) / (2 * a)
        escreva ("A equação possui duas raízes reais:")
        escreva ("x1 =", x1)
        escreva ("x2 =", x2)
    }
        
  }
}

//exercicio 06

programa {

	inclua biblioteca Matematica --> mat
	
  funcao inicio() {
  		real a, b, c, resultado, arredondado

escreva ("Informe o comprimento do lado a: ")
leia(a)

escreva ("Informe o comprimento do lado b:")
leia(b)

c = (a * a + b * b)
resultado = mat.raiz(c, 2)
arredondado = mat.arredondar(resultado, 2)

escreva("O agricultor precisa comprar ", arredondado, " metros de cerca para fechar o triângulo.")
  }
}

//exercicio 07

programa {

  inclua biblioteca Matematica --> mat

  funcao inicio() {

    real cotacao
    escreva ("Digite a atual cotação do dólar: ")
    leia (cotacao)

    real valor
    escreva ("Digite o valor à ser convertido para real: ")
    leia (valor)

    real resultado = cotacao * valor
    real arredondado
    arredondado = mat.arredondar(resultado, 2)
    escreva(resultado)

  }
}

// exercicio 08 

programa {

  funcao inicio() {
    
    // criando as variaveis: 
    real horas, minutos, minutos_passados 

    //recebendo os valores nas variaveis: 
    escreva ("Digite a hora (0-23): ")
    leia (horas)
    escreva ("Digite os minutos (0-59): ")
    leia (minutos)

    //realizando o calculo: 

    minutos_passados = horas * 60 + minutos

    escreva ("Passaram-se " + minutos_passados + " minutos desde o início do dia.") 

  }
}

// exercicio 09 

programa {

  funcao inicio() {

    inteiro numero1, numero2, resto

    escreva("Digite o primeiro número: ")
    leia(numero1)

    escreva("Digite o segundo número: ")
    leia(numero2)

    resto = numero1 % numero2

    escreva ("O resto da divisão de ", numero1, " por ", numero2, " é ", resto)
  }
}

// exercicio 10 

programa {

  funcao inicio() {
    
    real valor_gasto, valor_total, gorjeta

    escreva("Digite o valor do gasto: R$")
    leia (valor_gasto)

    //Calcula os 10% do garçom
    gorjeta = valor_gasto * 0.10

    valor_total = valor_gasto + gorjeta

    escreva("O valor total a ser pago, incluindo 10% de gorjeta, é R$", valor_total)
  }
}

//exercicio 11


programa {

  funcao inicio() {
    
    inteiro numero1, numero2, soma, subtracao, multiplicacao, divisao, resto_divisao

    escreva("Digite o primeiro número inteiro: ")
    leia(numero1)
    escreva("Digite o segundo número inteiro: ")
    leia(numero2)

    soma = numero1 + numero2
    subtracao = numero1 - numero2
    multiplicacao = numero1 * numero2

    divisao = numero1 / numero2
    resto_divisao = numero1 % numero2

    escreva("Soma: ", soma,"\n")
    escreva("Subtração: ", subtracao,"\n")
    escreva("Multiplicação: ", multiplicacao,"\n")
    escreva("Divisão: ", divisao,"\n")
    escreva("Resto da divisão: ", resto_divisao,"\n")
  
  }
}

//exercicio 12 

programa {

  funcao inicio() {

    real n, resultado

    escreva("Digite um  número: ")
    leia (n)

    resultado = n*3

    escreva("O triplo de ", n ," é ", resultado, ".")

  }
}

//exercicio 13

programa {

  funcao inicio() {

    real altura, base, area

    escreva("Digite o valor da altura do triângulo: ")
    leia (altura)
    escreva("Digite o valor da bâse do triângulo: ")
    leia (base)

    area = base*altura/2

    escreva("A área do triângulo é ", area,".")
    
  }
}

//FIM// 
