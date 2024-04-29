# 355: Design Twitter
# https://leetcode.com/problems/design-twitter

import heapq
from collections import defaultdict

from typing import List

# SOLUTION
class Twitter:
    def __init__(self):
        self.count = 0
        self.tweet_map = defaultdict(list)
        self.follow_map = defaultdict(set)

    def postTweet(self, user_id: int, tweet_id: int) -> None:
        self.tweet_map[user_id].append([self.count, tweet_id])
        self.count -= 1

    def getNewsFeed(self, user_id: int) -> List[int]:
        result = []
        min_heap = []

        self.follow_map[user_id].add(user_id)
        for followee_id in self.follow_map[user_id]:
            if followee_id in self.tweet_map:
                index = len(self.tweet_map[followee_id]) - 1
                count, tweet_id = self.tweet_map[followee_id][index]
                heapq.heappush(min_heap, [count, tweet_id, followee_id, index - 1])

        while min_heap and len(result) < 10:
            count, tweet_id, followee_id, index = heapq.heappop(min_heap)
            result.append(tweet_id)
            if index >= 0:
                count, tweet_id = self.tweet_map[followee_id][index]
                heapq.heappush(min_heap, [count, tweet_id, followee_id, index - 1])
        
        return result

    def follow(self, follower_id: int, followee_id: int) -> None:
        self.follow_map[follower_id].add(followee_id)

    def unfollow(self, follower_id: int, followee_id: int) -> None:
        if followee_id in self.follow_map[follower_id]:
            self.follow_map[follower_id].remove(followee_id)


if __name__ == "__main__":
    o = Twitter()

    # OPERATIONS
    o.postTweet(1, 5)
    print(o.getNewsFeed(1))
    o.follow(1, 2)
    o.postTweet(2, 6)
    print(o.getNewsFeed(1))
    o.unfollow(1, 2)
    print(o.getNewsFeed(1))
