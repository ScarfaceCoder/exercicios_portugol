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