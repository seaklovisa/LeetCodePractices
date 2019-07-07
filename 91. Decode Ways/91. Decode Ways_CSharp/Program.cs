using System;

namespace _91.Decode_Ways_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            var result = solution.NumDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253");
            Console.ReadKey();
        }
    }

    public class Solution
    {

        public int NumDecodings(string s)
        {
            int totalcount = 0;

            totalcount += loopStr(s,0);

            return totalcount;
        }

        public int DFS(string s, int index)
        {


        }

        public int loopStr(string s,int indexer)
        {
            int index = indexer;
            int count = 0;
            if (index + 1 > s.Length)
            {
                 return 1;
            }
            
            
            string part = "";
            //有二格空間
            if(index + 2 <= s.Length)
            {
                part = s.Substring(index, 2);
            }
            else if(index + 1 <= s.Length)
            {
                part = s.Substring(index, 1);
            }

            int intPart = Convert.ToInt32(part);

            if (part.Length == 1)
            {
                if (intPart == 0) return 0;

                count += loopStr(s, ++index);
            }
            else if(part.Length == 2)
            {
                if(intPart > 26)
                {
                    count += loopStr(s, ++index);
                }
                //無法分解成一個一個 ex 10 無法分成 1  0
                else if(intPart == 10 || intPart == 20)
                {
                    count += loopStr(s, index += 2);
                }
                else if(intPart < 10)
                {
                    return 0;
                }
                else
                {
                    count += loopStr(s, ++index);
                    index = indexer;
                    count += loopStr(s, index += 2);
                }
            }

            return count;
        }
    }
}
