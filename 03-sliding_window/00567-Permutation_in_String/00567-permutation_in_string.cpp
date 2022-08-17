// 567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    bool checkInclusion(string s1, string s2) {
        if (s1.size() > s2.size()) return false;
        
        vector<int> s1Count(26); for (auto c : s1) s1Count[c-'a']++;
        vector<int> s2Count(26);
        int i = 0;
        int j = 0;
        
        while (j < s2.size()) {
            s2Count[s2[j]-'a']++;
            
            if (j-i+1 == s1.size()) 
                if (s1Count == s2Count) 
                    return true;
            if (j-i+1 < s1.size()) 
                j++;
            else {
                s2Count[s2[i]-'a']--;
                i++;
                j++;
            }
        }

        return false;
    }
};


int main() {
    Solution o;

    // INPUT
    string s1 = "ab";
    string s2 = "eidbaooo";

    // OUTPUT
    auto result = o.checkInclusion(s1, s2);
    cout<<(result ? "true" : "false")<<endl;

    return 0;
}