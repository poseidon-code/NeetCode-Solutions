// 49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    // SOLUTION
    public List<List<String>> groupAnagrams(String[] strs) {
        if (strs == null || strs.length == 0) return new ArrayList<>();

        Map<String, List<String>> tmap = new HashMap<>();

        for (String s : strs) {
            char[] ca = new char[26];
            for (char c : s.toCharArray()) ca[c - 'a']++;
            String keyStr = String.valueOf(ca);
            if (!tmap.containsKey(keyStr)) tmap.put(keyStr, new ArrayList<>());
            tmap.get(keyStr).add(s);
        }

        return new ArrayList<>(tmap.values());
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        String[] strs = {"eat","tea","tan","ate","nat","bat"};

        // OUTPUT
        var result = o.groupAnagrams(strs);
        System.out.print(result);
    }
}
