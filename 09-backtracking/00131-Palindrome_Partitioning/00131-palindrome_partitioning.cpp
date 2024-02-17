// 131: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning



#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    void backtrack(string s, int start, vector<string>& current, vector<vector<string>>& result) {
        if (start == s.size()) {
            result.push_back(current);
            return;
        }

        for (int i=start; i<s.size(); i++) {
            if (isPalindrome(s, start, i)) {
                string str = s.substr(start, i - start + 1);
                current.push_back(str);
                backtrack(s, i + 1, current, result);
                current.pop_back();
            }
        }
    }

    bool isPalindrome(string s, int left, int right) {
        while (left < right) {
            if (s[left] != s[right]) return false;
            left++;
            right--;
        }

        return true;
    }

public:
    // SOLUTOIN
    vector<vector<string>> partition(string s) {
        vector<string> current;
        vector<vector<string>> result;
        backtrack(s, 0, current, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    string s = "aab";

    // OUTPUT
    auto result = o.partition(s);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << y << ",";
        } cout << "\b],";
    } cout << "\b]" << endl;

    return 0;
}
