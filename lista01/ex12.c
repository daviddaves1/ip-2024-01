#include <stdio.h>

int main(void) {
    int hours=0, hours_exced=0;

    printf("Informe o numero de horas que o locatario usou a charrete: ");
    scanf("%d", &hours);

    float valor_pagar;
    hours_exced = hours % 3;                      

    if (hours_exced != 0) {
        valor_pagar = ((hours/3) * 10) + (hours_exced * 5);
    } else {
        valor_pagar = (hours/3) * 10;
    }

    printf("O VALOR A PAGAR E: R$ %.2f\n", valor_pagar);

    return 0;
}
