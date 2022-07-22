// 242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/

#include <iostream>

using namespace std;

class Solution {
public:
    // SOLUTION
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) return false;

        int n = s.length();
        int counts[26] = {0};
        
        for (int i=0; i<n; i++) {
            counts[s[i] - 'a']++;
            counts[t[i] - 'a']--;
        }
        
        for (int i=0; i<26; i++)
            if (counts[i]) return false;
        
        return true;
    }
};

int main() {
    Solution o;

    // INPUT
    string s = "rat";
    string t = "cat";

    // OUTPUT
    auto result = o.isAnagram(s, t);
    cout << (result ? "true" : "false") << endl;

    return 0;
}
