// 125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/


#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
    // SOLUTION
    bool isPalindrome(string s) {
        int l = 0;
        int r = s.size() - 1;

        while (l<r) {
            while (!isalnum(s[l]) && l<r) l++;
            while (!isalnum(s[r]) && l<r) r--;
            if (tolower(s[l]) != tolower(s[r])) return false;
            
            l++;
            r--;
        }

        return true;
    }
};


int main() {
    Solution o;
    
    // INPUT
    string s = "race a car";

    // OUTPUT
    auto result = o.isPalindrome(s);
    cout<<(result ? "true" : "false")<<endl;

    return 0;
}