// 853: Car Fleet
// https://leetcode.com/problems/car-fleet/


import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.Stack;


class Solution {
    static public class car {
        public car(int p, int s) {this.p = p; this.s = s;};
        public int p;
        public int s;
    }

    // SOLUTION
    public int carFleet(int target, int[] position, int[] speed) {
        ArrayList<car> cars = new ArrayList<>();
        int n = position.length;
        for (int i=0; i<n; i++)
            cars.add(new car(position[i], speed[i]));

        Collections.sort(cars, new Comparator<car>() {
            public int compare(car a, car b) {
                return a.p - b.p;
            }
        });

        Stack<Double> s = new Stack<>();
        for (int i=0; i<n; i++) {
            double time = (target - cars.get(i).p) / (double)cars.get(i).s;
            while (!s.isEmpty() && time >= s.peek()) s.pop();
            s.push(time);
        }

        return s.size();
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int target = 12;
        int[] position = {10,8,0,5,3};
        int[] speed = {2,4,1,1,3};

        // OUTPUT
        var result = o.carFleet(target, position, speed);
        System.out.println(result);
    }
}