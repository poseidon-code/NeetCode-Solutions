// 424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // SOLUTION
    int characterReplacement(string s, int k) {
        vector<int> count(26);
        int maxCount = 0;
        int i = 0;
        int j = 0;

        int result = 0;
        
        while (j<s.size()) {
            count[s[j]-'A']++;
            maxCount = max(maxCount, count[s[j]-'A']);
            
            if (j-i+1-maxCount>k) {
                count[s[i]-'A']--;
                i++;
            }

            result = max(result, j-i+1);
            j++;
        }
        
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    string s = "ABAB";
    int k = 2;

    // OUTPUT
    auto result = o.characterReplacement(s, k);
    cout<<result<<endl;

    return 0;
}