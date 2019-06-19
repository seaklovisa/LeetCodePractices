// 59. Spiral Matrix II_CPP.cpp : 此檔案包含 'main' 函式。程式會於該處開始執行及結束執行。
//

#include <iostream>
#include <vector>
using namespace std;

class Solution {

public:
	vector<vector<int>> generateMatrix(int n) {

		vector <vector<int>> result(n,vector<int>(n));


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

			for (int i = downLimit; i >= upLimit; i--)
			{
				result[i][leftLimit] = ++val;
			}
			if (++leftLimit > rightLimit) break;
		}

		return result;
	}
};

int main()
{
	Solution solution;
	int n = 4;
	vector<vector<int>> result =  solution.generateMatrix(n);

	

    std::cout << "Hello World!\n";
	return 0;
}



// 執行程式: Ctrl + F5 或 [偵錯] > [啟動但不偵錯] 功能表
// 偵錯程式: F5 或 [偵錯] > [啟動偵錯] 功能表

// 開始使用的提示: 
//   1. 使用 [方案總管] 視窗，新增/管理檔案
//   2. 使用 [Team Explorer] 視窗，連線到原始檔控制
//   3. 使用 [輸出] 視窗，參閱組建輸出與其他訊息
//   4. 使用 [錯誤清單] 視窗，檢視錯誤
//   5. 前往 [專案] > [新增項目]，建立新的程式碼檔案，或是前往 [專案] > [新增現有項目]，將現有程式碼檔案新增至專案
//   6. 之後要再次開啟此專案時，請前往 [檔案] > [開啟] > [專案]，然後選取 .sln 檔案
