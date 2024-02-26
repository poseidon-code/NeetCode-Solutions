// 17: Letter Combinations Of A Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number


#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
private:
    void backtrack(string digits, int index, unordered_map<char, string>& m, string& current, vector<string>& result) {
        if (index == digits.size()) {
            result.push_back(current);
            return;
        }

        string s = m[digits[index]];
        for (int i=0; i<s.size(); i++) {
            current.push_back(s[i]);
            backtrack(digits, index + 1, m, current, result);
            current.pop_back();
        }
    }


public:
    // SOLUTOIN
   vector<string> letterCombinations(string digits) {
        if (digits.empty()) {
            return {};
        }
        
        unordered_map<char, string> m = {
            {'2', "abc"},
            {'3', "def"},
            {'4', "ghi"},
            {'5', "jkl"},
            {'6', "mno"},
            {'7', "pqrs"},
            {'8', "tuv"},
            {'9', "wxyz"}
        };
        
        string current;
        vector<string> result;
        backtrack(digits, 0, m, current, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    string digits = "23";

    // OUTPUT
    auto result = o.letterCombinations(digits);

    cout << "["; for (auto x : result) {
        cout << "\"" << x << "\", ";
    } cout << "\b\b]" << endl;

    return 0;
}
