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