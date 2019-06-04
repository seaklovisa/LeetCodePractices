using System;
using System.Collections;

namespace UTF_8_Validation
{
    class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();
            //int[] data = {240,162,138,147,17 }; //true
            //int[] data = { 240, 162, 138, 147, 145 };//false
            //int[] data = {32,196,147,225,184,165,246,149,170,129,204,153,243,188,141,147,0,217,149,234,176,176,243,178,133,144,213,181,193,187,238,137,134,218,155,33,231,134,162,243,184,144,131,71,201,131,244,133,189,140,242,178,128,156,207,154,230,165,181,240,181,134,180,227,129,199,172,226,158,164,214,183,224,137,141,20,194,188,232,177,151,242,157,180,153,200,189,239,153,186,240,153,181,154};
            int[] data = {  227, 129, 199, 172, 226, 158, 164, 214, 183, 224, 137, 141, 20, 194, 188, 232, 177, 151, 242, 157, 180, 153, 200, 189, 239, 153, 186, 240, 153, 181, 154 };
            var result = solution.ValidUtf8(data);

            Console.WriteLine("Hello World!");
        }
    }

    public class Solution
    {
        public bool ValidUtf8(int[] data)
        {
            bool result = true;
            byte[] temp = Array.ConvertAll(data, (x) =>
            {
                return Convert.ToByte(x);
            });

            BitArray bitArray = null;
            int checkCount = 0;

            for(int i = 0;i < data.Length; i++)
            {
                bitArray = new BitArray(new byte[] { Convert.ToByte(data[i]) });

                if (checkCount > 0)
                {
                    if(bitArray[7] == true && bitArray[6] == false)
                    {
                        checkCount--;
                        continue;
                    }
                    else
                    {
                        result = false;
                        break;
                    }
                }

                if (bitArray[7] == false)
                    continue;

                if(bitArray[7] == true && bitArray[6] == true && bitArray[5] == true && bitArray[4] == true && bitArray[3] == false)
                {
                    checkCount = 3;
                }
                else if(bitArray[7] == true && bitArray[6] == true && bitArray[5] == true && bitArray[4] == false)
                {
                    checkCount = 2;
                }
                else if(bitArray[7] == true && bitArray[6] == true && bitArray[5] == false)
                {
                    checkCount = 1;
                }
                else
                {
                    result = false;
                    break;
                }
            }


            if(checkCount > 0)
            {
                result = false;
            }

            return result;
        }
    }
}
