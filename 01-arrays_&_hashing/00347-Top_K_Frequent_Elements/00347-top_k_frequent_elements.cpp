// 347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/


#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> tmap;
        for (int num : nums) ++tmap[num];
        
        vector<vector<int>> buckets(nums.size() + 1); 
        for (auto p : tmap) buckets[p.second].push_back(p.first);
        
        vector<int> result;
        for (int i=buckets.size()-1; i>=0 && result.size()<k; --i) {
            for (int num : buckets[i]) {
                result.push_back(num);
                if (result.size() == k) break;
            }
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> nums = {1,1,1,2,2,3};
    int k = 2;

    // OUTPUT
    auto result = o.topKFrequent(nums, k);
    cout << "["; for (auto v : result) cout << v << ","; cout << "\b]" << endl;
    
    return 0;
}
