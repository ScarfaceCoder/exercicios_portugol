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