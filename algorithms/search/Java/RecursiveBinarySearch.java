/* Binary Search Algorithm (recursive approach)
 * by Roberto Tatasciore
 */

public class RecursiveBinarySearch {
    
    private static int find(int key, int low, int high, int[] values) {
        if (low > high) return -1;
        int mid = (low + high) / 2;
        if (key < values[mid]) return find(key, low, mid - 1, values);
        if (key > values[mid]) return find(key, mid + 1, high, values);
        return mid;
    }
    
    // helper method
    static int find(int key, int[] values) {
        return find(key, 0, values.length - 1, values);
    }
    
}
