/* Recursive Power Function
 * by Roberto Tatasciore
 */

public class RecursivePow {
    
    static int pow(int base, int exp) {
        if (exp == 0) return 1; // base case
        return base * pow(base, exp - 1);
    }
    
}