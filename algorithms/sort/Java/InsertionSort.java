/* Insertion Sort Algorithm
 * by Roberto Tatasciore
 */

/* Insertion Sort is efficient for small data sets
 * and for data sets that are already partially sorted
 * since it's a quadratic algorithm.
 *
 * To sort huge data sets, consider using other sorting algorithms
 * i.e. mergesort, quicksort
 * 
 */

public class InsertionSort {
    
    static void sort(int[] values) {
        for (int i = 1; i < values.length; i++) {
            int aux = values[i];
            int j = i - 1;
            while (j >= 0 && values[j] > aux) {
                values[j + 1] = values[j];
                j--;
            }
            values[j + 1] = aux;
        }
    }
    
}
