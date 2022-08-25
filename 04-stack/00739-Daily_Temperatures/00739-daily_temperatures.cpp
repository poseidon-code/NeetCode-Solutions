// 739: Daily Temperstures
// https://leetcode.com/problems/daily-temperatures/


#include <iostream>
#include <vector>
#include <stack>

using namespace std;


class Solution {
public:
    // SOLUTION
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        stack<int> s;
        vector<int> result(temperatures.size());

        for (int i=0; i<temperatures.size(); i++) {
            while (!s.empty() && temperatures[i]>temperatures[s.top()]) {
                result[s.top()] = i - s.top();
                s.pop();
            }
            s.push(i);
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> temperatures = {73,74,75,71,69,72,76,73};

    // OUTPUT
    auto result = o.dailyTemperatures(temperatures);
    cout<<"["; for (auto v : result) cout<<v<<" "; cout<<"\b]"<<endl;

    return 0;
}