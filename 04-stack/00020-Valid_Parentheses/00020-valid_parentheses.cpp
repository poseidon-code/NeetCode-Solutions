// 20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/


#include <iostream>
#include <stack>

using namespace std;


class Solution {
public:
    // SOLUTION
    bool isValid (string s) {
        stack<char> parentheses;

        for (auto &c : s) {
            switch(c) {
                case '{': parentheses.push('}'); break;
                case '[': parentheses.push(']'); break;
                case '(': parentheses.push(')'); break;
                default:
                    if (parentheses.size()==0 || c!=parentheses.top())
                        return false;
                    else parentheses.pop();
            }
        }

        return parentheses.empty();
    }
};


int main() {
    Solution o;

    // INPUT
    string s = "()[]{}";

    // OUTPUT
    auto result = o.isValid(s);
    cout<<(result ? "true" : "false")<<endl;

    return 0;
}