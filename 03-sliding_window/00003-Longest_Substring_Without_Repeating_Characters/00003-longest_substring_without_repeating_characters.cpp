// 3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/


#include <iostream>
#include <unordered_set>

using namespace std;

class Solution {
public:
    // SOLUTION
    int lengthOfLongestSubstring(string s) {
        unordered_set<char> letters;
        int i=0, j=0;
        int result = 0;

        while (j<s.size()) {
            if (letters.find(s[j]) == letters.end()) {
                letters.insert(s[j++]);
                result = max(result, j-i);
            } else {
                letters.erase(s[i++]);
            }
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    string s = "abcabcbb";

    // OUTPUT
    auto result = o.lengthOfLongestSubstring(s);
    cout<<result<<endl;

    return 0;
}