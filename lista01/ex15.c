#include <stdio.h>

int main(void) {
    int N;

    printf("Informe o valor de N: ");
    scanf("%d", &N);

    if (N <= 5 || N >= 2000) {
        printf("Valor incorrespondente!\n");
    } else {
        printf("SAIDA:\n");
        for (int i=1; i<=N; i++) {
            if (i%2==0) {
                printf("%d\n", i*i);
            }
        }
    }
}
