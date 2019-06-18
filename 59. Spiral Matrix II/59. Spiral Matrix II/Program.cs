using System;

namespace _59._Spiral_Matrix_II
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            int n = 1;
            var result = solution.GenerateMatrix(n);
            Console.WriteLine("Hello World!");
        }
    }

    public class Solution
    {
        public int[][] GenerateMatrix(int n)
        {
            int[][] result = new int[n][];
            for(int i = 0; i <n; i++)
            {
                result[i] = new int[n];
            }
            int upLimit = 0, downLimit = n - 1, leftLimit = 0, rightLimit = n - 1;
            int val = 0;

            while (true)
            {
                for (int i = leftLimit; i <= rightLimit; i++)
                {
                    result[leftLimit][i] = ++val;
                }
                if (++upLimit > downLimit) break;

                for (int i = upLimit; i <= downLimit; i++)
                {
                    result[i][rightLimit] = ++val;
                }
                if (--rightLimit < leftLimit) break;

                for (int i = rightLimit; i >= leftLimit; i--)
                {
                    result[downLimit][i] = ++val;
                }
                if (--downLimit < upLimit) break;

                for(int i = downLimit; i >= upLimit; i--)
                {
                    result[i][leftLimit] = ++val;
                }
                if (++leftLimit > rightLimit) break;
            }

            return result;
        }
    }
}
