// 22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    vector<string> generateParentheses(int n) {
        vector<string> result;
        generate(n, 0, 0, "", result);
        return result;
    }

    void generate(int n, int open, int close, string s, vector<string>& result) {
        if (open==n && close==n) {
            result.push_back(s);
            return;
        }
        if (open < n)
            generate(n, open+1, close, s+'(', result);
        if (open > close)
            generate(n, open, close+1, s+')', result);
    }
};


int main() {
    Solution o;
    
    // INPUT
    int n = 3;

    // OUTPUT
    auto result = o.generateParentheses(n);
    cout<<"["; for (auto v : result) cout<<"\""<<v<<"\" "; cout<<"\b]"<<endl;

    return 0;
}