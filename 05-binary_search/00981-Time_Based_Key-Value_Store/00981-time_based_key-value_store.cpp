// 981: Time Based Key-Value Store
// https://leetcode.com/problems/time-based-key-value-store/


#include <iostream>
#include <unordered_map>
#include <vector>
#include <string>

using namespace std;


// SOLUTION
class TimeMap {
private:
    unordered_map<string, vector<pair<int, string>>> tm;

public:
    TimeMap() {}

    void set(string key, string value, int timestamp) {
        tm[key].push_back({timestamp, value});
    }
    
    string get(string key, int timestamp) {
        if (tm.find(key) == tm.end()) return "";

        int l = 0;
        int h = tm[key].size() - 1;

        while (l <= h) {
            int m = l + (h - l) / 2;

            if (tm[key][m].first < timestamp)
                l = m + 1;
            else if (tm[key][m].first > timestamp)
                h = m - 1;
            else
                return tm[key][m].second;
        }
        
        if (h >= 0) return tm[key][h].second;

        return "";
    }
};


int main() {
    // OPERATIONS
    TimeMap* o = new TimeMap();
    o->set("foo", "bar", 1);
    cout << o->get("foo", 1) << endl;
    cout << o->get("foo", 3) << endl;
    o->set("foo", "bar2", 4);
    cout << o->get("foo", 4) << endl;
    cout << o->get("foo", 5) << endl;

    return 0;
}