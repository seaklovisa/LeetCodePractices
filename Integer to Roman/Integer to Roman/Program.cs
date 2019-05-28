using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Integer_to_Roman
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            var result = solution.IntToRoman(3);
            
            Console.ReadLine();
        }
    }

    public class Solution
    {
        /// <summary>
        /*
        Symbol       Value
        I             1
        V             5
        X             10
        L             50
        C             100
        D             500
        M             1000
        */
        /// </summary>
        /// <param name="num"></param>
        /// <returns></returns>
        public string IntToRoman(int num)
        {
            string result = string.Empty;
            //千位
            var thun = num / 1000;
            //百位
            var hun = (num = (num - (thun * 1000))) / 100;
            //十位
            var ten = (num = (num - (hun * 100))) / 10;
            //個位
            var remain = num % 10;

            result += string.Join("", Enumerable.Repeat("M", thun));

            if (hun == 9)
                result += "CM";
            else if (hun == 4)
                result += "CD";
            else
            {
                //500的重複
                result += string.Join("", Enumerable.Repeat("D", hun / 5));
                //小於500的重複
                result += string.Join("", Enumerable.Repeat("C", hun % 5));
            }

            if (ten == 9)
                result += "XC";
            else if (ten == 4)
                result += "XL";
            else
            {
                //50的重複
                result += string.Join("", Enumerable.Repeat("L", ten / 5));
                //小於50的重複
                result += string.Join("", Enumerable.Repeat("X", ten % 5));
            }

            if (remain == 9)
                result += "IX";
            else if (remain == 4)
                result += "IV";
            else
            {
                //5的重複
                result += string.Join("", Enumerable.Repeat("V", remain / 5));
                //小於5的重複
                result += string.Join("", Enumerable.Repeat("I", remain % 5));
            }

            return result;
        }
    }
}
