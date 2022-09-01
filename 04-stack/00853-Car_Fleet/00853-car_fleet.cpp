// 853: Car Fleet
// https://leetcode.com/problems/car-fleet/


#include <iostream>
#include <vector>
#include <algorithm>
#include <stack>

using namespace std;


class Solution {
    class car {
    public:
        car(int p, int s): p(p), s(s) {};
        int p;
        int s;
    };
    
public:
    // SOLUTION
    int carFleet(int target, vector<int>& position, vector<int>& speed) {
        vector<car> cars;
        int n = position.size();
        for (int i=0; i<n; i++)
            cars.emplace_back(position[i], speed[i]);

        sort(cars.begin(), cars.end(), [](const car& a, car& b) {return a.p < b.p;});

        stack<float> s;
        for (int i=0; i<n; i++) {
            float time = (target - cars[i].p) / (float)cars[i].s;
            while (!s.empty() && time >= s.top()) s.pop();
            s.push(time);
        }

        return s.size();
    }
};


int main() {
    Solution o;

    // INPUT
    int target = 12;
    vector<int> position = {10,8,0,5,3};
    vector<int> speed = {2,4,1,1,3};

    // OUTPUT
    auto result = o.carFleet(target, position, speed);
    cout<<result<<endl;

    return 0;
}