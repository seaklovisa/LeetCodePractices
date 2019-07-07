using System;
using System.Net;

namespace _1108.Defanging_an_IP_Address_CSharp
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            var result = solution.DefangIPaddr("1.1.1.1");
            Console.WriteLine("Hello World!");
        }

        public class Solution
        {
            public string DefangIPaddr(string address)
            {
                IPAddress addr = IPAddress.Parse(address);

                string[] temp = addr.ToString().Split('.');

                string result = $"{temp[0]}[.]{temp[1]}[.]{temp[2]}[.]{temp[3]}";

                return result;
            }
        }
    }
}
