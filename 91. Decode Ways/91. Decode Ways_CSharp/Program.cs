using System;

namespace _91.Decode_Ways_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            var result = solution.NumDecodings("12");
            Console.ReadKey();
        }
    }

    public class Solution
    {

        public int NumDecodings(string s)
        {
            if (s.StartsWith("0") && s.Contains("0")) return 0;
            if (s.Length == 1) return 1;
           

            int totalcount = 0;

            totalcount += loopStr(s,0);

            return totalcount;
        }

        public int loopStr(string s,int indexer)
        {
            int index = indexer;
            int count = 0;
            if (index + 1 >= s.Length)
            {
                if (s.Substring(index, 1) == "0") return 0;
                else return 1;
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
                if(intPart > 26 || intPart  == 0)
                {
                    count += loopStr(s, ++index);
                }
                //無法分解成一個一個 ex 10 無法分成 1  0
                else if(intPart == 10 || intPart == 20)
                {
                    count += loopStr(s, index += 2);
                }
                else
                {
                    count += loopStr(s, ++index);
                    count += loopStr(s, index += 2);
                }
            }

            return count;
        }
    }
}
