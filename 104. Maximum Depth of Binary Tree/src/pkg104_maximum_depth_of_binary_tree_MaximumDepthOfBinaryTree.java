
/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author seako
 */
public class pkg104_maximum_depth_of_binary_tree_MaximumDepthOfBinaryTree {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        Solution solution = new Solution();
        
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        
        root.left.left = new TreeNode(4);
        root.right.right = new TreeNode(5);
        
        root.right.right.right = new TreeNode(6);
        
        
        int result = solution.maxDepth(root);
        
        System.out.println(result);
    }
    
}

