// 150: Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/


#include <iostream>
#include <vector>
#include <stack>

using namespace std;


class Solution {
public:
    // SOLUTION
    int evalRPN(vector<string>& tokens) {
        stack<int> result;

        for (auto s : tokens) {
            if (s.size()>1 || isdigit(s[0])) {
                result.push(stoi(s));
            } else {
                auto x2 = result.top(); result.pop();
                auto x1 = result.top(); result.pop();
                
                switch (s[0]) {
                    case '+': x1+=x2; break;
                    case '-': x1-=x2; break;
                    case '*': x1*=x2; break;
                    case '/': x1/=x2; break;
                }

                result.push(x1);
            }
        }

        return result.top();
    }
};



int main() {
    Solution o;

    // INPUT
    vector<string> tokens = {"2","1","+","3","*"};

    // OUTPUT
    auto result = o.evalRPN(tokens);
    cout<<result<<endl;

    return 0;
}