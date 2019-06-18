using System;

namespace _318._Maximum_Product_of_Word_Lengths_2
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            string[] words = new string[] { "a", "aa", "aaa", "aaaa" };
            //string[] words = new string[] { "a", "ab", "abc", "d", "cd", "bcd", "abcd" };
            //string[] words = new string[] { "abcw", "baz", "foo", "bar", "xtfn", "abcdef" };
            var result = solution.MaxProduct(words);
            Console.WriteLine("Hello World!");
        }
    }

    public class Solution
    {
        int[] mask = new int[26];

        public int MaxProduct(string[] words)
        {
            int n = words.Length;
            int[] die
            for(int i = 1; i <= words.Length; i++)
            {

            }
        }

        public int Mask()
        {

        }
    }
}
