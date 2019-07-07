// 652. Find Duplicate Subtrees_CPP.cpp : 此檔案包含 'main' 函式。程式會於該處開始執行及結束執行。
//

#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <sstream>
using namespace std;


struct TreeNode {
	int val;
	TreeNode* left;
	TreeNode* right;
	TreeNode(int x) :val(x), left(NULL), right(NULL) {}
};

class Solution {
	vector<TreeNode*> result;

public:
	vector<TreeNode*> findDuplicateSubtrees(TreeNode* root) {
		
		map<string, int> duplicated;

		string sss = TreeTraverse(root, &duplicated);
		return result;
	}

	string TreeTraverse(TreeNode* root, map<string, int>* duplicated) {
		if (root == NULL)
			return "";
		vector<string> temp;
		if (root != NULL) temp.push_back(to_string(root->val));

		string leftTree = TreeTraverse(root->left, duplicated);
		temp.push_back(leftTree);

		string rightTree = TreeTraverse(root->right, duplicated);
		temp.push_back(rightTree);

		string treeList = string_join(&temp);
		if (false == treeList.empty())
		{
			if (duplicated->count(treeList) == 1)
				duplicated->at(treeList) = duplicated->at(treeList)++;
			else
				duplicated->insert_or_assign(treeList, 1);
			if (duplicated->at(treeList) == 2) result.push_back(root);
		}

		return treeList;
	}

	string string_join(vector<string>* list) {
		string result;
		stringstream resultstream;
		if (list->size() == 0) return result;
		if (list->size() == 1) {
			result = list->at(0);
			return result;
		}

		for (int i = 0; i < list->size() -1; i++) {
			result += list->at(i);
			result += ",";
		}

		result += list->at(list->size() - 1);

		return result;

	}
};

int main()
{
	Solution solution;

	TreeNode root(1);
	root.left = new TreeNode(2);
	root.left->left = new TreeNode(4);
	root.right = new TreeNode(3);
	root.right->left = new TreeNode(2);
	root.right->right = new TreeNode(4);
	root.right->left->left = new TreeNode(4);

	vector<TreeNode*> result = solution.findDuplicateSubtrees(&root);
    std::cout << "Hello World!\n";
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
