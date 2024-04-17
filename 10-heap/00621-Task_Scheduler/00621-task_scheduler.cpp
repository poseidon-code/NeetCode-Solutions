// 621: Task Scheduler
// https://leetcode.com/problems/task-scheduler


#include <vector>
#include <iostream>
#include <unordered_map>
#include <queue>

using namespace std;

class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        unordered_map<char, int> counts;
        for (char t : tasks) counts[t]++;

        priority_queue<int> pq;
        for (pair<char, int> count : counts)
            pq.push(count.second);

        int result = 0;
        int cycle = n + 1;

        while (!pq.empty()) {
            int time = 0;
            vector<int> t;

            for (int i = 0; i < cycle; i++) {
                if (!pq.empty()) {
                    t.push_back(pq.top());
                    pq.pop();
                    time++;
                }
            }

            for (int cnt : t)
                if (--cnt)
                    pq.push(cnt);

            result += !pq.empty() ? cycle : time;
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<char> tasks = {'A','C','A','B','D','B'};
    int n = 1;

    // OUTPUT
    auto result = o.leastInterval(tasks, n);
    cout << result << endl;

    return 0;
}
