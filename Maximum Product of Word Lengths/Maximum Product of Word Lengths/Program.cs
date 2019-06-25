using System;
using System.Collections.Generic;
using System.Linq;
namespace Maximum_Product_of_Word_Lengths
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            string[] words = new string[] { };

            var result = solution.MaxProduct(words);
            Console.WriteLine("Hello World!");
        }
    }

    public class Solution
    {
        public int MaxProduct(string[] words)
        {
            HashSet<char[]> aaa = new HashSet<char[]>();
            aaa.Add("abcdef".ToCharArray());
            aaa.Add("xtfn".ToCharArray());
            return 1;
        }
    }
}
