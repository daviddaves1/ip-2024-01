#include <stdio.h>

int main(void){

    float sala_min=0, kw=0, valor_kw=0, consumo_res=0, new_value=0;

    printf("\n\n");
    printf("Informe o valor do salario minimo: ");
    scanf("%f", &sala_min);

    printf("Informe quantos kWh a residencia gasta: ");
    scanf("%f", &kw);

        valor_kw = (0.70 * sala_min) /100;
        consumo_res = (valor_kw * kw);
        new_value = consumo_res * (1 - 0.10);

    printf("\n");
    printf("Valor de cada kWh:   R$ %.2f\n", kw);
    printf("Valor a ser pago pelo consumo residencial:   R$ %.2f\n", consumo_res);
    printf("Novo valor a ser pago com desconto de 10 por cento:   R$: %.2f\n", new_value);
}