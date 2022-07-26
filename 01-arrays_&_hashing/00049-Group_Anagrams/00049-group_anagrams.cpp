// 49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

#include <iostream>
#include <vector>
#include <unordered_map>
#include <algorithm>

using namespace std;

class Solution {
public:
    // SOLUTION
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> result;
        unordered_map<string, vector<string>> tmap;
        
        for(int  i=0; i<strs.size(); i++) {
            string str = strs[i];
            sort(strs[i].begin(), strs[i].end());
            tmap[strs[i]].push_back(str);
        }
        
        for(auto i=tmap.begin(); i!=tmap.end(); i++)
            result.push_back(i->second);
        
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<string> strs = {"eat","tea","tan","ate","nat","bat"};

    // OUTPUT
    auto result = o.groupAnagrams(strs);
    cout<<"[";
    for (auto g : result) {
        cout<< "[";
        for (auto s : g) {
            cout << "\"" << s << "\",";
        }
        cout << "\b],";
    }
    cout << "\b]" << endl;

    return 0;
}
