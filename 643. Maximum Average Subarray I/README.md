643. Maximum Average Subarray I
Easy
Topics
premium lock icon
Companies
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

 

Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000
 

Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
 

Accepted
1,046,598/2.2M
Acceptance Rate
46.6%


---

---
description: Strong-hire 乾跑：用不變量驗證使用者解法（事件驅動、一般性正確性、含複雜度與抓 bug），並控制篇幅避免 token 膨脹
globs:
  - "**/*.py"
  - "**/*.md"
alwaysApply: false
---

# Strong-Hire Dry Run Protocol（僅在使用者要求乾跑/驗證解法時套用）

## 觸發條件
當使用者提供「題目 + 解法」並要求 dry run / correctness / complexity / 驗證正確性 / 找 bug 時，必須遵守本規則。
若使用者不是在要求乾跑或驗證解法，請不要強制使用此格式。

## 篇幅控制（省 token：強制）
- Representative Testcases：2–3 組為預設（除非使用者指定更多）。
- Event-driven Dry Run：
  - 每個 testcase 最多 **6 個事件**（必要時可到 8，但要有理由）。
  - 每個事件的 state_snapshot 最多 **5 個欄位**（只保留關鍵狀態）。
- 禁止輸出完整資料結構 dump：完整 heap、完整 dp 表、完整鄰接邊清單、完整 window 逐字追蹤。
  - 若需要，只能列「局部」：例如 heap top、被 relax 的那條邊、dp 的某一列/某幾格、window 的關鍵計數等。
- 禁止逐行逐迭代唸程式。只講轉折事件（turning points）。

---

## 輸出格式（固定順序；標題不可省略）
1) Spec Checkpoints
2) State Semantics
3) Invariants
4) Representative Testcases
5) Event-driven Dry Run
6) Correctness Argument
7) Complexity
8) Bugs / Fixes (if any)

---

## 1) Spec Checkpoints（2–5 點，可檢核）
- 用 2–5 點把題目轉成「可驗證條件」：
  - 功能正確（包含重複需求/計數/限制）
  - 最優性（最短/最小/最大）若題目要求
  - 無解與邊界輸入的行為（若適用）
- 不要只是重述題目句子；要能被反例推翻的那種條件。

---

## 2) State Semantics（列出核心狀態與語意）
列出核心變數/資料結構（通常 3–8 個即可），每個都要：
- meaning：它代表什麼（語意）
- updates：什麼情況會更新它（新增/移除/比較/回溯/鬆弛等）

---

## 3) Invariants（至少 1 個 main；可 0–2 個 aux）
- 必須至少提出 1 個 **Main invariant**（直接支撐 correctness）。
- 可選 0–2 個 **Aux invariants**（支撐最優/不漏/終止/攤提）。
- 每個 invariant 必須包含三段式（簡短即可）：
  - init：為何初始成立
  - preserve：每次更新後為何仍成立
  - use：何時/如何用它推出規格成立（或最優/完備/終止）

> Invariant 必須可檢核（能對上某個狀態），禁止空泛語句。

---

## 4) Representative Testcases（2–3 組，目的導向）
- 每個 testcase 必須標註 purpose：
  - normal_turning_point：能觸發核心轉折（例如第一次合法、第一次更新最優、第一次 finalize 等）
  - edge_case：邊界（空、單元素、極小、t>s 等）
  - risk_point：重複需求、權重、去重、分支/剪枝等常見出錯點
  - no_solution：明確無解（若題目允許）
- testcase 不求多；求能覆蓋「最容易出錯」與「核心轉折」。

---

## 5) Event-driven Dry Run（只追 turning points）
### Turning points 的通用定義（最重要）
只記錄會改變下列任一項的事件：
- validity/可行性（合法↔不合法、約束滿足↔不滿足）
- best/答案候選（更新最優解）
- 核心狀態類別切換（例如：節點被 finalize、顏色改變、dp cell 被確定、遞迴返回、剪枝觸發）
- 進度與終止（指標/層級/子問題規模確實前進）

### 每個事件必須包含（簡短、固定欄位）
- trigger：事件是什麼（例如：加入元素、移除元素、pop、relax 成功、dp 轉移、enter/return/prune）
- state_snapshot：只列必要欄位（<= 5 個）
- invariant_alignment：一句話說明不變量仍成立 / 或指出在此處即將被破壞
- answer_update（若有）：did_update + why_valid（為何此刻可更新）+ new_answer（可省略細節）

### Examples（僅為例子，不要求列舉全部）
- sliding window：valid 變化、答案更新、收縮停點
- graph：pop、relax 成功、狀態/顏色變更、union 成功/失敗
- DP：base、某格候選比較、依賴已就緒
- recursion/backtracking：enter/return、choose→recurse→unchoose(restore)、prune、record solution

---

## 6) Correctness Argument（一般性，用 invariants 串起來）
必須包含三段（每段 1–3 句即可）：
- Safety/Validity：為何過程不會違反規格（用 invariant）
- Optimality 或 Completeness：為何最優 / 不漏不重（視題型）
- Termination/Progress：為何會停（單調指標/子問題規模下降/有限狀態）

附註（必要時才提）：
- DP：state 定義 + 轉移完整性 + fill order +（口頭版）歸納
- Dijkstra：需聲明非負權前提，並說明 finalize 合理性

---

## 7) Complexity（time/space + why）
- Time：Big-O + 為什麼（用“總次數”角度，例如指標最多走 m 次、每條邊 relax 次數、每格 dp 計算成本）
- Space：Big-O + 主要來源（表格/圖/遞迴深度/哈希表）
- 若有攤提（amortized），用 1 句話說明攤提理由。

---

