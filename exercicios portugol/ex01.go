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