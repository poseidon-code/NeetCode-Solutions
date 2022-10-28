// 981: Time Based Key-Value Store
// https://leetcode.com/problems/time-based-key-value-store/


import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;


class Solution {
    static class Data {
        String val;
        int time;
        Data(String val, int time) {
            this.val = val;
            this.time = time;
        }
    }

    // SOLUTION
    static class TimeMap {
        HashMap<String, List<Data>> tm;

        public TimeMap() {
            tm = new HashMap<>();
        }

        public void set(String key, String value, int timestamp) {
            if (!tm.containsKey(key)) tm.put(key, new ArrayList<Data>());
            tm.get(key).add(new Data(value, timestamp));
        }
        
        public String get(String key, int timestamp) {
            if (!tm.containsKey(key)) return "";
            List<Data> list = tm.get(key);
            
            int l = 0;
            int h = list.size() - 1;

            while (l < h) {
                int m = l + (h-l+1) / 2;
                
                if (list.get(m).time <= timestamp) l = m;
                else h = m - 1;
            }

            return list.get(l).time <= timestamp ? list.get(l).val : "";
        }
    }

    public static void main(String[] args) {
        // OPERATIONS
        TimeMap o = new TimeMap();
        o.set("foo", "bar", 1);
        System.out.println(o.get("foo", 1));
        System.out.println(o.get("foo", 3));
        o.set("foo", "bar2", 4);
        System.out.println(o.get("foo", 4));
        System.out.println(o.get("foo", 5));
    }
}