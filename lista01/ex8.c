#include <stdio.h>
#include <math.h>
#define PI 3.14159

int main(void) {
    float raio=0, altura=0, area_circ=0, area_lateral=0, custo_total=0;

    printf("Informe as informações da lata abaixo:\n\n");

    printf("Raio: ");
    scanf("%f", &raio);
    printf("Altura: ");
    scanf("%f", &altura);

    area_circ = PI * pow(raio, 2);
    area_lateral = 2 * PI * raio * altura;
    custo_total = (2 * area_circ + area_lateral) * 100;

    printf("O VALOR DO CUSTO E:   R$ %.2f\n", custo_total);

}