## 8) Bugs / Fixes（若有）
若發現違反 spec 或 invariant：
1) earliest_failure_point：最早失效點（指出具體語句/步驟）
2) minimal_counterexample：最小反例
3) minimal_fix：最小 diff 修正（prefer minimal）
4) why_fix_works：修正後 invariant 如何恢復並滿足 spec

若未發現問題：簡短列出「在以下假設下未發現問題」與假設清單。

---

## chatgpt給的格式
## 我就會用你貼的「Spec Checkpoints → ... → Bugs/Fixes」那套格式輸出，篇幅會被控制（2–3 testcases、每組 6 個 turning points 左右）。

Problem: LeetCode 643. Maximum Average Subarray I
Language: Go
My solution:
sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)

Please: dry run + correctness + complexity + find bugs


---
Spec(規格)
1. 回傳連續子陣列的最小長度 使sum >= target
2. 子陣列長度不固定 可變
3. 若不存在任何符合條件的子陣列 回傳0
4. nums.length <= 10的5次方

State Semantics(核心語意)
k:固定視窗長度
sum:視窗內的總合
maxSum:所有長度等於k的視窗中的最大總和
i:視窗右邊界

Invariants(不變量)
每個loop結束後 sum = nums[i-k+1 .. i]的總合
    - init(初始化) 為什麼一開始成立：初始化前k個元素後成立
    - preserve(維持) 為什麼每次更新後還成立：每次+nums[i] & -nums[i-k]，視窗右移1格後仍成立
    - use(用途) 為什麼它能保證答案正確：確保每個長度為k的連續子陣列都被正確比較

Representative Testcases(有代表性的testcases)
    - normal(會跑到slide window轉折):nums = [1,12,-5,-6,50,3] k=4 result=12.75
    - edge(最小邊界 k=1): nums[5] k=1 result=5.0

    ```
        很好的一個問題，這句話**真的很多人聽不懂，但它其實是 sliding window 的靈魂**。
我用**白話 + 小例子 + 對照錯誤理解**一次講清楚。

---

## 一句話先給答案

> **「會真的跑到 sliding window 轉折點」＝
> 測資必須讓視窗發生「狀態改變」，而不是一路平順滑過去。**

也就是：

* 不是只跑初始化
* 不是每次都一樣
* 而是**真的觸發「關鍵判斷」**

---

## 什麼叫「轉折點（turning point）」？

在 sliding window 中，**轉折點 = 程式邏輯「會改變決策」的地方**。

不是 index +1 而已，而是發生下面其中一種：

### 常見 sliding window 轉折點

1. **第一次形成合法 window**
2. **第一次更新答案**
3. **條件從成立 → 不成立（或反過來）**
4. **視窗開始縮（while 進入）**
5. **視窗縮到不能再縮（while 停止）**

👉 只要發生其中一個，就是「轉折點」。

---

## 用你的 643 來看（固定長度）

### 題目回顧

* window 長度固定為 k
* 每次只會「右移一格」

### 轉折點是什麼？

其實只有兩種：

1️⃣ **初始化第一個 window**

```go
for i := 0; i < k; i++ { sum += nums[i] }
```

➡️ 從「沒有 window」→「有第一個合法 window」

2️⃣ **maxSum 被更新**

```go
if sum > maxSum { maxSum = sum }
```

➡️ 出現比以前更大的候選答案

### 所以我選的測資為什麼「會跑到轉折點」？

```text
[1,12,-5,-6,50,3], k=4
```

* 第一個 window：sum = 2
* 第二個 window：sum = 51 → **maxSum 更新**

👉 **有轉折（答案真的改變）**

---

## 對照一個「沒有轉折點」的爛測資

```text
nums = [5,5,5,5], k = 2
```

每個 window：

* sum 永遠 = 10
* maxSum 永遠不變

❌ 這種測資：

* 跑了
* 但**完全沒驗證邏輯**
* 有 bug 也看不出來

---

## 再看 209（可變長度）會更明顯

### 題目特性

* sum < target → 擴窗
* sum ≥ target → 縮窗

### 轉折點就更多了：

1️⃣ sum 第一次 ≥ target
2️⃣ 進入 while 開始縮
3️⃣ 縮到剛好不合法，while 停止
4️⃣ minLen 更新

### 為什麼官方範例是好測資？

```text
target=7
nums=[2,3,1,2,4,3]
```

因為它：

* 會先超過
* 再縮
* 再超過
* 再縮
* **還會更新更小答案**

👉 所有關鍵轉折都跑到

---

## 用一句面試官腦內翻譯給你

當我寫：

> 「會真的跑到 sliding window 轉折點」

面試官腦內是在想：

> 「這個測資能不能逼你的程式
> 走到 if / while 的關鍵分支？」

---

## 一個你可以直接用的判斷法（超實用）

選測資時問自己三個問題：

1️⃣ 會不會進 if？
2️⃣ 會不會更新答案？
3️⃣ 會不會走到 while（如果有）？

如果三個都是「會」
👉 **這就是會跑到轉折點的好測資**

---

## 最後一句話（幫你內化）

> **轉折點 = 演算法「真的做出決定」的時刻**
> 測資要讓那些時刻發生，dry run 才有意義。

如果你願意，下一步我可以：

* 拿一題新的 sliding window 題
* 讓你自己挑測資
* 我幫你判斷「有沒有跑到轉折點」

你要不要試一次？


    ```






1. 讓所有人知道這個題目是什麼意思
2. 邊講邊敲一些字
3. 


澄清題意
MAIN IDEA 從EXAMPLE觀察到的 EDGE或是自已的範例
PSUDO CODE 討論範例 OR 文字描述 心裡想的說出來


