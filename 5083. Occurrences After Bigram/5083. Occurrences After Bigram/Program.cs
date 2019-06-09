using System;
using System.Text.RegularExpressions;
using System.Collections.Generic;

namespace _5083._Occurrences_After_Bigram
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            var result = solution.FindOcurrences("we will we will rock you", "we", "will");
            
        }
    }

    public class Solution
    {
        public string[] FindOcurrences(string text, string first, string second)
        {
            List<string> result = new List<string>();
            Regex regex = new Regex($@"{first} {second} [a-z]+");

            Match matched = null;

            while ((matched = regex.Match(text)).Success)
            {
                string[] splited = matched.Captures[0].Value.Split(' ');
                result.Add(splited[2]);

                int len = text.IndexOf(matched.Captures[0].Value) + matched.Captures[0].Value.Length - splited[2].Length;
                text = text.Remove(0, len);
            }

            return result.ToArray();
        }
    }
}
