package sbtract_the_Product_and_Sum_of_Digits_of_an_Integer.Leetcode;

public class Subtract_the_Product_and_Sum_of_Digits_of_an_Integer {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
        
		Solution solution = new Solution();
		
		int result = solution.subtractProductAndSum(4421);
		
		System.out.println(result);
	}

}

class Solution {
    public int subtractProductAndSum(int n) {
        int productResult = 1;
        int sumResult = 0;
    	do {
    	    int currentval = n % 10;
    	    productResult *= currentval;
    	    sumResult += currentval;
    	    n /= 10;
    	}while(n > 0);
    	
    	return productResult - sumResult;
    }
}
