
#如何取得中位數
使用兩個head 
maxheap 存較小的左半邊 top() 會是較小中的最大的 
minheap 存較大的右半邊 top() 會是較大中的最大的
left.size() == right.size or left.size() == right.size+1
left.top() <= right.top()
ex 以window [-1 1 3] 及 [2 4 6 8]
[-1 1] [3] => 中位數是 left.top()
[2 4] [8 6] => 中位數就是 (left.top() + right.top()) / 2


# 視窗滑動
加進一個新數
left.size() == right.size or left.size() == right.size+1
left.top() <= right.top()
if new <= left.top() 放left else 放 right
移除一個舊數
重新調整兩個heap到合法狀態


# 完整流程圖
初始化兩個 heap
for i in range(nums):
    加 nums[i]
    balance()

    if i >= k:
        標記 nums[i-k] 為待刪
        balance()

    if i >= k-1:
        計算中位數