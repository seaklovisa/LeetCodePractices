from typing import List


class Solution:
    
    def generate(self, numRows: int) -> List[List[int]]:
        result = []
        for i in range(1,numRows + 1):
            if i == 1:
                result.append([1])
                continue
            temp = []
            for j in range(1,i + 1):
                if j == 1 or j == i:
                    temp.append(1)
                    continue
                temp.append(result[i - 1 - 1][j - 1 - 1] + result[i - 1 - 1][j - 1])
            result.append(temp)
        return result

def main():
    aaa = Solution.generate(5,5)
    print("hi")

main()