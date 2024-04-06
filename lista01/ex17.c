#include <stdio.h>

int main(void) {
    int x, y;

    printf("Diga os valores de X e Y: ");
    scanf("%d%d", &x, &y);

    if (x % 2==0) {
        for (int i=0; i<=y; i++){
            int receive = x + 2 * i;
            printf("%d\n", receive);
        }
        printf("\n\n");
    } else if(x % 2!=0){
        printf("O PRIMEIRO NUMERO NAO E PAR\n");
    }

}
