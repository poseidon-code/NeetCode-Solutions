// 76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/


#include <iostream>
#include <unordered_map>

using namespace std;


class Solution {
public:
    // SOLUTION
    string minWindow(string s, string t) {
        unordered_map<char, int> m;
        for (int i=0; i<t.size(); i++) m[t[i]]++;

        int i = 0;
        int j = 0;
        int counter = t.size();
        
        int minStart = 0;
        int minLength = INT16_MAX;
        
        while (j < s.size()) {
            if (m[s[j]] > 0) counter--;
            m[s[j]]--;
            j++;

            while (counter==0) {
                if (j - i < minLength) {
                    minStart = i;
                    minLength = j-i;
                }
                
                m[s[i]]++;
                if (m[s[i]] > 0) counter++;
                i++;
            }
        }
        
        if (minLength != INT16_MAX)
            return s.substr(minStart, minLength);
        return "";
    }
};


int main() {
    Solution o;

    // INPUT
    string s = "ADOBECODEBANC";
    string t = "ABC";

    // OUTPUT
    auto result = o.minWindow(s, t);
    cout<<result<<endl;

    return 0;
}