using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _1120.Maximum_Average_Subtree_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            TreeNode root = new TreeNode(2);
            root.right = new TreeNode(1);

            var result = solution.MaximumAverageSubtree(root);
            Console.WriteLine("Hello World!");
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
        List<double> avgs = new List<double>();


        public double MaximumAverageSubtree(TreeNode root)
        {
            int counter = 0;
            double summer = 0;
            var result = TreeTraverse(root, ref counter, ref summer);

            return avgs.Max();
        }

        public double? TreeTraverse(TreeNode subTreeRoot, ref int counter, ref double summer)
        {
            if (subTreeRoot == null)
            {
                return null;
            }
            int leftCounter = 0;
            int rightCounter = 0;
            double leftSummer = 0;
            double rightSummer = 0;

            double? leftSum = TreeTraverse(subTreeRoot.left, ref leftCounter, ref leftSummer);

            if (false == leftSum.HasValue)
            {
                leftSum = 0;
            }
            else
                counter++;

            double? rightSum = TreeTraverse(subTreeRoot.right, ref rightCounter, ref rightSummer);

            if (false == rightSum.HasValue)
            {
                rightSum = 0;
            }
            else
                counter++;

            counter = (leftCounter + rightCounter + 1);
            summer = (leftSummer + rightSummer + subTreeRoot.val);
            avgs.Add(summer / counter);

            return subTreeRoot.val;
        }
    }
}
