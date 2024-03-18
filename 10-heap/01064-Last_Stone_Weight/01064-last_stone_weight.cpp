// 1064: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight


#include <iostream>
#include <queue>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int lastStoneWeight(vector<int> stones) {
        priority_queue<int> pq(stones.begin(),stones.end());
        while (pq.size() > 1) {
            int y = pq.top();
            pq.pop();
            int x = pq.top();
            pq.pop();

            if (x != y) pq.push(y-x);
        }

        return pq.empty() ? 0 : pq.top();
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> stones = {2,7,4,1,8,1};

    // OUTPUT
    auto result = o.lastStoneWeight(stones);
    cout << result << endl;

    return 0;
}
