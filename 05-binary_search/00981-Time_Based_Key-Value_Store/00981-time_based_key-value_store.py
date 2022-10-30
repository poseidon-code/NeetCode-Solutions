# 981: Time Based Key-Value Store
# https://leetcode.com/problems/time-based-key-value-store/


import collections

class TimeMap(object):
    def __init__(self):
        self.tm = collections.defaultdict(list)

    def set(self, key, value, timestamp):
        self.tm[key].append([timestamp, value])

    def get(self, key, timestamp):
        tlist = self.tm[key]

        l = 0
        h = len(tlist)
        
        while l < h:
            m = (l + h) // 2
            
            if tlist[m][0] <= timestamp:
                l = m + 1
            elif tlist[m][0] > timestamp:
                h = m
        
        return "" if h == 0 else tlist[h - 1][1]



if __name__ == "__main__":
    # OPERATIONS
    o = TimeMap()
    o.set("foo", "bar", 1);
    print(o.get("foo", 1));
    print(o.get("foo", 3));
    o.set("foo", "bar2", 4);
    print(o.get("foo", 4));
    print(o.get("foo", 5));
