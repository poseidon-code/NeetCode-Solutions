// 347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/


import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


class Solution {
    // SOLUTION
    public int[] topKFrequent(int[] nums, int k) {
        List<Integer>[] buckets = new List[nums.length + 1];
        Map<Integer, Integer> tmap = new HashMap<Integer, Integer>();
    
        for (int n : nums) {
            tmap.put(n, tmap.getOrDefault(n, 0) + 1);
        }
    
        for (int key : tmap.keySet()) {
            int frequency = tmap.get(key);
            if (buckets[frequency] == null) {
                buckets[frequency] = new ArrayList<>();
            }
            buckets[frequency].add(key);
        }
    
        List<Integer> result = new ArrayList<>();
        for (int pos = buckets.length - 1; pos >= 0 && result.size() < k; pos--) {
            if (buckets[pos] != null) {
                result.addAll(buckets[pos]);
            }
        }

        return result.stream().mapToInt(i -> i).toArray();
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int nums[] = {1,1,1,2,2,3};
        int k = 2;

        // OUTPUT
        var result = o.topKFrequent(nums, k);
        System.out.print("["); for (var v : result) System.out.print(v+","); System.out.println("\b]");
    }
}
