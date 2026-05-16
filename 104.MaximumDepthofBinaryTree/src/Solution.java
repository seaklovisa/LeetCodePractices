/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author seako
 */
public class Solution {
    int depth = 0;
    int currentMaxDepth = 0;
    public int maxDepth(TreeNode root) {
        if(root == null){
            return 0;
        }
        
        int leftDepth = 1;
        int rightDepth = 1;
        
        leftDepth += maxDepth(root.left);
        
        rightDepth += maxDepth(root.right);
        
        if(leftDepth > rightDepth){
                return leftDepth;
            }
        else{
            return rightDepth;
        }
        
    }
}
