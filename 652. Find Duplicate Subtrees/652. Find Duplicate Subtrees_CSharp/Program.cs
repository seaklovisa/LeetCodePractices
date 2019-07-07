using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _652.Find_Duplicate_Subtrees_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {

            //[0,0,0,0,null,null,0,null,null,null,0]
            Solution solution = new Solution();

            //TreeNode root = new TreeNode(0);
            //root.left = new TreeNode(0);
            //root.right = new TreeNode(0);
            //root.left.left = new TreeNode(0);

            //root.right.right = new TreeNode(0);
            //root.right.right.right = new TreeNode(0);


            TreeNode root = new TreeNode(1);
            root.left = new TreeNode(2);
            root.left.left = new TreeNode(4);
            root.right = new TreeNode(3);
            root.right.left = new TreeNode(2);
            root.right.right = new TreeNode(4);
            root.right.left.left = new TreeNode(4);

            var result = solution.FindDuplicateSubtrees(root);
            Console.ReadLine();
        }
    }


    public class TreeNode
    {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }

    public class Solution
    {
        public IList<TreeNode> FindDuplicateSubtrees(TreeNode root)
        {
            Dictionary<string, int> duplicated = new Dictionary<string, int>();

            List<TreeNode> result = new List<TreeNode>();
            var sss = TreeTraverse(root, ref result, ref duplicated);
            return result;
        }

        public string TreeTraverse(TreeNode root, ref List<TreeNode> result, ref Dictionary<string, int> duplicated)
        {
            if (root == null)
                return "";
            List<string> temp = new List<string>();
            if (root != null) temp.Add(root.val.ToString());

            string leftTree = TreeTraverse(root.left, ref result, ref duplicated);
            temp.Add(leftTree);

            string rightTree = TreeTraverse(root.right, ref result, ref duplicated);
            temp.Add(rightTree);

            //根,左子樹,右子樹會形成key,value則為次數
            var treeList = string.Join(",", temp);
            if (false == string.IsNullOrEmpty(treeList))
            {
                if (duplicated.ContainsKey(treeList))
                    duplicated[treeList]++;
                else
                    duplicated.Add(treeList, 1);
                //發現重複就加入結果集內
                if (duplicated[treeList] == 2) result.Add(root);
            }


            return treeList;
        }
    }
}
