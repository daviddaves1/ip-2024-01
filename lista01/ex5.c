#include <stdio.h>

int main(void) {

    int conta_cliente = 0;
    float consumo_mc = 0;
    char tipo_cons;

    printf("\n\n");
    printf("Informe os dados abaixo:\n\n");

    printf("Conta do cliente: ");
    scanf("%d", &conta_cliente);
    printf("Consumo de agua por metros cubicos: ");
    scanf("%f", &consumo_mc);
    printf("Tipo de consumidor: (C, I ou R)\n");
    scanf(" %c", &tipo_cons);

    if (tipo_cons == 'R') {
        float valor_res = 5 + (consumo_mc * 0.05);
        printf("CONTA: %d\n", conta_cliente);
        printf("VALOR DA CONTA: %.2f\n", valor_res);
    }
    else if (tipo_cons == 'C') {
        float valor_com = consumo_mc * 0.25;
        if (consumo_mc >= 80) {
            valor_com = valor_com + 500;
        }
        printf("CONTA: %d\n", conta_cliente);
        printf("VALOR DA CONTA: %.2f\n", valor_com);
    }
    else if (tipo_cons == 'I') {
        float valor_ind = consumo_mc * 0.04;
        if (consumo_mc >= 100) {
            valor_ind = valor_ind + 800;
        }
        printf("CONTA: %d\n", conta_cliente);
        printf("VALOR DA CONTA: %.2f\n", valor_ind);
    }

    return 0;
}
