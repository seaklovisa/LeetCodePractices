import java.util.ArrayList;
import java.util.List;

public class Solution {

  public static void main(String[] args){
    subsets(new int[]{1,2,3});
  }


  public static List<List<Integer>> subsets(int[] nums){
    List<List<Integer>> result = new ArrayList<>();
    result.add(new ArrayList<>());

    for(int num : nums){
      List<List<Integer>> subSets = new ArrayList<>();
      for(List<Integer> item : result ){
        List<Integer> subset = new ArrayList<>(item);
        subset.add(num);
        subSets.add(subset);
      }
      result.addAll(subSets);
    }

    return result;
  }
}
