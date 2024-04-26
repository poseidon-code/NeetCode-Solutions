// 355: Design Twitter
// https://leetcode.com/problems/design-twitter

#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <iostream>

using namespace std;

// SOLUTION
class Twitter {
private:
    vector<pair<int, int>> posts;
    unordered_map<int, unordered_set<int>> follow_map;

public:
    Twitter(){}

    void postTweet(int user_id, int tweet_id) {
        posts.push_back({user_id, tweet_id});
    }

    vector<int> getNewsFeed(int user_id) {
        int count = 10;
        vector<int> result;

        for (int i = posts.size() - 1; i >= 0; i--) {
            if (count == 0) {
                break;
            }

            int following_id = posts[i].first;
            int tweet_id = posts[i].second;
            unordered_set<int> following = follow_map[user_id];

            if (following.find(following_id) != following.end() || following_id == user_id) {
                result.push_back(tweet_id);
                count--;
            }
        }

        return result;
    }

    void follow(int follower_id, int followee_id) {
        follow_map[follower_id].insert(followee_id);
    }

    void unfollow(int follower_id, int followee_id) {
        follow_map[follower_id].erase(followee_id);
    }
};


int main() {
    Twitter *o = new Twitter();

    // OPERATIONS
    o->postTweet(1, 5);
    auto results = o->getNewsFeed(1);
    cout << "["; for (auto v: results) cout << v << ","; cout << "\b]\n";
	o->follow(1, 2);
	o->postTweet(2, 6);
	results = o->getNewsFeed(1);
    cout << "["; for (auto v: results) cout << v << ","; cout << "\b]\n";
	o->unfollow(1, 2);
	results = o->getNewsFeed(1);
    cout << "["; for (auto v: results) cout << v << ","; cout << "\b]\n";

    return 0;
}
