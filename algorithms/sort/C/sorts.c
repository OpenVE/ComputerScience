void bubblesort(int *, int);
void selectionsort(int *, int);
void insertionsort(int *, int);

int main() {
	return 0;
}

void bubblesort(int *arr, int len) {
	int i, j, aux;
	
	len /= 4;
	for (i = 0; i < len - 1; i++)
		for (j = i + 1; j < len; j++)
			if (*(arr + i) > *(arr + j)) {
				aux = *(arr + i);
				*(arr + i) = *(arr + j);
				*(arr + j) = aux;
			}
}

void selectionsort(int *arr, int len) {
	int i, j, min, aux;
	
	len /= 4;
	for (i = 0; i < len; i++) {
		min = i;
		for (j = i + 1; j < len; j++)
			if (*(arr + j) < *(arr + min)) min = j;
		aux = *(arr + i);
		*(arr + i) = *(arr + min);
		*(arr + min) = aux;
	}
}

void insertionsort(int *arr, int len) {
	int i, j, aux;
	
	len /= 4;
	for (i = 1; i < len; i++) {
		aux = *(arr + i);
		j = i - 1;
		while (j >= 0 && *(arr + j) > aux) {
			*(arr + (j + 1)) = *(arr + j);
			j--;
		}
		*(arr + (j + 1)) = aux;
	}
}
