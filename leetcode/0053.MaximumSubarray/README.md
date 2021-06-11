# Question

https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

# Thoughts

## Python

假設有以下陣列：[🟥 🟢 🟦 ...]

🟥 第一次迭代
- globalMax = 🟥
- localMax = 🟥
    - localMax 用來記錄可以與下一個數字相加的最大值

> 以下簡稱 localMax 為 local；globalMax 為 global

🟢 第二次迭代
- 比較 🟢 vs. local+🟢 (此時 local = 🟥)
  - 情境2-1：若 local+🟢 < 🟢 ，可知道 local < 0，且連續最大值 local 應重新指派成 🟢
  - 情境2-2：若 local+🟢 > 🟢 ，那麼就知道 local > 0，且可重新指派連續最大值 `local = local+🟢`
- ⚠️ ***僅比較 local 與 🟢 沒意義***，因為也不能因為 local > 🟢 就留給下一次迭代用 (因為不連續了)
- 決定好 local 後，跟 global 比一下，紀錄曾經看過最大的連續總和 (sum of subarray)
  - 此步是我的思路盲點，local 一直變，又需要確保 local 可以與 current 相加，故需要額外的 global 儲存曾經看過最大值

🟦 第三次迭代
- 如果前一次迭代是因為情境2-1被重新指派，那麼情況就又回到上述
- 如果前一次迭代是因為情境2-2，則考慮以下：
  - 依照上述邏輯，比較 (🟥+🟢)+🟦 vs. 🟦
    - 情境3-1：🟥+🟢+🟦 < 🟦, 表示 🟦 就足夠了，於第二次迭代中的情境一相同
      - *此情況 🟦 也永遠>> 🟥+🟢，故不用管原本的最大值了*
      - `local = 🟦`
    - 情境3-2：🟥+🟢+🟦 > 🟦，那麼此時連續最大值為 🟥+🟢+🟦
      - *當然，🟥+🟢+🟦 也永遠>> 🟥+🟢，故不用管原本的最大值了*
      - `local = 🟥+🟢+🟦`
  - 但此時思考一個情況：有沒有可能 🟢+🟦 才是最大值？ (因為這也是一個連續子陣列)
    - 情境一得到一個事實是： 🟦 > 🟥+🟢+🟦
      - 想要證明 🟦 > 🟢+🟦
      - 移項得到 🟦-🟥 > 🟢+🟦
      - 又從第二次迭代中得知若有 `local = 🟥+🟢` 的可能，🟥 必須 > 0
      - 故 🟦 >> 🟢+🟦
    - 情境二的事實是 🟥+🟢+🟦 > 🟦, 🟥+🟢>0
      - 想要證明 🟥+🟢+🟦 > 🟢+🟦
      - 又已知前一個迭代若有 🟥+🟢 的可能，🟥 必須 > 0
      - 故 🟥+🟢+🟦 > 🟢+🟦
    - ***總結：🟢+🟦 不可能是連續最大值了，故不用考慮***

故整理一下得知，每次迭代需要比較的是：`localMax = max(localMax, current+localMax)`

並且再與 global maximum 比較紀錄：`globalMax = max(localMax, globalMax)`