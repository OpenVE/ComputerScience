/* Binary Search Algorithm (iterative approach)
 * by Roberto Tatasciore
 */

public class BinarySearch {
    
    // the array "values" needs to be sorted before calling this method
    static int find(int key, int[] values) {
        int low = 0; // index of first element in "values"
        int high = values.length - 1; // index of last element in "values"
        while (low <= high) {
            int mid = (low + high) / 2; // half of array
            if (key < values[mid]) high = mid - 1; // halve the search
            else if (key > values[mid]) low = mid + 1; // halve the search
            else return mid; // key == values[mid] found key in array, return index
        }
        return -1; // key not found in array
    }
    
}