// 973: K Closest Points To Origin
// https://leetcode.com/problems/k-closest-points-to-origin


#include <iostream>
#include <queue>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        vector<vector<int>> triples;
        for (auto& p : points)
            triples.push_back({p[0] * p[0] + p[1] * p[1], p[0], p[1]});

        priority_queue<vector<int>, vector<vector<int>>, greater<vector<int>>> pq(triples.begin(), triples.end());
        vector<vector<int>> result;

        while (k--){
            vector<int> e = pq.top();
            pq.pop();
            result.push_back({e[1], e[2]});
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<vector<int>> points = {{1,3},{-2,2}};
    int k = 1;

    // OUTPUT
    auto result = o.kClosest(points, k);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}