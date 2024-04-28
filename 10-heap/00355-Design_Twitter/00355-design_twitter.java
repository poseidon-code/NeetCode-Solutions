// 355: Design Twitter
// https://leetcode.com/problems/design-twitter

import java.util.HashMap;
import java.util.HashSet;
import java.util.ArrayList;
import java.util.List;
import java.util.PriorityQueue;



class Twitter {
    int count;
    HashMap<Integer, List<int[]>> tweet_map;
    HashMap<Integer, HashSet<Integer>> follower_map;

    public Twitter() {
        count = 0;
        tweet_map = new HashMap<>();
        follower_map = new HashMap<>();
    }

    public void postTweet(int user_id, int tweet_id) {
        tweet_map.computeIfAbsent(user_id,
            k -> new ArrayList<>()
        );

        tweet_map.computeIfPresent(user_id, (k, v) -> {
            v.add(new int[]{count, tweet_id});
            return v;
        });

        ++count;
    }

    public List<Integer> getNewsFeed(int user_id) {
        List<Integer> result = new ArrayList<>();

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(b[0], a[0]));

        follower_map.computeIfAbsent(user_id, k -> new HashSet<>());
        follower_map.get(user_id).add(user_id);

        follower_map.get(user_id).forEach((followee_id) -> {
            if (tweet_map.containsKey(followee_id)) {
                int index = tweet_map.get(followee_id).size() - 1;
                int[] tweet = tweet_map.get(followee_id).get(index);

                pq.offer(new int[]{tweet[0], tweet[1], followee_id, --index});
            }
        });

        while (!pq.isEmpty() && result.size() < 10) {
            int[] data = pq.poll();
            result.add(data[1]);

            if (data[3] >= 0) {
                int[] tweet = tweet_map.get(data[2]).get(data[3]);
                pq.offer(new int[]{tweet[0], tweet[1], data[2], --data[3]});
            }
        }

        return result;
    }

    public void follow(int followerId, int followee_id) {
        follower_map.computeIfAbsent(followerId, k -> new HashSet<>());

        follower_map.computeIfPresent(followerId, (k, v) -> {
            v.add(followee_id);
            return v;
        });
    }

    public void unfollow(int followerId, int followee_id) {
        HashSet<Integer> set = follower_map.computeIfAbsent(followerId, k -> new HashSet<>());

        if (set.contains(followee_id)) {
            set.remove(followee_id);
        }
    }

    public static void main(String[] args) {
        Twitter o = new Twitter();

        // OPERATIONS
        o.postTweet(1, 5);
        System.out.println(o.getNewsFeed(1));
        o.follow(1, 2);
        o.postTweet(2, 6);
        System.out.println(o.getNewsFeed(1));
        o.unfollow(1, 2);
        System.out.println(o.getNewsFeed(1));
    }
}
